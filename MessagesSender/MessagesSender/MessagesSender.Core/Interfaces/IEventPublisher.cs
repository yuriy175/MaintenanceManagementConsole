using System;

namespace MessagesSender.Core.Interfaces
{
    public interface IEventPublisher
    {
        void MqttCommandArrived(string command);
        void RegisterMqttCommandArrivedEvent(Action<string> handler);

        void ActivateCommandArrived();
        void RegisterActivateCommandArrivedEvent(Action handler);

        void DeactivateCommandArrived();
        void RegisterDeactivateCommandArrivedEvent(Action handler);

        void RunTVCommandArrived();
        void RegisterRunTVCommandArrivedEvent(Action handler);

        void ReconnectCommandArrived();
        void RegisterReconnectCommandArrivedEvent(Action handler);
    }
}
