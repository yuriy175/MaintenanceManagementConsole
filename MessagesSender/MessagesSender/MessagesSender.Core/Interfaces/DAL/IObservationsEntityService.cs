using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Atlas.Common.Core;
using Atlas.Common.DAL.Model;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// Interface to Observations database
    /// </summary>
    public interface IObservationsEntityService
    {
        /// <summary>
        /// Get study info.
        /// </summary>
        /// <param name="studyId">study id</param>
        /// <returns>study info</returns>
        Task<(int StudyId, string StudyDicomUid, string StudyName)?> GetStudyInfoByIdAsync(int studyId);

        /// <summary>
        /// Get all image count.
        /// </summary>
        /// <returns>image count</returns>
        Task<int> GetImageCountAsync();

        /// <summary>
        /// Get today's images with types.
        /// </summary>
        /// <returns>images</returns>
        Task<IEnumerable<(int Id, ImageTypes Type)>> GetTodayImagesWithTypesAsync();
    }
}
