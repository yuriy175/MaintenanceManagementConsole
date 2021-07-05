using System;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;
using Atlas.Acquisitions.Common.Core;
using Atlas.Acquisitions.Common.Core.Model;
using Atlas.Common.Core.Interfaces;
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
        private const string MainTopicName = "MainTopic";

        private readonly IConfigurationService _configurationService = null;
        private readonly IConfigEntityService _dbConfigEntityService;
        private readonly ISettingsEntityService _dbSettingsEntityService;  
        private readonly ILogger _logger;

        private readonly Regex _notAllowedSymbolsRegex = new Regex("[^a-zA-Z0-9]");

        private string _mainTopic = string.Empty;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="dbConfigEntityService">config database connector</param>
        /// <param name="dbSettingsEntityService">settings database connector</param>
        /// <param name="logger">logger</param>
        public TopicService(
            IConfigurationService configurationService,
            IConfigEntityService dbConfigEntityService,
            ISettingsEntityService dbSettingsEntityService,
            ILogger logger)
        {
            _configurationService = configurationService;
            _dbConfigEntityService = dbConfigEntityService;
            _dbSettingsEntityService = dbSettingsEntityService;
            _logger = logger;

            GetMainTopicAsync().Wait();

            _logger.Information("Topic service started");
        }

        /// <summary>
        /// gets main topic
        /// </summary>
        /// <returns>result</returns>
        public async Task<string> GetTopicAsync()
        {
            return _mainTopic;
        }

        private async Task GetMainTopicAsync()
        {
            string Strip(string text) => _notAllowedSymbolsRegex.Replace(text, string.Empty);

            var atlasMainTopic = string.Empty;
            var localMainTopic = _configurationService.Get(MainTopicName, string.Empty);
            try
            {
                var equipInfo = await _dbSettingsEntityService.GetEquipmentInfoAsync();
                atlasMainTopic = $"{Strip(equipInfo.Name)}/{Strip(equipInfo.Number)}" +
                    (string.IsNullOrEmpty(equipInfo.HddNumber) ? string.Empty : $"_{Strip(equipInfo.HddNumber)}");
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "GetMainTopicAsync");
            }

            if (localMainTopic == string.Empty && atlasMainTopic == string.Empty)
            {
                _logger.Error("!! No main topics FOUND");
                throw new Exception();
            }

            if (atlasMainTopic == string.Empty)
            {
                _logger.Information("Main topics from sqllite");
                _mainTopic = localMainTopic;
                return;
            }

            _mainTopic = atlasMainTopic;
            if (localMainTopic != atlasMainTopic)
            {
                _ = _dbConfigEntityService.UpsertConfigParamAsync(MainTopicName, atlasMainTopic);
            }
        }
    }
}
