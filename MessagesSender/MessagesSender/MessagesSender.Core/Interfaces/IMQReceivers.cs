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
    /// mq receiver base interface
    /// </summary>
    public interface IMQReceiver : IDisposable
    {
    }

    /// <summary>
    /// mqtt receiver interface
    /// </summary>
    public interface IMqttReceiver : IMQReceiver
    {
        /// <summary>
        /// creates sender
        /// </summary>
        /// <param name="equipInfo">equipment info</param>
        /// <returns>result</returns>
        Task<bool> CreateAsync((string Name, string Number) equipInfo);
    }
}
