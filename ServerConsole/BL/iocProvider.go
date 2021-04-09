package BL

import (
	"../DAL"
	"../Interfaces"
	"../Models"
)

type IIoCProvider interface {
	GetMqttReceiverService() IMqttReceiverService
	GetWebSocketService() IWebSocketService
	GetDalService() DAL.IDalService
	GetHttpService() IHttpService
	GetWebSocket() IWebSock
	GetMqttClient() IMqttClient
}

type types struct {
	_authService         Interfaces.IAuthService
	_mqttReceiverService IMqttReceiverService
	_webSocketService    IWebSocketService
	_dalService          DAL.IDalService
	_httpService         IHttpService
	_dalCh               chan *Models.RawMqttMessage
	_webSockCh           chan *Models.RawMqttMessage
}

var _types = &types{}

func InitIoc() IIoCProvider {
	dalCh := make(chan *Models.RawMqttMessage)
	webSockCh := make(chan *Models.RawMqttMessage)

	authService := AuthServiceNew()
	webSocketService := WebSocketServiceNew(_types, webSockCh)
	mqttReceiverService := MqttReceiverServiceNew(_types, webSocketService, dalCh, webSockCh)
	dalService := DAL.DalServiceNew(authService, dalCh)
	httpService := HttpServiceNew(mqttReceiverService, webSocketService, dalService, authService)

	_types._authService = authService
	_types._mqttReceiverService = mqttReceiverService
	_types._webSocketService = webSocketService
	_types._dalService = dalService
	_types._httpService = httpService
	_types._dalCh = dalCh
	_types._webSockCh = webSockCh

	return _types
}

func (t *types) GetMqttReceiverService() IMqttReceiverService {
	return t._mqttReceiverService
}

func (t *types) GetWebSocketService() IWebSocketService {
	return t._webSocketService
}

func (t *types) GetDalService() DAL.IDalService {
	return t._dalService
}

func (t *types) GetHttpService() IHttpService {
	return t._httpService
}

func (t *types) GetWebSocket() IWebSock {
	return WebSockNew(t._webSocketService)
}

func (t *types) GetMqttClient() IMqttClient {
	return MqttClientNew(t._mqttReceiverService, t._webSocketService, t._dalCh, t._webSockCh)
}
