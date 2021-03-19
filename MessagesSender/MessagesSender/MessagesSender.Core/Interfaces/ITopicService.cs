using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// topic service interface
    /// </summary>
    public interface ITopicService
    {
        /// <summary>
        /// gets main topic
        /// </summary>
        /// <returns>result</returns>
        Task<string> GetTopicAsync();
    }
}
