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
        private readonly IMQCommunicationService _mqService;
        private readonly IWorkqueueSender _wqSender;
        private readonly IMqttSender _mqttSender;
        private readonly IMqttReceiver _mqttReceiver;
        private readonly IHardwareStateService _hwStateService;

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
        /// <param name="mqService">MQ service</param>
        /// <param name="wqSender">work queue sender</param>
        /// <param name="mqttSender">mqtt sender</param>
        /// <param name="mqttReceiver">mqtt receiver</param>
        /// <param name="hwStateService">hardware state service</param>
        public Service(
            ISettingsEntityService dbSettingsEntityService,
            IObservationsEntityService dbObservationsEntityService,
            ILogger logger,
            IMQCommunicationService mqService,
            IWorkqueueSender wqSender,
            IMqttSender mqttSender,
            IMqttReceiver mqttReceiver,
            IHardwareStateService hwStateService)
        {
            _dbSettingsEntityService = dbSettingsEntityService;
            _dbObservationsEntityService = dbObservationsEntityService;
            _logger = logger;
            _mqService = mqService;
            _wqSender = wqSender;
            _mqttSender = mqttSender;
            _mqttReceiver = mqttReceiver;
            _hwStateService = hwStateService;

            new Action[]
                {
                    () => _ = SubscribeMQRecevers(),
                    async () =>
                    {
                        await GetEquipmentInfoAsync();
                        await _mqttSender.CreateAsync(_equipmentInfo);
                        await _mqttReceiver.CreateAsync(_equipmentInfo);
                        await OnServiceStateChangedAsync(true);
                    },
                    () => _ = GetEquipmentIPAsync(),
                }.RunTasksAsync();

            _logger.Information("Main service started");
        }

        public void Dispose()
        {
            OnServiceStateChangedAsync(false);
        }

        private Task SubscribeMQRecevers()
        {
            return Task.Run(() =>
            {
                _mqService.Subscribe<MQCommands, int>(
                    (MQCommands.StudyInWork, async data => OnStudyInWorkAsync(data)));

                // _mqService.Subscribe<MQCommands, (int Id, string Name, string Type, DeviceConnectionState Connection)>(
                //        (MQCommands.HwConnectionStateArrived, state => OnConnectionStateArrivedAsync(state)));
                _mqService.Subscribe<MQCommands, (int Id, GeneratorState State)>(
                    (MQCommands.GeneratorStateArrived, state => OnGeneratorState(state)));

                _mqService.Subscribe<MQCommands, (int Id, StandState State)>(
                    (MQCommands.StandStateArrived, state => OnStandState(state)));

                _mqService.Subscribe<MQCommands, (int Id, CollimatorState State)>(
                    (MQCommands.CollimatorStateArrived, state => OnCollimatorState(state)));

                //_mqService.Subscribe<MQCommands, (int detectorId, string detectorName, DetectorState state)>(
                //    (MQCommands.DetectorStateArrived, state => OnDetectorStateChanged(state)));

                _mqService.Subscribe<MQCommands, int>(
                        (MQCommands.NewImageCreated, async imageId => OnNewImageCreatedAsync(imageId)));
            });
        }

        private async Task GetEquipmentInfoAsync()
        {
            _equipmentInfo = await _dbSettingsEntityService.GetEquipmentInfoAsync();
        }

        private async Task GetEquipmentIPAsync()
        {
            if (!System.Net.NetworkInformation.NetworkInterface.GetIsNetworkAvailable())
            {
                return;
            }
            IPHostEntry host = Dns.GetHostEntry(Dns.GetHostName());
            _ipAddress = host
               .AddressList
               .FirstOrDefault(ip => ip.AddressFamily == AddressFamily.InterNetwork);
        }

        private async Task<bool> OnStudyInWorkAsync(int studyId)
        {
            var studyProps = await _dbObservationsEntityService.GetStudyInfoByIdAsync(studyId);
            if (!studyProps.HasValue)
            {
                _logger.Error($"no study found for {studyId}");
                return false;
            }

            _ = SendInfoAsync(
                _mqttSender,
                MQCommands.StudyInWork,
                new { studyProps.Value.StudyId, studyProps.Value.StudyDicomUid, studyProps.Value.StudyName });

            return true;
        }

        private async Task<bool> OnNewImageCreatedAsync(int imageId)
        {
            // _ = SendInfoAsync(_mqttSender, MQCommands.NewImageCreated, imageId);
            return true;
        }

        private async Task<bool> OnConnectionStateArrivedAsync(
            (int Id, string Name, string Type, DeviceConnectionState Connection) state)
        {
            _ = SendInfoAsync(
                _mqttSender,
                MQCommands.HwConnectionStateArrived, 
                new { state.Id, state.Name, state.Type, state.Connection });
            return true;
        }

        private async Task<bool> OnServiceStateChangedAsync(bool isOn)
        {
            _ = SendInfoAsync(
                _wqSender,
                isOn ? MQMessages.InstanceOn : MQMessages.InstanceOff,
                new { });
            return true;
        }

        private async Task SendInfoAsync<TMsgType, T>(IMQSenderBase sender, TMsgType msgType, T info)
        {
            if (string.IsNullOrEmpty(_equipmentInfo.Number) || string.IsNullOrEmpty(_equipmentInfo.Name))
            {
                _logger.Error($"wrong equipment props {_equipmentInfo.Number} {_equipmentInfo.Name}");
                return;
            }

            await sender.SendAsync(
				msgType,
				new { 
                    _equipmentInfo.Number, 
                    _equipmentInfo.Name, 
                    ipAddress = _ipAddress?.ToString(),
                    msgType = msgType.ToString(),
                    info,
                });
        }

        private void OnStandState((int Id, StandState State) state)
        {
            var standState = _hwStateService.GetStandState(state.State);
            if (standState != null)
            {
                _ = SendInfoAsync( _mqttSender, MQCommands.StandStateArrived, standState);
            }
        }

        private void OnGeneratorState((int Id, GeneratorState State) state)
        {
            var standState = _hwStateService.GetGeneratorState(state.State);
            if (standState != null)
            {
                _ = SendInfoAsync(_mqttSender, MQCommands.GeneratorStateArrived, standState);
            }
        }

        private void OnCollimatorState((int Id, CollimatorState State) state)
        {
            var standState = _hwStateService.GetCollimatorState(state.State);
            if (standState != null)
            {
                _ = SendInfoAsync(_mqttSender, MQCommands.CollimatorStateArrived, standState);
            }
        }
    }
}
