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

            CreateConnectionProps();

            _logger.Information("HddWatchService started");
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

            if (!_connectionProps.HasValue) 
            {
                return false;
            }

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
					_ = _sendingService.SendInfoToMqttAsync(MQMessages.HddDrivesInfo, hddDrives);
				}

				await Task.Yield();
				if (!_isActivated)
				{
					break;
				}

				var ramInfo = await GetRamInfoAsync();
				if (ramInfo.HasValue)
				{
					_ = _sendingService.SendInfoToMqttAsync(MQMessages.MemoryInfo,
						new { ramInfo.Value.AvailableSize, TotalMemory = ramInfo.Value.TotalSize });
				}

				await Task.Yield();
				if (!_isActivated)
				{
					break;
				}

				var cpuInfo = await GetCpuInfoAsync();
				if (ramInfo.HasValue)
				{
					_ = _sendingService.SendInfoToMqttAsync(MQMessages.CPUInfo,
						new { cpuInfo.Value.Model, cpuInfo.Value.CPU_Load });
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

        private async Task<(string Model, long CPU_Load)?> GetCpuInfoAsync()
        {
            try
            {
                await Task.Yield();
                float prevLoad = _totalCpu.NextValue();
                float prevIdle = _idleCpu.NextValue();
                System.Threading.Thread.Sleep(1000); //This avoid that answer always 0
                float load = _totalCpu.NextValue();
                float idle = _idleCpu.NextValue();
                var diff = (load - idle) / Environment.ProcessorCount;

                Console.WriteLine("total "+ load + " idle" + idle + " cpu load " + diff);

                return (string.Empty, (long)diff);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "GetRamInfoAsync error");
                return (string.Empty, 0);
            }
        }        

        private async Task<(long TotalSize, long AvailableSize)?> GetRamInfoAsync()
        {
            try
            {
                var gcMemoryInfo = GC.GetGCMemoryInfo();

                var ramCounter = new PerformanceCounter("Memory", "Available MBytes");
                return ((long)(gcMemoryInfo.TotalAvailableMemoryBytes / Megabyte), (long)(ramCounter.NextValue()));
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "GetRamInfoAsync error");
                return (0, 0);
            }
        }
        #endregion

        private void CreateConnectionProps()
        {
            var connectionString = _configurationService.Get<string>(MessagesSenderModel.Constants.RabbitMQConnectionStringName, null);
            try
            {
                _connectionProps = ConnectionPropsCreator.Create(connectionString);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "Rabbit MQ work queue wrong connection string");
            }
        }

        private void RunCommand(string exePath, string args)
        {
            var path = Path.Combine(Path.GetDirectoryName(Assembly.GetExecutingAssembly().Location), exePath);
            var processStartInfo = new ProcessStartInfo(
                path, 
                // @"D:\Gits\MaintenanceManagementConsole\MessagesSender\MessagesSender\bin\Debug\netcoreapp3.1\MQTTsysinfo\MQTTsysinfo.exe",
                args);
            processStartInfo.WorkingDirectory = Path.GetDirectoryName(path);
            
            var process = Process.Start(processStartInfo);
        }
    }
}
