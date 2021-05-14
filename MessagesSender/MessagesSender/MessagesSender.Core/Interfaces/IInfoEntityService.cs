using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Atlas.Common.DAL.Model;
using MessagesSender.DAL.Model;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// Interface to info database
    /// </summary>
    public interface IInfoEntityService
    {
		/// <summary>
		/// get app parameter by name
		/// </summary>
		/// <returns>all data</returns>
		Task<(IEnumerable<HardDrive> HardDrives, IEnumerable<Lan> Lans)> GetAllDataAsync();
    }
}
