using System;
using System.Collections.Generic;
using System.Text;

namespace MessagesSender.Core.Model
{
    public enum MQMessages
    {
        InstanceOn = 5000,
        InstanceOff,
        HddDrivesInfo,
        CPUInfo,
        MemoryInfo,
    }
}
