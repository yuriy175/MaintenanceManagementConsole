﻿using System;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;
using Atlas.Acquisitions.Common.Core;
using Atlas.Acquisitions.Common.Core.Model;
using Atlas.Common.Impls.Helpers;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using Atlas.Remoting.Core.Interfaces;
using MessagesSender.Core.Interfaces;
using MessagesSender.Core.Model;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using Serilog;

namespace MessagesSender.BL
{
    /// <summary>
    /// sending service interface implementation
    /// </summary>
    public class SendingService : ISendingService
    {
        private readonly IObservationsEntityService _dbObservationsEntityService;
        private readonly IEventPublisher _eventPublisher;
        private readonly ILogger _logger;
        private readonly IMQCommunicationService _mqService;
        private readonly IWorkqueueSender _wqSender;
        private readonly IMqttSender _mqttSender;
        private readonly IOfflineService _offlineService;
        private readonly ITopicService _topicService;

        private IPAddress _ipAddress = null;
        private (string Name, string Number, string HddNumber) _equipmentInfo = (null, null, null);

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="dbObservationsEntityService">observations database connector</param>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="mqService">MQ service</param>
        /// <param name="wqSender">work queue sender</param>
        /// <param name="mqttSender">mqtt sender</param>
        /// <param name="offlineService">offline service</param>
        /// <param name="topicService">topic service</param>
        public SendingService(
            IObservationsEntityService dbObservationsEntityService,
            ILogger logger,
            IEventPublisher eventPublisher,
            IMQCommunicationService mqService,
            IWorkqueueSender wqSender,
            IMqttSender mqttSender,
            IOfflineService offlineService,
            ITopicService topicService)
        {
            _dbObservationsEntityService = dbObservationsEntityService;
            _logger = logger;
            _mqService = mqService;
            _wqSender = wqSender;
            _mqttSender = mqttSender;
            _offlineService = offlineService;
            _eventPublisher = eventPublisher;
            _topicService = topicService;

            _eventPublisher.RegisterServerReadyCommandArrivedEvent(() => OnServerReadyArrivedAsync());

            _logger.Information("SendingService started");
        }

        /// <summary>
        /// creates service
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> CreateAsync()
        {
            await Task.WhenAll(new[]
                {
                    Task.Run(async () =>
                    {
                        // await GetEquipmentInfoAsync();
                        await _mqttSender.CreateAsync();
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
        /// <param name="msgType">messge type</param>
        /// <param name="info">info</param>
        /// <returns>result</returns>
        public async Task<bool> SendInfoToWorkQueueAsync<TMsgType, T>(TMsgType msgType, T info)
        {
            /*if (string.IsNullOrEmpty(_equipmentInfo.Number) || string.IsNullOrEmpty(_equipmentInfo.Name))
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
                */
            return true;
        }

        /// <summary>
        /// sends info to mqtt
        /// </summary>
        /// <typeparam name="TMsgType">message type</typeparam>
        /// <typeparam name="T">info type</typeparam>
        /// <param name="msgType">messge type</param>
        /// <param name="info">info</param>
        /// <returns>result</returns>
        public async Task<bool> SendInfoToMqttAsync<TMsgType, T>(TMsgType msgType, T info)
        {
            var result = await SendInfoAsync(_mqttSender, msgType, info);
            if (!result)
            {
                await _offlineService.CheckInfoAsync(msgType, info);
            }

            return result;
        }

        /// <summary>
        /// sends info to common mqtt
        /// </summary>
        /// <typeparam name="T">info type</typeparam>
        /// <param name="msgType">message type</param>
        /// <param name="info">info</param>
        /// <returns>result</returns>
        public async Task<bool> SendInfoToCommonMqttAsync<T>(MQMessages msgType, T info)
        {
            var result = await _mqttSender.SendCommonAsync(msgType, info);
            if (!result)
            {
                await _offlineService.CheckInfoAsync(msgType, info);
            }

            return result;
        }

        private async Task<bool> SendInfoAsync<TMsgType, T>(IMQSenderBase sender, TMsgType msgType, T info)
        {
            return await sender.SendAsync(
                msgType,
                info);
        }

        /*private async Task GetEquipmentInfoAsync()
        {
            _equipmentInfo = await _topicService.GetEquipmentInfoAsync();
        }*/

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

        private Task OnServerReadyArrivedAsync() => SendOfflinedInfosAsync();

        private async Task SendOfflinedInfosAsync()
        {
            var infos = (await _offlineService.GetInfosAsync())?.ToList();
            if (infos == null)
            {
                _logger.Information("no offline info");
                return;
            }

            var result = false;
            foreach (var info in infos)
            {
                result = await SendInfoAsync(_mqttSender, info.MsgType, info.Msg) || result;
            }

            if (result)
            {
                await _offlineService.ClearInfosAsync();
            }
        }
    }
}
