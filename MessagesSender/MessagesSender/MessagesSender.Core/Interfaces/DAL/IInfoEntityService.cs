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
        /// <param name="news">news data</param>
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
            )> GetSystemDataAsync(Dictionary<string, int[]> news);

        /// <summary>
        /// get software tables data
        /// </summary>
        /// <param name="news">news data</param>
        /// <returns>software tables data</returns>
        Task<(
            IEnumerable<AtlasSW> Atlas,
            IEnumerable<Dependency> Dependencies,
            IEnumerable<Error> Errors,
            IEnumerable<OsInfo> OsInfos,
            IEnumerable<SqlDatabase> SqlDatabases,
            IEnumerable<SqlService> SqlServices
            )> GetSoftwareDataAsync(Dictionary<string, int[]> news);

        /// <summary>
        /// get atlas tables data
        /// </summary>
        /// <param name="news">news data</param>
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
            )> GetAtlasDataAsync(Dictionary<string, int[]> news);

        /// <summary>
        /// get hospital table data
        /// </summary>
        /// <param name="news">news data</param>
        /// <returns>hospital table data</returns>
        Task<IEnumerable<HospitalInfo>> GetHospitalDataAsync(Dictionary<string, int[]> news);

        /// <summary>
        /// get news table data
        /// </summary>
        /// <returns>news table data</returns>
        Task<IEnumerable<News>> GetNewsDataAsync();

        /// <summary>
        /// set news table data sent
        /// </summary>
        /// <returns>result</returns>
        Task<bool> SetNewsDataSentAsync();
    }
}
