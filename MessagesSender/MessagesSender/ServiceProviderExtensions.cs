﻿using System;
using Atlas.Common.Core.Interfaces;
using Atlas.Common.Impls;
using Atlas.Remoting.Core.Interfaces;
using Atlas.Remoting.Impls;
using MessagesSender.BL;
using MessagesSender.BL.Helpers;
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
        /// Add application services.
        /// </summary>
        /// <param name="services">service collection.</param>
        /// <returns>updated service collection.</returns>
        public static IServiceCollection AddTopicService(this IServiceCollection services)
        {
            return services.AddSingleton(
               typeof(ITopicService),
               typeof(TopicService));
        }

        /// <summary>
        /// Add application services.
        /// </summary>
        /// <param name="services">service collection.</param>
        /// <returns>updated service collection.</returns>
        public static IServiceCollection AddAppServices(this IServiceCollection services)
        {
            services.AddSingleton(
               typeof(IEventPublisher),
               typeof(EventPublisher));
            
            services.AddTopicService();

            services.AddSingleton(
               typeof(IZipService),
               typeof(ZipService));

            services.AddSingleton(
               typeof(IEmailSender),
               typeof(EmailSender));

            services.AddSingleton(
               typeof(ICommandService),
               typeof(CommandService));

            services.AddSingleton(
               typeof(ISendingService),
               typeof(SendingService));
            
            services.AddSingleton(
               typeof(IHardwareStateService),
               typeof(HardwareStateService));

            services.AddSingleton(
               typeof(ISystemWatchService),
               typeof(SystemWatchService));

            services.AddSingleton(
               typeof(IDicomStateService),
               typeof(DicomStateService));

            services.AddSingleton(
               typeof(IImagesWatchService),
               typeof(ImagesWatchService));

            services.AddSingleton(
               typeof(IStudyingWatchService),
               typeof(StudyingWatchService));

            services.AddSingleton(
               typeof(ISoftwareWatchService),
               typeof(SoftwareWatchService));

            services.AddSingleton(
               typeof(IRemoteControlService),
               typeof(RemoteControlService));

            services.AddSingleton(
               typeof(IHospitalInfoService),
               typeof(HospitalInfoService));

            services.AddSingleton(
               typeof(IDBDataService),
               typeof(DBDataService));

            services.AddSingleton(
               typeof(IOfflineService),
               typeof(OfflineService));

            services.AddSingleton(
               typeof(IKeepAliveService),
               typeof(KeepAliveService));

            return services.AddSingleton(
               typeof(IMainService),
               typeof(MainService));
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
                    try
                    {
                        var settingsEntityService =
                            new CommonDAL.Impls.SettingsEntityService(configService, logger);
                    }
                    catch (Exception ex)
                    {
                        logger.Error(ex, "AddConfigurationService");
                    }

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
                typeof(BL.LoggerBuilder));

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

            services.AddSingleton(
                typeof(IInfoEntityService),
                typeof(InfoEntityService));

            services.AddSingleton(
                typeof(IObservationsEntityService),
                typeof(ObservationsEntityService));

            services.AddSingleton(
                typeof(IMasterEntityService),
                typeof(MasterEntityService));

            services.AddSingleton(
                typeof(IConfigEntityService),
                typeof(ConfigEntityService));

            return services;
        }

        /// <summary>
        /// Add remoting services.
        /// </summary>
        /// <param name="services">service collection.</param>
        /// <returns>updated service collection.</returns>
        public static IServiceCollection AddRemotingServices(this IServiceCollection services)
        {
            services.AddSingleton(
               typeof(IFtpClient),
               typeof(FtpClient));

            services.AddSingleton(
               typeof(IWebClientService),
               typeof(WebClientService));
            
            services.AddMQTTRemotingServices();    

            return services.AddSingleton(
               typeof(IWorkqueueSender),
               typeof(RabbitMQWorkqueueSender)); 
        }

        /// <summary>
        /// Add remoting services.
        /// </summary>
        /// <param name="services">service collection.</param>
        /// <returns>updated service collection.</returns>
        public static IServiceCollection AddMQTTRemotingServices(this IServiceCollection services)
        {
            return services.AddSingleton(
               typeof(IMqttSender),
               typeof(RabbitMQTTSender));
        }

        /// <summary>
        /// Add chat message application services.
        /// </summary>
        /// <param name="services">service collection.</param>
        /// <returns>updated service collection.</returns>
        public static IServiceCollection AddChatMessageAppServices(this IServiceCollection services)
        {
            services.AddTopicService();

            /*services.AddSingleton(
               typeof(ISendingService),
               typeof(SendingService));*/

            services.AddSingleton(
                typeof(IEventPublisher),
                typeof(EventPublisher));

            return services.AddSingleton(
               typeof(IMainChatMessageService),
               typeof(MainChatMessageService));
        }
    }
}
