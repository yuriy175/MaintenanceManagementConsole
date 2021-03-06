﻿using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;
using Atlas.Common.Core.Interfaces;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using MessagesSender.Core.Interfaces;
using MessagesSender.Core.Model;
using MQTTnet;
using MQTTnet.Client;
using MQTTnet.Extensions.ManagedClient;
using Newtonsoft.Json;
using RabbitMQ.Client;
using Serilog;

namespace MessagesSender.BL.Remoting
{
    /// <summary>
    /// RabbitMQ mqtt sender class
    /// </summary>
    public class RabbitMQTTSender : RabbitMQTTBase, IMqttSender
    {
        private const string CommonTopic = "Subscribe";
        private const string CommandSubTopic = "/command";
        private const string BroadcastCommandsTopic = "Broadcast/command";

        private const int ConnectWaitingAttempts = 5;

        private readonly ILogger _logger;
        private readonly IEventPublisher _eventPublisher;
        private readonly ITopicService _topicService;

        private readonly Dictionary<string, string> _topicMap = new Dictionary<string, string>
        {
            { MQCommands.StudyInWork.ToString(), "/study" },
            { MQCommands.GeneratorStateArrived.ToString(), "/generator/state" },
            { MQCommands.StandStateArrived.ToString(), "/stand/state" },
            { MQCommands.CollimatorStateArrived.ToString(), "/collimator/state" },
            { MQCommands.DetectorStateArrived.ToString(), "/detector/state" },
            { MQCommands.AecStateArrived.ToString(), "/aec/state" },
            { MQMessages.HddDrivesInfo.ToString(), "/ARM/Hardware/HDD" },
            { MQMessages.CPUInfo.ToString(), "/ARM/Hardware/Processor" },
            { MQMessages.MemoryInfo.ToString(), "/ARM/Hardware/Memory" },
            { MQMessages.AllDBInfo.ToString(), "/ARM/AllDBInfo" },
            { MQCommands.SetOrganAuto.ToString(), "/organauto" },
            { MQCommands.ProcessDoseArrived.ToString(), "/dosimeter/state" },
            { MQMessages.DicomInfo.ToString(), "/dicom" },
            { MQMessages.SoftwareInfo.ToString(), "/ARM/Software" },
            { MQMessages.SoftwareMsgInfo.ToString(), "/ARM/Software/msg" },            
            { MQMessages.RemoteAccess.ToString(), "/remoteaccess" },
            { MQMessages.ImagesInfo.ToString(), "/images" },
            { MQMessages.Events.ToString(), "/events" },
            { MQMessages.HospitalInfo.ToString(), "/hospital" },            
        };

        /// <summary>
        /// Initializes a new instance of the <see cref="RabbitMQBase"/> class.
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="topicService">topic service</param>
        public RabbitMQTTSender(
            IConfigurationService configurationService,
            ILogger logger,
            IEventPublisher eventPublisher,
            ITopicService topicService)
            : base(configurationService, logger)
        {
            _logger = logger;
            _eventPublisher = eventPublisher;
            _topicService = topicService;
        }

        /// <summary>
        /// sends a message
        /// </summary>
        /// <typeparam name="TMsg">message type</typeparam>
        /// <typeparam name="T">entity type</typeparam>
        /// <param name="msgType">messge type</param>
        /// <param name="payload">entity</param>
        /// <returns>result</returns>
        public async Task<bool> SendAsync<TMsg, T>(TMsg msgType, T payload)
        {
            if (!Created)
            {
                return false;
            }

            var msgTypeKey = msgType.ToString();
            var subtopic = _topicMap.ContainsKey(msgTypeKey) ? _topicMap[msgTypeKey] : string.Empty;
            if (string.IsNullOrEmpty(subtopic))
            {
                return false;
            }

            var content = JsonConvert.SerializeObject(
                payload,
                new JsonSerializerSettings { NullValueHandling = NullValueHandling.Ignore });

            return await SendAsync(msgType, payload, $"{Topic}{subtopic}", content);
        }

