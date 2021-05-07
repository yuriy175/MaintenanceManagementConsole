package BL

import (
	"../DAL"
	"../Interfaces"
	"../Models"
	"../Utils"
)

type types struct {
	_authService         Interfaces.IAuthService
	_mqttReceiverService Interfaces.IMqttReceiverService
	_webSocketService    Interfaces.IWebSocketService
	_dalService          Interfaces.IDalService
	_httpService         Interfaces.IHttpService
	_topicStorage        Interfaces.ITopicStorage
	_settingsService     Interfaces.ISettingsService
	_equipsService       Interfaces.IEquipsService
	_dalCh               chan *Models.RawMqttMessage
	_webSockCh           chan *Models.RawMqttMessage
	_equipsCh            chan *Models.RawMqttMessage
}

var _types = &types{}

func InitIoc() Interfaces.IIoCProvider {
	dalCh := make(chan *Models.RawMqttMessage)
	webSockCh := make(chan *Models.RawMqttMessage)
	equipsCh := make(chan *Models.RawMqttMessage)

	authService := AuthServiceNew()
	topicStorage := Utils.TopicStorageNew()
	settingsService := SettingsServiceNew()

	dalService := DAL.DalServiceNew(authService, settingsService, dalCh)
	equipsService := EquipsServiceNew(dalService, equipsCh)
	webSocketService := WebSocketServiceNew(_types, webSockCh)
	mqttReceiverService := MqttReceiverServiceNew(_types, webSocketService, dalService, equipsService, topicStorage, dalCh, webSockCh)
	httpService := HttpServiceNew(mqttReceiverService, webSocketService, dalService, equipsService, authService)

	_types._authService = authService
	_types._mqttReceiverService = mqttReceiverService
	_types._webSocketService = webSocketService
	_types._dalService = dalService
	_types._equipsService = equipsService
	_types._httpService = httpService
	_types._topicStorage = topicStorage
	_types._settingsService = settingsService
	_types._dalCh = dalCh
	_types._webSockCh = webSockCh
	_types._equipsCh = equipsCh

	return _types
}

func (t *types) GetMqttReceiverService() Interfaces.IMqttReceiverService {
	return t._mqttReceiverService
}

func (t *types) GetWebSocketService() Interfaces.IWebSocketService {
	return t._webSocketService
}

func (t *types) GetDalService() Interfaces.IDalService {
	return t._dalService
}

func (t *types) GetEquipsService() Interfaces.IEquipsService {
	return t._equipsService
}

func (t *types) GetHttpService() Interfaces.IHttpService {
	return t._httpService
}

func (t *types) GetWebSocket() Interfaces.IWebSock {
	return WebSockNew(t._webSocketService)
}

func (t *types) GetMqttClient() Interfaces.IMqttClient {
	return MqttClientNew(t._settingsService, t._mqttReceiverService, t._webSocketService, t._dalCh, t._webSockCh, t._equipsCh)
}
