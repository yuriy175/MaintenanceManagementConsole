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
using MessagesSender.MessagesSender.BL.Helpers;
using System.Drawing.Imaging;

namespace MessagesSender.BL
{
    public class RemoteControlService : IRemoteControlService
    {
		private enum XilogModes { Normal, Detailed }
		private enum XilogLevels { Info, Verbose }

		private const string InstallPathName = "InstallPath";
		private const string EtlXilibLogsEnabledName = "EtlXilibLogsEnabled";
		private const string LogFolderPathName = "Logs";
		private const string XiLogFolderPathName = @"Logs\XiLogs";
		private const string XiLogFileName = @"xilogs.zip";
		private const string TaskManPath = @"C:\Windows\System32\Taskmgr.exe";
		private const string TeamViewerProcessName = @"TeamViewer";
		private const string TeamViewerPath = @"c:\Program Files (x86)\TeamViewer\TeamViewer.exe";//@"D:\res\SnapShot_src\SnapShot\SnapShot\bin\Debug\SnapShot.exe";
																				 //@"C:\Windows\System32\notepad.exe";
		private const string XilogsFolder = @".\XiLogs\xilogs.exe";
		private const string TeamViewerImagePath = @".\tvImage.jpeg";
		private const string XilogsCommandLineFormat = "{0} {1} {2} \"{3}\"";

		private readonly IConfigurationService _configurationService;
        private readonly ISettingsEntityService _dbSettingsEntityService;
        private readonly IObservationsEntityService _dbObservationsEntityService;
        private readonly ILogger _logger;
        private readonly ISendingService _sendingService;
        private readonly IEventPublisher _eventPublisher;
		private readonly IZipService _zipService;
		private readonly IFtpClient _ftpClient;
		private readonly ITopicService _topicService;
		private readonly IEmailSender _emailSender;

		private readonly AsyncLock _xilibSync = new AsyncLock();
		private readonly string _installPath = string.Empty;
		private bool _xilibState = false;

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
		/// <param name="emailSender"></param>
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
			IEmailSender emailSender,
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
			_emailSender = emailSender;
			_ftpClient = ftpClient;
			_topicService = topicService;

			_eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
			_eventPublisher.RegisterRunTVCommandArrivedEvent(() => RunTeamViewerAsync());
            _eventPublisher.RegisterRunTaskManCommandEvent(() => RunTaskManagerAsync());
            _eventPublisher.RegisterSendAtlasLogsCommandArrivedEvent(() => SendAtlasLogsAsync());
            _eventPublisher.RegisterXilibLogsOnCommandArrivedEvent(() => XilibLogsOnAsync());

			_installPath = _configurationService.Get<string>(InstallPathName, @"C:\Program Files\Atlas\bin");

			Task.Run(async () =>
			{
				var appParam = await _dbSettingsEntityService.GetAppParamAsync(EtlXilibLogsEnabledName);
				_xilibState = appParam == null ? false : Convert.ToBoolean(appParam.ParamValue);
			});

			_logger.Information("RemoteControlService started");
        }

        /// <summary>
        /// runs TeamViewer
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> RunTeamViewerAsync()
        {
			Process process = Process.GetProcesses().FirstOrDefault(p => p.ProcessName == TeamViewerProcessName);
			Process oldProc = null;

			if ((process != null && !WindowSnapshotHelper.IsValidUIWnd(process.MainWindowHandle)) || process == null)
			{
				oldProc = process;
				process = Process.Start(TeamViewerPath);
			}

			await Task.Delay(2000);

			var tvFile = Path.Combine(_installPath, LogFolderPathName, TeamViewerImagePath);
			var hWnd = WindowSnapshotHelper.GetWindowHandler(oldProc ?? process);
			
			var image = WindowSnapshotHelper.MakeSnapshot(hWnd, false, Win32API.WindowShowStyle.Restore);
			if (image != null)
			{
				image.Save(tvFile, ImageFormat.Jpeg);
			}
			
			var result = await _emailSender.SendTeamViewerAsync(tvFile);
			await SendTeamViewerStateAsync(result);

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
			var result = false;
			try
            {
                var zip = await _zipService.ZipFolderAsync(Path.Combine(installPath, LogFolderPathName));
				// await EmailSender.SendAtlasLogsAsync(zip);
				await _ftpClient.SendAsync(zip);


				File.Delete(zip);
                Directory.Delete(Path.Combine(Path.GetDirectoryName(zip), Path.GetFileNameWithoutExtension(zip)), true);

				result = true;
			}
            catch (Exception ex)
            {
                _logger.Error(ex, "SendAtlasLogs error: ");
            }

			await SendAtlasStateAsync(result);

			return true;
        }

