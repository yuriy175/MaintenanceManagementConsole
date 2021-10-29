using System;
using System.Collections.Generic;
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
    /// keep alive info service 
    /// </summary>
    public class KeepAliveService : IKeepAliveService
    {
        private const int KeepAlivePeriod = 4000;

        private readonly ILogger _logger;
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;

        private bool _isRunning = false;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        public KeepAliveService(
            ILogger logger,
            IEventPublisher eventPublisher,
            ISendingService sendingService)
        {
            _logger = logger;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;

            _eventPublisher.RegisterServerReadyCommandArrivedEvent(() => OnServerReadyArrivedAsync());

            _logger.Information("KeepAliveService started");
        }

        private async Task<bool> OnServerReadyArrivedAsync()
        {
            if (!_isRunning)
            {
                _isRunning = true;

				try
				{
					while (true)
					{
						await _sendingService.SendInfoToCommonMqttAsync(MQMessages.KeepAlive, new { });
						await Task.Delay(KeepAlivePeriod);
					}
				}
				catch (Exception ex)
				{
					_logger.Error(ex, "KeepAlive error");
				}

				_isRunning = false;
				return false;
			}

            return true;
        }
    }
}
