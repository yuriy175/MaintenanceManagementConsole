package Interfaces

type IIoCProvider interface {
	GetMqttReceiverService() IMqttReceiverService
	GetWebSocketService() IWebSocketService
	GetDalService() IDalService
	GetHttpService() IHttpService
	GetWebSocket() IWebSock
	GetMqttClient() IMqttClient
}
