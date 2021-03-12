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
    /// software watch service 
    /// </summary>
    public class SoftwareWatchService : ISoftwareWatchService
    {
        private readonly ILogger _logger; 
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        public SoftwareWatchService(
            ILogger logger,
            IEventPublisher eventPublisher,
            ISendingService sendingService)
        {
            _logger = logger;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;

            _logger.Information("SoftwareWatchService started");

            var t = FileVersionInfo.GetVersionInfo(@"C:\Program Files\Atlas\Services\XiLibs\XiLibNet.dll").FileVersion;
        }
    }
}
