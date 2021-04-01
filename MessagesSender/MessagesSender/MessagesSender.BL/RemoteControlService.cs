using System;
using System.Threading.Tasks;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using Atlas.Remoting.Core.Interfaces;
using MessagesSender.Core.Interfaces;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;
using Serilog;
using Atlas.Common.Impls.Helpers;
using System.Net;
using System.Linq;
using System.Net.Sockets;
using Atlas.Acquisitions.Common.Core;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System.Text;
using MessagesSender.Core.Model;
using Atlas.Acquisitions.Common.Core.Model;
using Atlas.Common.Core.Interfaces;
using MessagesSender.BL.Helpers;
using System.Diagnostics;
using System.IO.Compression;
using System.IO;
using System.Reflection;

namespace MessagesSender.BL
{
    public class RemoteControlService : IRemoteControlService
    {
		private enum XilogModes { Normal, Detailed }
		private enum XilogLevels { Info, Verbose }

		private const string InstallPathName = "InstallPath";
        private const string LogFolderPathName = "Logs";
		private const string XiLogFolderPathName = @"Logs\XiLogs";
		private const string TaskManPath = @"C:\Windows\System32\Taskmgr.exe";
		private const string XilogsFolder = @".\XiLogs\xilogs.exe";
		private const string XilogsCommandLineFormat = "1 {0} {1} \"{2}\"";

		private readonly IConfigurationService _configurationService;
        private readonly ISettingsEntityService _dbSettingsEntityService;
        private readonly IObservationsEntityService _dbObservationsEntityService;
        private readonly ILogger _logger;
        private readonly ISendingService _sendingService;
        private readonly IEventPublisher _eventPublisher;
		private readonly IZipService _zipService;
		private readonly IFtpClient _ftpClient;
		private readonly ITopicService _topicService;

		/// <summary>
		/// public constructor
		/// </summary>
		/// <param name="configurationService">configuration service</param>
		/// <param name="dbSettingsEntityService">settings database connector</param>
		/// <param name="dbObservationsEntityService">observations database connector</param>
		/// <param name="logger">logger</param>
		/// <param name="eventPublisher">event publisher service</param>
		/// <param name="sendingService">sending service</param>
		/// <param name="zipService">zip service</param>
		/// <param name="ftpClient">ftp client</param>
		/// <param name="topicService">topic service</param>
		public RemoteControlService(
            IConfigurationService configurationService,
            ISettingsEntityService dbSettingsEntityService,
            IObservationsEntityService dbObservationsEntityService,
            ILogger logger,
            IEventPublisher eventPublisher,
            ISendingService sendingService,
			IZipService zipService,
			IFtpClient ftpClient,
			ITopicService topicService)
        {
            _dbSettingsEntityService = dbSettingsEntityService;
            _dbObservationsEntityService = dbObservationsEntityService;
            _logger = logger;
            _sendingService = sendingService;
            _eventPublisher = eventPublisher;
            _configurationService = configurationService;
			_zipService = zipService;
			_ftpClient = ftpClient;
			_topicService = topicService;

			_eventPublisher.RegisterRunTVCommandArrivedEvent(() => RunTeamViewerAsync());
            _eventPublisher.RegisterRunTaskManCommandEvent(() => RunTaskManagerAsync());
            _eventPublisher.RegisterSendAtlasLogsCommandArrivedEvent(() => SendAtlasLogsAsync());
            _eventPublisher.RegisterXilibLogsOnCommandArrivedEvent(() => XilibLogsOnAsync());

            _logger.Information("RemoteControlService started");
        }

        /// <summary>
        /// runs TeamViewer
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> RunTeamViewerAsync()
        {
            return true;
        }

        /// <summary>
        /// runs TaskManager
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> RunTaskManagerAsync()
        {
            RegistryManager.SetPolicies(false);
            var process = Process.Start(TaskManPath);

            return true;
        }

        /// <summary>
        /// sends Atlas logs to email
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> SendAtlasLogsAsync()
        {
            var installPath = _configurationService.Get<string>(InstallPathName, @"C:\Program Files\Atlas\bin");
            try
            {
                var zip = await _zipService.ZipFolderAsync(Path.Combine(installPath, LogFolderPathName));
				// await EmailSender.SendAtlasLogsAsync(zip);
				await _ftpClient.SendAsync(zip);


				File.Delete(zip);
                Directory.Delete(Path.Combine(Path.GetDirectoryName(zip), Path.GetFileNameWithoutExtension(zip)), true);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "SendAtlasLogs error: ");
            }

            return true;
        }

		private string _xilog = string.Empty;
        /// <summary>
        /// turns on XilibLogs
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> XilibLogsOnAsync()
        {
			var installPath = _configurationService.Get<string>(InstallPathName, @"C:\Program Files\Atlas\bin");
			
			try
			{
				if (string.IsNullOrEmpty(_xilog))
				{
					var xilog = Path.Combine(installPath, XiLogFolderPathName);
					if (!Directory.Exists(xilog))
					{
						Directory.CreateDirectory(xilog);
					}

					xilog += @"\xilogs" + DateTime.Now.ToString("_dd_MM_yyyy_HH_mm_ss") + ".etl";

					RunCommand(
						XilogsFolder,
						string.Format(
							XilogsCommandLineFormat,
							//xilogs.exe [1-Run; 0-Stop] [Mode: Normal, Detailed] [Level: Info, Verbose] [path to the log]
							XilogModes.Normal,
							XilogLevels.Info,
							xilog
							));

					_xilog = xilog;
				}
				else
				{
					RunCommand(XilogsFolder, string.Empty);

					await Task.Yield();

					int i = 0;
					var zip = string.Empty;
					while (i < 5)
					{
						try
						{
							zip = await _zipService.ZipFileAsync(_xilog);
							break;
						}
						catch (Exception ex)
						{
							Console.WriteLine(ex.Message);

							await Task.Delay(1000);
							++i;
						}
					}

					await _ftpClient.SendAsync(zip);

					//File.Delete(zip);
					//Directory.Delete(Path.Combine(Path.GetDirectoryName(zip), Path.GetFileNameWithoutExtension(zip)), true);

					_xilog = string.Empty;
				}
			}
			catch (Exception ex)
			{
				_logger.Error(ex, "XilibLogsOnAsync error: ");
			}

			return true;
        }

		private void RunCommand(string exePath, string args)
		{
			var path = Path.Combine(Path.GetDirectoryName(Assembly.GetExecutingAssembly().Location), exePath);
			var processStartInfo = new ProcessStartInfo(
				path,
				//@"C:\Gits\MaintenanceManagementConsole\MessagesSender\MessagesSender\bin\Debug\netcoreapp3.1\XiLogs\xilogs.exe",
				args
				//"1 Detailed Verbose \"C:\\Program Files\\Atlas\\bin\\Logs\\XiLogs\\loggs4.etl\""
				);
			processStartInfo.WorkingDirectory = Path.GetDirectoryName(path);

			var process = Process.Start(processStartInfo);
		}
	}
}
