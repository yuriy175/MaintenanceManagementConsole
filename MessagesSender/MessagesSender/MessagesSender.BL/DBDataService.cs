using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Reflection;
using System.Text;
using System.Threading.Tasks;
using Atlas.Acquisitions.Common.Core;
using Atlas.Acquisitions.Common.Core.Model;
using Atlas.Common.Core.Interfaces;
using Atlas.Common.Impls.Helpers;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using Atlas.Remoting.Core.Interfaces;
using MessagesSender.BL.BusWrappers.Helpers;
using MessagesSender.Core.Interfaces;
using MessagesSender.Core.Model;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using Serilog;
using MessagesSenderModel = MessagesSender.Core.Model;

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

            // _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            // _eventPublisher.RegisterDeactivateCommandArrivedEvent(() => OnDeactivateArrivedAsync());
            _logger.Information("DBDataService started");
        }

        /// <summary>
        /// Updates db info
        /// </summary>
        /// <returns>result</returns>
        public Task<bool> UpdateDBInfoAsync()
        {
            return SendAllDBDataAsync();
        }

        private void OnDeactivateArrivedAsync()
        {            
        }

        private async Task<bool> SendAllDBDataAsync()
        {
            var newsData = await _dbInfoEntityService.GetNewsDataAsync();
            if (!newsData.Any())
            {
                return true;
            }

            var newTables = newsData
                .GroupBy(n => n.Tbl)
                .ToDictionary(n => n.Key, n => n.Select(e => e.RowId).ToArray());

            var atlasTask = _dbInfoEntityService.GetAtlasDataAsync(newTables);
            var hospitalTask = _dbInfoEntityService.GetHospitalDataAsync(newTables);
            var softwareTask = _dbInfoEntityService.GetSoftwareDataAsync(newTables);
            var systemTask = _dbInfoEntityService.GetSystemDataAsync(newTables);

            await Task.WhenAll(new[] { atlasTask as Task, hospitalTask, softwareTask, systemTask });

            var atlasData = await atlasTask;
            var hospitalData = await hospitalTask;
            var softwareData = await softwareTask;
            var systemData = await systemTask;

            _ = _sendingService.SendInfoToMqttAsync(
                MQMessages.AllDBInfo,
                    new
                    {
                        Hospital = new { HospitalInfo = hospitalData?.ToArray() },
                        Software = new
                        {
                            softwareData.Atlas,
                            softwareData.Dependencies,
                            softwareData.Errors,
                            softwareData.OsInfos,
                            softwareData.SqlDatabases,
                            softwareData.SqlServices
                        },
                        System = new
                        {
                            systemData.HardDrives,
                            systemData.Lans,
                            systemData.LogicalDisks,
                            systemData.Modems,
                            systemData.Monitors,
                            systemData.Motherboards,
                            systemData.Printers,
                            systemData.Screens,
                            systemData.VideoAdapters
                        },
                        Atlas = new
                        {
                            atlasData.AppParams,
                            atlasData.AspNetUsers,
                            atlasData.Detectors,
                            atlasData.DetectorProcessings,
                            atlasData.DicomServices,
                            atlasData.DicomPrinters,
                            atlasData.HardwareParams,
                            atlasData.RasterParams
                        }
                    });

            await _dbInfoEntityService.SetNewsDataSentAsync();

            return true;
        }
    }
}
