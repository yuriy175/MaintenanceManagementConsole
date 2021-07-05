using System;
using System.Collections.Generic;
using System.IO;
using System.Text;
using Atlas.Common.Core.Interfaces;
using Atlas.Common.DAL.Impls;
using Microsoft.Extensions.Configuration;
using Serilog;
using Serilog.Events;

namespace MessagesSender.BL
{
    /// <summary>
    /// logger builder interface
    /// </summary>
    public class LoggerBuilder : ILoggerBuilder
    {
        private const string DefaultLogFile = @"Logs\msgSender.log";
        private const long LogSizeLimitBytes = 50000000;

        /// <summary>
        /// Initializes a new instance of the <see cref="LoggerBuilder"/> class.
        /// </summary>
        /// <param name="logger">logger</param>
        public LoggerBuilder()
        {
            /*var config = new ConfigurationBuilder()
                .SetBasePath(Directory.GetCurrentDirectory())
                .AddJsonFile(Path.Combine(
                        Path.GetDirectoryName(
                            typeof(ILoggerBuilder).Assembly.Location), "dbsettings.json"))
                .Build();

            _logEntityService = new LogEntityService(config.GetConnectionString("SettingsConnection"));
            */
        }

        /// <summary>
        /// builds logger
        /// </summary>
        /// <param name="name">log name</param>
        /// <returns>logger</returns>
        public ILogger Build(string name)
        {
            /*var installPath = _logEntityService.GetInstallPath();
            var logParam = _logEntityService.GetLogParamByName(name);
            if (logParam == null)
                return CreateDefaultLogger();
                */

            var logger = new LoggerConfiguration()
                .WriteTo.File(
                    DefaultLogFile, // Path.Combine(installPath, logParam.FilePath),
                    restrictedToMinimumLevel: LogEventLevel.Information,
                    rollingInterval: RollingInterval.Infinite,
                    rollOnFileSizeLimit: true,
                    fileSizeLimitBytes: LogSizeLimitBytes,
                    retainedFileCountLimit: 2
                    )
                .CreateLogger();

            return logger;
        }

        private ILogger CreateDefaultLogger() => new LoggerConfiguration()
              .WriteTo.RollingFile(DefaultLogFile)
              .CreateLogger();
    }
}
