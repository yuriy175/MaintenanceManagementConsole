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

namespace MessagesSender.BL
{
    public class Service : IService
    {
        private readonly ISettingsEntityService _dbSettingsEntityService;
        private readonly ILogger _logger;
        private readonly IMQCommunicationService _mqService;
        private readonly IWorkqueueSender _wqSender;

        private IPAddress _ipAddress = null;
        private (string Name, string Number) _equipmentInfo = (null, null);

        private enum MessageType
        {
            StudyInWork = 1,
            ConnectionState,

        }

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="dbSettingsEntityService">settings database connector</param>
        /// <param name="logger">logger</param>
        /// <param name="mqService">MQ service</param>
        /// <param name="wqSender">work queue sender</param>
        public Service(
            ISettingsEntityService dbSettingsEntityService,
            ILogger logger,
            IMQCommunicationService mqService,
            IWorkqueueSender wqSender)
        {
            _dbSettingsEntityService = dbSettingsEntityService;
            _logger = logger;
            _mqService = mqService;
            _wqSender = wqSender;

            new Action[]
                {
                    () => _ = SubscribeMQRecevers(),
                    () => _ = GetEquipmentInfoAsync(),
                    () => _ = GetEquipmentIPAsync(),
                }.RunTasksAsync();

            _logger.Information("Main service started");
        }

        private async Task SubscribeMQRecevers()
        {
            _mqService.Subscribe<MQCommands, int>(
                    (MQCommands.StudyInWork, async data => OnStudyInWorkAsync(data)));

            _mqService.Subscribe<MQCommands, (int Id, string Name, string Type, DeviceConnectionState Connection)>(
                    (MQCommands.HwConnectionStateArrived, state => OnConnectionStateArrivedAsync(state)));

            _mqService.Subscribe<MQCommands, int>(
                    (MQCommands.NewImageCreated, async imageId => OnNewImageCreatedAsync(imageId)));
        }

        private async Task GetEquipmentInfoAsync()
        {
            _equipmentInfo = await _dbSettingsEntityService.GetEquipmentInfo();
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

        private async Task<bool> OnStudyInWorkAsync(int studyId)
        {
            _ = SendInfoAsync(MQCommands.StudyInWork, studyId);
            return true;
        }

        private async Task<bool> OnNewImageCreatedAsync(int imageId)
        {
            _ = SendInfoAsync(MQCommands.NewImageCreated, imageId);
            return true;
        }

        private async Task<bool> OnConnectionStateArrivedAsync(
            (int Id, string Name, string Type, DeviceConnectionState Connection) state)
        {
            _ = SendInfoAsync(MQCommands.HwConnectionStateArrived, 
                new { state.Id, state.Name, state.Type, state.Connection });
            return true;
        }

        private async Task SendInfoAsync<T>(MQCommands msgType, T info)
        {
            if (string.IsNullOrEmpty(_equipmentInfo.Number) || string.IsNullOrEmpty(_equipmentInfo.Name))
            {
                _logger.Error($"wrong equipment props {_equipmentInfo.Number} {_equipmentInfo.Name}");
                return;
            }

            await _wqSender.SendAsync(
                new { 
                    _equipmentInfo.Number, 
                    _equipmentInfo.Name, 
                    ipAddress = _ipAddress?.ToString(),
                    msgType = msgType.ToString(),
                    info,
                });
        }
    }
}
