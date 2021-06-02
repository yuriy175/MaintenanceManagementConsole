using System;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Atlas.Common.Core.Interfaces;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using MessagesSender.BL.BusWrappers.Helpers;
using MessagesSender.Core.Interfaces;
using MessagesSender.Core.Model;
using MQTTnet;
using MQTTnet.Client;
using MQTTnet.Client.Options;
using MQTTnet.Extensions.ManagedClient;
using Newtonsoft.Json;
using RabbitMQ.Client;
using Serilog;

namespace MessagesSender.BL.Remoting
{
    /// <summary>
    /// RabbitMQ base mqtt class
    /// </summary>
    public abstract class RabbitMQTTBase
    {
        private readonly IConfigurationService _configurationService;
        private readonly ILogger _logger;
        private readonly string _clientId = Guid.NewGuid().ToString();

        private (string HostName, string UserName, string Password)? _connectionProps;
        private IManagedMqttClient _mqttClient = null;
        private string _topic = string.Empty;

        /// <summary>
        /// Initializes a new instance of the <see cref="RabbitMQBase"/> class.
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        public RabbitMQTTBase(
            IConfigurationService configurationService,
            ILogger logger)
        {
            _configurationService = configurationService;
            _logger = logger;

            _configurationService.AddConfigFile(
                Path.Combine(
                    Path.GetDirectoryName(
                        typeof(IWorkqueueSender).Assembly.Location), "consoleMQsettings.json"));
        }

        /// <summary>
        /// if channel created
        /// </summary>
        protected bool Created { get; set; }

        /// <summary>
        /// if channel created
        /// </summary>
        protected IManagedMqttClient Client => _mqttClient;

        /// <summary>
        /// if channel created
        /// </summary>
        protected string Topic => _topic;

        /// <summary>
        /// creates sender
        /// </summary>
        /// <param name="equipInfo">equipment info</param>
        /// <returns>result</returns>        
        public virtual async Task<bool> CreateAsync((string Name, string Number) equipInfo)
        {
            if (string.IsNullOrEmpty(equipInfo.Name) || string.IsNullOrEmpty(equipInfo.Number))
            {
                return false;
            }

            _topic = await GetTopicAsync(equipInfo);

            await CreateConnection(new ConnectionFactory());
            Created = _mqttClient != null;

            return Created;
        }

        /// <inheritdoc/>
        public virtual void Dispose()
        {
            using (_mqttClient)
            {
            }
        }

        protected abstract Task<string> GetTopicAsync((string Name, string Number) equipInfo);

        protected virtual async Task<IManagedMqttClient> CreateConnection(ConnectionFactory connectionFactory)
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
                    Console.WriteLine("Connected successfully with MQTT Brokers. " + connectionFactory.HostName);
                });

                _mqttClient.UseDisconnectedHandler(e =>
                {
                    Console.WriteLine("Disconnected from MQTT Brokers." + connectionFactory.HostName);
                });
                                
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
                using (_mqttClient)
                {
                }

                _mqttClient = null;

                _logger.Error(ex, $"MQ connection error: { connectionFactory.HostName}, {connectionFactory.UserName}, {connectionFactory.Password}."); ;
                return _mqttClient;
            }

            return _mqttClient;
        }

        private void CreateConnectionProps()
        {
            var connectionString = _configurationService.Get<string>(Constants.RabbitMQConnectionStringName, null);
            try
            {
                _connectionProps = ConnectionPropsCreator.Create(connectionString);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "Rabbit MQ work queue wrong connection string");
            }
        }
    }
}














