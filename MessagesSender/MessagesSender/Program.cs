using Atlas.Common.Core.Interfaces;
using Atlas.Remoting.Helpers;
using Atlas.Remoting.Impls;
using MessagesSender.Core.Interfaces;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using System;
using System.Threading.Tasks;

namespace MessagesSender
{
    class Program
    {
        //static void Main(string[] args)
        //{
        //    Console.WriteLine("Hello World!");
        //}

        static void Main(string[] args)
        {
            using IHost host = CreateHostBuilder(args).Build();

            Configure(host.Services);

            host.Run();
        }

        static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .ConfigureServices((_, services) =>
                    services.AddConfigurationService()
                        .AddLoggerService("MsgSender")
                        .AddMQRemotingServices<MQCommunicationService>()
                        .AddEntityServices()
                        .AddRemotingServices()
                        .AddAppServices());

        static void Configure(IServiceProvider services)
        {
            using IServiceScope serviceScope = services.CreateScope();
            IServiceProvider provider = serviceScope.ServiceProvider;

            IService service = provider.GetRequiredService<IService>();
        }
    }
}
