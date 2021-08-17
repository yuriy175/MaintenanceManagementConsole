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
        private const string AtlasExeName = "Atlas.Acquisition"; // "Atlas.System";
        private const string XilibModuleName = @"XiLibs\XiLibNet.dll";

        private readonly IConfigurationService _configurationService;
        private readonly ILogger _logger; 
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;
        private readonly IMQCommunicationService _mqService;
        private readonly IMasterEntityService _dbMasterEntityService;

        private readonly string _installExeName = $"{AtlasExeName}.exe";

        private (string Version, string XilibVersion) _versions = (string.Empty, string.Empty);
        private EventLogWatcher _appWatcher = null;
        private EventLogWatcher _sysWatcher = null;
        private bool _isActivated = false;
        private (string UserName, string PersonInfo, IEnumerable<string> UserRoles)? _currentUserProps = null;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="dbMasterEntityService">master database connector</param>
        /// <param name="sendingService">sending service</param>
        /// <param name="mqService">MQ service</param>
        public SoftwareWatchService(
            IConfigurationService configurationService,
            ILogger logger,
            IEventPublisher eventPublisher,
            IMasterEntityService dbMasterEntityService,
            ISendingService sendingService,
            IMQCommunicationService mqService)
        {
            _configurationService = configurationService;
            _logger = logger;
            _eventPublisher = eventPublisher;
            _dbMasterEntityService = dbMasterEntityService;
            _sendingService = sendingService;
            _mqService = mqService;

            _versions = GetVersions();

            _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            _eventPublisher.RegisterDeactivateCommandArrivedEvent(() => OnDeactivateArrivedAsync());
            _eventPublisher.RegisterUpdateDBInfoCommandArrivedEvent(() => OnUpdateDBInfoAsync());

            new Action[]
                {
                    () => _ = SubscribeMQRecevers(),
                    () => SubscribeSystemEvents(),
                }.RunTasksAsync();

            _logger.Information("SoftwareWatchService started");
        }

        /// <summary>
        /// Dispose resources
        /// </summary>
        public void Dispose()
        {
            // Stop listening to events
            if (_appWatcher != null)
            {
                _appWatcher.Enabled = false;
            }

            if (_sysWatcher != null)
            {
                _sysWatcher.Enabled = false;
            }

            _appWatcher?.Dispose();
            _sysWatcher?.Dispose();
        }

        private void OnDeactivateArrivedAsync()
        {
            _isActivated = false;
        }

        private Task<bool> OnActivateArrivedAsync()
        {
            _isActivated = true;
            var atlasRunning = IsAtlasRunning();

            SendAtlasStatusAsync(atlasRunning);
            OnUpdateDBInfoAsync();

            return Task.FromResult(true);
            /*_sendingService.SendInfoToMqttAsync(
                MQMessages.SoftwareInfo,
                new
                {
                    SettingsDB = true,
                    ObservationsDB = true,
                    _versions.Version,
                    _versions.XilibVersion,
                });*/
        }

        private (string Version, string XilibVersion) GetVersions()
        {
            try
            {
                var version = _configurationService.Get<string>(AtlasInstanceVersionName, string.Empty);
                if (string.IsNullOrEmpty(version))
                {
                    var installPath = _configurationService.Get<string>(InstallPathName, @"C:\Program Files\Atlas\bin");
                    version = FileVersionInfo.GetVersionInfo(Path.Combine(installPath, _installExeName)).FileVersion;
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
            EventLogWatcher Subscribe(string path)
            {
                var subscriptionQuery = new EventLogQuery(path, PathType.LogName);

                var watcher = new EventLogWatcher(subscriptionQuery);
                watcher.EventRecordWritten +=
                    new EventHandler<EventRecordWrittenEventArgs>(EventLogEventRead);

                watcher.Enabled = true;

                return watcher;
            }

            try
            {
                _appWatcher = Subscribe("Application");
                _sysWatcher = Subscribe("System");
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
                    SendCommonErrorAsync(eventRecord.LevelDisplayName, eventRecord.ProviderName, eventRecord.FormatDescription());
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
                _mqService.Subscribe<MQCommands, (string, string, IEnumerable<string>)>(
                        (MQCommands.UserLoggedIn, userProps => OnUserLogInAsync(userProps)));

                _mqService.Subscribe<MQCommands, int>(
                        (MQCommands.Message, code => SendAtlasErrorAsync(
                            "Ошибка Атлас", code, string.Empty)));

                _mqService.Subscribe<MQCommands, bool>(
                        (MQCommands.AppStarted, code => SendAtlasStatusAsync(true)));

                _mqService.Subscribe<MQCommands, bool>(
                        (MQCommands.ExitAll, exitAlways => OnAtlasShutdownAsync()));

                _mqService.Subscribe<MQCommands, string>(
                        (MQCommands.UserLoggedOut, userName => OnUserLogOutAsync()));
            });
        }

        private async Task<bool> OnUpdateDBInfoAsync()
        {
            var dbStates = await _dbMasterEntityService.GetDatabasesStatesAsync();            
            SendDBStatesAsync(dbStates);

            return true;
        }

        private async Task OnUserLogOutAsync()
        {
            _currentUserProps = null;
        }

        private async Task OnUserLogInAsync((string UserName, string PersonInfo, IEnumerable<string> UserRoles) userProps)
        {
            _currentUserProps = userProps;
            await SendUserLogInAsync();
        }

        private async Task OnAtlasShutdownAsync()
        {
            _currentUserProps = null;
            SendAtlasShutdownAsync();
        }

        private async Task SendUserLogInAsync()
        {
            if (_currentUserProps == null)
            {
                return;
            }

            var userProps = _currentUserProps.Value;
            _ = _sendingService.SendInfoToMqttAsync(
                    MQMessages.SoftwareMsgInfo,
                    new
                    {
                        AtlasUser = new
                        {
                            User = userProps.UserName,
                            Role = userProps.UserRoles?.FirstOrDefault(),
                        },
                    });
        }

        private async Task SendAtlasStatusAsync(bool atlasRunning)
        {
            var userProps = _currentUserProps;
            _ = _sendingService.SendInfoToMqttAsync(
                    MQMessages.SoftwareMsgInfo,
                    new
                    {
                        AtlasStatus = new
                        {
                            AtlasRunning = atlasRunning,
                            AtlasUser = new
                            {
                                User = userProps?.UserName,
                                Role = userProps?.UserRoles?.FirstOrDefault(),
                            },
                        },
                    });
        }

        private async Task SendAtlasShutdownAsync()
        {
            _ = _sendingService.SendInfoToMqttAsync(
                    MQMessages.SoftwareMsgInfo,
                    new
                    {
                       SimpleMsgType = MQMessages.AtlasExited.ToString(),
                    });
        }

        private void SendCommonErrorAsync(string level, string code, string description)
        {
            _ = _sendingService.SendInfoToMqttAsync(
                    MQMessages.SoftwareMsgInfo,
                    new
                    {
                        ErrorDescriptions = new[]
                        {
                            new
                            {
                                Level = level,
                                Code = code,
                                Description = description,
                            },
                        },
                    });
        }

        private void SendDBStatesAsync(IEnumerable<(string Name, string State)> dbStates)
        {
            _ = _sendingService.SendInfoToMqttAsync(
                    MQMessages.SoftwareMsgInfo,
                    new
                    {
                        DBStates = dbStates == null ?
                            new[] { new { Name = "Все БД", Status = "OFFLINE", } }
                            :
                            dbStates.Select(s =>
                                new 
                                {
                                    Name = s.Name,
                                    Status = s.State,
                                }),
                    });
        }

        private void SendAtlasErrorAsync(string level, int code, string description)
        {
            _ = _sendingService.SendInfoToMqttAsync(
                    MQMessages.SoftwareMsgInfo,
                    new
                    {
                        AtlasErrorDescriptions = new[]
                        {
                            new
                            {
                                Level = level,
                                Code = code,
                                Description = description,
                            },
                        },
                    });
        }

        private bool IsAtlasRunning()
        {
            return Process.GetProcesses().FirstOrDefault(p => p.ProcessName == AtlasExeName) != null;
        }
    }
}
