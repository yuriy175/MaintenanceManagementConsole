using System;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Atlas.Common.Core.Interfaces;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using MessagesSender.BL.BusWrappers.Helpers;
using MessagesSender.Core.Interfaces;
using Newtonsoft.Json;
using RabbitMQ.Client;
using Serilog;

namespace MessagesSender.BL.Remoting
{
    /// <summary>
    /// RabbitMQ work queue sender class
    /// </summary>
    public class RabbitMQWorkqueueSender : IWorkqueueSender
    {
        /// <summary>
        /// RabbitMQ connection string name
        /// </summary>
        protected const string RabbitMQConnectionStringName = "ConsoleRabbitMQConnectionString";

        private readonly IConfigurationService _configurationService;
        private readonly ILogger _logger;
        private readonly string _queueName = "SystemInfoQueue"; // string.Empty;

        private (string HostName, string UserName, string Password)? _connectionProps;
        private IConnection _connection = null;
        private IModel _channel = null;

        /// <summary>
        /// Initializes a new instance of the <see cref="RabbitMQBase"/> class.
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        public RabbitMQWorkqueueSender(
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
        /// channel
        /// </summary>
        protected IModel Channel => _channel;

        /// <summary>
        /// exchange name
        /// </summary>
        protected string QueueName => _queueName;

        /// <summary>
        /// if channel created
        /// </summary>
        protected bool Created { get; set; }

        /// <summary>
        /// sends a message
        /// </summary>
        /// <typeparam name="TMsg">message type</typeparam>
        /// <typeparam name="T">entity type</typeparam>
        /// <param name="msgType">messge type</param>
        /// <param name="payload">entity</param>
        /// <returns>result</returns>
        public Task<bool> SendAsync<TMsg, T>(TMsg msgType, T payload)
        {
            if (!Created)
            {
                return Task.FromResult(false);
            }

            IBasicProperties basicProperties = _channel.CreateBasicProperties();
            basicProperties.Persistent = false;
            var content = JsonConvert.SerializeObject(
                payload,
                new JsonSerializerSettings { NullValueHandling = NullValueHandling.Ignore });
            var body = Encoding.UTF8.GetBytes(content);

            _channel.BasicPublish(
                exchange: string.Empty,
                routingKey: _queueName,
                basicProperties: basicProperties,
                body: body);

            return Task.FromResult(true);
        }

        /// <inheritdoc/>
        public void Dispose()
        {
            using (_channel)
            {
            }

            using (_connection)
            {
            }
        }

        /// <summary>
        /// creates queue
        /// </summary>
        /// <param name="exchangeName">queue name</param>
        /// <returns>result</returns>
        protected virtual bool CreateAsync()
        {
            _connection = CreateConnection(new ConnectionFactory());
            if (_connection == null)
            {
                _logger.Error("No connection");
                return false;
            }

            _channel = _connection.CreateModel();
            if (_channel == null)
            {
                _logger.Error("No channel");
                return false;
            }

            _channel.QueueDeclare(
                queue: _queueName,
                durable: false,
                exclusive: false,
                autoDelete: false,
                arguments: null);
            Created = true;

            return Created;
        }

        private IConnection CreateConnection(ConnectionFactory connectionFactory)
        {
            // Server=medprom.ml;User=user;Password=medtex
            CreateConnectionProps();
            connectionFactory.HostName = _connectionProps?.HostName ?? "localhost";
            connectionFactory.UserName = _connectionProps?.UserName ?? "guest";
            connectionFactory.Password = _connectionProps?.Password ?? "guest";

            try
            {
                return connectionFactory.CreateConnection();
            }
            catch (Exception ex)
            {
                _logger.Error(ex, $"MQ connection error: {connectionFactory.HostName}, {connectionFactory.UserName}, {connectionFactory.Password}."); 
                return null;
            }
        }

        private void CreateConnectionProps()
        {
            var connectionString = _configurationService.Get<string>(RabbitMQConnectionStringName, null);
            try
            {
                _connectionProps = ConnectionPropsCreator.Create(connectionString);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "Rabbit MQ mqtt wrong connection string");
            }
        }
    }
}
