using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// remote control service interface
    /// </summary>
    public interface IRemoteControlService
    {
        /// <summary>
        /// runs TeamViewer
        /// </summary>
        /// <returns>result</returns>
        Task<bool> RunTeamViewerAsync();

        /// <summary>
        /// runs TaskManager
        /// </summary>
        /// <returns>result</returns>
        Task<bool> RunTaskManagerAsync();

        /// <summary>
        /// sends Atlas logs to email
        /// </summary>
        /// <returns>result</returns>
        Task<bool> SendAtlasLogsAsync();

        /// <summary>
        /// turns on XilibLogs
        /// </summary>
        /// <returns>result</returns>
        Task<bool> XilibLogsOnAsync();
    }
}
