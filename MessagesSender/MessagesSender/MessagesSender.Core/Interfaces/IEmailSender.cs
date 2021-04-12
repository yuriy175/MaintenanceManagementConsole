using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MessagesSender.Core.Interfaces
{
	/// <summary>
	/// email service interface
	/// </summary>
	public interface IEmailSender
	{
		/// <summary>
		/// sends teamviewer file
		/// </summary>
		/// <param name="tvPath">file path</param>
		/// <returns>result</returns>
		Task<bool> SendTeamViewerAsync(string tvPath);
	}
}
