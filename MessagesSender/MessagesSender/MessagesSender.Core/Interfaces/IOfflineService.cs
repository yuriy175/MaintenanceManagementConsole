using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// offline service interface
    /// </summary>
    public interface IOfflineService
    {
        /// <summary>
        /// Check info for offline persistence
        /// </summary>
        /// <typeparam name="TMsgType">message type</typeparam>
        /// <typeparam name="T">info type</typeparam>
        /// <param name="msgType">messge type</param>
        /// <param name="info">info</param>
        /// <returns>result</returns>
        Task<bool> CheckInfoAsync<TMsgType, T>(TMsgType msgType, T info);

        /// <summary>
        /// Get offlined infos
        /// </summary>
        /// <returns>offlined infos</returns>
        Task<IEnumerable<(string MsgType, object Msg)>> GetInfosAsync();

        /// <summary>
        /// Clears offline events
        /// </summary>
        /// <returns>result</returns>
        Task<bool> ClearInfosAsync();
    }
}
