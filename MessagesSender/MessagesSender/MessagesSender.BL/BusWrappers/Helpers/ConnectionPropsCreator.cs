using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace MessagesSender.BL.BusWrappers.Helpers
{
    /// <summary>
    /// Connection properties creator
    /// </summary>
    internal static class ConnectionPropsCreator
    {
        private const string ConnectionStringValuesSeparator = ";";
        private const string ConnectionStringValueSeparator = "=";
        private const string ConnectionStringServerName = "Server";
        private const string ConnectionStringUserName = "User";
        private const string ConnectionStringPasswordName = "Password";

        /// <summary>
        /// creates connection properties
        /// </summary>
        /// <param name="connectionString">connection string</param>
        /// <returns>connection props</returns>
        public static (string HostName, string UserName, string Password)? Create(string connectionString)
        {
            if (!string.IsNullOrEmpty(connectionString))
            {
                var props = connectionString.Split(new[] { ConnectionStringValuesSeparator }, StringSplitOptions.RemoveEmptyEntries)
                    .Select(s =>
                    {
                        var pair = s.Split(new[] { ConnectionStringValueSeparator }, StringSplitOptions.RemoveEmptyEntries).ToArray();
                        return new { Key = pair.First(), Value = pair.Last() };
                    })
                    .ToDictionary(s => s.Key, s => s.Value);

                return (props[ConnectionStringServerName], props[ConnectionStringUserName], props[ConnectionStringPasswordName]);
            }

            return null;
        }
    }
}
