﻿using System;
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
    /// hardware state service
    /// </summary>
    public class HardwareStateService : IHardwareStateService
    {
        private const int GeneratorId = 1;
        private const int StandId = 1;

        private const int StandConnectedValue = 4;
        private const int GeneratorConnectedValue = 4;
        private const int CollimatorConnectedValue = 2;
        private const int DosimeterConnectedValue = 2;

        private readonly ILogger _logger;
        private readonly IMQCommunicationService _mqService;
        private readonly IWebClientService _webClientService;
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;

        private bool _isActivated = false;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="mqService">MQ service</param>
        /// <param name="webClientService">web client service</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        public HardwareStateService(
            ILogger logger,
            IMQCommunicationService mqService,
            IWebClientService webClientService,
            IEventPublisher eventPublisher,
            ISendingService sendingService)
        {
            _logger = logger;
            _mqService = mqService;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;
            _webClientService = webClientService;

            _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            _eventPublisher.RegisterDeactivateCommandArrivedEvent(() => OnDeactivateArrivedAsync());

            SubscribeMQRecevers();

            _logger.Information("HardwareStateService started");
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

                // _mqService.Subscribe<MQCommands, (int detectorId, int? detectorField)>(
                //   (MQCommands.DetectorField, state => OnDetectorFieldChanged(state)));
                _mqService.Subscribe<MQCommands, (int Id, DosimeterState State)>(
                            (MQCommands.ProcessDoseArrived, state => OnDosimeterState(state)));

                _mqService.Subscribe<MQCommands, (int Id, string Type, AecState State)>(
                            (MQCommands.AecStateArrived, state => OnAecState(state)));
            });
        }

        private void OnStandState((int Id, StandState State) state)
        {
            if (!AlwaysSendStandState(state.State) && 
                (!_isActivated || !CanSendStandState(state.State)))
            {
                return;
            }

            if (state.State.State.HasValue)
            {
                state.State.State = (int)(state.State.State.Value < StandConnectedValue ?
                    ConnectionStates.Disconnected : ConnectionStates.Connected);
            }

            _sendingService.SendInfoToMqttAsync(
                MQCommands.StandStateArrived,
                new { state.Id, state.State });
        }

        private void OnGeneratorState((int Id, GeneratorState State) state)
        {
            if (!AlwaysSendGeneratorState(state.State) &&
                (!_isActivated || !CanSendGeneratorState(state.State)))
            {
                return;
            }

            if (state.State.State.HasValue)
            {
                state.State.State = (int)(state.State.State.Value < GeneratorConnectedValue ?
                    ConnectionStates.Disconnected : ConnectionStates.Connected);
            }

            _sendingService.SendInfoToMqttAsync(
                MQCommands.GeneratorStateArrived, 
                new { state.Id, state.State });
        }

        private void OnCollimatorState((int Id, CollimatorState State) state)
        {
            if (!_isActivated || !CanSendCollimatorStandState(state.State))
            {
                return;
            }

            if (state.State.State.HasValue)
            {
                state.State.State = (uint)(state.State.State.Value < CollimatorConnectedValue ?
                    ConnectionStates.Disconnected : ConnectionStates.Connected);
            }

            _sendingService.SendInfoToMqttAsync(
                MQCommands.CollimatorStateArrived,
                new { state.Id, state.State });
        }

        private void OnDosimeterState((int Id, DosimeterState State) state)
        {
            if (!_isActivated)
            {
                return;
            }

            if (state.State.State.HasValue)
            {
                state.State.State = (uint)(state.State.State.Value < DosimeterConnectedValue ?
                    ConnectionStates.Disconnected : ConnectionStates.Connected);
            }

            _sendingService.SendInfoToMqttAsync(
                MQCommands.ProcessDoseArrived,
                new { state.Id, state.State });
        }

        private void OnAecState((int Id, string Type, AecState State) state)
        {
            if (_isActivated)
            {
                _sendingService.SendInfoToMqttAsync(
                    MQCommands.AecStateArrived,
                    new { state.Id, state.Type, state.State });
            }
        }

        private void OnDetectorStateChanged((int DetectorId, string DetectorName, DetectorState State) state)
        {
            if (_isActivated)
            {
                _sendingService.SendInfoToMqttAsync(
                    MQCommands.DetectorStateArrived,
                    new { state.DetectorId, state.DetectorName, state.State.State });
            }
        }

        private void OnDetectorFieldChanged((int DetectorId, int? DetectorField) state)
        {
            if (_isActivated)
            {
                _sendingService.SendInfoToMqttAsync(
                    MQCommands.DetectorStateArrived,
                    new { state.DetectorId, state.DetectorField });
            }
        }

        private void OnDeactivateArrivedAsync()
        {
            _isActivated = false;
        }

        private async Task<bool> OnActivateArrivedAsync()
        {
            _isActivated = true;

            var generatorState = await _webClientService.SendAsync<GeneratorState>(
                "Exposition",
                "RequestGeneratorState",
                new Dictionary<string, string> { });

            if (CanSendGeneratorState(generatorState))
            {
                OnGeneratorState((GeneratorId, generatorState));
            }

            await Task.Yield();

            var standState = await _webClientService.SendAsync<StandState>(
                "Exposition",
                "RequestStandState",
                new Dictionary<string, string> { });

            if (CanSendStandState(standState))
            {
                OnStandState((StandId, standState));
            }

            await Task.Yield();

            var result = await _webClientService.SendAsync<bool>(
                "Detectors",
                "RequestDetectorState",
                new Dictionary<string, string> { });

            await Task.Yield();

            result = await _webClientService.SendAsync<bool>(
                "Dosimeter",
                "RequestDose",
                new Dictionary<string, string> { { "dosimeterId", "1" } });

            return true;
        }

        private bool CanSendGeneratorState(GeneratorState state) =>
            state != null && (
                state.State.HasValue ||
                state.Error != null ||
                state.Kv.HasValue ||
                state.Mas.HasValue ||
                state.Workstation.HasValue ||
                state.HeatStatus.HasValue ||
                state.PedalPressed.HasValue ||
                state.Focus.HasValue
            );

        private bool CanSendStandState(StandState state) =>
            state != null && (
                state.State.HasValue ||
                state.Error != null ||
                state.RasterState.HasValue ||
                state.Position_Current.HasValue ||

                state.Mode.HasValue ||
                state.Tube_Incline.HasValue ||
                state.Deck_Incline.HasValue ||
                state.Camera_Incline.HasValue ||
                state.Ffd_Current.HasValue ||
                state.Deck_Height.HasValue
            );

        private bool CanSendCollimatorStandState(CollimatorState state) =>
            state != null && (
                state.State.HasValue ||
                !string.IsNullOrEmpty(state.FilterPresentation)
            );

        private bool AlwaysSendGeneratorState(GeneratorState state) =>
            state != null && (
                state.Error != null 
            );

        private bool AlwaysSendStandState(StandState state) =>
            state != null && (
                state.Error != null
            );
    }
}
