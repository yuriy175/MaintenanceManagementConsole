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
using System.Collections.Generic;
using MessagesSender.Core.Model;

namespace MessagesSender.BL.Remoting
{
	/// <summary>
	/// RabbitMQ mqtt sender class
	/// </summary>
	public class RabbitMQTTSender : RabbitMQTTBase, IMqttSender
	{
		private const string CommandSubTopic = "/command";

		private const int ConnectWaitingAttempts = 5;

		private readonly ILogger _logger;
		private readonly Dictionary<string, string> _topicMap = new Dictionary<string, string>
		{
			{ MQCommands.StudyInWork.ToString(), "/study"},
			{ MQCommands.GeneratorStateArrived.ToString(), "/generator/state"},
			{ MQCommands.StandStateArrived.ToString(), "/stand/state"},
			{ MQCommands.CollimatorStateArrived.ToString(), "/collimator/state"},
			// { MQCommands.CollimatorStateArrived.ToString(), "/detector/state"},
			{ MQMessages.HddDrivesInfo.ToString(), "/hdd"},
		};

		/// <summary>
		/// Initializes a new instance of the <see cref="RabbitMQBase"/> class.
		/// </summary>
		/// <param name="configurationService">configuration service</param>
		/// <param name="logger">logger</param>
		public RabbitMQTTSender(
            IConfigurationService configurationService,
            ILogger logger) : base(configurationService, logger)
        {
			_logger = logger;
		}

		/// <summary>
		/// sends a message
		/// </summary>
		/// <typeparam name="TMsg">message type</typeparam>
		/// <typeparam name="T">entity type</typeparam>
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

			await CheckConnectedAsync();

			_ = Task.Run(async () =>
            {
                var content = JsonConvert.SerializeObject(payload);
                var res = await Client.PublishAsync(new MqttApplicationMessageBuilder()
                    .WithTopic($"{Topic}{subtopic}")
                    .WithPayload(Encoding.UTF8.GetBytes(content))
                    .WithQualityOfServiceLevel((MQTTnet.Protocol.MqttQualityOfServiceLevel)0) // qos)
                    .WithRetainFlag(false) // retainFlag)
                    .Build());

                Console.WriteLine($"Sent from SendAsync. {Topic} {res.ReasonCode} {content}");
                var tt = res;
            });

            return true;
        }

		public Action<string> OnCommandArrived { get; set; } = command => { };

		protected override string GetTopic((string Name, string Number) equipInfo)
            => $"{equipInfo.Name}/{equipInfo.Number}";

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
							OnCommandArrived(command);
							Console.WriteLine($"Topic: {topic}. Message Received: {command}");
						}
					}
					catch (Exception ex)
					{
						Console.WriteLine(ex.Message, ex);
					}
				});

				await Client.SubscribeAsync(new TopicFilterBuilder()
					.WithTopic(Topic + CommandSubTopic)
					.WithQualityOfServiceLevel((MQTTnet.Protocol.MqttQualityOfServiceLevel)0) // qos)
					.Build());
			}
			catch (Exception ex)
			{
				using (Client) ;

				// _logger.Error(ex, $"MQ connection error: { connectionFactory.HostName}, {connectionFactory.UserName}, {connectionFactory.Password}."); ;
				return null;
			}

			return null;
		}

		private async Task CheckConnectedAsync()
		{
			int attempts = 0;
			while (!Client.IsConnected && attempts++ < ConnectWaitingAttempts)
			{
				Console.WriteLine($"Sent to not connected {Topic}");
				await Task.Delay(Client.Options.ConnectionCheckInterval);
			}
		}
	}
}
