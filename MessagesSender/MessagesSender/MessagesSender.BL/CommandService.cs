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
    /// <summary>
    /// command service
    /// </summary>
    public class CommandService : ICommandService
    {
        private const string ActivateCommandName = "activate";
        private const string DeactivateCommandName = "deactivate";
        private const string RunTeamViewerCommandName = "runTV";
        private const string RunTaskManCommandName = "runTaskMan";
        private const string SendAtlasLogsCommandName = "sendAtlasLogs";
        private const string XilibLogsOnCommandName = "xilibLogsOn";
        private const string EquipLogsOnCommandName = "equipLogsOn";
        private const string ReconnectCommandName = "reconnect";
        private const string EquipInfoCommandName = "equipInfo";
        private const string ServerReadyCommandName = "serverReady";        
        private const string UpdateDBInfoCommandName = "updateDBInfo";
        private const string RecreateDBInfoCommandName = "recreateDBInfo";

        private readonly ILogger _logger;
        private readonly IEventPublisher _eventPublisher;

        private readonly Dictionary<string, Action<MqttCommand>> _commandMap = 
            new Dictionary<string, Action<MqttCommand>>
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

            _commandMap = new Dictionary<string, Action<MqttCommand>>
            {
                { ActivateCommandName, (command) => OnActivateCommand() },
                { DeactivateCommandName, (command) => OnDeactivateCommand() },
                { RunTeamViewerCommandName, (command) => OnRunTVCommandAsync() },
                { RunTaskManCommandName, (command) => OnRunTaskManCommandAsync() },
                { SendAtlasLogsCommandName, (command) => OnSendAtlasLogsCommandAsync() },
                { XilibLogsOnCommandName, (command) => OnXilibLogsOnCommandAsync() },
                { ReconnectCommandName, (command) => OnReconnectCommand() },
                { EquipInfoCommandName, (command) => OnEquipInfoCommand() },
                { ServerReadyCommandName, (command) => OnServerReadyCommand() },
                { UpdateDBInfoCommandName, (command) => OnUpdateDBInfoCommand() },
                { RecreateDBInfoCommandName, (command) => OnRecreateDBInfoCommand() },
                { EquipLogsOnCommandName, (command) => OnEquipLogsOnCommandAsync(command.Parameters) },                
            };

            _eventPublisher.RegisterMqttCommandArrivedEvent(command => OnCommandArrivedAsync(command));

            _logger.Information("CommandService started");
        }

        private Task<bool> OnCommandArrivedAsync(MqttCommand command)
        {
            try
            {
                _commandMap[command.Type](command);
                return Task.FromResult(true);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, $"command error {command.Type}");
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

        private void OnEquipLogsOnCommandAsync(Dictionary<string, string> parameters)
        {
            _eventPublisher.EquipLogsOnCommandArrived(parameters);
        }

        private void OnEquipInfoCommand()
        {
            _eventPublisher.GetHospitalInfoCommandArrived();
        }

        private void OnServerReadyCommand()
        {
            _eventPublisher.ServerReadyCommandArrived();
        }        

        private void OnUpdateDBInfoCommand()
        {
            _eventPublisher.UpdateDBInfoCommandArrived();
        }

        private void OnRecreateDBInfoCommand()
        {
            _eventPublisher.RecreateDBInfoCommandArrived();
        }        
    }
}
