using System;
using System.Collections.Generic;
using MessagesSender.Core.Model;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// event publisher interface
    /// </summary>
    public interface IEventPublisher
    {
        /// <summary>
        /// new command arrived
        /// </summary>
        /// <param name="command">mqtt command</param>
        void MqttCommandArrived(MqttCommand command);

        /// <summary>
        /// register any command handler
        /// </summary>
        /// <param name="handler">command handler</param>
        void RegisterMqttCommandArrivedEvent(Action<MqttCommand> handler);

        /// <summary>
        /// Activate command arrived
        /// </summary>
        void ActivateCommandArrived();

        /// <summary>
        /// register Activate command handler
        /// </summary>
        /// <param name="handler">command handler</param>
        void RegisterActivateCommandArrivedEvent(Action handler);

        /// <summary>
        /// Deactivate command arrived
        /// </summary>
        void DeactivateCommandArrived();

        /// <summary>
        /// register Deactivate command handler
        /// </summary>
        /// <param name="handler">command handler</param>
        void RegisterDeactivateCommandArrivedEvent(Action handler);

        /// <summary>
        /// RunTV command arrived
        /// </summary>
        void RunTVCommandArrived();

        /// <summary>
        /// register RunTV command handler
        /// </summary>
        /// <param name="handler">command handler</param>
        void RegisterRunTVCommandArrivedEvent(Action handler);

        /// <summary>
        /// SendAtlasLogs command arrived
        /// </summary>
        void SendAtlasLogsCommandArrived();

        /// <summary>
        /// register SendAtlasLogs command handler
        /// </summary>
        /// <param name="handler">command handler</param>
        void RegisterSendAtlasLogsCommandArrivedEvent(Action handler);

        /// <summary>
        /// RunTaskMan command arrived
        /// </summary>
        void RunTaskManCommandArrived();

        /// <summary>
        /// register RunTaskMan command handler
        /// </summary>
        /// <param name="handler">command handler</param>
        void RegisterRunTaskManCommandEvent(Action handler);

        /// <summary>
        /// XilibLogsOn command arrived
        /// </summary>
        void XilibLogsOnCommandArrived();

        /// <summary>
        /// register XilibLogsOn command handler
        /// </summary>
        /// <param name="handler">command handler</param>
        void RegisterXilibLogsOnCommandArrivedEvent(Action handler);

        /// <summary>
        /// EquipLogsOn command arrived
        /// </summary>
        /// <param name="parameters">command parameters</param>
        void EquipLogsOnCommandArrived(Dictionary<string, string> parameters);

        /// <summary>
        /// register EquipLogsOn command handler
        /// </summary>
        /// <param name="handler">command handler</param>
        void RegisterEquipLogsOnCommandArrivedEvent(Action<Dictionary<string, string>> handler);

        /// <summary>
        /// Reconnect command arrived
        /// </summary>
        void ReconnectCommandArrived();

        /// <summary>
        /// register Reconnect command handler
        /// </summary>
        /// <param name="handler">command handler</param>
        void RegisterReconnectCommandArrivedEvent(Action handler);

        /// <summary>
        /// GetHospitalInfo command arrived
        /// </summary>
        void GetHospitalInfoCommandArrived();

        /// <summary>
        /// register GetHospitalInfo command handler
        /// </summary>
        /// <param name="handler">command handler</param>
        void RegisterGetHospitalInfoCommandArrivedEvent(Action handler);

        /// <summary>
        /// UpdateDBInfo command arrived
        /// </summary>
        void UpdateDBInfoCommandArrived();

        /// <summary>
        /// register UpdateDBInfo command handler
        /// </summary>
        /// <param name="handler">command handler</param>    
        void RegisterUpdateDBInfoCommandArrivedEvent(Action handler);

        /// <summary>
        /// RecreateDBInfo command arrived
        /// </summary>
        void RecreateDBInfoCommandArrived();

        /// <summary>
        /// register RecreateDBInfo command handler
        /// </summary>
        /// <param name="handler">command handler</param>    
        void RegisterRecreateDBInfoCommandArrivedEvent(Action handler);

        /// <summary>
        /// ServerReady command arrived
        /// </summary>
        void ServerReadyCommandArrived();

        /// <summary>
        /// register ServerReady command handler
        /// </summary>
        /// <param name="handler">command handler</param>    
        void RegisterServerReadyCommandArrivedEvent(Action handler);        
    }
}
