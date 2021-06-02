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
    /// dicom state service 
    /// </summary>
    public class DicomStateService : IDicomStateService
    {
        private const int PACSServiceRole = 1;
        private const int WorkListServiceRole = 4;

        private readonly ILogger _logger;
        private readonly ISettingsEntityService _dbSettingsEntityService;
        private readonly IWebClientService _webClientService;
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;

        private bool _isActivated = false;
        private IEnumerable<(int Id, string Name, string IP, int ServiceRole)> _dicomServices = null;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="dbSettingsEntityService">settings database connector</param>
        /// <param name="webClientService">web client service</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        public DicomStateService(
            ILogger logger,
            ISettingsEntityService dbSettingsEntityService,
            IWebClientService webClientService,
            IEventPublisher eventPublisher,
            ISendingService sendingService)
        {
            _logger = logger;
            _dbSettingsEntityService = dbSettingsEntityService;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;
            _webClientService = webClientService;

            _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            _eventPublisher.RegisterDeactivateCommandArrivedEvent(() => OnDeactivateArrivedAsync());

            GetDicomServicesAsync();

            _logger.Information("DicomStateService started");
        }

        private async Task GetDicomServicesAsync()
        {
            _dicomServices = await _dbSettingsEntityService.GetDicomServicesAsync();
        }

        private void OnDeactivateArrivedAsync()
        {
            _isActivated = false;
        }

        private async Task<bool> OnActivateArrivedAsync()
        {
            _isActivated = true;

            await SendDicomServicesAsync();

            _dicomServices.Where(d => (d.ServiceRole & PACSServiceRole) > 0 || (d.ServiceRole & WorkListServiceRole) > 0)
                .ToList()
                .ForEach(d =>
                {
                    Task.Run(async () =>
                    {
                        var state = await _webClientService.SendAsync<bool>(
                            "Verify",
                            "CheckService",
                            new Dictionary<string, string> { { "serviceId", d.Id.ToString() } });

                        await SendDicomServiceStateAsync(d, state);
                    });
                });

            return true;
        }

        private async Task SendDicomServicesAsync()
        {
            if (_isActivated && _dicomServices != null)
            {
                await _sendingService.SendInfoToMqttAsync(
                    MQMessages.DicomInfo,
                    new
                    {
                        PACS = _dicomServices.Where(d => (d.ServiceRole & PACSServiceRole) > 0)
                            .Select(d => new { d.Name, d.IP }),
                        WorkList = _dicomServices.Where(d => (d.ServiceRole & WorkListServiceRole) > 0)
                            .Select(d => new { d.Name, d.IP }),
                    });
            }
        }

        private async Task SendDicomServiceStateAsync((int Id, string Name, string IP, int ServiceRole) dicomService, bool state)
        {
            if (_isActivated)
            {
                var isWL = (dicomService.ServiceRole & WorkListServiceRole) > 0;
                await _sendingService.SendInfoToMqttAsync(
                    MQMessages.DicomInfo,                    
                    new
                    {
                        WorkList = isWL ? new[] { new { dicomService.Name, dicomService.IP, State = state } } : null,
                        PACS = isWL ? null : new[] { new { dicomService.Name, dicomService.IP, State = state } },
                    });
            }
        }
    }
}
