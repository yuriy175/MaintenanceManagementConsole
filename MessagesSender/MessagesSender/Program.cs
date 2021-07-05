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
        /// <summary>
        /// Main function
        /// </summary>
        /// <param name="args">command line argumentes</param>
        public static void Main(string[] args)
        {
            using var mutex = new Mutex(false, "MessagesSender");
            if (!mutex.WaitOne(0, false))
            {
                Console.WriteLine("Instance already running");
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
            IService service = provider.GetRequiredService<IService>();
        }
    }
}