		private string _xilog = string.Empty;
        /// <summary>
        /// turns on XilibLogs
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> XilibLogsOnAsync()
        {
			using (await _xilibSync.LockAsync())
			{
				try
				{
					if (!_xilibState)
					{
						var xilog = Path.Combine(_installPath, XiLogFolderPathName);
						if (!Directory.Exists(xilog))
						{
							Directory.CreateDirectory(xilog);
						}

						xilog += @"\xilogs" + DateTime.Now.ToString("_dd_MM_yyyy_HH_mm_ss") + ".etl";
						await _dbSettingsEntityService.UpsertAppParamAsync(EtlXilibLogsEnabledName, true);

						RunCommand(
							XilogsFolder,
							string.Format(
								XilogsCommandLineFormat,
								//xilogs.exe [1-Run; 0-Stop] [Mode: Normal, Detailed] [Level: Info, Verbose] [path to the log]
								"1",
								XilogModes.Normal,
								XilogLevels.Info,
								xilog
								));

						_xilibState = true;
						await SendXilogsStateAsync();
						_xilog = xilog;
					}
					else
					{
						RunCommand(
							XilogsFolder,
							string.Format(
								XilogsCommandLineFormat,
								//xilogs.exe [1-Run; 0-Stop] [Mode: Normal, Detailed] [Level: Info, Verbose] [path to the log]
								"0",
								XilogModes.Normal,
								XilogLevels.Info,
								_xilog
								));

						await Task.Yield();

						await _dbSettingsEntityService.UpsertAppParamAsync(EtlXilibLogsEnabledName, false);
						int i = 0;
						var zip = Path.Combine(Path.GetDirectoryName(_xilog), XiLogFileName); // Path.GetFileNameWithoutExtension(_xilog) + ".zip");
						if (!File.Exists(zip))
						{
							_logger.Error("XilibLogsOnAsync error: No zip file found");
							return false;
						}

						var result = await _ftpClient.SendAsync(zip, $"{await _topicService.GetTopicAsync()}_{Path.GetFileName(_xilog)}");

						_xilibState = false;
						await SendXilogsStateAsync(result);

						File.Delete(zip);

						_xilog = string.Empty;
					}
				}
				catch (Exception ex)
				{
					_logger.Error(ex, "XilibLogsOnAsync error: ");
					return false;
				}
			}

			return true;
        }

		private async Task<bool> OnActivateArrivedAsync()
		{
			await SendXilogsStateAsync();

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

			process.WaitForExit();
		}

		private async Task SendXilogsStateAsync(bool? ftpSendResult = null)
		{
			await _sendingService.SendInfoToMqttAsync(
				MQMessages.RemoteAccess,
				new
				{
					Xilogs = new { on = _xilibState},
					FtpSendResult = ftpSendResult
				});
		}

		private async Task SendAtlasStateAsync(bool? ftpSendResult = null)
		{
			await _sendingService.SendInfoToMqttAsync(
				MQMessages.RemoteAccess,
				new
				{
					FtpSendResult = ftpSendResult
				});
		}

		private async Task SendTeamViewerStateAsync(bool? emailSendResult = null)
		{
			await _sendingService.SendInfoToMqttAsync(
				MQMessages.RemoteAccess,
				new
				{
					EmailSendResult = emailSendResult
				});
		}
	}
}
