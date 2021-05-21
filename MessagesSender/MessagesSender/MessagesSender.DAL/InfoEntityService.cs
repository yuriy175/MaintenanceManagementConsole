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
            )> GetSystemDataAsync()
		{
            var hardDrives = await GetManyAction<HardDrive>(context => context.HardDrives);
            var lans = await GetManyAction<Lan>(context => context.Lans);
            var logicalDisks = await GetManyAction<LogicalDisk>(context => context.LogicalDisks);
            var modems = await GetManyAction<Modem>(context => context.Modems);
            var monitors = await GetManyAction<Monitor>(context => context.Monitors);
            var motherboards = await GetManyAction<Motherboard>(context => context.Motherboards);
            var printers = await GetManyAction<Printer>(context => context.Printers);
            var screens = await GetManyAction<Screen>(context => context.Screens);
            var videoadapters = await GetManyAction<VideoAdapter>(context => context.VideoAdapters);

            return (hardDrives, lans, logicalDisks, modems, monitors, motherboards, printers, screens, videoadapters);
        }

        /// <summary>
        /// get software tables data
        /// </summary>
        /// <returns>software tables data</returns>
        public async Task<(
            IEnumerable<AtlasSW> Atlas, 
            IEnumerable<Dependency> Dependencies,
            IEnumerable<Error> Errors,
            IEnumerable<OsInfo> OsInfos,
            IEnumerable<SqlDatabase> SqlDatabases, 
            IEnumerable<SqlService> SqlServices
            )> GetSoftwareDataAsync()
        {
            var atlas = await GetManyAction<AtlasSW>(context => context.Atlases);
            var dependencies = await GetManyAction<Dependency>(context => context.Dependencies);
            var errors = await GetManyAction<Error>(context => context.Errors);
            var osInfos = await GetManyAction<OsInfo>(context => context.OsInfos);
            var sqlDatabases = await GetManyAction<SqlDatabase>(context => context.SqlDatabases);
            var sqlServices = await GetManyAction<SqlService>(context => context.SqlServices);

            return (atlas, dependencies, errors, osInfos, sqlDatabases, sqlServices);
        }

        /// <summary>
        /// get atlas tables data
        /// </summary>
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
            )> GetAtlasDataAsync()
        {
            var appParams = await GetManyAction<AppParam>(context => context.AppParams);
			// var aspnetusers = new AspNetUser[] { }; // await GetManyAction<AspNetUser>(context => context.AspNetUsers);
			var aspnetusers = await GetManyAction<AspNetUser>(context => context.AspNetUsers);
			var detectors = await GetManyAction<Detector>(context => context.Detectors);
            var detectorprocessings = await GetManyAction<DetectorProcessing>(context => context.DetectorProcessings);
            var dicoms = await GetManyAction<DicomService>(context => context.DicomServices);
            var dicomPrinters = await GetManyAction<DicomPrinter>(context => context.DicomPrinters);
            var hardwareparams = await GetManyAction<HardwareParam>(context => context.HardwareParams);
            var rasterParams = await GetManyAction<RasterParam>(context => context.RasterParams);
            
            return (appParams, aspnetusers, detectors, detectorprocessings, dicoms, dicomPrinters, 
                hardwareparams, rasterParams);
        }

        /// <summary>
        /// get hospital table data
        /// </summary>
        /// <returns>hospital table data</returns>
        public async Task<IEnumerable<HospitalInfo>> GetHospitalDataAsync()
        {
            return await GetManyAction<HospitalInfo>(context => context.HospitalInfos);
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
    }
}
