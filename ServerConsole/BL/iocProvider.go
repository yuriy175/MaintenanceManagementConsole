package bl

import (
	"../dal"
	"../interfaces"
	"../models"
	Models "../models"
	"../utils"
)

type types struct {
	_authService         interfaces.IAuthService
	_mqttReceiverService interfaces.IMqttReceiverService
	_webSocketService    interfaces.IWebSocketService
	_dalService          interfaces.IDalService
	_httpService         interfaces.IHttpService
	_topicStorage        interfaces.ITopicStorage
	_settingsService     interfaces.ISettingsService
	_equipsService       interfaces.IEquipsService
	_dalCh               chan *Models.RawMqttMessage
	_webSockCh           chan *Models.RawMqttMessage
	_equipsCh            chan *Models.RawMqttMessage
}

var _types = &types{}

func InitIoc() interfaces.IIoCProvider {
	dalCh := make(chan *models.RawMqttMessage)
	webSockCh := make(chan *models.RawMqttMessage)
	equipsCh := make(chan *models.RawMqttMessage)

	authService := AuthServiceNew()
	topicStorage := utils.TopicStorageNew()
	settingsService := SettingsServiceNew()

	dalService := dal.DalServiceNew(authService, settingsService, dalCh)
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

func (t *types) GetMqttReceiverService() interfaces.IMqttReceiverService {
	return t._mqttReceiverService
}

func (t *types) GetWebSocketService() interfaces.IWebSocketService {
	return t._webSocketService
}

func (t *types) GetDalService() interfaces.IDalService {
	return t._dalService
}

func (t *types) GetEquipsService() interfaces.IEquipsService {
	return t._equipsService
}

func (t *types) GetHttpService() interfaces.IHttpService {
	return t._httpService
}

func (t *types) GetWebSocket() interfaces.IWebSock {
	return WebSockNew(t._webSocketService)
}

func (t *types) GetMqttClient() interfaces.IMqttClient {
	return MqttClientNew(t._settingsService, t._mqttReceiverService, t._webSocketService, t._dalCh, t._webSockCh, t._equipsCh)
}
