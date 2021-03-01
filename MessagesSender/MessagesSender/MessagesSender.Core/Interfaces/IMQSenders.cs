using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using RabbitMQ.Client;
using Serilog;
using System;
using System.Threading.Tasks;
using Newtonsoft.Json;
using System.Text;

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
		/// <param name="payload">entity</param>
		/// <returns>result</returns>
		Task<bool> SendAsync<TMsg, T>(TMsg msgType, T payload);
    }

    /// <summary>
    /// work queue sender interface
    /// </summary>
    public interface IWorkqueueSender : IMQSenderBase
    {
    }

    /// <summary>
    /// mqtt sender interface
    /// </summary>
    public interface IMqttSender : IMQSenderBase
    {
        Action<string> OnCommandArrived { get; set; }

        /// <summary>
        /// creates sender
        /// </summary>
        /// <param name="equipInfo">equipment info</param>
        /// <returns>result</returns>
        Task<bool> CreateAsync((string Name, string Number) equipInfo);
    }
}
