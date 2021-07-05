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
    /// IMasterEntityService implementation
    /// </summary>
    public class MasterEntityService
        : EntityServiceBase<MasterContext>, IMasterEntityService
    {
        private readonly IConfigurationService _configurationService = null;
        private readonly ILogger _logger;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger.</param>
        public MasterEntityService(
            IConfigurationService configurationService,
            ILogger logger)
            : base(logger)
        {
            _configurationService = configurationService;
            _logger = logger;
        }

        /// <summary>
        /// Get databases states.
        /// </summary>
        /// <returns>databases states</returns>
        public async Task<IEnumerable<(string Name, string State)>> GetDatabasesStatesAsync()
        {
            var query = "select name, state_desc from sys.databases";

            try
            {
                var states = (await ExecuteQueryAsync(query))?
                    .Select(i => (i.FirstOrDefault().ToString(), i.LastOrDefault().ToString()));

                return states?.ToArray();
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "GetDatabasesStatesAsync");
                return null;
            }
        }

        /// <summary>
        /// Create context.
        /// </summary>
        /// <param name="logger">logger.</param>
        /// <returns>settings context.</returns>
        protected override MasterContext CreateContext() =>
            MasterContext.Create(
                    _configurationService?["ConnectionStrings", "MasterConnection"],
                    _logger);
    }
}
