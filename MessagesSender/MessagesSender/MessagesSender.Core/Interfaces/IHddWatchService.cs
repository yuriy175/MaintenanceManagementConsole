using Atlas.Acquisitions.Common.Core.Model;
using MessagesSender.Core.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// hdd watch service interface
    /// </summary>
    public interface IHddWatchService
    {
        /// <summary>
        /// gets hdd drives info
        /// </summary>
        /// <returns>drives info</returns>
        Task<IEnumerable<VolumeInfo>> GetDriveInfosAsync();
    }
}
