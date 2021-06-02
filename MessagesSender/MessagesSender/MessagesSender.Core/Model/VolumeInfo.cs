using System;
using System.Collections.Generic;
using System.Text;

namespace MessagesSender.Core.Model
{
    /// <summary>
    /// Volume info
    /// </summary>
    public class VolumeInfo
    {
        /// <summary>
        /// volume name 
        /// </summary>
        public string Letter { get; set; }

        /// <summary>
        /// volume free size
        /// </summary>
        public long FreeSize { get; set; }

        /// <summary>
        /// volume total size
        /// </summary>
        public long TotalSize { get; set; }
    }
}
