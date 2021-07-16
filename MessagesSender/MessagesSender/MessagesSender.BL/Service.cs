using System;
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
    /// main service interface implementation
    /// </summary>
    public class Service : IService, IDisposable
    {
        private readonly IObservationsEntityService _dbObservationsEntityService;
        private readonly ILogger _logger;
        private readonly ISendingService _sendingService;
        private readonly IMQCommunicationService _mqService;
        private readonly ICommandService _commandService;
        private readonly ISystemWatchService _systemWatchService;
        private readonly IStudyingWatchService _studyingWatchService;
        private readonly IHardwareStateService _hwStateService;
        private readonly ISoftwareWatchService _softwareWatchService;
        private readonly IEventPublisher _eventPublisher;
        private readonly IDicomStateService _dicomStateService;
        private readonly IRemoteControlService _remoteControlService;
        private readonly IImagesWatchService _imagesWatchService;
        private readonly IHospitalInfoService _hospitalInfoService;
        private readonly IDBDataService _dbDataService;
        private readonly IOfflineService _offlineService;
        private readonly IKeepAliveService _keepAliveService;

        private IPAddress _ipAddress = null;
        private (string Name, string Number) _equipmentInfo = (null, null);

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="dbObservationsEntityService">observations database connector</param>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        /// <param name="mqService">MQ service</param>
        /// <param name="commandService">command service</param>
        /// <param name="systemWatchService">system watch service</param>
        /// <param name="studyingWatchService">studying watch service</param>
        /// <param name="hwStateService">hardware state service</param>
        /// <param name="softwareWatchService">software watch service</param>
        /// <param name="dicomStateService">dicom state service</param>
        /// <param name="remoteControlService">remote control service</param>
        /// <param name="imagesWatchService">images watch service</param>
        /// <param name="hospitalInfoService">hospital info service</param>
        /// <param name="dbDataService">db raw data service</param>
        /// <param name="offlineService">offline service</param>
        public Service(
            IObservationsEntityService dbObservationsEntityService,
            ILogger logger,
            IEventPublisher eventPublisher,
            ISendingService sendingService,
            IMQCommunicationService mqService,
            ICommandService commandService,
            ISystemWatchService systemWatchService,
            IStudyingWatchService studyingWatchService,
            IHardwareStateService hwStateService,
            ISoftwareWatchService softwareWatchService,
            IDicomStateService dicomStateService,
            IRemoteControlService remoteControlService,
            IImagesWatchService imagesWatchService,
            IHospitalInfoService hospitalInfoService,
            IDBDataService dbDataService,
            IOfflineService offlineService,
            IKeepAliveService keepAliveService)
        {
            _dbObservationsEntityService = dbObservationsEntityService;
            _logger = logger;
            _sendingService = sendingService;
            _mqService = mqService;
            _commandService = commandService;
            _systemWatchService = systemWatchService;
            _studyingWatchService = studyingWatchService;
            _hwStateService = hwStateService;
            _eventPublisher = eventPublisher;
            _softwareWatchService = softwareWatchService;
            _dicomStateService = dicomStateService;
            _remoteControlService = remoteControlService;
            _imagesWatchService = imagesWatchService;
            _hospitalInfoService = hospitalInfoService;
            _dbDataService = dbDataService;
            _offlineService = offlineService;
            _keepAliveService = keepAliveService;

            new Action[]
                {
                    () => _ = SubscribeMQRecevers(),
                    async () =>
                    {
                        await _sendingService.CreateAsync();
                        await OnServiceStateChangedAsync(true);
                    },
                    () => _eventPublisher.RegisterReconnectCommandArrivedEvent(() => OnReconnectArrivedAsync())
                }.RunTasksAsync();

            _logger.Information("Main service started");
        }
        
        private enum MessageType
        {
            StudyInWork = 1,
            ConnectionState,
        }

        /// <inheritdoc/>
        public void Dispose()
        {
            _ = OnServiceStateChangedAsync(false).Result;
        }

        private Task SubscribeMQRecevers()
        {
            return Task.Run(() => { });
        }

        private async Task<bool> OnConnectionStateArrivedAsync(
            (int Id, string Name, string Type, DeviceConnectionState Connection) state)
        {
            return await _sendingService.SendInfoToMqttAsync(
                MQCommands.HwConnectionStateArrived,
                new { state.Id, state.Name, state.Type, state.Connection });
        }

        private async Task<bool> OnServiceStateChangedAsync(bool isOn)
        {
            var result = await _sendingService.SendInfoToCommonMqttAsync(
               isOn ? MQMessages.InstanceOn : MQMessages.InstanceOff,
               new { });

            return result;
        }

        private async Task<bool> OnReconnectArrivedAsync()
        {
            return await OnServiceStateChangedAsync(true);
        }
    }
}
