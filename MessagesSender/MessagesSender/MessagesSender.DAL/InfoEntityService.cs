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
		/// <param name="appParam">app parameter name</param>
		/// <returns>app parameter</returns>
		public async Task<AppParam> GetAppParamAsync(string appParam)
		{
            var tt2 = await GetManyAction<Dependency>(context => context.Dependencies);

            var tt3 = await GetManyAction<AspNetUser>(context => context.AspNetUsers);

            var tt31 = await GetManyAction<AtlasSW>(context => context.Atlases);
            var tt311 = await GetManyAction<HospitalInfo>(context => context.HospitalInfos);

            var tt =  await GetManyAction<AppParam>(context => context.AppParams);
            return tt?.FirstOrDefault();
            // await GetAction<AppParam>(
            //    context => context.AppParams);
            // return await GetAction<AppParams>(
            //				context => context.AppParams.FirstOrDefault(x => x.ParamName == appParam));
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
