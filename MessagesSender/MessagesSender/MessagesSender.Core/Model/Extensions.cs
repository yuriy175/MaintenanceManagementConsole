using System;
using System.Collections.Generic;
using System.Text;

namespace MessagesSender.Core.Model
{
    /// <summary>
    /// Extensions helper
    /// </summary>
    internal static class Extensions
    {
        public static bool IsStateMQMessage(this MQMessages msg) => msg == MQMessages.InstanceOn || msg == MQMessages.InstanceOff;
    }
}
