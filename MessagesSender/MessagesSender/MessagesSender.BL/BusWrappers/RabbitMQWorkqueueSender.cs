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

namespace MessagesSender.BL.Remoting
{
    /// <summary>
    /// RabbitMQ mqtt sender class
    /// </summary>
    public class RabbitMQTTSender : IWorkqueueSender
    {
        protected const string RabbitMQConnectionStringName = "ConsoleRabbitMQConnectionString";
        protected const string ConnectionStringValuesSeparator = ";";
        protected const string ConnectionStringValueSeparator = "=";
        protected const string ConnectionStringServerName = "Server";
        protected const string ConnectionStringUserName = "User";
        protected const string ConnectionStringPasswordName = "Password";
        protected const string Topic = "topic/test";

        private readonly IConfigurationService _configurationService;
        private readonly ILogger _logger;
        private readonly string _clientId = Guid.NewGuid().ToString();

        private (string HostName, string UserName, string Password)? _connectionProps;
        private IManagedMqttClient _mqttClient = null;

        /// <summary>
        /// Initializes a new instance of the <see cref="RabbitMQBase"/> class.
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        public RabbitMQTTSender(
            IConfigurationService configurationService,
            ILogger logger)
        {
            _configurationService = configurationService;
            _logger = logger;

            _configurationService.AddConfigFile(
                Path.Combine(
                    Path.GetDirectoryName(
                        typeof(IWorkqueueSender).Assembly.Location), "consoleMQsettings.json"));

            CreateAsync();
        }

        /// <summary>
        /// if channel created
        /// </summary>
        protected bool Created { get; set; }

        /// <summary>
        /// creates queue
        /// </summary>
        /// <param name="exchangeName">queue name</param>
        /// <returns>result</returns>
        protected virtual async Task<bool> CreateAsync()
        {
            await CreateConnection(new ConnectionFactory());
            Created = true;

            return Created;
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
                var res = await _mqttClient.PublishAsync(new MqttApplicationMessageBuilder()
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

        /// <inheritdoc/>
        public void Dispose()
        {
            using (_mqttClient)
            {
            }
        }
        
        private async Task<IConnection> CreateConnection(ConnectionFactory connectionFactory)
        {
            CreateConnectionProps();
            connectionFactory.HostName = _connectionProps?.HostName ?? "localhost";
            connectionFactory.UserName = _connectionProps?.UserName ?? "guest";
            connectionFactory.Password = _connectionProps?.Password ?? "guest";

            try
            {
                var messageBuilder = new MqttClientOptionsBuilder()
                    .WithClientId(_clientId)
                    .WithCredentials(connectionFactory.UserName, connectionFactory.Password)
                    .WithTcpServer(connectionFactory.HostName, 1883)
                    .WithCleanSession();

                var options = false // mqttSecure
                  ? messageBuilder
                    .WithTls()
                    .Build()
                  : messageBuilder
                    .Build();

                var managedOptions = new ManagedMqttClientOptionsBuilder()
                  .WithAutoReconnectDelay(TimeSpan.FromSeconds(5))
                  .WithClientOptions(options)
                  .Build();

                _mqttClient = new MqttFactory().CreateManagedMqttClient();

                await _mqttClient.StartAsync(managedOptions);

                _mqttClient.UseConnectedHandler(e =>
                {
                    Console.WriteLine("Connected successfully with MQTT Brokers.");
                });

                _mqttClient.UseDisconnectedHandler(e =>
                {
                    Console.WriteLine("Disconnected from MQTT Brokers.");
                });

                /*
                _mqttClient.UseApplicationMessageReceivedHandler(e =>
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

                await _mqttClient.SubscribeAsync(new TopicFilterBuilder()
    .WithTopic(Topic)
    .WithQualityOfServiceLevel((MQTTnet.Protocol.MqttQualityOfServiceLevel)0) // qos)
    .Build());
                
                
                /*var options = new ManagedMqttClientOptionsBuilder()
                       .WithAutoReconnectDelay(TimeSpan.FromSeconds(5))
                       .WithClientOptions(new MqttClientOptionsBuilder()
                           .WithClientId("epo")
                           //.WithTcpServer("mskorp.tk")
                           //.WithCredentials("epo", "medtex")
                           //.WithTls()
                           .WithTcpServer(connectionFactory.HostName)
                           .WithCredentials(connectionFactory.UserName, connectionFactory.Password)
                           .Build())
                       .Build();

                _mqttClient = new MqttFactory().CreateManagedMqttClient();
                    // await mqttClient.SubscribeAsync(new TopicFilterBuilder().WithTopic("epotopic").Build());
                    //await _mqttClient.StartAsync(options);
                    var result = await _mqttClient.InternalClient.ConnectAsync(options.ClientOptions).ConfigureAwait(false);
                    await _mqttClient.StartAsync(options);
                    var msg = new MqttApplicationMessage
                    {
                        Topic = "topic/test", // "test", // "topic/test", // "epotopic",
                        Payload = Encoding.UTF8.GetBytes("Ohrenet"),
                    };
                    await _mqttClient.SubscribeAsync("topic/test");
                    var res = await _mqttClient.PublishAsync(msg);
                /*
                var result = await mqttClient.InternalClient.ConnectAsync(options.ClientOptions).ConfigureAwait(false);
                var result2 = await mqttClient.InternalClient.PublishAsync(new MqttApplicationMessage
                {
                    Topic = "epotopic",
                    Payload = Encoding.UTF8.GetBytes("Ohrenet"),
                }).ConfigureAwait(false);
                */
            }
            catch (Exception ex)
            {
                _logger.Error(ex, $"MQ connection error: { connectionFactory.HostName}, {connectionFactory.UserName}, {connectionFactory.Password}."); ;
                return null;
            }

            return null;
        }

        private void CreateConnectionProps()
        {
            var connectionString = _configurationService.Get<string>(RabbitMQConnectionStringName, null);
            if (!string.IsNullOrEmpty(connectionString))
            {
                try
                {
                    var props = connectionString.Split(new[] { ConnectionStringValuesSeparator }, StringSplitOptions.RemoveEmptyEntries)
                        .Select(s =>
                        {
                            var pair = s.Split(new[] { ConnectionStringValueSeparator }, StringSplitOptions.RemoveEmptyEntries).ToArray();
                            return new { Key = pair.First(), Value = pair.Last() };
                        })
                        .ToDictionary(s => s.Key, s => s.Value);

                    _connectionProps = (props[ConnectionStringServerName], props[ConnectionStringUserName], props[ConnectionStringPasswordName]);
                }
                catch (Exception ex)
                {
                    _logger.Error(ex, "Rabbit MQ wrong connection string");
                }
            }
        }
    }
}