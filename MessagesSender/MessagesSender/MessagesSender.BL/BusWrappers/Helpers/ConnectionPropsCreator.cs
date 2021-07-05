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
        private const string ConnectionStringPortName = "Port";
        private const string ConnectionStringSecuredName = "Secured";

        private const string ConnectionStringSmtpName = "Smtp";
        private const string ConnectionStringEmailFromName = "EmailFrom";
        private const string ConnectionStringEmailToName = "EmailTo";

        /// <summary>
        /// creates mqtt connection properties
        /// </summary>
        /// <param name="connectionString">connection string</param>
        /// <returns>connection props</returns>
        public static (string HostName, int Port, string UserName, string Password, bool Secured)? CreateMqttProps(string connectionString)
        {
            var props = CreateProps(connectionString);
            if (props == null ||
                !int.TryParse(props[ConnectionStringPortName], out int port) ||
                !bool.TryParse(props[ConnectionStringSecuredName], out bool secured))
            {
                return null as (string, int, string, string, bool)?;
            }

            return (
                props[ConnectionStringServerName], 
                port, 
                props[ConnectionStringUserName], 
                props[ConnectionStringPasswordName], 
                secured
                );
        }

        /// <summary>
        /// creates ftp connection properties
        /// </summary>
        /// <param name="connectionString">connection string</param>
        /// <returns>connection props</returns>
        public static (string HostName, string UserName, string Password)? CreateFtpProps(string connectionString)
        {
            var props = CreateProps(connectionString);
            return props == null ?
                null as (string, string, string)? :
                (
                    props[ConnectionStringServerName], 
                    props[ConnectionStringUserName], 
                    props[ConnectionStringPasswordName]
                );
        }

        /// <summary>
        /// creates email connection properties
        /// </summary>
        /// <param name="connectionString">connection string</param>
        /// <returns>connection props</returns>
        public static (string Smtp, string EmailFrom, string EmailTo, string Login, string Password)? CreateEmailProps(string connectionString)
        {
            var props = CreateProps(connectionString);
            return props == null ?
                null as (string, string, string, string, string)? :
                (
                    props[ConnectionStringSmtpName],
                    props[ConnectionStringEmailFromName],
                    props[ConnectionStringEmailToName],
                    props[ConnectionStringUserName],
                    props[ConnectionStringPasswordName]
                );
        }

        private static Dictionary<string, string> CreateProps(string connectionString)
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

                return props;
            }

            return null;
        }
    }
}
