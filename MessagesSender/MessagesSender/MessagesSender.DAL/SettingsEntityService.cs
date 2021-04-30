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
        /// Get dicom info.
        /// </summary>
        /// <returns>dicom info</returns>
        public async Task<IEnumerable<(int Id, string Name, string IP, int ServiceRole)>> GetDicomServicesAsync()
        {
            var dicomServices = await GetManyAction<DicomServices>(
                context => context.DicomServices
                    .Where(p => p.Status)
                    .OrderBy(p => p.Id));

            return dicomServices?.Select(d => (d.Id, d.LogicalName, d.IPAddress, d.ServiceRole));
        }

        /// <summary>
        /// Get hospital info info.
        /// </summary>
        /// <returns>equipment info</returns>
        public async Task<(string Name, string Address, double Latitude, double Longitude)?> GetHospitalInfoAsync()
        {
            var sysInfo = await GetAction<SysInfo>(
                            context => context.SysInfo.FirstOrDefault());

            return sysInfo == null ? 
                null as (string, string, double, double)? : 
                (sysInfo.HospitalName, sysInfo.Address, 10, 20);
        }

        /// <summary>
        /// inserts or updates app parameter
        /// </summary>
        /// <param name="appParam">app parameter name</param>
        /// <param name="value">value</param>
        /// <returns>new app parameter</returns>
        public async Task<AppParams> UpsertAppParamAsync<T>(string appParam, T value)
		{
			AppParams appParams = null;
			var paramVal = value.ToString();

			await UpsertAction(
				context =>
				{
					return context.AppParams.FirstOrDefault(a => a.ParamName == appParam);
				},
				context =>
				{
					appParams = new AppParams
					{
						ParamName = appParam,
						ParamValue = paramVal,
						Comment = appParam,
					};
					context.AppParams.Add(appParams);
				},
				(context, dbAppParams) =>
				{
					appParams = dbAppParams;
					dbAppParams.ParamValue = paramVal;
				});

			return appParams;
		}

		/// <summary>
		/// get app parameter by name
		/// </summary>
		/// <param name="appParam">app parameter name</param>
		/// <returns>app parameter</returns>
		public async Task<AppParams> GetAppParamAsync(string appParam)
		{
			return await GetAction<AppParams>(
							context => context.AppParams.FirstOrDefault(x => x.ParamName == appParam));
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
