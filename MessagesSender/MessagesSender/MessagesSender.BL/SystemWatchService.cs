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
    /// system watch service interface
    /// </summary>
    public class SystemWatchService : ISystemWatchService
    {
        private const string MQTTsysinfoFolder = @".\MQTTsysinfo\MQTTsysinfo.exe";
        private const string MQTTsysinfoCommandLine = "{0} {1} {2} {3}";
        private const long Megabyte = 1024 * 1024;
        private const long Gigabyte = 1024 * 1024 * 1024;
        private const int WatchInterval = 5000;

        private readonly IConfigurationService _configurationService;
        private readonly ILogger _logger; 
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;
        private readonly ITopicService _topicService;

        private readonly PerformanceCounter _totalCpu = new PerformanceCounter("Process", "% Processor Time", "_Total");
        private readonly PerformanceCounter _idleCpu = new PerformanceCounter("Process", "% Processor Time", "Idle");

        private (string HostName, string UserName, string Password)? _connectionProps;
        private bool _isActivated = false;
        private bool _isStartupSent = false;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        /// <param name="topicService">topic service</param>
        public SystemWatchService(
            IConfigurationService configurationService,
            ILogger logger,
            IEventPublisher eventPublisher,
            ISendingService sendingService,
            ITopicService topicService)
        {
            _configurationService = configurationService;
            _logger = logger;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;
            _topicService = topicService;

            _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            _eventPublisher.RegisterDeactivateCommandArrivedEvent(() => OnDeactivateArrivedAsync());

            _eventPublisher.RegisterServerReadyCommandArrivedEvent(() => SendSystemStartupEventAsync());
            _logger.Information("HddWatchService started");
        }

        /// <summary>
        /// Dispose object
        /// </summary>
        public void Dispose()
        {
            _idleCpu?.Dispose();
            _totalCpu?.Dispose();
        }

        private void OnDeactivateArrivedAsync()
        {
            _isActivated = false;
        }

        private async Task<bool> OnActivateArrivedAsync()
        {
            if (_isActivated)
            {
                return true;
            }

            /*if (!_connectionProps.HasValue) 
            {
                return false;
            }*/

            _isActivated = true;

            /*RunCommand(
                MQTTsysinfoFolder, 
                string.Format(
                    MQTTsysinfoCommandLine,
                    //mprom.ml client1 medtex KRT/TESTARM 
                    _connectionProps.Value.HostName,
                    _connectionProps.Value.UserName,
                    _connectionProps.Value.Password,
                    await _topicService.GetTopicAsync()
                    ));*/

            while (_isActivated)
            {
                var hddDrives = await GetDriveInfosAsync();
                if (hddDrives != null)
                {
                    _ = _sendingService.SendInfoToMqttAsync(
                        MQMessages.HddDrivesInfo, 
                        new { HDD = hddDrives });
                }

                await Task.Yield();
                if (!_isActivated)
                {
                    break;
                }

                var ramInfo = await GetRamInfoAsync();
                if (ramInfo.HasValue)
                {
                    _ = _sendingService.SendInfoToMqttAsync(
                        MQMessages.MemoryInfo,
                        new
                        {
                            Memory = new
                            { 
                                MemoryFreeGb = ramInfo.Value.AvailableSize, 
                                MemoryTotalGb = ramInfo.Value.TotalSize,
                            },
                        });
                }

                await Task.Yield();
                if (!_isActivated)
                {
                    break;
                }

                var cpuInfo = await GetCpuInfoAsync();
                if (ramInfo.HasValue)
                {
                    _ = _sendingService.SendInfoToMqttAsync(
                        MQMessages.CPUInfo,
                        new
                        {
                            Processor = new { cpuInfo.Value.Model, cpuInfo.Value.CPULoad }
                        });
                }

                if (!_isActivated)
                {
                    break;
                }

                await Task.Delay(WatchInterval);
            }

            return false;
        }

        #region depricated region

        /// <summary>
        /// gets hdd drives info
        /// </summary>
        /// <returns>drives info</returns>
        private async Task<IEnumerable<MessagesSenderModel.VolumeInfo>> GetDriveInfosAsync()
        {
            return DriveInfo.GetDrives()
                .Where(d => d.IsReady)
                .Select(d => new MessagesSenderModel.VolumeInfo
                {
                    Letter = d.Name,
                    FreeSize = (long)(d.TotalFreeSpace / Gigabyte),
                    TotalSize = (long)(d.TotalSize / Gigabyte),
                }).ToArray();
        }

        private async Task<(string Model, long CPULoad)?> GetCpuInfoAsync()
        {
            try
            {
                await Task.Yield();
                float prevLoad = _totalCpu.NextValue();
                float prevIdle = _idleCpu.NextValue();
                System.Threading.Thread.Sleep(1000); 
                float load = _totalCpu.NextValue();
                float idle = _idleCpu.NextValue();
                var diff = Math.Abs(load - idle) / Environment.ProcessorCount;

                Console.WriteLine("total " + load + " idle" + idle + " cpu load " + diff);

                return (string.Empty, (long)diff);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "GetRamInfoAsync error");
                return (string.Empty, 0);
            }
        }        

        private async Task<(float TotalSize, float AvailableSize)?> GetRamInfoAsync()
        {
            try
            {
                var gcMemoryInfo = GC.GetGCMemoryInfo();

                var ramCounter = new PerformanceCounter("Memory", "Available MBytes");
                return ((float)Math.Round((double)gcMemoryInfo.TotalAvailableMemoryBytes / Gigabyte, 2),
                    (float)Math.Round(ramCounter.NextValue() / 1024, 2));
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "GetRamInfoAsync error");
                return (0, 0);
            }
        }
        #endregion

        private async Task<bool> SendSystemStartupEventAsync()
        {
            if (_isStartupSent)
            {
                return true;
            }

            var ticks = Environment.TickCount;
            var startupTime = DateTime.Now - TimeSpan.FromMilliseconds(ticks);
            var eventLog = new EventLog("System");

            var mostRecentWake =
                EnumerateLog(eventLog, "Microsoft-Windows-Kernel-Power", 41)
                .OrderByDescending(item => item.TimeGenerated)
                .LastOrDefault();

            _ = _sendingService.SendInfoToMqttAsync(
                        MQMessages.StartupInfo,
                        new
                        {
                            StartupTime = startupTime,
                            KernelPower41 = mostRecentWake == null ? 
                                null as DateTime? : mostRecentWake.TimeGenerated,
                        });

            _isStartupSent = true;

            return true;
        }

        private IEnumerable<EventLogEntry> EnumerateLog(EventLog log, string source, int? code = null)
        {
            foreach (EventLogEntry entry in log.Entries)
                if (entry.Source == source && (code.HasValue ? entry.EventID == code : true ))
                    yield return entry;
        }
    }
}
