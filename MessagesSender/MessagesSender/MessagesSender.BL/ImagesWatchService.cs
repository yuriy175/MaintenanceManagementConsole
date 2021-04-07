using System;
using System.Threading.Tasks;
using Atlas.Remoting.BusWrappers.RabbitMQ.Model;
using Atlas.Remoting.Core.Interfaces;
using MessagesSender.Core.Interfaces;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;
using Serilog;
using Atlas.Common.Impls.Helpers;
using System.Net;
using System.Linq;
using System.Net.Sockets;
using Atlas.Acquisitions.Common.Core;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System.Text;
using MessagesSender.Core.Model;
using Atlas.Acquisitions.Common.Core.Model;
using System.Collections.Generic;
using Atlas.Common.Core;

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
                _mqService.Subscribe<MQCommands, int>(
                        (MQCommands.NewImageCreated, async imageId => OnNewImageCreatedAsync(imageId)));
            });
        }
		
        private async Task<bool> OnNewImageCreatedAsync(int imageId)
        {
			++_imageCount;
			SendImagesInfoAsync();

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
					} : null,
				});
		}
	}
}
