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
        private readonly ILogger _logger;

        private ConnectionStates _standState = ConnectionStates.Disconnected;
        private ConnectionStates _generatorState = ConnectionStates.Disconnected;
        private ConnectionStates _collimatorState = ConnectionStates.Disconnected;
        private ConnectionStates _detectorState = ConnectionStates.Disconnected;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        public HardwareStateService(
            ILogger logger)
        {
            _logger = logger;
            _logger.Information("HardwareStateService started");
        }

        /// <summary>
        /// gets stand state for sending to mqtt
        /// </summary>
        /// <param name="state">state</param>
        /// <returns>state object</returns>
        public object GetStandState(StandState state) =>
            GetHardwareState(state?.State, StandConnectedValue, ref _standState);

        /// <summary>
        /// gets generator state for sending to mqtt
        /// </summary>
        /// <param name="state">state</param>
        /// <returns>state object</returns>
        public object GetGeneratorState(GeneratorState state) =>
            GetHardwareState(state?.State, GeneratorConnectedValue, ref _generatorState);

        /// <summary>
        /// gets collimator state for sending to mqtt
        /// </summary>
        /// <param name="state">state</param>
        /// <returns>state object</returns>
        public object GetCollimatorState(CollimatorState state) =>
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
    }
}
