using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// zip service interface
    /// </summary>
    public interface IZipService
	{
		/// <summary>
		/// zips folder
		/// </summary>
		/// <param name="folder">folder to zip</param>
		/// <returns>zip file path</returns>
		Task<string> ZipFolderAsync(string folder);

		/// <summary>
		/// zips file
		/// </summary>
		/// <param name="filePath">file to zip</param>
		/// <returns>zipped file path</returns>
		Task<string> ZipFileAsync(string filePath);
	}
}
