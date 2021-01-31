using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Atlas.Common.DAL.Model;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// Interface to Settings database
    /// </summary>
    public interface ISettingsEntityService
    {
        /// <summary>
        /// Get equipment info.
        /// </summary>
        /// <returns>equipment info</returns>
        Task<(string Name, string Number)> GetEquipmentInfo();
    }
}
