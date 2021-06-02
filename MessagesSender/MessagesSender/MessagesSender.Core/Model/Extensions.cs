using System;
using System.Collections.Generic;
using System.Text;

namespace MessagesSender.Core.Model
{
    static class Extensions
    {
        public static bool IsStateMQMessage(this MQMessages msg) => msg == MQMessages.InstanceOn || msg == MQMessages.InstanceOff;
    }
}
