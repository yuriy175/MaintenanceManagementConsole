using System;
using System.Collections.Generic;
using System.Linq;
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

        /// <summary>
        /// parses string command to object command
        /// </summary>
        /// <param name="command">string command</param>
        /// <returns>command</returns>
        public static MqttCommand Parse(this string command)
        {
            if (string.IsNullOrEmpty(command))
            {
                return new MqttCommand { Type = string.Empty };
            }

            var typeParts = command.Split(MqttCommand.TypeSeparator);
            var @params = typeParts.Length < 2 ? null :
                typeParts[1].Split(MqttCommand.ParamSeparator)
                .Select(p => p.Split(MqttCommand.ParamValueSeparator))
                .ToDictionary(p => p.FirstOrDefault(), p => p.LastOrDefault());

            return new MqttCommand 
            { 
                Type = typeParts?.FirstOrDefault() ?? string.Empty, 
                Parameters = @params,
            };
        }
    }
}
