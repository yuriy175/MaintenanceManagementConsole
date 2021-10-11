using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;
using Atlas.Acquisitions.Common.Core;
using Atlas.Acquisitions.Common.Core.Model;
using Atlas.Common.Core;
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
    /// studying watch service
    /// </summary>
    public class ImagesWatchService : IImagesWatchService
    {
        private readonly ILogger _logger;
        private readonly IObservationsEntityService _dbObservationsEntityService;
        private readonly IMQCommunicationService _mqService;
        private readonly IEventPublisher _eventPublisher;
        private readonly ISendingService _sendingService;
        private readonly IWebClientService _webClientService;

        private bool _isActivated = false;
        private int? _imageCount = null;
        private List<(int Id, ImageTypes Type)> _todayImages = null;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="dbObservationsEntityService">observations database connector</param>
        /// <param name="mqService">MQ service</param>
        /// <param name="eventPublisher">event publisher service</param>
        /// <param name="sendingService">sending service</param>
        /// <param name="webClientService">web client service</param>
        public ImagesWatchService(
            ILogger logger,
            IObservationsEntityService dbObservationsEntityService,
            IMQCommunicationService mqService,
            IEventPublisher eventPublisher,
            ISendingService sendingService,
            IWebClientService webClientService)
        {
            _logger = logger;
            _dbObservationsEntityService = dbObservationsEntityService;
            _mqService = mqService;
            _eventPublisher = eventPublisher;
            _sendingService = sendingService;
            _webClientService = webClientService;

            _eventPublisher.RegisterActivateCommandArrivedEvent(() => OnActivateArrivedAsync());
            _eventPublisher.RegisterDeactivateCommandArrivedEvent(() => OnDeactivateArrivedAsync());

            SubscribeMQRecevers();

            _logger.Information("StudyingWatchService started");
        }

        private Task SubscribeMQRecevers()
        {
            return Task.Run(() =>
            {
                _mqService.Subscribe<MQCommands, (int Id, string DicomUID, ImageTypes ImageType)>(
                        (MQCommands.NewImageCreated, async image => await OnNewImageCreatedAsync(image)));
            });
        }
        
        private async Task<bool> OnNewImageCreatedAsync((int Id, string DicomUID, ImageTypes ImageType) image)
        {
            if (!_isActivated)
            {
                return true;
            }

            ++_imageCount;
            _ = SendImagesInfoAsync();

            return true;
        }
        
        private void OnDeactivateArrivedAsync()
        {
            _isActivated = false;
        }

        private async Task<bool> OnActivateArrivedAsync()
        {
            _isActivated = true;
            _imageCount = await _dbObservationsEntityService.GetImageCountAsync();
            _todayImages = (await _dbObservationsEntityService.GetTodayImagesWithTypesAsync())?.ToList();

            SendImagesInfoAsync();

            return false;
        }

        private async Task<bool> SendImagesInfoAsync()
        {
            var imageCount = _imageCount;
            var todayImages = _todayImages;

            return await _sendingService.SendInfoToMqttAsync(
                MQMessages.ImagesInfo,
                new
                {
                    ImageCount = imageCount,
                    Today = todayImages != null ? new
                    {
                        SingleGraphy = todayImages.Count(i => i.Type == ImageTypes.Graphy),
                        Scopy = todayImages.Count(i => i.Type == ImageTypes.Scopy),
                        Stitch = todayImages.Count(i => i.Type == ImageTypes.Stitch),
                    }
                    : null,
                });
        }
    }
}
