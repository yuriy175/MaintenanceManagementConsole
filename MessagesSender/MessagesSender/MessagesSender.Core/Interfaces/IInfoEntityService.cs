using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Atlas.Common.DAL.Model;
using MessagesSender.DAL.Model;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// Interface to info database
    /// </summary>
    public interface IInfoEntityService
    {
        /// <summary>
        /// get system tables data
        /// </summary>
        /// <returns>system tables data</returns>
        Task<(
            IEnumerable<HardDrive> HardDrives,
            IEnumerable<Lan> Lans,
            IEnumerable<LogicalDisk> LogicalDisks,
            IEnumerable<Modem> Modems,
            IEnumerable<Monitor> Monitors,
            IEnumerable<Motherboard> Motherboards,
            IEnumerable<Printer> Printers,
            IEnumerable<Screen> Screens,
            IEnumerable<VideoAdapter> VideoAdapters
            )> GetSystemDataAsync();

        /// <summary>
        /// get software tables data
        /// </summary>
        /// <returns>software tables data</returns>
        Task<(
            IEnumerable<AtlasSW> Atlas,
            IEnumerable<Dependency> Dependencies,
            IEnumerable<Error> Errors,
            IEnumerable<OsInfo> OsInfos,
            IEnumerable<SqlDatabase> SqlDatabases,
            IEnumerable<SqlService> SqlServices
            )> GetSoftwareDataAsync();

        /// <summary>
        /// get atlas tables data
        /// </summary>
        /// <returns>atlas tables data</returns>
        Task<(
            IEnumerable<AppParam> AppParams,
            IEnumerable<AspNetUser> AspNetUsers,
            IEnumerable<Detector> Detectors,
            IEnumerable<DetectorProcessing> DetectorProcessings,
            IEnumerable<DicomService> DicomServices,
            IEnumerable<DicomPrinter> DicomPrinters,
            IEnumerable<HardwareParam> HardwareParams,
            IEnumerable<RasterParam> RasterParams
            )> GetAtlasDataAsync();

        /// <summary>
        /// get hospital table data
        /// </summary>
        /// <returns>hospital table data</returns>
        Task<IEnumerable<HospitalInfo>> GetHospitalDataAsync();
    }
}
