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

namespace MessagesSender.BL
{
    /// <summary>
    /// system watch service interface
    /// </summary>
    public class SystemWatchService : ISystemWatchService
    {
        private const long Megabyte = 1024 * 1024;
        private const long Gigabyte = 1024 * 1024 * 1024;

        private readonly ILogger _logger; 
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        public SystemWatchService(
            ILogger logger,
            IEventPublisher eventPublisher,
            ISendingService sendingService)
        {
            _logger = logger;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;

            _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            _eventPublisher.RegisterDeactivateCommandArrivedEvent(() => OnDeactivateArrivedAsync());

            _logger.Information("HddWatchService started");
        }

        private void OnDeactivateArrivedAsync()
        {
            
        }

        private async Task<bool> OnActivateArrivedAsync()
        {
            var hddDrives = await GetDriveInfosAsync();
            if (hddDrives != null)
            {
                _sendingService.SendInfoToMqttAsync(MQMessages.HddDrivesInfo, hddDrives);
            }

            await Task.Yield();

            var ramInfo = await GetRamInfoAsync();
            if (ramInfo != null)
            {
                _sendingService.SendInfoToMqttAsync(MQMessages.MemoryInfo, ramInfo);
            }

            return false;
        }

        /// <summary>
        /// gets hdd drives info
        /// </summary>
        /// <returns>drives info</returns>
        private async Task<IEnumerable<VolumeInfo>> GetDriveInfosAsync()
        {
            return DriveInfo.GetDrives()
                .Where(d => d.IsReady)
                .Select(d => new VolumeInfo
                {
                    Letter = d.Name,
                    FreeSize = (long)(d.TotalFreeSpace / Gigabyte),
                    TotalSize = (long)(d.TotalSize / Gigabyte),
                }).ToArray();
        }


        //var total_cpu = new PerformanceCounter("Process", "% Processor Time", "_Total");
        //float t = total_cpu.NextValue();
        //System.Threading.Thread.Sleep(1000); //This avoid that answer always 0
        //t = (int)total_cpu.NextValue();

        private async Task<(long TotalSize, long AvailableSize)?> GetRamInfoAsync()
        {
            try
            {
                var gcMemoryInfo = GC.GetGCMemoryInfo();

                var ramCounter = new PerformanceCounter("Memory", "Available MBytes");
                return ((long)(gcMemoryInfo.TotalAvailableMemoryBytes / Megabyte), (long)(ramCounter.NextValue() / Megabyte));
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "GetRamInfoAsync error");
                return (0, 0);
            }
        }
    }
}
