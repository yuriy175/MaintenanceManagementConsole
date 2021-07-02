package interfaces

// IIoCProvider describes IoC provider interface
type IIoCProvider interface {
	// GetLogger returns ILogger service
	GetLogger() ILogger 

	// GetMqttReceiverService returns IMqttReceiverService service
	GetMqttReceiverService() IMqttReceiverService

	// GetWebSocketService returns IWebSocketService service
	GetWebSocketService() IWebSocketService

	// GetDalService returns IDalService service
	GetDalService() IDalService

	// GetEquipsService returns IEquipsService service
	GetEquipsService() IEquipsService

	// GetEventsService returns IEventsService service
	GetEventsService() IEventsService

	// GetChatService returns IHttpService service
	GetChatService() IChatService

	// GetHTTPService returns IHttpService service
	GetHTTPService() IHttpService

	// GetWebSocket returns a new IWebSock instance
	GetWebSocket() IWebSock

	// GetMqttClient returns a new IMqttClient instance
	GetMqttClient() IMqttClient
}

