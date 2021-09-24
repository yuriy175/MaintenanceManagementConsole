using System;
using System.Diagnostics;
using System.Threading;
using System.Threading.Tasks;
using Atlas.Common.Core.Interfaces;
using Atlas.Remoting.Helpers;
using Atlas.Remoting.Impls;
using MessagesSender.Core.Interfaces;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

namespace MessagesSender
{
    /// <summary>
    /// Program
    /// </summary>
    public class Program
    {
        private const string MessageModeCommandArgName = "-m";

        /// <summary>
        /// Main function
        /// </summary>
        /// <param name="args">command line argumentes</param>
        public static void Main(string[] args)
        {
            if (args.Length > 1)
            {
                StartForSpecialService(args);
                return;
            }

            using var mutex = new Mutex(false, "MessagesSender");
            if (!mutex.WaitOne(0, false))
            {
                Console.WriteLine("Instance already running");

                Task.Run(async () =>
                {
                    await Task.Delay(1000);
                    Environment.Exit(0);
                });

                Console.ReadKey();

                return;
            }

            using IHost host = CreateHostBuilder(args).Build();

            Configure(host.Services);

            host.Run();
        }

        private static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .ConfigureServices((_, services) =>
                    services.AddConfigurationService()
                        .AddLoggerService("MsgSender")
                        .AddMQRemotingServices<MQCommunicationService>()
                        .AddEntityServices()
                        .AddRemotingServices()
                        .AddAppServices());

        private static void Configure(IServiceProvider services)
        {
            using IServiceScope serviceScope = services.CreateScope();
            IServiceProvider provider = serviceScope.ServiceProvider;

            var topicService = provider.GetRequiredService<ITopicService>(); 
            var service = provider.GetRequiredService<IMainService>();
        }

        private static void StartForSpecialService(string[] args)
        {
            if (args.Length < 1)
            {
                return;
            }

            var argNumber = 0;

            while (argNumber < args.Length)
            {
                switch (args[argNumber])
                {
                    case MessageModeCommandArgName:
                        {
                            if (args.Length < argNumber + 2)
                            {
                                Console.WriteLine("-m option requires at least 2 args");
                                return;
                            }

                            ++argNumber;

                            using IHost host = CreateMessageModeHostBuilder(args).Build();

                            var service = ConfigureMessageMode(host.Services);
                            _ = service.SendChatMessageAsync(args[argNumber]).Result;

                            break;
                        }
                }

                ++argNumber;
            }
        }

        private static IHostBuilder CreateMessageModeHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .ConfigureServices((_, services) =>
                    services.AddConfigurationService()
                        .AddLoggerService("MsgSender")
                        .AddMQRemotingServices<MQCommunicationService>()
                        .AddEntityServices()
                        .AddMQTTRemotingServices()
                        .AddChatMessageAppServices());

        private static IMainChatMessageService ConfigureMessageMode(IServiceProvider services)
        {
            using IServiceScope serviceScope = services.CreateScope();
            IServiceProvider provider = serviceScope.ServiceProvider;

            var topicService = provider.GetRequiredService<ITopicService>();
            return provider.GetRequiredService<IMainChatMessageService>();
        }
    }
}
