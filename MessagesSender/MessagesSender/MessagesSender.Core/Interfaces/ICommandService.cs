using Atlas.Acquisitions.Common.Core.Model;
using MessagesSender.Core.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// command service interface
    /// </summary>
    public interface ICommandService
    {
        /// <summary>
        /// command handler
        /// </summary>
        /// <param name="command">command</param>
        /// <returns>result</returns>
        Task<(string MsgType, object Info)?> OnCommandArrivedAsync(string command);
    }
}
