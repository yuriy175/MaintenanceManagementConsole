using System;
using System.Collections.Generic;
using System.Linq;
using MessagesSender.Core.Interfaces;

namespace MessagesSender.BL
{
    /// <summary>
    /// event publisher interface implememntation
    /// </summary>
    public class EventPublisher : IEventPublisher
    {
        private readonly Dictionary<EventTypes, List<Delegate>> _eventHandlers =
            Array.ConvertAll((int[])Enum.GetValues(typeof(EventTypes)), Convert.ToInt32)
            .Select(i => (EventTypes)i)
            .ToDictionary(t => t, t => new List<Delegate> { });
        
        private enum EventTypes
        { 
            CommandArrived, 
            Activate, 
            Deactivate, 
            RunTV, 
            RunTaskManager, 
            SendAtlasLogs, 
            XilibLogsOn, 
            Reconnect,
            GetHospitalInfo,
        }

        /// <summary>
        /// register any command handler
        /// </summary>
        /// <param name="handler">command handler</param>        
        public void RegisterMqttCommandArrivedEvent(Action<string> handler)
        {
            _eventHandlers[EventTypes.CommandArrived].Add(handler);
        }

        /// <summary>
        /// new command arrived
        /// </summary>
        /// <param name="command">mqtt command</param>
        public void MqttCommandArrived(string command)
        {
            _eventHandlers[EventTypes.CommandArrived].ForEach(h => (h as Action<string>)(command));
        }

        /// <summary>
        /// Activate command arrived
        /// </summary>
        public void ActivateCommandArrived()
        {
            _eventHandlers[EventTypes.Activate].ForEach(h => (h as Action)());
        }

        /// <summary>
        /// register Activate command handler
        /// </summary>
        /// <param name="handler">command handler</param>    
        public void RegisterActivateCommandArrivedEvent(Action handler)
        {
            _eventHandlers[EventTypes.Activate].Add(handler);
        }

        /// <summary>
        /// Deactivate command arrived
        /// </summary>
        public void DeactivateCommandArrived()
        {
            _eventHandlers[EventTypes.Deactivate].ForEach(h => (h as Action)());
        }

        /// <summary>
        /// register Deactivate command handler
        /// </summary>
        /// <param name="handler">command handler</param>   
        public void RegisterDeactivateCommandArrivedEvent(Action handler)
        {
            _eventHandlers[EventTypes.Deactivate].Add(handler);
        }

        /// <summary>
        /// RunTV command arrived
        /// </summary>
        public void RunTVCommandArrived()
        {
            _eventHandlers[EventTypes.RunTV].ForEach(h => (h as Action)());
        }

        /// <summary>
        /// register RunTV command handler
        /// </summary>
        /// <param name="handler">command handler</param>    
        public void RegisterRunTVCommandArrivedEvent(Action handler)
        {
            _eventHandlers[EventTypes.RunTV].Add(handler);
        }

        /// <summary>
        /// Reconnect command arrived
        /// </summary>
        public void ReconnectCommandArrived()
        {
            _eventHandlers[EventTypes.Reconnect].ForEach(h => (h as Action)());
        }

        /// <summary>
        /// register Reconnect command handler
        /// </summary>
        /// <param name="handler">command handler</param>    
        public void RegisterReconnectCommandArrivedEvent(Action handler)
        {
            _eventHandlers[EventTypes.Reconnect].Add(handler);
        }

        /// <summary>
        /// SendAtlasLogs command arrived
        /// </summary>
        public void SendAtlasLogsCommandArrived()
        {
            _eventHandlers[EventTypes.SendAtlasLogs].ForEach(h => (h as Action)());
        }

        /// <summary>
        /// register SendAtlasLogs command handler
        /// </summary>
        /// <param name="handler">command handler</param>    
        public void RegisterSendAtlasLogsCommandArrivedEvent(Action handler)
        {
            _eventHandlers[EventTypes.SendAtlasLogs].Add(handler);
        }

        /// <summary>
        /// RunTaskMan command arrived
        /// </summary>
        public void RunTaskManCommandArrived()
        {
            _eventHandlers[EventTypes.RunTaskManager].ForEach(h => (h as Action)());
        }

        /// <summary>
        /// register RunTaskMan command handler
        /// </summary>
        /// <param name="handler">command handler</param>    
        public void RegisterRunTaskManCommandEvent(Action handler)
        {
            _eventHandlers[EventTypes.RunTaskManager].Add(handler);
        }

        /// <summary>
        /// XilibLogsOn command arrived
        /// </summary>
        public void XilibLogsOnCommandArrived()
        {
            _eventHandlers[EventTypes.XilibLogsOn].ForEach(h => (h as Action)());
        }

        /// <summary>
        /// register XilibLogsOn command handler
        /// </summary>
        /// <param name="handler">command handler</param>    
        public void RegisterXilibLogsOnCommandArrivedEvent(Action handler)
        {
            _eventHandlers[EventTypes.XilibLogsOn].Add(handler);
        }

        /// <summary>
        /// GetHospitalInfo command arrived
        /// </summary>
        public void GetHospitalInfoCommandArrived()
        {
            _eventHandlers[EventTypes.GetHospitalInfo].ForEach(h => (h as Action)());
        }

        /// <summary>
        /// register GetHospitalInfo command handler
        /// </summary>
        /// <param name="handler">command handler</param>    
        public void RegisterGetHospitalInfoCommandArrivedEvent(Action handler)
        {
            _eventHandlers[EventTypes.GetHospitalInfo].Add(handler);
        }
    }
}
