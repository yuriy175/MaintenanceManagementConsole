using System;
using System.Diagnostics;
using System.Drawing.Imaging;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Atlas.Common.Core.Interfaces;
using Atlas.Common.Impls.Helpers;
using MessagesSender.BL.Helpers;
using MessagesSender.Core.Interfaces;
using MessagesSender.Core.Model;
using MessagesSender.MessagesSender.BL.Helpers;
using Serilog;

namespace MessagesSender.BL
{
    /// <summary>
    /// remote control service interface implementation
    /// </summary>
    public class RemoteControlService : IRemoteControlService
    {
        private const string InstallPathName = "InstallPath";
        private const string EtlXilibLogsEnabledName = "EtlXilibLogsEnabled";
        private const string LogFolderPathName = "Logs";
        private const string XiLogFolderPathName = @"Logs\XiLogs";
        private const string XiLogFileName = @"xilogs.zip";
        private const string TaskManPath = @"C:\Windows\System32\Taskmgr.exe";
        private const string TeamViewerProcessName = @"TeamViewer";
        private const string TeamViewerPath = @"c:\Program Files\TeamViewer\TeamViewer.exe";
        private const string AdditionalTeamViewerPath = @"c:\Program Files (x86)\TeamViewer\TeamViewer.exe";

        // @"D:\res\SnapShot_src\SnapShot\SnapShot\bin\Debug\SnapShot.exe";
        // @"C:\Windows\System32\notepad.exe";
        private const string XilogsFolder = @".\XiLogs\xilogs.exe";

        // private const string TeamViewerImagePath = @".\tvImage.jpeg";
        private const string TeamViewerImagePath = @"tvImage.jpeg";
        private const string XilogsCommandLineFormat = "{0} {1} {2} \"{3}\"";
        private const string SqlInfoExePath = @".\SqlInfo\sqlinfo.exe";
        private const string SqlInfoCommandLine = "";

        private readonly IConfigurationService _configurationService;
        private readonly IConfigEntityService _dbConfigEntityService;
        private readonly IObservationsEntityService _dbObservationsEntityService;
        private readonly ILogger _logger;
        private readonly ISendingService _sendingService;
        private readonly IEventPublisher _eventPublisher;
        private readonly IZipService _zipService;
        private readonly IFtpClient _ftpClient;
        private readonly ITopicService _topicService;
        private readonly IEmailSender _emailSender;
        private readonly IDBDataService _dbDataService;

        private readonly AsyncLock _xilibSync = new AsyncLock();
        private readonly string _installPath = string.Empty;
        private bool _xilibState = false;
        private string _xilog = string.Empty;
        private bool _isDBInfoUpdating = false;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="dbConfigEntityService">settings database connector</param>
        /// <param name="dbObservationsEntityService">observations database connector</param>
        /// <param name="logger">logger</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        /// <param name="zipService">zip service</param>
        /// <param name="emailSender">email sender</param>
        /// <param name="ftpClient">ftp client</param>
        /// <param name="topicService">topic service</param>
        /// <param name="dbDataService">db raw data service</param>
        public RemoteControlService(
            IConfigurationService configurationService,
            IConfigEntityService dbConfigEntityService,
            IObservationsEntityService dbObservationsEntityService,
            ILogger logger,
            IEventPublisher eventPublisher,
            ISendingService sendingService,
            IZipService zipService,
            IEmailSender emailSender,
            IFtpClient ftpClient,
            ITopicService topicService,
            IDBDataService dbDataService)
        {
            _dbConfigEntityService = dbConfigEntityService;
            _dbObservationsEntityService = dbObservationsEntityService;
            _logger = logger;
            _sendingService = sendingService;
            _eventPublisher = eventPublisher;
            _configurationService = configurationService;
            _zipService = zipService;
            _emailSender = emailSender;
            _ftpClient = ftpClient;
            _topicService = topicService;
            _dbDataService = dbDataService;

            _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            _eventPublisher.RegisterRunTVCommandArrivedEvent(() => RunTeamViewerAsync());
            _eventPublisher.RegisterRunTaskManCommandEvent(() => RunTaskManagerAsync());
            _eventPublisher.RegisterSendAtlasLogsCommandArrivedEvent(() => SendAtlasLogsAsync());
            _eventPublisher.RegisterXilibLogsOnCommandArrivedEvent(() => XilibLogsOnAsync());
            _eventPublisher.RegisterUpdateDBInfoCommandArrivedEvent(() => OnUpdateDBInfoAsync());

            _installPath = _configurationService.Get<string>(InstallPathName, @"C:\Program Files\Atlas\bin");

            new Action[]
                {
                    async () => _ = await OnUpdateDBInfoAsync(),
                    async () =>
                    {
                        var configParam = await _dbConfigEntityService.GetConfigParamAsync(EtlXilibLogsEnabledName);
                        _xilibState = configParam == null ? false : Convert.ToBoolean(configParam.ParamValue);
                    },
                }.RunTasksAsync();
        }

