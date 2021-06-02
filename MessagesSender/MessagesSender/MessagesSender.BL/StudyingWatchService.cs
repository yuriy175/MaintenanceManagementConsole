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
    /// studying watch service
    /// </summary>
    public class StudyingWatchService : IStudyingWatchService
    {
        private readonly ILogger _logger;
        private readonly IObservationsEntityService _dbObservationsEntityService;
        private readonly IMQCommunicationService _mqService;
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;
        private readonly IWebClientService _webClientService;

        private bool _isActivated = false;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="dbObservationsEntityService">observations database connector</param>
        /// <param name="mqService">MQ service</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        /// <param name="webClientService">web client service</param>
        public StudyingWatchService(
            ILogger logger,
            IObservationsEntityService dbObservationsEntityService,
            IMQCommunicationService mqService,
            IEventPublisher eventPublisher,
            ISendingService sendingService,
            IWebClientService webClientService)
        {
            _logger = logger;
            _dbObservationsEntityService = dbObservationsEntityService;
            _mqService = mqService;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;
            _webClientService = webClientService;

            _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            _eventPublisher.RegisterDeactivateCommandArrivedEvent(() => OnDeactivateArrivedAsync());

            SubscribeMQRecevers();

            _logger.Information("StudyingWatchService started");
        }

        private Task SubscribeMQRecevers()
        {
            return Task.Run(() =>
            {
                _mqService.Subscribe<MQCommands, int>(
                    (MQCommands.StudyInWork, async data => OnStudyInWorkAsync(data)));

                _mqService.Subscribe<MQCommands, (OrganAuto OrganAuto, int LogicalWsId)>(
                    (MQCommands.SetOrganAuto, async organAuto => await OnOrganAutoAsync(organAuto)));
            });
        }

        private async Task<bool> OnStudyInWorkAsync(int studyId)
        {
            var studyProps = await _dbObservationsEntityService.GetStudyInfoByIdAsync(studyId);
            if (!studyProps.HasValue)
            {
                _logger.Error($"no study found for {studyId}");
                return false;
            }

            return await _sendingService.SendInfoToMqttAsync(
                MQCommands.StudyInWork,
                new { studyProps.Value.StudyId, studyProps.Value.StudyDicomUid, studyProps.Value.StudyName });
        }

        private async Task<bool> OnOrganAutoAsync((OrganAuto OrganAuto, int LogicalWsId) organAuto)
        {
            if (organAuto.OrganAuto == null)
            {
                _logger.Error("OnOrganAutoAsync error : no OrganAuto arrived");
                return false;
            }

            return await _sendingService.SendInfoToMqttAsync(
                MQCommands.SetOrganAuto,
                new
                { 
                    organAuto.OrganAuto.Name, 
                    organAuto.OrganAuto.Laterality,
                    organAuto.OrganAuto.Projection,
                    organAuto.OrganAuto.Direction,
                    organAuto.OrganAuto.AgeId,
                    organAuto.OrganAuto.Constitution,
                });
        }

        private void OnDeactivateArrivedAsync()
        {
            _isActivated = false;
        }

        private async Task<bool> OnActivateArrivedAsync()
        {
            _isActivated = true;

            var organAuto = await _webClientService.SendAsync<OrganAuto>(
                "OrganAutoManipulation",
                "GetCurrentOrganAuto",
                new Dictionary<string, string> { });

            if (organAuto != null)
            {
                return await OnOrganAutoAsync((organAuto, 1));
            }

            return false;
        }
    }
}
