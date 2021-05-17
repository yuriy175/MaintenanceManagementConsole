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
        /// get app parameter by name
        /// </summary>
        /// <returns>all data</returns>
        public async Task<(IEnumerable<HardDrive> HardDrives, IEnumerable<Lan> Lans)> GetAllDataAsync()
		{
            var hardDrives = await GetManyAction<HardDrive>(context => context.HardDrives);
            var lans = await GetManyAction<Lan>(context => context.Lans);
            var logicalDisks = await GetManyAction<LogicalDisk>(context => context.LogicalDisks);
            var modems = await GetManyAction<Modem>(context => context.Modems);
            var monitors = await GetManyAction<Monitor>(context => context.Monitors);
            var motherboards = await GetManyAction<Motherboard>(context => context.Motherboards);

            return (hardDrives, lans);
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
