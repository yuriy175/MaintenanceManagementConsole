using System;
using System.Threading.Tasks;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using Atlas.Remoting.Core.Interfaces;
using MessagesSender.Core.Interfaces;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;
using Serilog;
using Atlas.Common.Impls.Helpers;
using System.Net;
using System.Linq;
using System.Net.Sockets;
using Atlas.Acquisitions.Common.Core;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System.Text;
using MessagesSender.Core.Model;
using Atlas.Acquisitions.Common.Core.Model;
using System.Collections.Generic;
using System.IO;

namespace MessagesSender.BL
{
    public class CommandService : ICommandService
    {
        private const string ActivateCommandName = "activate";
		private const string RunTemViewerCommandName = "runTV";

		private readonly ILogger _logger;
        private readonly IHddWatchService _hddWatchService;
		private readonly IEventPublisher _eventPublisher;
		private readonly ISendingService _sendingService;

		private readonly Dictionary<string, Func<Task<bool>>> _commandMap = new Dictionary<string, Func<Task<bool>>>
        {
        };

		/// <summary>
		/// public constructor
		/// </summary>
		/// <param name="logger">logger</param>
		/// <param name="hddWatchService">hdd watch service</param>
		/// <param name="eventPublisher">event publisher service</param>
		/// <param name="sendingService">sending service</param>
		public CommandService(
            ILogger logger,
            IHddWatchService hddWatchService,
			IEventPublisher eventPublisher,
			ISendingService sendingService)
        {
            _logger = logger;
            _hddWatchService = hddWatchService;
			_sendingService = sendingService;
			_eventPublisher = eventPublisher;

			_commandMap = new Dictionary<string, Func<Task<bool>>>
            {
                { ActivateCommandName, async () => await OnActivateCommandAsync()},
				{ RunTemViewerCommandName, async () => await OnRunTVCommandAsync()},				
			};

			_eventPublisher.RegisterMqttCommandArrivedEvent(command => OnCommandArrivedAsync(command));

			_logger.Information("CommandService started");
        }

		private Task<bool> OnCommandArrivedAsync(string command)
		{
			try
			{
				return _commandMap[command]();
			}
			catch (Exception ex)
			{
				_logger.Error(ex, $"command error {command}");
			}

			return Task.FromResult(false);
		}

        /// <summary>
        /// activate command handler
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> OnActivateCommandAsync()
        {
            var hddDrives = await _hddWatchService.GetDriveInfosAsync();
            if (hddDrives != null)
            {
				_sendingService.SendInfoToMqttAsync(MQMessages.HddDrivesInfo, hddDrives);
            }

            return false;
        }

		private async Task<bool> OnRunTVCommandAsync()
		{
			return true;
		}
	}
}
