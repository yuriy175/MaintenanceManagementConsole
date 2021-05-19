package interfaces

// IoC provider interface
type IIoCProvider interface {
	GetMqttReceiverService() IMqttReceiverService
	GetWebSocketService() IWebSocketService
	GetDalService() IDalService
	GetEquipsService() IEquipsService
	GetHttpService() IHttpService
	GetWebSocket() IWebSock
	GetMqttClient() IMqttClient
}
