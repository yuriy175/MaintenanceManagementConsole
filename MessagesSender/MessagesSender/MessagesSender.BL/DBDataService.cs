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
using MessagesSenderModel = MessagesSender.Core.Model;
using Atlas.Acquisitions.Common.Core.Model;
using System.Collections.Generic;
using System.IO;
using System.Diagnostics;
using Atlas.Common.Core.Interfaces;
using MessagesSender.BL.BusWrappers.Helpers;
using System.Reflection;
using MessagesSender.Core.Model;

namespace MessagesSender.BL
{
    /// <summary>
    /// db raw data service
    /// </summary>
    public class DBDataService : IDBDataService
    {
        private readonly IConfigurationService _configurationService;
        private readonly ILogger _logger; 
        private readonly IEventPublisher _eventPublisher;
        private readonly IInfoEntityService _dbInfoEntityService;
        private readonly ISendingService _sendingService;
        private readonly ITopicService _topicService;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="dbInfoEntityService">info database connector</param>
        /// <param name="sendingService">sending service</param>
        /// <param name="topicService">topic service</param>
        public DBDataService(
            IConfigurationService configurationService,
            ILogger logger,
            IEventPublisher eventPublisher,
            IInfoEntityService dbInfoEntityService,
            ISendingService sendingService,
            ITopicService topicService)
        {
            _configurationService = configurationService;
            _logger = logger;
            _eventPublisher = eventPublisher;
            _dbInfoEntityService = dbInfoEntityService;
            _sendingService = sendingService;
            _topicService = topicService;

            _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            _eventPublisher.RegisterDeactivateCommandArrivedEvent(() => OnDeactivateArrivedAsync());

            _logger.Information("DBDataService started");
        }

        private void OnDeactivateArrivedAsync()
        {
            
        }

        private Task<bool> OnActivateArrivedAsync()
        {
            return SendAllDBDataAsync();
        }

        private async Task<bool> SendAllDBDataAsync()
        {
            var atlasData = _dbInfoEntityService.GetAtlasDataAsync();
            var hospitalData = _dbInfoEntityService.GetHospitalDataAsync();
            var softwareData = _dbInfoEntityService.GetSoftwareDataAsync();
            var systemData = _dbInfoEntityService.GetSystemDataAsync();

            await Task.WhenAll(new[] { atlasData as Task, hospitalData, softwareData, systemData });

            _ = _sendingService.SendInfoToMqttAsync(MQMessages.AllDBInfo,
                    new
                    {
                        Atlas = await atlasData,
                        HospitalD = await hospitalData,
                        Software = await softwareData,
                        System = await systemData
                    });


            return true;
        }
    }
}
