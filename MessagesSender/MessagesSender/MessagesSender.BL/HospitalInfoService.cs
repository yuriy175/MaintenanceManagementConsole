using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;
using Atlas.Acquisitions.Common.Core;
using Atlas.Acquisitions.Common.Core.Model;
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
    /// hospital info service 
    /// </summary>
    public class HospitalInfoService : IHospitalInfoService
    {
        private readonly ILogger _logger;
        private readonly ISettingsEntityService _dbSettingsEntityService;
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;

        private (string Name, string Address, double? Latitude, double? Longitude)? _hospitalInfo = null;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="dbSettingsEntityService">settings database connector</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        public HospitalInfoService(
            ILogger logger,
            ISettingsEntityService dbSettingsEntityService,
            IEventPublisher eventPublisher,
            ISendingService sendingService)
        {
            _logger = logger;
            _dbSettingsEntityService = dbSettingsEntityService;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;

            _eventPublisher.RegisterServerReadyCommandArrivedEvent(() => OnHospitalInfoArrivedAsync());
            
            GetHospitalInfoAsync();

            _logger.Information("HospitalInfoService started");
        }

        private async Task GetHospitalInfoAsync()
        {
            try
            { 
                _hospitalInfo = await _dbSettingsEntityService.GetHospitalInfoAsync();
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "GetHospitalInfoAsync");
            }
        }

        private async Task<bool> OnHospitalInfoArrivedAsync()
        {
            if (_hospitalInfo != null)
            {
                var hospitalInfo = _hospitalInfo.Value;
                await _sendingService.SendInfoToMqttAsync(
                    MQMessages.HospitalInfo,
                    new
                    {
                        HospitalName = hospitalInfo.Name,
                        HospitalAddress = hospitalInfo.Address,
                        HospitalLongitude = hospitalInfo.Longitude,
                        HospitalLatitude = hospitalInfo.Latitude,
                    });
            }

            return true;
        }
    }
}
