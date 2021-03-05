using System;
using System.Collections.Generic;
using System.Linq;
using MessagesSender.Core.Interfaces;

namespace MessagesSender.BL
{
    public class EventPublisher : IEventPublisher
    {
        enum EventTypes { CommandArrived, Activate, Deactivate, RunTV }

        private readonly Dictionary<EventTypes, List<Delegate>> _eventHandlers =
            Array.ConvertAll((int[])Enum.GetValues(typeof(EventTypes)), Convert.ToInt32)
            .Select(i => (EventTypes)i)
            .ToDictionary(t => t, t => new List<Delegate> { });
		
        public void RegisterMqttCommandArrivedEvent(Action<string> handler)
        {
            _eventHandlers[EventTypes.CommandArrived].Add(handler);
        }

		public void MqttCommandArrived(string command)
        {
            _eventHandlers[EventTypes.CommandArrived].ForEach(h => (h as Action<string>)(command));
        }

        public void ActivateCommandArrived()
        {
            _eventHandlers[EventTypes.Activate].ForEach(h => (h as Action)());
        }

        public void RegisterActivateCommandArrivedEvent(Action handler)
        {
            _eventHandlers[EventTypes.Activate].Add(handler);
        }

        public void DeactivateCommandArrived()
        {
            _eventHandlers[EventTypes.Deactivate].ForEach(h => (h as Action)());
        }
        public void RegisterDeactivateCommandArrivedEvent(Action handler)
        {
            _eventHandlers[EventTypes.Deactivate].Add(handler);
        }

        public void RunTVCommandArrived()
        {
            _eventHandlers[EventTypes.RunTV].ForEach(h => (h as Action)());
        }

        public void RegisterRunTVCommandArrivedEvent(Action handler)
        {
            _eventHandlers[EventTypes.RunTV].Add(handler);
        }
    }
}
