using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Atlas.Acquisitions.Common.Core.Model;
using MessagesSender.Core.Model;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// db raw data service interface
    /// </summary>
    public interface IDBDataService
    {
        /// <summary>
        /// Updates db info
        /// </summary>
        /// <param name="recreate">if to update or recreate db</param>
        /// <returns>result</returns>
        Task<bool> UpdateDBInfoAsync(bool recreate = false);
    }
}
