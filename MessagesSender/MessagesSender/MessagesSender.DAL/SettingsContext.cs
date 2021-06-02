using MessagesSender.DAL.Model;
using Microsoft.EntityFrameworkCore;
using Serilog;
using CommonDAL = Atlas.Common.DAL;

namespace MessagesSender.DAL
{
    /// <summary>
    /// Settings db cntext.
    /// </summary>
    public partial class SettingsContext : CommonDAL.SettingsContext
    {
        private static DbContextOptions<SettingsContext> _options = null;

        /// <summary>
        /// Initializes a new instance of the <see cref="SettingsContext"/> class.
        /// </summary>
        /// <param name="options">options.</param>
        /// <param name="logger">logger.</param>
        public SettingsContext(
            DbContextOptions options,
            ILogger logger)
        : base(options, logger)
        {
        }

        /// <summary>
        /// set of EquipmentDicomParam.
        /// </summary>
        public virtual DbSet<EquipmentDicomParam> EquipmentDicomParams { get; set; }

        /// <summary>
        /// Creates SettingsContext.
        /// </summary>
        /// <param name="connectionString">connection string to db</param>
        /// <param name="logger">logger.</param>
        /// <returns>context</returns>
        public static new SettingsContext Create(
            string connectionString,
            ILogger logger)
        {
            if (_options == null)
                _options = CreateOptions<SettingsContext>(connectionString, logger);

            return new SettingsContext(_options, logger);
        }

        /// <inheritdoc/>
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            base.OnModelCreating(modelBuilder);
        }
    }
}
