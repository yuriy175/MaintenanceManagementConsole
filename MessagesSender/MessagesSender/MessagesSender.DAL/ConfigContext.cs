using System;
using MessagesSender.DAL.Model;
using Microsoft.EntityFrameworkCore;
using Serilog;
using CommonDAL = Atlas.Common.DAL;
using MessagesSenderDAL = MessagesSender.DAL.Model;

namespace MessagesSender.DAL
{
    /// <summary>
    /// config sqllite db context.
    /// </summary>
    public partial class ConfigContext : CommonDAL.DbContextBase
    {
        private static DbContextOptions<ConfigContext> _options = null;

        /// <summary>
        /// Initializes a new instance of the <see cref="ConfigContext"/> class.
        /// </summary>
        /// <param name="options">options.</param>
        /// <param name="logger">logger.</param>
        public ConfigContext(
            DbContextOptions options,
            ILogger logger)
        : base(options, logger)
        {
            try
            {
                var created = Database.EnsureCreated();
                logger.Information($"ConfigContext created {created}");
            }
            catch (Exception ex)
            {
                logger.Error(ex, "Database.EnsureCreated error");
            }
        }

        /// <summary>
        /// set of OfflineEvent
        /// </summary>
        public virtual DbSet<OfflineEvent> OfflineEvents { get; set; }

        /// <summary>
        /// set of ConfigParam
        /// </summary>
        public virtual DbSet<ConfigParam> ConfigParams { get; set; }        

        /// <summary>
        /// Creates ConfigContext.
        /// </summary>
        /// <param name="connectionString">connection string to db</param>
        /// <param name="logger">logger.</param>
        /// <returns>context</returns>
        public static new ConfigContext Create(
            string connectionString,
            ILogger logger)
        {
            if (_options == null)
                _options = CreateOptions<ConfigContext>(connectionString, logger);

            return new ConfigContext(_options, logger);
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
            base.OnModelCreating(modelBuilder);
        }
    }
}
