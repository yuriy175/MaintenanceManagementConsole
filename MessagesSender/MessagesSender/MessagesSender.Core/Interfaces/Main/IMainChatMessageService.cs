using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// main chat message service interface
    /// </summary>
    public interface IMainChatMessageService
    {
        /// <summary>
        /// Sends a chat message
        /// </summary>
        /// <param name="message">chat message</param>
        /// <returns>result</returns>
        Task<bool> SendChatMessageAsync(string message);
    }
}
