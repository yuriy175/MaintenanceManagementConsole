using MessagesSender.DAL.Model;
using Microsoft.EntityFrameworkCore;
using Serilog;
using CommonDAL = Atlas.Common.DAL;
using MessagesSenderDAL = MessagesSender.DAL.Model;

namespace MessagesSender.DAL
{    
    /// <summary>
    /// mssql master db context.
    /// </summary>
    public partial class MasterContext : CommonDAL.DbContextBase
    {
        private static DbContextOptions<MasterContext> _options = null;

        /// <summary>
        /// Initializes a new instance of the <see cref="MasterContext"/> class.
        /// </summary>
        /// <param name="options">options.</param>
        /// <param name="logger">logger.</param>
        public MasterContext(
            DbContextOptions options,
            ILogger logger)
        : base(options, logger)
        {
        }

        /// <summary>
        /// Creates SettingsContext.
        /// </summary>
        /// <param name="connectionString">connection string to db</param>
        /// <param name="logger">logger.</param>
        /// <returns>context</returns>
        public static new MasterContext Create(
            string connectionString,
            ILogger logger)
        {
            if (_options == null)
                _options = CreateOptions<MasterContext>(connectionString, logger);

            return new MasterContext(_options, logger);
        }
        
        /// <inheritdoc/>
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            base.OnModelCreating(modelBuilder);
        }
    }
}
