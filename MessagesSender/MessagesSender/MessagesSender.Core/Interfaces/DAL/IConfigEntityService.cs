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

        /// <summary>
        /// get config parameters
        /// </summary>
        /// <returns>config parameters</returns>
        Task<IEnumerable<ConfigParam>> GetConfigParamAsync();

        /// <summary>
        /// inserts or updates config parameter
        /// </summary>
        /// <typeparam name="T">value type</typeparam>
        /// <param name="configParamName">config parameter name</param>
        /// <param name="value">value</param>
        /// <returns>new config parameter</returns>
        Task<ConfigParam> UpsertConfigParamAsync<T>(string configParamName, T value);

        /// <summary>
        /// get config parameter by name
        /// </summary>
        /// <param name="configParamName">config parameter name</param>
        /// <returns>config parameter</returns>
        Task<ConfigParam> GetConfigParamAsync(string configParamName);
    }
}
