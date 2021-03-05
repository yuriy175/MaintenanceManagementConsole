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

        private bool _isActivated = false;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="dbObservationsEntityService">observations database connector</param>
        /// <param name="mqService">MQ service</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        public StudyingWatchService(
            ILogger logger,
            IObservationsEntityService dbObservationsEntityService,
            IMQCommunicationService mqService,
            IEventPublisher eventPublisher,
            ISendingService sendingService)
        {
            _logger = logger;
            _dbObservationsEntityService = dbObservationsEntityService;
            _mqService = mqService;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;

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

                _mqService.Subscribe<MQCommands, int>(
                        (MQCommands.NewImageCreated, async imageId => OnNewImageCreatedAsync(imageId)));

                _mqService.Subscribe<MQCommands, (OrganAuto OrganAuto, int LogicalWsId)>(
                    (MQCommands.SetOrganAuto, async organAuto => await OnOrganAutoAsync(organAuto)));
            });
        }

        private async Task<bool> OnStudyInWorkAsync(int studyId)
        {
            if (!_isActivated)
            {
                return true;
            }

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

        private async Task<bool> OnNewImageCreatedAsync(int imageId)
        {
            // _ = SendInfoAsync(_mqttSender, MQCommands.NewImageCreated, imageId);
            return true;
        }

        private async Task<bool> OnOrganAutoAsync((OrganAuto OrganAuto, int LogicalWsId) organAuto)
        {
            if (!_isActivated)
            {
                return true;
            }

            if (organAuto.OrganAuto == null)
            {
                _logger.Error("OnOrganAutoAsync error : no OrganAuto arrived");
                return false;
            }

            return await _sendingService.SendInfoToMqttAsync(
                MQCommands.SetOrganAuto,
                new { 
                    organAuto.OrganAuto.Name, 
                    organAuto.OrganAuto.Laterality,
                    organAuto.OrganAuto.Projection,
                    organAuto.OrganAuto.Direction,
                    organAuto.OrganAuto.AgeId,
                    organAuto.OrganAuto.Constitution
                });
        }

        private void OnDeactivateArrivedAsync()
        {
            _isActivated = false;
        }

        private async Task<bool> OnActivateArrivedAsync()
        {
            _isActivated = true;
            return true;
        }
    }
}