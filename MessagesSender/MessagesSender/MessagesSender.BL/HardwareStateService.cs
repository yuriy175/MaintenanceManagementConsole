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
    public class HardwareStateService : IHardwareStateService
    {
        private const int StandConnectedValue = 4;
        private const int GeneratorConnectedValue = 4;
        private const int CollimatorConnectedValue = 2;
        private const int DosimeterConnectedValue = 2;

        private readonly ILogger _logger;
        private readonly IMQCommunicationService _mqService;
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;

        private ConnectionStates _standState = ConnectionStates.Disconnected;
        private ConnectionStates _generatorState = ConnectionStates.Disconnected;
        private ConnectionStates _collimatorState = ConnectionStates.Disconnected;
        private ConnectionStates _detectorState = ConnectionStates.Disconnected;

        private bool _isActivated = false;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="mqService">MQ service</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        public HardwareStateService(
            ILogger logger,
            IMQCommunicationService mqService,
            IEventPublisher eventPublisher,
            ISendingService sendingService)
        {
            _logger = logger;
            _mqService = mqService;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;

            _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            _eventPublisher.RegisterDeactivateCommandArrivedEvent(() => OnDeactivateArrivedAsync());

            SubscribeMQRecevers();

            _logger.Information("HardwareStateService started");
        }

        private object GetStandState(StandState state) =>
            GetHardwareState(state?.State, StandConnectedValue, ref _standState);

        private object GetGeneratorState(GeneratorState state) =>
            GetHardwareState(state?.State, GeneratorConnectedValue, ref _generatorState);

        private object GetCollimatorState(CollimatorState state) =>
            GetHardwareState((int?)state?.State, CollimatorConnectedValue, ref _collimatorState);

        private object GetHardwareState(int? state, int hwConnectedValue, ref ConnectionStates hwConnectionStates)
        {
            var connectionState = ConnectionStates.Disconnected;
            if (state.HasValue)
            {
                connectionState = state.Value <= hwConnectedValue ?
                    ConnectionStates.Disconnected : ConnectionStates.Connected;
            }

            if (connectionState != _standState)
            {
                hwConnectionStates = connectionState;
                return new { State = _standState };
            }

            return null;
        }

        private Task SubscribeMQRecevers()
        {
            return Task.Run(() =>
            {
                _mqService.Subscribe<MQCommands, (int Id, GeneratorState State)>(
                    (MQCommands.GeneratorStateArrived, state => OnGeneratorState(state)));

                _mqService.Subscribe<MQCommands, (int Id, StandState State)>(
                    (MQCommands.StandStateArrived, state => OnStandState(state)));

                _mqService.Subscribe<MQCommands, (int Id, CollimatorState State)>(
                    (MQCommands.CollimatorStateArrived, state => OnCollimatorState(state)));

                _mqService.Subscribe<MQCommands, (int detectorId, string detectorName, DetectorState state)>(
                    (MQCommands.DetectorStateArrived, state => OnDetectorStateChanged(state)));
            });
        }

        private void OnStandState((int Id, StandState State) state)
        {
            var standState = GetStandState(state.State);
            if (standState != null)
            {
                _sendingService.SendInfoToMqttAsync(MQCommands.StandStateArrived, standState);
            }
        }

        private void OnGeneratorState((int Id, GeneratorState State) state)
        {
            var standState = GetGeneratorState(state.State);
            if (standState != null)
            {
                _sendingService.SendInfoToMqttAsync(MQCommands.GeneratorStateArrived, standState);
            }
        }

        private void OnCollimatorState((int Id, CollimatorState State) state)
        {
            var standState = GetCollimatorState(state.State);
            if (standState != null)
            {
                _sendingService.SendInfoToMqttAsync(MQCommands.CollimatorStateArrived, standState);
            }
        }

        private void OnDetectorStateChanged((int DetectorId, string DetectorName, DetectorState State) state)
        {
            if (_isActivated)
            {
                _sendingService.SendInfoToMqttAsync(MQCommands.DetectorStateArrived, state);
            }
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
