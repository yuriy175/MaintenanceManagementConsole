using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Atlas.Common.Core;
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
    /// IObservationsEntityService implementation.
    /// </summary>
    public class ObservationsEntityService :
        ObservationsEntityServiceBase<ObservationContext>, IObservationsEntityService
    {
        private readonly IConfigurationService _configurationService = null;
        private readonly ILogger _logger;

        /// <summary>
        /// public constructor.
        /// </summary>
        /// <param name="configurationService">configuration service object.</param>
        /// <param name="logger">logger.</param>
        public ObservationsEntityService(
            IConfigurationService configurationService,
            ILogger logger)
            : base(configurationService, logger)
        {
            _configurationService = configurationService;
            _logger = logger;
        }

        /// <summary>
        /// Get study info.
        /// </summary>
        /// <param name="studyId">study id</param>
        /// <returns>study info</returns>
        public async Task<(int StudyId, string StudyDicomUid, string StudyName)?> GetStudyInfoByIdAsync(int studyId)
        {
            var study = await GetAction<Study>(
                context => context.Studies.FirstOrDefault(s => s.Id == studyId));

            return study == null ? null as (int, string, string)? : (study.Id, study.DicomUid, study.Name);
        }

		/// <summary>
		/// Get all image count.
		/// </summary>
		/// <returns>image count</returns>
		public async Task<int> GetImageCountAsync()
		{
			using (var context = CreateContext())
			{
				return await context.Images.CountAsync();
			}
		}

		/// <summary>
		/// Get today's images with types.
		/// </summary>
		/// <returns>images</returns>
		public async Task<IEnumerable<(int Id, ImageTypes Type)>> GetTodayImagesWithTypesAsync()
		{
			var todayStart = new DateTime(DateTime.Now.Year, DateTime.Now.Month, DateTime.Now.Day, 0, 0, 0);
			var todayEnd = new DateTime(todayStart.Year, todayStart.Month, todayStart.Day, 23, 59, 59);
			var images = await GetManyAction<Image>(
				context => context.Images
					.Where(i => todayStart <= i.DateCreation && i.DateCreation <= todayEnd));

			return images?.Select(i => (i.Id, (ImageTypes)i.ImageType));
		}

		/// <summary>
		/// Create context
		/// </summary>
		/// <param name="logger">logger.</param>
		/// <returns>observation context.</returns>
		protected override ObservationContext CreateContext()
        {
            return ObservationContext.Create(
                _configurationService?["ConnectionStrings", "ObservationConnection"],
                _logger) as ObservationContext;
        }
    }
}
