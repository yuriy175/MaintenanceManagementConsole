using System;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;
using Atlas.Acquisitions.Common.Core;
using Atlas.Acquisitions.Common.Core.Model;
using Atlas.Common.Impls.Helpers;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using Atlas.Remoting.Core.Interfaces;
using MessagesSender.Core.Interfaces;
using MessagesSender.Core.Model;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using Serilog;

namespace MessagesSender.BL
{
    /// <summary>
    /// main chat message service interface implementation
    /// </summary>
    public class MainChatMessageService : IMainChatMessageService
    {
        private const string TechUserName = "tech"; 

        private readonly ILogger _logger;
        private readonly IMqttSender _mqttSender;

        private IPAddress _ipAddress = null;
        private (string Name, string Number) _equipmentInfo = (null, null);

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="mqttSender">mqtt sender</param>
        public MainChatMessageService(
            ILogger logger,
            IMqttSender mqttSender)
        {
            _logger = logger;
            _mqttSender = mqttSender;

            /*new Action[]
                {
                    () => _ = SubscribeMQRecevers(),
                    async () =>
                    {
                        await _sendingService.CreateAsync();
                        await OnServiceStateChangedAsync(true);
                    },
                    () => _eventPublisher.RegisterReconnectCommandArrivedEvent(() => OnReconnectArrivedAsync())
                }.RunTasksAsync();
*/
            _logger.Information("Main chat message service started");
        }

        /// <summary>
        /// Sends a chat message
        /// </summary>
        /// <param name="message">chat message</param>
        /// <returns>result</returns>
        public async Task<bool> SendChatMessageAsync(string message)
        {
            return await _mqttSender.CreateAsync() && await _mqttSender.SendAsync(
                MQMessages.Chat,
                new
                {
                    Message = message,
                    User = TechUserName,
                    IsInternal = true,
                });
        }
    }
}
