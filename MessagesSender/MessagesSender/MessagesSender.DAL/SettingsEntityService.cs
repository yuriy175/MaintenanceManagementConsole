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
    /// ISettingsEntityService implementation
    /// </summary>
    public class SettingsEntityService
        : EntityServiceBase<SettingsContext>, ISettingsEntityService
    {
        private readonly IConfigurationService _configurationService = null;
        private readonly ILogger _logger;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger.</param>
        public SettingsEntityService(
            IConfigurationService configurationService,
            ILogger logger)
            : base(logger)
        {
            _configurationService = configurationService;
            _logger = logger;
        }

        /// <summary>
        /// Get equipment info.
        /// </summary>
        /// <returns>equipment info</returns>
        public async Task<(string Name, string Number)> GetEquipmentInfoAsync()
        {
            var dicomTags = new[] { "(0008,1090)", "(0018,1000)" };
            var dicomParams = await GetManyAction<EquipmentDicomParam>(
                context => context.EquipmentDicomParams
                    .Where(p => dicomTags.Contains(p.DicomAttribute))
                    .OrderBy(p => p.Id));

            return dicomParams.Count() != 2 ? 
                (null, null) : 
                (dicomParams.FirstOrDefault().Value, dicomParams.LastOrDefault().Value);
        }

        /// <summary>
        /// Create context.
        /// </summary>
        /// <param name="logger">logger.</param>
        /// <returns>settings context.</returns>
        protected override SettingsContext CreateContext() =>
            SettingsContext.Create(
                    _configurationService?["ConnectionStrings", "SettingsConnection"],
                    _logger);
    }
}
