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
using System.Diagnostics;
using Atlas.Common.Core.Interfaces;
using System.Diagnostics.Eventing.Reader;

namespace MessagesSender.BL
{
    /// <summary>
    /// software watch service 
    /// </summary>
    public class SoftwareWatchService : ISoftwareWatchService
    {
        private const string AtlasInstanceVersionName = "AtlasInstanceVersion";
        private const string ServicesFolderName = "ServicesFolder";
        private const string InstallPathName = "InstallPath";
        private const string InstallExeName = "Atlas.System.exe";
        private const string XilibModuleName = @"XiLibs\XiLibNet.dll";

        private readonly IConfigurationService _configurationService;
        private readonly ILogger _logger; 
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;

        private (string Version, string XilibVersion) _versions = (string.Empty, string.Empty);
        private EventLogWatcher _watcher = null;
        
        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        public SoftwareWatchService(
            IConfigurationService configurationService,
            ILogger logger,
            IEventPublisher eventPublisher,
            ISendingService sendingService)
        {
            _configurationService = configurationService;
            _logger = logger;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;

            _versions = GetVersions();

            _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            _eventPublisher.RegisterDeactivateCommandArrivedEvent(() => OnDeactivateArrivedAsync());

            Task.Run(() =>
            {
                SubscribeSystemEvents();
            });        

            _logger.Information("SoftwareWatchService started");
        }

        private void OnDeactivateArrivedAsync()
        {            
        }

        private Task<bool> OnActivateArrivedAsync()
        {
            return _sendingService.SendInfoToMqttAsync(MQMessages.SoftwareInfo,
                new
                {
                    SettingsDB = true,
                    ObservationsDB = true,
                    _versions.Version,
                    _versions.XilibVersion,
                });
        }

        private (string Version, string XilibVersion) GetVersions()
        {
            try
            {
                var version = _configurationService.Get<string>(AtlasInstanceVersionName, @"");
                if (string.IsNullOrEmpty(version))
                {
                    var installPath = _configurationService.Get<string>(InstallPathName, @"C:\Program Files\Atlas\bin");
                    version = FileVersionInfo.GetVersionInfo(Path.Combine(installPath, InstallExeName)).FileVersion;
                }

                var servicesFolder = _configurationService.Get<string>(ServicesFolderName, @"C:\Program Files\Atlas\Services");
                var xilibVersion = FileVersionInfo.GetVersionInfo(Path.Combine(servicesFolder, XilibModuleName)).FileVersion;
                
                return (version, xilibVersion);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "GetVersions error ");

                return (string.Empty, string.Empty);
            }
        }

        private void SubscribeSystemEvents()
        {
            try
            {
                EventLogQuery subscriptionQuery = new EventLogQuery(
                    //"Security",
                    "Application",
                    PathType.LogName
                    //,"*[System/EventID=4624]"
                    );

                _watcher = new EventLogWatcher(subscriptionQuery);

                // Make the watcher listen to the EventRecordWritten
                // events.  When this event happens, the callback method
                // (EventLogEventRead) is called.
                _watcher.EventRecordWritten +=
                    new EventHandler<EventRecordWrittenEventArgs>(
                        EventLogEventRead);

                // Activate the subscription
                _watcher.Enabled = true;

                //for (int i = 0; i < 5; i++)
                //{
                //    // Wait for events to occur. 
                //    System.Threading.Thread.Sleep(10000);
                //}
            }
            catch (EventLogReadingException e)
            {
                //Log("Error reading the log: {0}", e.Message);
            }
            finally
            {
                // Stop listening to events
                //watcher.Enabled = false;

                //if (watcher != null)
                //{
                //    watcher.Dispose();
                //}
            }
            //Console.ReadKey();
        }

        // Callback method that gets executed when an event is
        // reported to the subscription.
        private void EventLogEventRead(object obj,
            EventRecordWrittenEventArgs arg)
        {
            var eventRecord = arg.EventRecord;
            // Make sure there was no error reading the event.
            if (eventRecord != null && eventRecord.Level == 2)
            {
                _sendingService.SendInfoToMqttAsync(MQMessages.SoftwareInfo,
                new
                {
                    ErrorDescriptions = new[] {
                        new {
                            Code = eventRecord.LevelDisplayName,
                            Description = eventRecord.ProviderName,
                        }
                    }
                });
            }
            else
            {
                _logger.Error("The event instance was null.");
            }
        }
    }
}
