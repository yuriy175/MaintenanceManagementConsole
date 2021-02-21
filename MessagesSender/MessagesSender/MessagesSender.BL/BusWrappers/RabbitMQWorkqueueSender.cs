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
using MessagesSender.BL.BusWrappers.Helpers;

namespace MessagesSender.BL.Remoting
{
    /// <summary>
    /// RabbitMQ work queue sender class
    /// </summary>
    public class RabbitMQWorkqueueSender : IWorkqueueSender
    {
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

            _channel.QueueDeclare(queue: _queueName,
                                 durable: false,
                                 exclusive: false,
                                 autoDelete: false,
                                 arguments: null);
            Created = true;

            return Created;
        }

		/// <summary>
		/// sends a message
		/// </summary>
		/// <typeparam name="TMsg">message type</typeparam>
		/// <typeparam name="T">entity type</typeparam>
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
            var content = JsonConvert.SerializeObject(payload);
            var body = Encoding.UTF8.GetBytes(content);

            _channel.BasicPublish(exchange: "",
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

        private IConnection CreateConnection(ConnectionFactory connectionFactory)
        {
            //Server=medprom.ml;User=user;Password=medtex
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
                _logger.Error(ex, $"MQ connection error: { connectionFactory.HostName}, {connectionFactory.UserName}, {connectionFactory.Password}."); ;
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












