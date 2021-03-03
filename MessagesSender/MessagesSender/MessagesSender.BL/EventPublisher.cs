using System;
using System.Collections.Generic;
using System.Linq;
using MessagesSender.Core.Interfaces;

namespace MessagesSender.BL
{
    public class EventPublisher : IEventPublisher
    {
        enum EventTypes { CommandArrived }

        private readonly Dictionary<EventTypes, List<Delegate>> _imageHandlers =
            Array.ConvertAll((int[])Enum.GetValues(typeof(EventTypes)), Convert.ToInt32)
            .Select(i => (EventTypes)i)
            .ToDictionary(t => t, t => new List<Delegate> { });
		
        public void RegisterMqttCommandArrivedEvent(Action<string> handler)
        {
            _imageHandlers[EventTypes.CommandArrived].Add(handler);
        }

		public void MqttCommandArrived(string command)
        {
            _imageHandlers[EventTypes.CommandArrived].ForEach(h => (h as Action<string>)(command));
        }
    }
}
