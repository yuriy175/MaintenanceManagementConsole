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

        private readonly ILogger _logger;
        private readonly IHddWatchService _hddWatchService;

        private readonly Dictionary<string, Func<Task<(string MsgType, object Info)?>>> _commandMap = 
            new Dictionary<string, Func<Task<(string MsgType, object Info)?>>>
        {
        };

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="hddWatchService">hdd watch service</param>
        public CommandService(
            ILogger logger,
            IHddWatchService hddWatchService)
        {
            _logger = logger;
            _hddWatchService = hddWatchService;

            _commandMap = new Dictionary<string, Func<Task<(string MsgType, object Info)?>>>
            {
                { ActivateCommandName, async () => await OnActivateCommandAsync()},
            };

            _logger.Information("CommandService started");
        }

        /// <summary>
        /// command handler
        /// </summary>
        /// <param name="command">command</param>
        /// <returns>result</returns>
        public async Task<(string MsgType, object Info)?> OnCommandArrivedAsync(string command)
        {
            try
            {
                return await _commandMap[command]();
            }
            catch (Exception ex)
            {
                _logger.Error(ex, $"command error {command}");
            }

            return null;
        }

        /// <summary>
        /// activate command handler
        /// </summary>
        /// <returns>result</returns>
        public async Task<(string MsgType, object Info)?> OnActivateCommandAsync()
        {
            var hddDrives = await _hddWatchService.GetDriveInfosAsync();
            if (hddDrives != null)
            {
                return (MQMessages.HddDrivesInfo.ToString(), hddDrives);
            }

            return null;
        }
    }
}
