using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Atlas.Common.Core.Interfaces;
using Atlas.Common.DAL;
using Atlas.Common.DAL.Helpers;
using Atlas.Common.DAL.Model;
using MessagesSender.Core.Interfaces;
using MessagesSender.DAL.Model;
using Microsoft.EntityFrameworkCore;
using Serilog;

namespace MessagesSender.DAL
{
    /// <summary>
    /// InfoEntityService implementation
    /// </summary>
    public class InfoEntityService
        : EntityServiceBase<InfoContext>, IInfoEntityService
    {
        private const string HardDriveTableName = "hard_drive";
        private const string PrintersTableName = "printers";
        private const string AppParamTableName = "appparam";
        private const string AspNetUsersTableName = "aspnetusers";
        private const string AtlasTableName = "atlas";
        private const string DependenciesTableName = "dependencies";
        private const string DetectorTableName = "detector";
        private const string DetectorProcessingsTableName = "detectorprocessings";
        private const string DicomPrintersTableName = "dicom_printers";
        private const string DicomServicesTableName = "dicomservices";
        private const string ErrorsTableName = "errors";
        private const string HardwareParamsTableName = "hardwareparams";
        private const string HospitalInfoTableName = "hospital_info";
        private const string LanTableName = "lan";
        private const string LogicalDisksTableName = "logical_disks";
        private const string MemoryTableName = "memory";
        private const string ModemTableName = "modem";
        private const string MonitorsTableName = "monitors";
        private const string MotherboardTableName = "motherboard";
        private const string OsInfoTableName = "os_info";
        private const string RasterParamsTableName = "rasterparams";
        private const string ScreensTableName = "screens";
        private const string SqlDatabasesTableName = "sql_databases";
        private const string SqlServiceTableName = "sql_service";
        private const string VideoadapterTableName = "videoadapter";

        private readonly IConfigurationService _configurationService = null;
        private readonly ILogger _logger;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger.</param>
        public InfoEntityService(
            IConfigurationService configurationService,
            ILogger logger)
            : base(logger)
        {
            _configurationService = configurationService;
            _logger = logger;
        }

        /// <summary>
        /// get system tables data
        /// </summary>
        /// <param name="news">news data</param>
        /// <returns>system tables data</returns>
        public async Task<(
            IEnumerable<HardDrive> HardDrives, 
            IEnumerable<Lan> Lans,
            IEnumerable<LogicalDisk> LogicalDisks,
            IEnumerable<Modem> Modems,
            IEnumerable<Monitor> Monitors,
            IEnumerable<Motherboard> Motherboards,
            IEnumerable<Printer> Printers,
            IEnumerable<Screen> Screens,
            IEnumerable<VideoAdapter> VideoAdapters
            )> GetSystemDataAsync(Dictionary<string, int[]> news)
		{
            var hardDrives = await GetNewRowsAsync(news, HardDriveTableName,
                (context, rows) => context.HardDrives.Where(e => rows.Contains(e.RowId)));
            var lans = await GetNewRowsAsync(news, LanTableName,
                (context, rows) => context.Lans.Where(e => rows.Contains(e.RowId)));
            var logicalDisks = await GetNewRowsAsync(news, LogicalDisksTableName,
                (context, rows) => context.LogicalDisks.Where(e => rows.Contains(e.RowId)));
            var modems = await GetNewRowsAsync(news, ModemTableName,
                (context, rows) => context.Modems.Where(e => rows.Contains(e.RowId)));
            var monitors = await GetNewRowsAsync(news, MonitorsTableName,
                (context, rows) => context.Monitors.Where(e => rows.Contains(e.RowId)));
            var motherboards = await GetNewRowsAsync(news, MotherboardTableName,
                (context, rows) => context.Motherboards.Where(e => rows.Contains(e.RowId)));
            var printers = await GetNewRowsAsync(news, PrintersTableName,
                (context, rows) => context.Printers.Where(e => rows.Contains(e.RowId)));

            var screens = await GetNewRowsAsync(news, ScreensTableName,
                (context, rows) => context.Screens.Where(e => rows.Contains(e.RowId)));
            var videoadapters = await GetNewRowsAsync(news, VideoadapterTableName,
                (context, rows) => context.VideoAdapters.Where(e => rows.Contains(e.RowId)));

            return (hardDrives, lans, logicalDisks, modems, monitors, motherboards, printers, screens, videoadapters);
        }

        /// <summary>
        /// get software tables data
        /// </summary>
        /// <param name="news">news data</param>
        /// <returns>software tables data</returns>
        public async Task<(
            IEnumerable<AtlasSW> Atlas, 
            IEnumerable<Dependency> Dependencies,
            IEnumerable<Error> Errors,
            IEnumerable<OsInfo> OsInfos,
            IEnumerable<SqlDatabase> SqlDatabases, 
            IEnumerable<SqlService> SqlServices
            )> GetSoftwareDataAsync(Dictionary<string, int[]> news)
        {
            var atlas = await GetNewRowsAsync(news, AtlasTableName,
                (context, rows) => context.Atlases.Where(e => rows.Contains(e.RowId)));
            var dependencies = await GetNewRowsAsync(news, DependenciesTableName,
                (context, rows) => context.Dependencies.Where(e => rows.Contains(e.RowId)));
            var errors = await GetNewRowsAsync(news, ErrorsTableName,
                (context, rows) => context.Errors.Where(e => rows.Contains(e.RowId)));
            var osInfos = await GetNewRowsAsync(news, OsInfoTableName,
                (context, rows) => context.OsInfos.Where(e => rows.Contains(e.RowId)));
            var sqlDatabases = await GetNewRowsAsync(news, SqlDatabasesTableName,
                (context, rows) => context.SqlDatabases.Where(e => rows.Contains(e.RowId)));
            var sqlServices = await GetNewRowsAsync(news, SqlServiceTableName,
                (context, rows) => context.SqlServices.Where(e => rows.Contains(e.RowId)));

            return (atlas, dependencies, errors, osInfos, sqlDatabases, sqlServices);
        }

        /// <summary>
        /// get atlas tables data
        /// </summary>
        /// <param name="news">news data</param>
        /// <returns>atlas tables data</returns>
        public async Task<(
            IEnumerable<AppParam> AppParams, 
            IEnumerable<AspNetUser> AspNetUsers,
            IEnumerable<Detector> Detectors,
            IEnumerable<DetectorProcessing> DetectorProcessings,
            IEnumerable<DicomService> DicomServices,
            IEnumerable<DicomPrinter> DicomPrinters,
            IEnumerable<HardwareParam> HardwareParams,
            IEnumerable<RasterParam> RasterParams
            )> GetAtlasDataAsync(Dictionary<string, int[]> news)
        {
            var appParams = await GetNewRowsAsync(news, AppParamTableName,
                (context, rows) => context.AppParams.Where(e => rows.Contains(e.RowId)));
            var aspnetusers = await GetNewRowsAsync(news, AspNetUsersTableName,
                (context, rows) => context.AspNetUsers.Where(e => rows.Contains(e.RowId)));
            var detectors = await GetNewRowsAsync(news, DetectorTableName,
                (context, rows) => context.Detectors.Where(e => rows.Contains(e.RowId)));
            var detectorprocessings = await GetNewRowsAsync(news, DetectorProcessingsTableName,
                (context, rows) => context.DetectorProcessings.Where(e => rows.Contains(e.RowId)));
            var dicoms = await GetNewRowsAsync(news, DicomServicesTableName,
                (context, rows) => context.DicomServices.Where(e => rows.Contains(e.RowId)));
            var dicomPrinters = await GetNewRowsAsync(news, DicomPrintersTableName,
                (context, rows) => context.DicomPrinters.Where(e => rows.Contains(e.RowId)));
            var hardwareparams = await GetNewRowsAsync(news, HardwareParamsTableName,
                (context, rows) => context.HardwareParams.Where(e => rows.Contains(e.RowId)));
            var rasterParams = await GetNewRowsAsync(news, RasterParamsTableName,
                (context, rows) => context.RasterParams.Where(e => rows.Contains(e.RowId)));

            return (appParams, aspnetusers, detectors, detectorprocessings, dicoms, dicomPrinters, 
                hardwareparams, rasterParams);
        }

        /// <summary>
        /// get hospital table data
        /// </summary>
        /// <param name="news">news data</param>
        /// <returns>hospital table data</returns>
        public async Task<IEnumerable<HospitalInfo>> GetHospitalDataAsync(
            Dictionary<string, int[]> news)
        {
            return await GetNewRowsAsync(news, HospitalInfoTableName,
                (context, rows) => context.HospitalInfos.Where(e => rows.Contains(e.RowId))); 
        }

        /// <summary>
        /// get news table data
        /// </summary>
        /// <returns>news table data</returns>
        public async Task<IEnumerable<News>> GetNewsDataAsync()
        {
            return await GetManyAction<News>(context => context.News.Where(n => n.Active));
        }

        /// <summary>
        /// set news table data sent
        /// </summary>
        /// <returns></returns>
        public async Task<bool> SetNewsDataSentAsync()
        {
            using (var context = CreateContext())
            {
                foreach(var news in context.News.Where(n => n.Active).ToList())
                {
                    news.Active = false;
                }
                
                await context.SaveChangesAsync();
                return true;
            }
        }

        /// <summary>
        /// Create context.
        /// </summary>
        /// <param name="logger">logger.</param>
        /// <returns>settings context.</returns>
        protected override InfoContext CreateContext() =>
            InfoContext.Create(
                    _configurationService?["ConnectionStrings", "InfoConnection"],
                    _logger);

        private async Task<IEnumerable<T>> GetNewRowsAsync<T>(
            Dictionary<string, int[]> news,
            string tableName,
            Func<InfoContext, int[], IQueryable<T>> action)
            where T : class
        {
            if (!news.ContainsKey(tableName))
            {
                return new T[]{ };
            }

            var rowIds = news[tableName];
            return await GetManyAction<T>(context => action(context, rowIds));
        }
    }
}
