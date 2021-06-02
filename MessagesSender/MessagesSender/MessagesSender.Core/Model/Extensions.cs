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
        /// <summary>
        /// Checks if message of state types
        /// </summary>
        /// <param name="msg">messgae</param>
        /// <returns>result</returns>
        public static bool IsStateMQMessage(this MQMessages msg) => msg == MQMessages.InstanceOn || msg == MQMessages.InstanceOff;
    }
}
