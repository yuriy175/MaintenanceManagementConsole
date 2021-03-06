﻿using MessagesSender.DAL.Model;
using Microsoft.EntityFrameworkCore;
using Serilog;
using CommonDAL = Atlas.Common.DAL;
using MessagesSenderDAL = MessagesSender.DAL.Model;

namespace MessagesSender.DAL
{
    /// <summary>
    /// Info sqllite db cntext.
    /// </summary>
    public partial class InfoContext : CommonDAL.DbContextBase
    {
        private static DbContextOptions<InfoContext> _options = null;

        /// <summary>
        /// Initializes a new instance of the <see cref="MasterContext"/> class.
        /// </summary>
        /// <param name="options">options.</param>
        /// <param name="logger">logger.</param>
        public InfoContext(
            DbContextOptions options,
            ILogger logger)
        : base(options, logger)
        {
        }

        /// <summary>
        /// set of AppParam
        /// </summary>
        public virtual DbSet<AppParam> AppParams { get; set; }

        /// <summary>
        /// set of Dependency
        /// </summary>
        public virtual DbSet<Dependency> Dependencies { get; set; }

        /// <summary>
        /// set of AspNetUser
        /// </summary>
        public virtual DbSet<AspNetUser> AspNetUsers { get; set; }

        /// <summary>
        /// set of AtlasSW
        /// </summary>
        public virtual DbSet<AtlasSW> Atlases { get; set; }

        /// <summary>
        /// set of Detector
        /// </summary>
        public virtual DbSet<Detector> Detectors { get; set; }

        /// <summary>
        /// set of Dicom
        /// </summary>
        public virtual DbSet<DicomService> DicomServices { get; set; }        

        /// <summary>
        /// set of DetectorProcessing
        /// </summary>
        public virtual DbSet<DetectorProcessing> DetectorProcessings { get; set; }

        /// <summary>
        /// set of DicomPrinters
        /// </summary>
        public virtual DbSet<DicomPrinter> DicomPrinters { get; set; }

        /// <summary>
        /// set of Error
        /// </summary>
        public virtual DbSet<Error> Errors { get; set; }

        /// <summary>
        /// set of Error
        /// </summary>
        public virtual DbSet<HardDrive> HardDrives { get; set; }

        /// <summary>
        /// set of Error
        /// </summary>
        public virtual DbSet<HardwareParam> HardwareParams { get; set; }

        /// <summary>
        /// set of Error
        /// </summary>
        public virtual DbSet<HospitalInfo> HospitalInfos { get; set; }

        /// <summary>
        /// set of Lan
        /// </summary>
        public virtual DbSet<Lan> Lans { get; set; }

        /// <summary>
        /// set of LogicalDisk
        /// </summary>
        public virtual DbSet<LogicalDisk> LogicalDisks { get; set; }

        /// <summary>
        /// set of Modem
        /// </summary>
        public virtual DbSet<Modem> Modems { get; set; }

        /// <summary>
        /// set of Monitor
        /// </summary>
        public virtual DbSet<Monitor> Monitors { get; set; }

        /// <summary>
        /// set of Motherboard
        /// </summary>
        public virtual DbSet<Motherboard> Motherboards { get; set; }

        /// <summary>
        /// set of OsInfo
        /// </summary>
        public virtual DbSet<OsInfo> OsInfos { get; set; }

        /// <summary>
        /// set of Printer
        /// </summary>
        public virtual DbSet<Printer> Printers { get; set; }

        /// <summary>
        /// set of RasterParam
        /// </summary>
        public virtual DbSet<RasterParam> RasterParams { get; set; }

        /// <summary>
        /// set of Screen
        /// </summary>
        public virtual DbSet<Screen> Screens { get; set; }

        /// <summary>
        /// set of SqlDatabase
        /// </summary>
        public virtual DbSet<SqlDatabase> SqlDatabases { get; set; }

        /// <summary>
        /// set of SqlService
        /// </summary>
        public virtual DbSet<SqlService> SqlServices { get; set; }

        /// <summary>
        /// set of VideoAdapter
        /// </summary>
        public virtual DbSet<VideoAdapter> VideoAdapters { get; set; }

        /// <summary>
        /// set of News
        /// </summary>
        public virtual DbSet<News> News { get; set; }

        /// <summary>
        /// Creates SettingsContext.
        /// </summary>
        /// <param name="connectionString">connection string to db</param>
        /// <param name="logger">logger.</param>
        /// <returns>context</returns>
        public static new InfoContext Create(
            string connectionString,
            ILogger logger)
        {
            if (_options == null)
                _options = CreateOptions<InfoContext>(connectionString, logger);

            return new InfoContext(_options, logger);
        }
        
        /// <summary>
        /// creates db options
        /// </summary>
        /// <typeparam name="T">db context type</typeparam>
        /// <param name="connectionString">connection string to db</param>
        /// <param name="logger">logger</param>
        /// <returns>db options</returns>
        protected static new DbContextOptions<T> CreateOptions<T>(
            string connectionString,
            ILogger logger)
            where T : CommonDAL.DbContextBase
        {
            return new DbContextOptionsBuilder<T>()
                       .UseSqlite(
                            connectionString,
                            options =>
                            {
                            }
                       )
                       .Options;
        }

        /// <inheritdoc/>
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<AppParam>().ToTable("appparam");
            modelBuilder.Entity<Dependency>().ToTable("dependencies");
            modelBuilder.Entity<AtlasSW>().ToTable("atlas");
            modelBuilder.Entity<Detector>().ToTable("detector");
            modelBuilder.Entity<DicomPrinter>().ToTable("dicom_printers");
            modelBuilder.Entity<HardDrive>().ToTable("hard_drive");
            modelBuilder.Entity<HospitalInfo>().ToTable("hospital_info");
            modelBuilder.Entity<Lan>().ToTable("lan");
            modelBuilder.Entity<LogicalDisk>().ToTable("logical_disks");
            modelBuilder.Entity<Modem>().ToTable("modem");
            modelBuilder.Entity<Motherboard>().ToTable("motherboard");
            modelBuilder.Entity<OsInfo>().ToTable("os_info");
            modelBuilder.Entity<SqlDatabase>().ToTable("sql_databases");
            modelBuilder.Entity<SqlService>().ToTable("sql_service");
            modelBuilder.Entity<VideoAdapter>().ToTable("videoadapter");
            modelBuilder.Entity<News>().ToTable("news");

            base.OnModelCreating(modelBuilder);
        }
    }
}
