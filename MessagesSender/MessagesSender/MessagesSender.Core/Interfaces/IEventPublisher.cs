using System;

namespace MessagesSender.Core.Interfaces
{
    public interface IEventPublisher
    {
        void MqttCommandArrived(string command);
        void RegisterMqttCommandArrivedEvent(Action<string> handler);
    }
}
