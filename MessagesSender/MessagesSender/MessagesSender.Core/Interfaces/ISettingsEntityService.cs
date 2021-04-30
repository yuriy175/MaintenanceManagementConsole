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

		/// <summary>
		/// inserts or updates app parameter
		/// </summary>
		/// <param name="appParam">app parameter name</param>
		/// <param name="value">value</param>
		/// <returns>new app parameter</returns>
		Task<AppParams> UpsertAppParamAsync<T>(string appParam, T value);

		/// <summary>
		/// get app parameter by name
		/// </summary>
		/// <param name="appParam">app parameter name</param>
		/// <returns>app parameter</returns>
		Task<AppParams> GetAppParamAsync(string appParam);

        /// <summary>
        /// Get hospital info info.
        /// </summary>
        /// <returns>equipment info</returns>
        Task<(string Name, string Address, double Latitude, double Longitude)?> GetHospitalInfoAsync();
    }
}