        /// <summary>
        /// sends a message to a common mqtt
        /// </summary>
        /// <typeparam name="T">entity type</typeparam>
        /// <param name="msgType">message type</param>
        /// <param name="payload">entity</param>
        /// <returns>result</returns>
        public async Task<bool> SendCommonAsync<T>(MQMessages msgType, T payload)
        {
            if (!Created)
            {
                return false;
            }

            var content = string.Empty;
            if (msgType.IsStateMQMessage())
            {
                content = "{\"" + Topic + "\" : " + (msgType == MQMessages.InstanceOff ? "\"off\"" : "\"on\"") + "}";
            }
            else
            {
                content = JsonConvert.SerializeObject(
                    payload,
                    new JsonSerializerSettings { NullValueHandling = NullValueHandling.Ignore });
            }

            return await SendAsync(msgType, payload, CommonTopic, content);
        }

        /// <inheritdoc/>
        protected override Task<string> GetTopicAsync()
            => _topicService.GetTopicAsync();

        /// <inheritdoc/>
        protected override async Task<IManagedMqttClient> CreateConnection(ConnectionFactory connectionFactory)
        {
            try
            {
                var client = await base.CreateConnection(connectionFactory);
                if (client == null)
                {
                    return null;
                }

                Client.UseApplicationMessageReceivedHandler(e =>
                {
                    try
                    {
                        string topic = e.ApplicationMessage.Topic;

                        if (string.IsNullOrWhiteSpace(topic) == false)
                        {
                            string command = Encoding.UTF8.GetString(e.ApplicationMessage.Payload);
                            _eventPublisher.MqttCommandArrived(command);
                            Console.WriteLine($"Topic: {topic}. Message Received: {command}");
                        }
                    }
                    catch (Exception ex)
                    {
                        Console.WriteLine(ex.Message, ex);
                    }
                });

                await Client.SubscribeAsync(
                    new TopicFilterBuilder()
                        .WithTopic(BroadcastCommandsTopic)
                        .WithQualityOfServiceLevel((MQTTnet.Protocol.MqttQualityOfServiceLevel)0) // qos)
                        .Build(),
                    new TopicFilterBuilder()
                        .WithTopic(Topic + CommandSubTopic)
                        .WithQualityOfServiceLevel((MQTTnet.Protocol.MqttQualityOfServiceLevel)0) // qos)
                        .Build());
            }
            catch (Exception ex)
            {
                using (Client)
                {
                }

                // _logger.Error(ex, $"MQ connection error: { connectionFactory.HostName}, {connectionFactory.UserName}, {connectionFactory.Password}."); ;
                return null;
            }

            return null;
        }

        private async Task<bool> CheckConnectedAsync()
        {
            int attempts = 0;
            while (!Client.IsConnected && attempts++ < ConnectWaitingAttempts)
            {
                Console.WriteLine($"Sent to not connected {Topic}");
                await Task.Delay(Client.Options.ConnectionCheckInterval);
            }

            return Client.IsConnected;
        }

        private async Task<bool> SendAsync<TMsg, T>(TMsg msgType, T payload, string topic, string content)
        {
            var connected = await CheckConnectedAsync();

            if (connected)
            {
                _ = Task.Run(async () =>
                {
                    var res = await Client.PublishAsync(new MqttApplicationMessageBuilder()
                        .WithTopic(topic)
                        .WithPayload(Encoding.UTF8.GetBytes(content))
                        .WithQualityOfServiceLevel((MQTTnet.Protocol.MqttQualityOfServiceLevel)0) // qos)
                        .WithRetainFlag(false) // retainFlag)
                        .Build());

                    Console.WriteLine($"Sent from SendAsync. {topic} {res.ReasonCode} {content}");
                    var tt = res;
                });
            }

            return connected;
        }
    }
}
