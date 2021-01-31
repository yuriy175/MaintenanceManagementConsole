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

namespace MessagesSender.BL
{
    public class Service : IService
    {
        private readonly ISettingsEntityService _dbSettingsEntityService;
        private readonly ILogger _logger;
        private readonly IMQCommunicationService _mqService;

        private IPAddress _ipAddress = null;
        private (string Name, string Number) _equipmentInfo = (null, null);

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="dbSettingsEntityService">settings database connector</param>
        /// <param name="logger">logger</param>
        /// <param name="mqService">MQ service</param>
        public Service(
            ISettingsEntityService dbSettingsEntityService,
            ILogger logger,
            IMQCommunicationService mqService)
        {
            _dbSettingsEntityService = dbSettingsEntityService;
            _logger = logger;
            _mqService = mqService;

            new Action[]
                {
                    () => _ = SubscribeMQRecevers(),
                    () => _ = GetEquipmentInfoAsync(),
                    () => _ = GetEquipmentIPAsync(),
                }.RunTasksAsync();

            _logger.Information("Main service started");
        }

        private async Task SubscribeMQRecevers()
        {
            /*_mqService.Subscribe<MQCommands, bool>(
                (MQCommands.StopShooting, async data => await OnStopShootingAsync(data)));

            _mqService.Subscribe<MQCommands, HardwarePedals>(
               (MQCommands.PedalPressed, async data => await OnPedalPressedAsync(data)));

            _mqService.Subscribe<MQCommands, (int Id, GeneratorState State)>(
                (MQCommands.GeneratorStateArrived, state => OnGeneratorState(state)));

            _mqService.Subscribe<MQCommands, (int Id, StandState State)>(
                (MQCommands.StandStateArrived, state => OnStandState(state)));*/
        }

        private async Task GetEquipmentInfoAsync()
        {
            _equipmentInfo = await _dbSettingsEntityService.GetEquipmentInfo();
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
    }
}
