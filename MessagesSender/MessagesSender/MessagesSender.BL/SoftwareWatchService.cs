using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Diagnostics.Eventing.Reader;
using System.IO;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;
using Atlas.Acquisitions.Common.Core;
using Atlas.Acquisitions.Common.Core.Model;
using Atlas.Common.Core.Interfaces;
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
    /// software watch service 
    /// </summary>
    public class SoftwareWatchService : ISoftwareWatchService, IDisposable
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
        private readonly IMQCommunicationService _mqService;

        private (string Version, string XilibVersion) _versions = (string.Empty, string.Empty);
        private EventLogWatcher _watcher = null;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        /// <param name="mqService">MQ service</param>
        public SoftwareWatchService(
            IConfigurationService configurationService,
            ILogger logger,
            IEventPublisher eventPublisher,
            ISendingService sendingService,
            IMQCommunicationService mqService)
        {
            _configurationService = configurationService;
            _logger = logger;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;
            _mqService = mqService;

            _versions = GetVersions();

            _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            _eventPublisher.RegisterDeactivateCommandArrivedEvent(() => OnDeactivateArrivedAsync());

            new Action[]
                {
                    () => _ = SubscribeMQRecevers(),
                    () => SubscribeSystemEvents(),
                }.RunTasksAsync();

            _logger.Information("SoftwareWatchService started");
        }

        public void Dispose()
        {
            // Stop listening to events
            if (_watcher != null)
            {
                _watcher.Enabled = false;
            }

            _watcher?.Dispose();
        }

        private void OnDeactivateArrivedAsync()
        {            
        }

        private Task<bool> OnActivateArrivedAsync()
        {
            return _sendingService.SendInfoToMqttAsync(
                MQMessages.SoftwareInfo,
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
                var version = _configurationService.Get<string>(AtlasInstanceVersionName, string.Empty);
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
                    "Application",
                    PathType.LogName
                    );

                _watcher = new EventLogWatcher(subscriptionQuery);

                // Make the watcher listen to the EventRecordWritten
                // events.  When this event happens, the callback method
                // (EventLogEventRead) is called.
                _watcher.EventRecordWritten +=
                    new EventHandler<EventRecordWrittenEventArgs>(EventLogEventRead);

                // Activate the subscription
                _watcher.Enabled = true;
            }
            catch (EventLogReadingException ex)
            {
                _logger.Error(ex, "Error reading the log: ");
            }
        }

        // Callback method that gets executed when an event is
        // reported to the subscription.
        private void EventLogEventRead(object obj, EventRecordWrittenEventArgs arg)
        {
            var eventRecord = arg.EventRecord;
            try
            {
                // Make sure there was no error reading the event.
                if (eventRecord != null && eventRecord.Level == 2)
                {
                    SendCommonErrorAsync(eventRecord.LevelDisplayName, eventRecord.ProviderName);
                }
                else
                {
                    _logger.Error("The event instance was null.");
                }
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "EventLogEventRead error");
            }
        }

        private Task SubscribeMQRecevers()
        {
            return Task.Run(() =>
            {
                _mqService.Subscribe<MQCommands, (string, IEnumerable<string>)>(
                        (MQCommands.UserLoggedIn, userProps => OnUserLogIn(userProps)));

                _mqService.Subscribe<MQCommands, int>(
                        (MQCommands.Message, code => SendAtlasErrorAsync(code, string.Empty)));

                // _mqService.Subscribe<MQCommands, string>(
                //        (MQCommands.UserLoggedOut, userName => OnUserLogOut(userName)));
            });
        }

        private async void OnUserLogIn((string UserName, IEnumerable<string> UserRoles) userProps)
        {
            _ = _sendingService.SendInfoToMqttAsync(
                    MQMessages.SoftwareInfo,
                    new
                    {
                        Atlas_User = new
                        {
                            User = userProps.UserName,
                            Role = userProps.UserRoles?.FirstOrDefault(),
                        },
                    });
        }

        private void SendCommonErrorAsync(string code, string description)
        {
            _ = _sendingService.SendInfoToMqttAsync(
                    MQMessages.SoftwareMsgInfo,
                    new
                    {
                        ErrorDescriptions = new[]
                        {
                            new
                            {
                                Code = code,
                                Description = description,
                            }
                        },
                    });
        }

        private void SendAtlasErrorAsync(int code, string description)
        {
            _ = _sendingService.SendInfoToMqttAsync(
                    MQMessages.SoftwareMsgInfo,
                    new
                    {
                        AtlasErrorDescriptions = new[]
                        {
                            new
                            {
                                Code = code,
                                Description = description,
                            }
                        },
                    });
        }
    }
}
