using System;
using System.IO;
using System.Linq;
using System.Security.Authentication;
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
        private const int MqttPort = 1883;

        // private const int MqttsPort = 11884;
        private readonly IConfigurationService _configurationService;
        private readonly ILogger _logger;
        private readonly string _clientId = Guid.NewGuid().ToString();

        // private readonly bool _useMqttSecure = false;
        private (string HostName, int Port, string UserName, string Password, bool Secured)? _connectionProps;
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

            /*_configurationService.AddConfigFile(
                Path.Combine(
                    Path.GetDirectoryName(
                        typeof(IWorkqueueSender).Assembly.Location), "consoleMQsettings.json"));
                        */
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
        public virtual async Task<bool> CreateAsync()
        {
            _topic = await GetTopicAsync();

            await CreateConnection(new ConnectionFactory());
            Created = _mqttClient != null;

            return Created;
        }

        /// <summary>
        /// Dispose resources
        /// </summary>
        public virtual void Dispose()
        {
            using (_mqttClient)
            {
            }
        }

        /// <summary>
        /// Get topic
        /// </summary>
        /// <param name="equipInfo">equipment info</param>
        /// <returns>topic</returns>
        protected abstract Task<string> GetTopicAsync();

        /// <summary>
        /// Create connection
        /// </summary>
        /// <param name="connectionFactory">connection factory</param>
        /// <returns>mqtt client</returns>
        protected virtual async Task<IManagedMqttClient> CreateConnection(ConnectionFactory connectionFactory)
        {
            CreateConnectionProps();
            connectionFactory.HostName = _connectionProps?.HostName ?? "localhost";
            connectionFactory.UserName = _connectionProps?.UserName ?? "guest";
            connectionFactory.Password = _connectionProps?.Password ?? "guest";
            connectionFactory.Port = _connectionProps?.Port ?? MqttPort;

            var secured = _connectionProps?.Secured ?? false;

            try
            {
                var messageBuilder = new MqttClientOptionsBuilder()
                    .WithClientId(_clientId)
                    .WithCredentials(connectionFactory.UserName, connectionFactory.Password)
                    .WithTcpServer(connectionFactory.HostName, connectionFactory.Port) // _useMqttSecure ? MqttsPort : MqttPort)
                    .WithCleanSession();

                var options = secured // _useMqttSecure
                  ? messageBuilder
                    .WithTls(new MqttClientOptionsBuilderTlsParameters()
                    {
                        UseTls = true,
                        SslProtocol = SslProtocols.Tls12,
                    })
                    .Build()
                  : messageBuilder
                    .Build();

                var managedOptions = new ManagedMqttClientOptionsBuilder()
                  .WithAutoReconnectDelay(TimeSpan.FromSeconds(5))
                  .WithClientOptions(options)
                  .Build();

                _mqttClient = new MqttFactory().CreateManagedMqttClient();

                _mqttClient.UseConnectedHandler(e =>
                {
                    Console.WriteLine("Connected successfully with MQTT Brokers. " + connectionFactory.HostName);
                });

                _mqttClient.UseDisconnectedHandler(e =>
                {
                    Console.WriteLine("Disconnected from MQTT Brokers." + connectionFactory.HostName);
                });    

                await _mqttClient.StartAsync(managedOptions);                
            }
            catch (Exception ex)
            {
                using (_mqttClient)
                {
                }

                _mqttClient = null;

                _logger.Error(ex, $"MQ connection error: {connectionFactory.HostName}, {connectionFactory.UserName}, {connectionFactory.Password}."); 
                return _mqttClient;
            }

            return _mqttClient;
        }

        private void CreateConnectionProps()
        {
            var connectionString = _configurationService.Get<string>(Constants.RabbitMQConnectionStringName, null);
            try
            {
                _connectionProps = ConnectionPropsCreator.CreateMqttProps(connectionString);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "Rabbit MQ work queue wrong connection string");
            }
        }
    }
}
