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

            FillConfig();
        }

        /// <summary>
        /// get config parameters
        /// </summary>
        /// <returns>config parameters</returns>
        public async Task<IEnumerable<ConfigParam>> GetConfigParamAsync()
        {
            return await GetManyAction<ConfigParam>(context => context.ConfigParams);
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
        /// inserts or updates config parameter
        /// </summary>
        /// <typeparam name="T">value type</typeparam>
        /// <param name="configParamName">config parameter name</param>
        /// <param name="value">value</param>
        /// <returns>new config parameter</returns>
        public async Task<ConfigParam> UpsertConfigParamAsync<T>(string configParamName, T value)
        {
            ConfigParam configParam = null;
            var paramVal = value.ToString();

            await UpsertAction(
                context =>
                {
                    return context.ConfigParams.FirstOrDefault(a => a.ParamName == configParamName);
                },
                context =>
                {
                    configParam = new ConfigParam
                    {
                        ParamName = configParamName,
                        ParamValue = paramVal,
                        Comment = configParamName,
                    };
                    context.ConfigParams.Add(configParam);
                },
                (context, dbConfigParam) =>
                {
                    configParam = dbConfigParam;
                    dbConfigParam.ParamValue = paramVal;
                });

            return configParam;
        }

        /// <summary>
        /// get config parameter by name
        /// </summary>
        /// <param name="configParamName">config parameter name</param>
        /// <returns>config parameter</returns>
        public async Task<ConfigParam> GetConfigParamAsync(string configParamName)
        {
            return await GetAction<ConfigParam>(
                            context => context.ConfigParams.FirstOrDefault(x => x.ParamName == configParamName));
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

        /// <summary>
        /// Fills configuration service from DB.
        /// </summary>
        private void FillConfig()
        {
            using (var context = CreateContext())
            {
                var configParams = context.ConfigParams.ToDictionary(a => a.ParamName.Trim(), a => a.ParamValue);
                _configurationService?.AppendConfigParams(configParams);
            }
        }
    }
}
