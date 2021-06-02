using System;
using System.Collections.Generic;
using System.IO;
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
    public class CommandService : ICommandService
    {
        private const string ActivateCommandName = "activate";
        private const string DeactivateCommandName = "deactivate";
        private const string RunTeamViewerCommandName = "runTV";
        private const string RunTaskManCommandName = "runTaskMan";
        private const string SendAtlasLogsCommandName = "sendAtlasLogs";
        private const string XilibLogsOnCommandName = "xilibLogsOn";
        private const string ReconnectCommandName = "reconnect";
        private const string EquipInfoCommandName = "equipInfo";

        private readonly ILogger _logger;
        private readonly IEventPublisher _eventPublisher;

        private readonly Dictionary<string, Action> _commandMap = new Dictionary<string, Action>
        {
        };

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        public CommandService(
            ILogger logger,
            IEventPublisher eventPublisher)
        {
            _logger = logger;
            _eventPublisher = eventPublisher;

            _commandMap = new Dictionary<string, Action>
            {
                { ActivateCommandName, () => OnActivateCommand() },
                { DeactivateCommandName, () => OnDeactivateCommand() },
                { RunTeamViewerCommandName, () => OnRunTVCommandAsync() },
                { RunTaskManCommandName, () => OnRunTaskManCommandAsync() },
                { SendAtlasLogsCommandName, () => OnSendAtlasLogsCommandAsync() },
                { XilibLogsOnCommandName, () => OnXilibLogsOnCommandAsync() },
                { ReconnectCommandName, () => OnReconnectCommand() },
                { EquipInfoCommandName, () => OnEquipInfoCommand() },
            };

            _eventPublisher.RegisterMqttCommandArrivedEvent(command => OnCommandArrivedAsync(command));

            _logger.Information("CommandService started");
        }

        private Task<bool> OnCommandArrivedAsync(string command)
        {
            try
            {
                _commandMap[command]();
                return Task.FromResult(true);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, $"command error {command}");
            }

            return Task.FromResult(false);
        }

        /// <summary>
        /// activate command handler
        /// </summary>
        private void OnActivateCommand()
        {
            _eventPublisher.ActivateCommandArrived();
        }

        private void OnDeactivateCommand()
        {
            _eventPublisher.DeactivateCommandArrived();
        }

        private void OnRunTVCommandAsync()
        {
            _eventPublisher.RunTVCommandArrived();
        }

        private void OnReconnectCommand()
        {
            _eventPublisher.ReconnectCommandArrived();
        }

        private void OnRunTaskManCommandAsync()
        {
            _eventPublisher.RunTaskManCommandArrived();
        }

        private void OnSendAtlasLogsCommandAsync()
        {
            _eventPublisher.SendAtlasLogsCommandArrived();
        }

        private void OnXilibLogsOnCommandAsync()
        {
            _eventPublisher.XilibLogsOnCommandArrived();
        }

        private void OnEquipInfoCommand()
        {
            _eventPublisher.GetHospitalInfoCommandArrived();
        }        
    }
}
