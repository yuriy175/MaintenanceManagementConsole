using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;
using System.Threading.Tasks;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using MessagesSender.Core.Interfaces;
using MessagesSender.Core.Model;
using MessagesSender.DAL.Model;
using Newtonsoft.Json;
using Serilog;

namespace MessagesSender.BL
{
    /// <summary>
    /// offline service interface implementation
    /// </summary>
    public class OfflineService : IOfflineService
    {
        private readonly IConfigEntityService _dbConfigEntityService;
        private readonly ILogger _logger;

        private readonly List<string> _importantTopics = new List<string>
        {
            MQMessages.SoftwareMsgInfo.ToString(),
            MQMessages.Events.ToString(),
            MQMessages.InstanceOn.ToString(),
        };

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="dbConfigEntityService">config database connector</param>
        /// <param name="logger">logger</param>
        public OfflineService(
            IConfigEntityService dbConfigEntityService,
            ILogger logger)
        {
            _dbConfigEntityService = dbConfigEntityService;
            _logger = logger;

            _logger.Information("Offline service started");
        }

        /// <summary>
        /// 
        /// Check info for offline persistence
        /// </summary>
        /// <typeparam name="TMsgType">message type</typeparam>
        /// <typeparam name="T">info type</typeparam>
        /// <param name="msgType">messge type</param>
        /// <param name="info">info</param>
        /// <returns>result</returns>
        public async Task<bool> CheckInfoAsync<TMsgType, T>(TMsgType msgType, T info)
        {
            if (!_importantTopics.Contains(msgType.ToString()))
            {
                return true;
            }
            
            var type = msgType.ToString();
            var instanceOnType = MQMessages.InstanceOn.ToString();

            OfflineEvent offlineEvent;
            offlineEvent = type switch
            {
                var _ when type == instanceOnType => new OfflineEvent
                {
                    MsgType = MQMessages.SoftwareMsgInfo.ToString(),
                    MsgDate = DateTime.Now,
                    Data = JsonConvert.SerializeObject(new
                    {
                        SimpleMsgType = MQMessages.InstanceOnOffline.ToString(),
                    }),
                },
                _ => new OfflineEvent
                {
                    MsgType = type,
                    MsgDate = DateTime.Now,
                    Data = JsonConvert.SerializeObject(info),
                }
            };

            return await _dbConfigEntityService.AddOfflineEventAsync(offlineEvent);
        }

        /// <summary>
        /// Get offlined infos
        /// </summary>
        /// <returns>offlined infos</returns>
        public async Task<IEnumerable<(string MsgType, object Msg)>> GetInfosAsync()
        {
            var infos = await _dbConfigEntityService.GetOfflineEventsAsync();

            return infos?.Select(i =>
            {
                object value = new
                {
                    OfflineMsg = new
                    {
                        Message = JsonConvert.DeserializeObject(i.Data),
                        DateTime = i.MsgDate,
                    },
                };

                return (i.MsgType, value);
            }).ToList();
        }

        /// <summary>
        /// Clears offline events
        /// </summary>
        /// <returns>result</returns>
        public Task<bool> ClearInfosAsync()
        {
            return _dbConfigEntityService.DeleteOfflineEventsAsync();
        }
    }
}
