using MessagesSender.Core.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// sending service interface
    /// </summary>
    public interface ISendingService
    {
        /// <summary>
        /// sends info to workqueue
        /// </summary>
        /// <typeparam name="TMsgType">message type</typeparam>
        /// <typeparam name="T">info type</typeparam>
        /// <param name="msgType">info type</param>
        /// <param name="info">info</param>
        /// <returns>result</returns>
        Task<bool> SendInfoToWorkQueueAsync<TMsgType, T>(TMsgType msgType, T info);

        /// <summary>
        /// sends info to mqtt
        /// </summary>
        /// <typeparam name="TMsgType">message type</typeparam>
        /// <typeparam name="T">info type</typeparam>
        /// <param name="msgType">info type</param>
        /// <param name="info">info</param>
        /// <returns>result</returns>
        Task<bool> SendInfoToMqttAsync<TMsgType, T>(TMsgType msgType, T info);

		/// <summary>
		/// sends info to common mqtt
		/// </summary>
		/// <typeparam name="T">info type</typeparam>
		/// <param name="msgType">info type</param>
		/// <param name="info">info</param>
		/// <returns>result</returns>
		Task<bool> SendInfoToCommonMqttAsync<T>(MQMessages msgType, T info);

		/// <summary>
		/// creates service
		/// </summary>
		/// <returns>result</returns>
		Task<bool> CreateAsync();
    }
}
