using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Atlas.Common.DAL.Model;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// Interface to master database
    /// </summary>
    public interface IMasterEntityService
    {
        /// <summary>
        /// Get databases states.
        /// </summary>
        /// <returns>databases states</returns>
        Task<IEnumerable<(string Name, string State)>> GetDatabasesStatesAsync();
    }
}
