using System;
using System.Collections.Generic;
using System.Text;

namespace MessagesSender.Core.Model
{
    /// <summary>
    /// MQ message types enumeration
    /// </summary>
    public enum MQMessages
    {
        /// <summary>
        /// InstanceOn message type
        /// </summary>
        InstanceOn = 5000,

        /// <summary>
        /// InstanceOff message type
        /// </summary>
        InstanceOff,

        /// <summary>
        /// HddDrivesInfo message type
        /// </summary>
        HddDrivesInfo,

        /// <summary>
        /// CPUInfo message type
        /// </summary>
        CPUInfo,

        /// <summary>
        /// MemoryInfo message type
        /// </summary>
        MemoryInfo,

        /// <summary>
        /// DicomInfo message type
        /// </summary>
        DicomInfo,

        /// <summary>
        /// SoftwareInfo message type
        /// </summary>
        SoftwareInfo,

        /// <summary>
        /// SoftwareMsgInfo message type
        /// </summary>
        SoftwareMsgInfo,

        /// <summary>
        /// RemoteAccess message type
        /// </summary>
        RemoteAccess,

        /// <summary>
        /// ImagesInfo message type
        /// </summary>
        ImagesInfo,

        /// <summary>
        /// HospitalInfo message type
        /// </summary>
        HospitalInfo,

        /// <summary>
        /// AllDBInfo message type
        /// </summary>
        AllDBInfo,

        /// <summary>
        /// Events message type
        /// </summary>
        Events,

        /// <summary>
        /// InstanceOn from offline message type
        /// </summary>
        InstanceOnOffline,

        /// <summary>
        /// AtlasExited message type
        /// </summary>
        AtlasExited,

        /// <summary>
        /// KeepAlive message type
        /// </summary>
        KeepAlive,

        /// <summary>
        /// Chat message type
        /// </summary>
        Chat,
    }
}
