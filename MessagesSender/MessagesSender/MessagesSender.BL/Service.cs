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
    public class Service : IService, IDisposable
    {
        private readonly ISettingsEntityService _dbSettingsEntityService;
        private readonly IObservationsEntityService _dbObservationsEntityService;        
        private readonly ILogger _logger;
        private readonly ISendingService _sendingService;
        private readonly IMQCommunicationService _mqService;
		private readonly ICommandService _commandService;
        private readonly ISystemWatchService _systemWatchService;
        private readonly IStudyingWatchService _studyingWatchService;

        private IPAddress _ipAddress = null;
        private (string Name, string Number) _equipmentInfo = (null, null);

        private enum MessageType
        {
            StudyInWork = 1,
            ConnectionState,
        }

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="dbSettingsEntityService">settings database connector</param>
        /// <param name="dbObservationsEntityService">observations database connector</param>
        /// <param name="logger">logger</param>
        /// <param name="sendingService">sending service</param>
        /// <param name="mqService">MQ service</param>
        /// <param name="commandService">command service</param>
        /// <param name="systemWatchService">system watch service</param>
        /// <param name="studyingWatchService">studying watch service</param>
        public Service(
            ISettingsEntityService dbSettingsEntityService,
            IObservationsEntityService dbObservationsEntityService,
            ILogger logger,
            ISendingService sendingService,
            IMQCommunicationService mqService,
			ICommandService commandService,
            ISystemWatchService systemWatchService,
            IStudyingWatchService studyingWatchService)
        {
            _dbSettingsEntityService = dbSettingsEntityService;
            _dbObservationsEntityService = dbObservationsEntityService;
            _logger = logger;
            _sendingService = sendingService;
            _mqService = mqService;
			_commandService = commandService;
            _systemWatchService = systemWatchService;
            _studyingWatchService = studyingWatchService;

            new Action[]
                {
                    () => _ = SubscribeMQRecevers(),
                    async () =>
                    {
                        await _sendingService.CreateAsync();
                        await OnServiceStateChangedAsync(true);
                    }
                }.RunTasksAsync();

            _logger.Information("Main service started");
        }

        public void Dispose()
        {
            _ = OnServiceStateChangedAsync(false).Result;
        }

        private Task SubscribeMQRecevers()
        {
            return Task.Run(() =>{});
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
			//return await _sendingService.SendInfoToWorkQueueAsync(
			//   isOn ? MQMessages.InstanceOn : MQMessages.InstanceOff,
			//   new { });
			return await _sendingService.SendInfoToCommonMqttAsync(
			   isOn ? MQMessages.InstanceOn : MQMessages.InstanceOff,
			   new { });
		}

    }
}
