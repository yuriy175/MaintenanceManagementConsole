using System;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;
using Atlas.Acquisitions.Common.Core;
using Atlas.Acquisitions.Common.Core.Model;
using Atlas.Common.Impls.Helpers;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using Atlas.Remoting.Core.Interfaces;
using MessagesSender.Core.Interfaces;
using MessagesSender.Core.Model;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using Serilog;

namespace MessagesSender.BL
{
    /// <summary>
    /// topic service interface implementation
    /// </summary>
    public class TopicService : ITopicService
    {
        private readonly ISettingsEntityService _dbSettingsEntityService;  
        private readonly ILogger _logger;

        private readonly Regex _notAllowedSymbolsRegex = new Regex("[^a-zA-Z0-9]");

        private string _mainTopic = string.Empty;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="dbSettingsEntityService">settings database connector</param>
        /// <param name="logger">logger</param>
        public TopicService(
            ISettingsEntityService dbSettingsEntityService,
            ILogger logger)
        {
            _dbSettingsEntityService = dbSettingsEntityService;
            _logger = logger;

            _logger.Information("Topic service started");
        }

        /// <summary>
        /// gets main topic
        /// </summary>
        /// <returns>result</returns>
        public async Task<string> GetTopicAsync()
        {
            if (string.IsNullOrEmpty(_mainTopic))
            {
                string Strip(string text) => _notAllowedSymbolsRegex.Replace(text, string.Empty);

                var equipInfo = await _dbSettingsEntityService.GetEquipmentInfoAsync();
                _mainTopic = $"{Strip(equipInfo.Name)}/{Strip(equipInfo.Number)}" +
                    (string.IsNullOrEmpty(equipInfo.HddNumber) ? string.Empty : $"_{Strip(equipInfo.HddNumber)}"); 
            }

            return _mainTopic;
        }
    }
}
