using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// ftp client interface
    /// </summary>
    public interface IFtpClient
	{
		/// <summary>
		/// send file content to ftp
		/// </summary>
		/// <param name="filePath">file path</param>
		/// <param name="destFileName">destination file name</param>
		/// <returns>result</returns>
		Task<bool> SendAsync(string filePath, string destination = null);
	}
}
