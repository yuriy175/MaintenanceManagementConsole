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

namespace MessagesSender.BL.Remoting
{
    /// <summary>
    /// RabbitMQ work queue sender class
    /// </summary>
    public class RabbitMQWorkqueueSender : IWorkqueueSender
    {
        protected const string RabbitMQConnectionStringName = "RabbitMQConnectionString";
        protected const string ConnectionStringValuesSeparator = ";";
        protected const string ConnectionStringValueSeparator = "=";
        protected const string ConnectionStringServerName = "Server";
        protected const string ConnectionStringUserName = "User";
        protected const string ConnectionStringPasswordName = "Password";

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
        /// <typeparam name="T">entity type</typeparam>
        /// <param name="payload">entity</param>
        /// <returns>result</returns>
        public Task<bool> SendAsync<T>(T payload)
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
