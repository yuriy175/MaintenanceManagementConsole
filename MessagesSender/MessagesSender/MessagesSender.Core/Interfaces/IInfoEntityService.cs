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
		/// <param name="appParam">app parameter name</param>
		/// <returns>app parameter</returns>
		Task<AppParam> GetAppParamAsync(string appParam);
    }
}
