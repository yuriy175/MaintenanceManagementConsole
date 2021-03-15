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
        Task<(string Name, string Number)> GetEquipmentInfoAsync();

        /// <summary>
        /// Get dicom info.
        /// </summary>
        /// <returns>dicom info</returns>
        Task<IEnumerable<(int Id, string Name, string IP, int ServiceRole)>> GetDicomServicesAsync();
    }
}
