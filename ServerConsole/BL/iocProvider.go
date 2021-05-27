package bl

import (
	"../dal"
	"../interfaces"
	"../models"
	Models "../models"
	"../utils"
)

// IoC provider implementation type
type types struct {
	// authorization service
	_authService interfaces.IAuthService

	// mqtt receiver service
	_mqttReceiverService interfaces.IMqttReceiverService
	_webSocketService    interfaces.IWebSocketService
	_dalService          interfaces.IDalService
	_httpService         interfaces.IHttpService
	_topicStorage        interfaces.ITopicStorage

	// settings service
	_settingsService interfaces.ISettingsService
	_equipsService   interfaces.IEquipsService

	// chanel for DAL communications
	_dalCh chan *Models.RawMqttMessage

	// chanel for communications with websocket services
	_webSockCh chan *Models.RawMqttMessage

	// chanel for communications with equipment service
	_equipsCh chan *Models.RawMqttMessage
}

// types instance
var _types = &types{}

// InitIoc initializes all services
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
	httpService := HTTPServiceNew(mqttReceiverService, webSocketService, dalService, equipsService, authService)

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

// GetMqttReceiverService returns IMqttReceiverService service
func (t *types) GetMqttReceiverService() interfaces.IMqttReceiverService {
	return t._mqttReceiverService
}

// GetWebSocketService returns IWebSocketService service
func (t *types) GetWebSocketService() interfaces.IWebSocketService {
	return t._webSocketService
}

// GetDalService returns IDalService service
func (t *types) GetDalService() interfaces.IDalService {
	return t._dalService
}

// GetEquipsService returns IEquipsService service
func (t *types) GetEquipsService() interfaces.IEquipsService {
	return t._equipsService
}

// GetHTTPService returns IHttpService service
func (t *types) GetHTTPService() interfaces.IHttpService {
	return t._httpService
}

// GetWebSocket returns a new IWebSock instance
func (t *types) GetWebSocket() interfaces.IWebSock {
	return WebSockNew(t._webSocketService)
}

// GetMqttClient returns a new IMqttClient instance
func (t *types) GetMqttClient() interfaces.IMqttClient {
	return MqttClientNew(t._settingsService, t._mqttReceiverService, t._webSocketService, t._dalCh, t._webSockCh, t._equipsCh)
}
