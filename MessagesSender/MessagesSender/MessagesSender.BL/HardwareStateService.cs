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
        private const int StandBlockedValue = 7;        

        private readonly ILogger _logger;
        private readonly IMQCommunicationService _mqService;
        private readonly IWebClientService _webClientService;
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;

        private bool _isActivated = false;
        private DetectorStates _detectorState = DetectorStates.Invalid;

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

                _mqService.Subscribe<MQCommands, (int detectorId, int? detectorField)>(
                   (MQCommands.DetectorField, state => OnDetectorFieldChanged(state)));

                _mqService.Subscribe<MQCommands, (int Id, DosimeterState State)>(
                            (MQCommands.ProcessDoseArrived, state => OnDosimeterState(state)));

                _mqService.Subscribe<MQCommands, (int Id, string Type, AecState State)>(
                            (MQCommands.AecStateArrived, state => OnAecState(state)));

                _mqService.Subscribe<MQCommands, (HardwareParams HardwareParam, string Value)>(
                    (MQCommands.SetHwParameter, value => OnSetParameter(value)));
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

                if (state.State.State == StandBlockedValue)
                {
                    SendHardwareErrorAsync("Сообщение от штатива", new[]
                    {
                        new DeviceError
                        {
                            Code = "Активирован аварийный выключатель",
                            Description = string.Empty,
                        },
                    });
                }
            }

            if (state.State.ErrorDescriptions != null)
            {
                SendHardwareErrorAsync("Ошибка штатива", state.State.ErrorDescriptions);
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

            if (state.State.ErrorDescriptions != null)
            {
                SendHardwareErrorAsync("Ошибка генератора", state.State.ErrorDescriptions);
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
            if (state.State.State != _detectorState) // _isActivated)
            {
                _sendingService.SendInfoToMqttAsync(
                    MQCommands.DetectorStateArrived,
                    new { state.DetectorId, state.DetectorName, state.State.State });
            }

            if (_detectorState == DetectorStates.Created && state.State.State == DetectorStates.CreationFailed)
            {
                SendHardwareErrorAsync("Ошибка детектора", new[] 
                { 
                    new DeviceError 
                    { 
                        Code = "Ошибка инициализации детектора",
                        Description = string.Empty,
                    }, 
                });
            }

            _detectorState = state.State.State;
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

        private void OnSetParameter((HardwareParams HardwareParam, string Value) parameter)
        {
            if (!_isActivated)
            {
                return;
            }

            if (parameter.HardwareParam == HardwareParams.FrameRate)
            {
                var value = JsonConvert.DeserializeObject(parameter.Value, typeof((ShootingModes, float))) as (ShootingModes Mode, float Value)?;
                if (value.HasValue)
                {
                    _sendingService.SendInfoToMqttAsync(
                        MQCommands.DetectorStateArrived,
                        new { DetectorId = 1, DetectorFrameRate = value.Value.Value });
                }
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

        private void SendHardwareErrorAsync(string level, DeviceError[] errors)
        {
            _ = _sendingService.SendInfoToMqttAsync(
                    MQMessages.SoftwareMsgInfo,
                    new
                    {
                        HardwareErrorDescriptions = errors.Select(e =>
                            new
                            {
                                Level = level,
                                Code = e.Code,
                                Description = e.Message,
                            }).ToArray(),
                    });
        }

        private bool CanSendGeneratorState(GeneratorState state) =>
            state != null && (
                state.State.HasValue ||
                state.Error != null ||
                state.Kv.HasValue ||
                state.Mas.HasValue ||
                state.Ma.HasValue ||
                state.Ms.HasValue ||
                state.Post_kv.HasValue ||
                state.Post_ma.HasValue ||
                state.Post_mas.HasValue ||
                state.Post_time.HasValue ||
                state.Workstation.HasValue ||
                state.HeatStatus.HasValue ||
                state.PedalPressed.HasValue ||
                state.Focus.HasValue ||
                state.Points_mode.HasValue ||
                state.Scopy_kv.HasValue ||
                state.Scopy_mas.HasValue ||
                state.Scopy_ma.HasValue ||
                state.Scopy_ms.HasValue ||
                state.Scopy_post_kv.HasValue ||
                state.Scopy_post_ms.HasValue ||
                state.Scopy_post_mas.HasValue ||
                state.Scopy_post_ma.HasValue ||
                state.Scopy_pps.HasValue ||
                state.Scopy_mode.HasValue
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
                state.Deck_Height.HasValue ||
                state.Uarm_Height.HasValue
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
