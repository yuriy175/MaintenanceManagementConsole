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

namespace MessagesSender.BL
{
    public class RemoteControlService : IRemoteControlService
    {
        private readonly ISettingsEntityService _dbSettingsEntityService;
        private readonly IObservationsEntityService _dbObservationsEntityService;
        private readonly ILogger _logger;
        private readonly ISendingService _sendingService;
        private readonly IEventPublisher _eventPublisher;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="dbSettingsEntityService">settings database connector</param>
        /// <param name="dbObservationsEntityService">observations database connector</param>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        public RemoteControlService(
            ISettingsEntityService dbSettingsEntityService,
            IObservationsEntityService dbObservationsEntityService,
            ILogger logger,
            IEventPublisher eventPublisher,
            ISendingService sendingService)
        {
            _dbSettingsEntityService = dbSettingsEntityService;
            _dbObservationsEntityService = dbObservationsEntityService;
            _logger = logger;
            _sendingService = sendingService;
            _eventPublisher = eventPublisher;

            _logger.Information("RemoteControlService started");
        }

        /// <summary>
        /// runs TeamViewer
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> RunTeamViewerAsync()
        {
            return true;
        }

        /// <summary>
        /// runs TaskManager
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> RunTaskManagerAsync()
        {
            return true;
        }

        /// <summary>
        /// sends Atlas logs to email
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> SendAtlasLogsAsync()
        {
            return true;
        }

        /// <summary>
        /// turns on XilibLogs
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> XilibLogsOnAsync()
        {
            return true;
        }
    }
}
