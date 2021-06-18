using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Atlas.Common.DAL.Model;
using MessagesSender.DAL.Model;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// Interface to config sqllite database
    /// </summary>
    public interface IConfigEntityService
    {
        /// <summary>
        /// Get offline events
        /// </summary>
        /// <returns>offline events</returns>
        Task<IEnumerable<OfflineEvent>> GetOfflineEventsAsync();

        /// <summary>
        /// adds offline event
        /// </summary>
        /// <param name="offlineEvent">offline event</param>
        /// <returns>true if success</returns>
        Task<bool> AddOfflineEventAsync(OfflineEvent offlineEvent);

        /// <summary>
        /// delete offline events
        /// </summary>
        /// <param name="before">before datetime</param>
        /// <returns>true if success</returns>
        Task<bool> DeleteOfflineEventsAsync(DateTime? before = null);
    }
}
