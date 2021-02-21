using Atlas.Acquisitions.Common.Core.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MessagesSender.Core.Interfaces
{
    /// <summary>
    /// hardware state service interface
    /// </summary>
    public interface IHardwareStateService
    {
        /// <summary>
        /// gets stand state for sending to mqtt
        /// </summary>
        /// <param name="state">state</param>
        /// <returns>state object</returns>
        object GetStandState(StandState state);

        /// <summary>
        /// gets generator state for sending to mqtt
        /// </summary>
        /// <param name="state">state</param>
        /// <returns>state object</returns>
        object GetGeneratorState(GeneratorState state);

        /// <summary>
        /// gets collimator state for sending to mqtt
        /// </summary>
        /// <param name="state">state</param>
        /// <returns>state object</returns>
        object GetCollimatorState(CollimatorState state);
    }
}
