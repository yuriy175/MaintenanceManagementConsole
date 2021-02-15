using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using RabbitMQ.Client;
using Serilog;
using System;
using System.Threading.Tasks;
using Newtonsoft.Json;
using System.Text;
using MessagesSender.Core.Interfaces;
using Atlas.Common.Core.Interfaces;
using System.Linq;
using System.IO;
using MQTTnet.Extensions.ManagedClient;
using MQTTnet;
using MQTTnet.Client;
using MQTTnet.Client.Options;
using MessagesSender.BL.BusWrappers.Helpers;

namespace MessagesSender.BL.Remoting
{
    /// <summary>
    /// RabbitMQ mqtt sender class
    /// </summary>
    public class RabbitMQTTSender : RabbitMQTTBase, IMqttSender
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="RabbitMQBase"/> class.
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        public RabbitMQTTSender(
            IConfigurationService configurationService,
            ILogger logger) : base(configurationService, logger)
        {
        }

        /// <summary>
        /// sends a message
        /// </summary>
        /// <typeparam name="T">entity type</typeparam>
        /// <param name="payload">entity</param>
        /// <returns>result</returns>
        public Task<bool> SendAsync<T>(T payload)
        {
            if (!Created)
            {
                return Task.FromResult(false);
            }

            _ = Task.Run(async () =>
            {
                var content = JsonConvert.SerializeObject(payload);
                var res = await Client.PublishAsync(new MqttApplicationMessageBuilder()
                    .WithTopic(Topic)
                    .WithPayload(Encoding.UTF8.GetBytes(content)) // "messa")) // payload)
                    .WithQualityOfServiceLevel((MQTTnet.Protocol.MqttQualityOfServiceLevel)0) // qos)
                    .WithRetainFlag(false) // retainFlag)
                    .Build());

                Console.WriteLine($"Sent from SendAsync. {Topic} {res.ReasonCode} {content}");
                var tt = res;
            });

            return Task.FromResult(true);
        }

        protected override string GetTopic((string Name, string Number) equipInfo)
            => $"{equipInfo.Name}/{equipInfo.Number}";         
    }
}