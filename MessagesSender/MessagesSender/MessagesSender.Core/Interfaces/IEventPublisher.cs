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

        void SendAtlasLogsCommandArrived();
        void RegisterSendAtlasLogsCommandArrivedEvent(Action handler);

        void RunTaskManCommandArrived();
        void RegisterRunTaskManCommandEvent(Action handler);

        void XilibLogsOnCommandArrived();
        void RegisterXilibLogsOnCommandArrivedEvent(Action handler);

        void ReconnectCommandArrived();
        void RegisterReconnectCommandArrivedEvent(Action handler);

        void GetHospitalInfoCommandArrived();
        void RegisterGetHospitalInfoCommandArrivedEvent(Action handler);
    }
}
