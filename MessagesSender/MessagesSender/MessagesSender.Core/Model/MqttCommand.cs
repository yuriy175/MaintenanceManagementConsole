using System;
using System.Collections.Generic;
using System.Text;

namespace MessagesSender.Core.Model
{
    /// <summary>
    /// mqtt command class
    /// </summary>
    public class MqttCommand
    {
        /// <summary>
        /// command type separator
        /// </summary>
        public const string TypeSeparator = "?";

        /// <summary>
        /// command parameter separator
        /// </summary>
        public const string ParamSeparator = "&";

        /// <summary>
        /// command parameter value separator
        /// </summary>
        public const string ParamValueSeparator = "=";

        /// <summary>
        /// command type
        /// </summary>
        public string Type { get; set; }

        /// <summary>
        /// command parameters
        /// </summary>
        public Dictionary<string, string> Parameters { get; set; }
    }
}
