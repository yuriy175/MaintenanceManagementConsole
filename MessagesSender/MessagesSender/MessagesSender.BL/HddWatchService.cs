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
    public class HddWatchService : IHddWatchService
    {
        private const long Megabyte = 1024 * 1024 * 1024;

        private readonly ILogger _logger;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        public HddWatchService(
            ILogger logger)
        {
            _logger = logger;
            _logger.Information("HddWatchService started");
        }

        /// <summary>
        /// gets hdd drives info
        /// </summary>
        /// <returns>drives info</returns>
        public async Task<IEnumerable<VolumeInfo>> GetDriveInfosAsync()
        {
            return DriveInfo.GetDrives()
                .Where(d => d.IsReady)
                .Select(d => new VolumeInfo
                {
                    Name = d.Name,
                    FreeSize = (long)(d.TotalFreeSpace / Megabyte),
                    TotalSize = (long)(d.TotalSize / Megabyte),
                }).ToArray();
        }
    }
}
