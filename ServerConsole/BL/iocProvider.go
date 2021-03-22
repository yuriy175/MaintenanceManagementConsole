package BL

import (
	"../DAL"
	"../Models"
)

type IIoCProvider interface {
	GetMqttReceiverService() IMqttReceiverService
	GetWebSocketService() IWebSocketService
	GetDalService() DAL.IDalService
	GetHttpService() IHttpService
	GetWebSocket() IWebSock
}

type types struct {
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

	webSocketService := WebSocketServiceNew(_types, webSockCh)
	mqttReceiverService := MqttReceiverServiceNew(webSocketService, dalCh, webSockCh)
	dalService := DAL.DalServiceNew(dalCh)
	httpService := HttpServiceNew(mqttReceiverService, webSocketService, dalService)

	_types = &types{mqttReceiverService, webSocketService, dalService, httpService, dalCh, webSockCh}

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
	return WebSockNew()
}
