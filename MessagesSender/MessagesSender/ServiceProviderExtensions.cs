using Atlas.Common.Core.Interfaces;
using Atlas.Common.Impls;
using MessagesSender.BL;
using MessagesSender.BL.Remoting;
using MessagesSender.Core.Interfaces;
using MessagesSender.DAL;
using Microsoft.Extensions.DependencyInjection;
using Serilog;
using CommonCore = Atlas.Common.Core.Interfaces;
using CommonDAL = Atlas.Common.DAL;

namespace MessagesSender
{
    /// <summary>
    /// Extension methods for service collection.
    /// </summary>
    public static class ServiceProviderExtensions
    {
        /// <summary>
        /// Add journal service.
        /// </summary>
        /// <param name="services">service collection.</param>
        /// <returns>updated service collection.</returns>
        public static IServiceCollection AddSystemService(this IServiceCollection services)
        {

            return services.AddSingleton(
               typeof(IService),
               typeof(Service));
        }

        /// <summary>
        /// Add configuration service.
        /// </summary>
        /// <param name="services">service collection.</param>
        /// <returns>updated service collection.</returns>
        public static IServiceCollection AddConfigurationService(this IServiceCollection services)
        {
            services.AddSingleton(typeof(ConfigurationService));
            return services.AddSingleton(
                typeof(IConfigurationService),
                provider =>
                {
                    var configService = provider.GetService<ConfigurationService>();
                    var logger = provider.GetService<ILogger>();
                    var settingsEntityService =
                        new CommonDAL.Impls.SettingsEntityService(configService, logger);
                    return configService;
                });
        }

        /// <summary>
        /// Add logger service.
        /// </summary>
        /// <param name="services">service collection.</param>
        /// <param name="logName">path to log.</param>
        /// <returns>updated service collection.</returns>
        public static IServiceCollection AddLoggerService(
            this IServiceCollection services, string logName)
        {
            services.AddSingleton(
                typeof(ILoggerBuilder),
                typeof(LoggerBuilder));

            return services.AddSingleton(
                typeof(ILogger),
                provider =>
                {
                    var loggerBuilder = provider.GetService<ILoggerBuilder>();
                    return loggerBuilder.Build(logName);
                });
        }

        /// <summary>
        /// Add entity service.
        /// </summary>
        /// <param name="services">service collection.</param>
        /// <returns>updated service collection.</returns>
        public static IServiceCollection AddEntityServices(this IServiceCollection services)
        {
            services.AddSingleton(
                typeof(ISettingsEntityService),
                typeof(SettingsEntityService));

            return services;
        }

        /// <summary>
        /// Add remoting services.
        /// </summary>
        /// <param name="services">service collection.</param>
        /// <returns>updated service collection.</returns>
        public static IServiceCollection AddRemotingServices(this IServiceCollection services)
        {
            return services.AddSingleton(
               typeof(IWorkqueueSender),
               typeof(RabbitMQWorkqueueSender));
        }
    }
}
