using MessagesSender.DAL.Model;
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
        /// Initializes a new instance of the <see cref="SettingsContext"/> class.
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
        public virtual DbSet<Dicom> Dicoms { get; set; }        

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
        /// Creates SettingsContext.
        /// </summary>
        /// <param name="connectionString">connection string to db</param>
        /// <param name="logger">logger.</param>
        /// <returns>context</returns>
        public new static InfoContext Create(
            string connectionString,
            ILogger logger)
        {
            if (_options == null)
                _options = CreateOptions<InfoContext>(connectionString, logger);

            return new InfoContext(_options, logger);
        }

        /// <inheritdoc/>
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<AppParam>().ToTable("appparam");
            modelBuilder.Entity<Dependency>().ToTable("dependencies");
            modelBuilder.Entity<AtlasSW>().ToTable("atlas");
            modelBuilder.Entity<Detector>().ToTable("detector");
            modelBuilder.Entity<Dicom>().ToTable("dicom");
            modelBuilder.Entity<DicomPrinter>().ToTable("dicom_printers");
            modelBuilder.Entity<HardDrive>().ToTable("hard_drive");
            modelBuilder.Entity<HospitalInfo>().ToTable("hospital_info");
            
            base.OnModelCreating(modelBuilder);
        }

        /// <summary>
        /// creates db options
        /// </summary>
        /// <typeparam name="T">db context type</typeparam>
        /// <param name="connectionString">connection string to db</param>
        /// <param name="logger">logger</param>
        /// <returns>db options</returns>
        protected new static DbContextOptions<T> CreateOptions<T>(
            string connectionString,
            ILogger logger)
            where T : CommonDAL.DbContextBase
        {
            return new DbContextOptionsBuilder<T>()
                       .UseSqlite( 
                                   // "Filename=TestDatabase.db", options =>
                                   connectionString, options =>
                                   {
                                // options.MigrationsAssembly(Assembly.GetExecutingAssembly().FullName);
                            }
                       )
                       .Options;
        }
    }
}
