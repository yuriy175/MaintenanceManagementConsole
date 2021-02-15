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
    /// RabbitMQ mqtt command receiver class
    /// </summary>
    public class RabbitMQTTReceiver : RabbitMQTTBase, IMqttReceiver
    {
        private readonly ILogger _logger;

        /// <summary>
        /// Initializes a new instance of the <see cref="RabbitMQBase"/> class.
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        public RabbitMQTTReceiver(
            IConfigurationService configurationService,
            ILogger logger) : base(configurationService, logger)
        {
            _logger = logger;
        }

        protected override string GetTopic((string Name, string Number) equipInfo)
            => $"{equipInfo.Name}/{equipInfo.Number}_commands";

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
                            string payload = Encoding.UTF8.GetString(e.ApplicationMessage.Payload);
                            Console.WriteLine($"Topic: {topic}. Message Received: {payload}");
                        }
                    }
                    catch (Exception ex)
                    {
                        Console.WriteLine(ex.Message, ex);
                    }
                });

                await Client.SubscribeAsync(new TopicFilterBuilder()
                    .WithTopic(Topic)
                    .WithQualityOfServiceLevel((MQTTnet.Protocol.MqttQualityOfServiceLevel)0) // qos)
                    .Build());
            }
            catch (Exception ex)
            {
                using (Client) ;

                _logger.Error(ex, $"MQ connection error: { connectionFactory.HostName}, {connectionFactory.UserName}, {connectionFactory.Password}."); ;
                return null;
            }

            return null;
        }

    }
}