        private enum XilogModes { Normal, Detailed }

        private enum XilogLevels { Info, Verbose }

        /// <summary>
        /// runs TeamViewer
        /// </summary>
        /// <returns>result</returns>
        public async Task<bool> RunTeamViewerAsync()
        {
            Process process = Process.GetProcesses().FirstOrDefault(p => p.ProcessName == TeamViewerProcessName);
            Process oldProc = null;

            var teamViewerPath = File.Exists(TeamViewerPath) ? TeamViewerPath : AdditionalTeamViewerPath;
            if ((process != null && !WindowSnapshotHelper.IsValidUIWnd(process.MainWindowHandle)) || process == null)
            {
                oldProc = process;
                process = Process.Start(teamViewerPath);
            }

            await Task.Delay(2000);

            // var tvFile = Path.Combine(_installPath, LogFolderPathName, TeamViewerImagePath);
            var tvFile = Path.Combine(Path.GetTempPath(), TeamViewerImagePath);
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
            try
            {
                var process = Process.Start(TaskManPath);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "RunTaskManagerAsync error");
                return false;
            }

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
                        await _dbConfigEntityService.UpsertConfigParamAsync(EtlXilibLogsEnabledName, true);

                        // xilogs.exe [1-Run; 0-Stop] [Mode: Normal, Detailed] [Level: Info, Verbose] [path to the log]
                        ProcessHelper.ProcessRunAndWait(
                            XilogsFolder,
                            string.Format(
                                XilogsCommandLineFormat,
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
                        ProcessHelper.ProcessRunAndWait(
                            XilogsFolder,
                            string.Format(
                                XilogsCommandLineFormat,

                                // xilogs.exe [1-Run; 0-Stop] [Mode: Normal, Detailed] [Level: Info, Verbose] [path to the log]
                                "0",
                                XilogModes.Normal,
                                XilogLevels.Info,
                                _xilog
                                ));

                        await Task.Yield();

                        await _dbConfigEntityService.UpsertConfigParamAsync(EtlXilibLogsEnabledName, false);
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

        private async Task<bool> OnUpdateDBInfoAsync()
        {
            if (_isDBInfoUpdating)
            {
                await SendDBInfoStateAsync();
                return true;
            }

            _isDBInfoUpdating = true;
            await SendDBInfoStateAsync();

            ProcessHelper.ProcessRunAndWait(SqlInfoExePath, SqlInfoCommandLine);
            await _dbDataService.UpdateDBInfoAsync();

            _isDBInfoUpdating = false;
            await SendDBInfoStateAsync();

            return true;
        }

        private async Task SendXilogsStateAsync(bool? ftpSendResult = null)
        {
            await _sendingService.SendInfoToMqttAsync(
                MQMessages.RemoteAccess,
                new
                {
                    Xilogs = new { on = _xilibState },
                    FtpSendResult = ftpSendResult,
                });
        }

        private async Task SendAtlasStateAsync(bool? ftpSendResult = null)
        {
            await _sendingService.SendInfoToMqttAsync(
                MQMessages.RemoteAccess,
                new
                {
                    FtpSendResult = ftpSendResult,
                });
        }

        private async Task SendTeamViewerStateAsync(bool? emailSendResult = null)
        {
            await _sendingService.SendInfoToMqttAsync(
                MQMessages.RemoteAccess,
                new
                {
                    EmailSendResult = emailSendResult,
                });
        }

        private async Task SendDBInfoStateAsync()
        {
            await _sendingService.SendInfoToMqttAsync(
                MQMessages.RemoteAccess,
                new
                {
                    DBInfoStateUpdating = _isDBInfoUpdating,
                });
        }
    }
}
