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
    public class ConfigEntityService
        : EntityServiceBase<ConfigContext>, IConfigEntityService
    {
        private readonly IConfigurationService _configurationService = null;
        private readonly ILogger _logger;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger.</param>
        public ConfigEntityService(
            IConfigurationService configurationService,
            ILogger logger)
            : base(logger)
        {
            _configurationService = configurationService;
            _logger = logger;
        }

        /// <summary>
        /// get offline events
        /// </summary>
        /// <returns>offline events</returns>
        public async Task<IEnumerable<OfflineEvent>> GetOfflineEventsAsync()
        {
            return await GetManyAction<OfflineEvent>(context => context.OfflineEvents);
        }

        /// <summary>
        /// adds offline event
        /// </summary>
        /// <param name="offlineEvent">offline event</param>
        /// <returns>true if success</returns>
        public async Task<bool> AddOfflineEventAsync(OfflineEvent offlineEvent)
        {
            return await AddAction(context => context.OfflineEvents.Add(offlineEvent));
        }

        /// <summary>
        /// delete offline events
        /// </summary>
        /// <param name="before">before datetime</param>
        /// <returns>true if success</returns>
        public async Task<bool> DeleteOfflineEventsAsync(DateTime? before)
        {
            if (!before.HasValue)
            {
                before = DateTime.Now;
            }

            var offlineEvents = await GetManyAction<OfflineEvent>(context => 
                context.OfflineEvents.Where(o => o.MsgDate < before));
            return await DeleteAction(context => context.OfflineEvents.RemoveRange(offlineEvents));
        }

        /// <summary>
        /// Create context.
        /// </summary>
        /// <param name="logger">logger.</param>
        /// <returns>settings context.</returns>
        protected override ConfigContext CreateContext() =>
            ConfigContext.Create(
                    _configurationService?["ConnectionStrings", "ConfigConnection"],
                    _logger);
    }
}
