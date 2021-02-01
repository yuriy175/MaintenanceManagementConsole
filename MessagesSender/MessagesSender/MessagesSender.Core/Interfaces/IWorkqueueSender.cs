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
    /// work queue sender interface
    /// </summary>
    public interface IWorkqueueSender : IDisposable
    {
        /// <summary>
        /// sends a message
        /// </summary>
        /// <typeparam name="T">entity type</typeparam>
        /// <param name="payload">entity</param>
        /// <returns>result</returns>
        Task<bool> SendAsync<T>(T payload);
    }
}
