package interfaces

// IIoCProvider describes IoC provider interface
type IIoCProvider interface {
	GetMqttReceiverService() IMqttReceiverService
	GetWebSocketService() IWebSocketService
	GetDalService() IDalService
	GetEquipsService() IEquipsService
	GetHTTPService() IHttpService
	GetWebSocket() IWebSock
	GetMqttClient() IMqttClient
}
