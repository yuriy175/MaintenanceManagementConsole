using System;
using System.Text;
using System.Threading.Tasks;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using MessagesSender.Core.Model;
using Newtonsoft.Json;
using RabbitMQ.Client;
using Serilog;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// mq sender base interface
    /// </summary>
    public interface IMQSenderBase : IDisposable
    {
        /// <summary>
        /// sends a message
        /// </summary>
        /// <typeparam name="TMsg">message type</typeparam>
        /// <typeparam name="T">entity type</typeparam>
        /// <param name="msgType">message type</param>
        /// <param name="payload">entity</param>
        /// <returns>result</returns>
        Task<bool> SendAsync<TMsg, T>(TMsg msgType, T payload);
    }

    /// <summary>
    /// mqtt sender interface
    /// </summary>
    public interface IMqttSender : IMQSenderBase
    {
        /// <summary>
        /// creates sender
        /// </summary>
        /// <param name="equipInfo">equipment info</param>
        /// <returns>result</returns>
        Task<bool> CreateAsync((string Name, string Number) equipInfo);

        /// <summary>
        /// sends a message to a common mqtt
        /// </summary>
        /// <typeparam name="T">entity type</typeparam>
        /// <param name="msgType">message type</param>
        /// <param name="payload">entity</param>
        /// <returns>result</returns>
        Task<bool> SendCommonAsync<T>(MQMessages msgType, T payload);
    }

    /// <summary>
    /// work queue sender interface
    /// </summary>
    public interface IWorkqueueSender : IMQSenderBase
    {
    }
}
