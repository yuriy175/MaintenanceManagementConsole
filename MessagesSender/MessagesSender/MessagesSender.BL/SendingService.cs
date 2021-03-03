using System;
using System.Threading.Tasks;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using Atlas.Remoting.Core.Interfaces;
using MessagesSender.Core.Interfaces;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;
using Serilog;
using Atlas.Common.Impls.Helpers;
using System.Net;
using System.Linq;
using System.Net.Sockets;
using Atlas.Acquisitions.Common.Core;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System.Text;
using MessagesSender.Core.Model;
using Atlas.Acquisitions.Common.Core.Model;

namespace MessagesSender.BL
{
    public class SendingService : ISendingService
    {
        private readonly ISettingsEntityService _dbSettingsEntityService;
        private readonly IObservationsEntityService _dbObservationsEntityService;        
        private readonly ILogger _logger;
        private readonly IMQCommunicationService _mqService;
        private readonly IWorkqueueSender _wqSender;
        private readonly IMqttSender _mqttSender;
        private readonly IHardwareStateService _hwStateService;

        private IPAddress _ipAddress = null;
        private (string Name, string Number) _equipmentInfo = (null, null);

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="dbSettingsEntityService">settings database connector</param>
        /// <param name="dbObservationsEntityService">observations database connector</param>
        /// <param name="logger">logger</param>
        /// <param name="mqService">MQ service</param>
        /// <param name="wqSender">work queue sender</param>
        /// <param name="mqttSender">mqtt sender</param>
        public SendingService(
            ISettingsEntityService dbSettingsEntityService,
            IObservationsEntityService dbObservationsEntityService,
            ILogger logger,
            IMQCommunicationService mqService,
            IWorkqueueSender wqSender,
            IMqttSender mqttSender)
        {
            _dbSettingsEntityService = dbSettingsEntityService;
            _dbObservationsEntityService = dbObservationsEntityService;
            _logger = logger;
            _mqService = mqService;
            _wqSender = wqSender;
            _mqttSender = mqttSender;            

            _logger.Information("Main service started");
        }

        /// <summary>
        /// creates service
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> CreateAsync()
        {
            await Task.WhenAll(new []
                {
                    Task.Run(async () =>
                    {
                        await GetEquipmentInfoAsync();
                        await _mqttSender.CreateAsync(_equipmentInfo);
                    }),
                    Task.Run(() => _ = GetEquipmentIPAsync()),
                });

            return true;
        }

        /// <summary>
        /// sends info to workqueue
        /// </summary>
        /// <typeparam name="TMsgType">message type</typeparam>
        /// <typeparam name="T">info type</typeparam>
        /// <param name="msgType">info type</param>
        /// <param name="info">info</param>
        /// <returns>result</returns>
        public async Task<bool> SendInfoToWorkQueueAsync<TMsgType, T>(TMsgType msgType, T info)
        {
			if (string.IsNullOrEmpty(_equipmentInfo.Number) || string.IsNullOrEmpty(_equipmentInfo.Name))
			{
				_logger.Error($"wrong equipment props {_equipmentInfo.Number} {_equipmentInfo.Name}");
				return false;
			}

			return await SendInfoAsync(_wqSender, msgType,
				new
				{
					_equipmentInfo.Number,
					_equipmentInfo.Name,
					ipAddress = _ipAddress?.ToString(),
					msgType = msgType.ToString(),
					info,
				});
        }

        /// <summary>
        /// sends info to mqtt
        /// </summary>
        /// <typeparam name="TMsgType">message type</typeparam>
        /// <typeparam name="T">info type</typeparam>
        /// <param name="msgType">info type</param>
        /// <param name="info">info</param>
        /// <returns>result</returns>
        public async Task<bool> SendInfoToMqttAsync<TMsgType, T>(TMsgType msgType, T info)
        {
            return await SendInfoAsync(_mqttSender, msgType, info);
        }

		/// <summary>
		/// sends info to common mqtt
		/// </summary>
		/// <typeparam name="T">info type</typeparam>
		/// <param name="msgType">info type</param>
		/// <param name="info">info</param>
		/// <returns>result</returns>
		public async Task<bool> SendInfoToCommonMqttAsync<T>(MQMessages msgType, T info)
		{
			return await _mqttSender.SendCommonAsync(msgType, info);
		}

		private async Task<bool> SendInfoAsync<TMsgType, T>(IMQSenderBase sender, TMsgType msgType, T info)
        {
			return await sender.SendAsync(
				msgType,
				info);
				/*new { 
                    _equipmentInfo.Number, 
                    _equipmentInfo.Name, 
                    ipAddress = _ipAddress?.ToString(),
                    msgType = msgType.ToString(),
                    info,
                });*/
		}

        private async Task GetEquipmentInfoAsync()
        {
            _equipmentInfo = await _dbSettingsEntityService.GetEquipmentInfoAsync();
        }

        private async Task GetEquipmentIPAsync()
        {
            if (!System.Net.NetworkInformation.NetworkInterface.GetIsNetworkAvailable())
            {
                return;
            }
            IPHostEntry host = Dns.GetHostEntry(Dns.GetHostName());
            _ipAddress = host
               .AddressList
               .FirstOrDefault(ip => ip.AddressFamily == AddressFamily.InterNetwork);
        }
    }
}
