package bl

import (
	"ServerConsole/dal"
	"ServerConsole/interfaces"
	"ServerConsole/models"
	Models "ServerConsole/models"
	"ServerConsole/utils"
)

// IoC provider implementation type
type types struct {
	//logger
	_log interfaces.ILogger

	// diagnostic service
	_diagnosticService interfaces.IDiagnosticService

	// authorization service
	_authService interfaces.IAuthService

	// mqtt receiver service
	_mqttReceiverService interfaces.IMqttReceiverService

	// web socket service
	_webSocketService interfaces.IWebSocketService

	// DAL service
	_dalService interfaces.IDalService

	// http service
	_httpService interfaces.IHttpService

	// topic storage
	_topicStorage interfaces.ITopicStorage

	// settings service
	_settingsService interfaces.ISettingsService

	// equipment service
	_equipsService interfaces.IEquipsService

	// events service
	_eventsService interfaces.IEventsService

	// chat service
	_chatService interfaces.IChatService

	// server state service
	_serverStateService interfaces.IServerStateService

	// chanel for DAL communications
	_dalCh chan *Models.RawMqttMessage

	// chanel for communications with websocket services
	_webSockCh chan *Models.RawMqttMessage

	// chanel for communications with equipment service
	_equipsCh chan *Models.RawMqttMessage

	// chanel for communications with events service
	_eventsCh chan *Models.RawMqttMessage

	// chanel for communications with events service (internal events)
	_internalEventsCh chan *models.MessageViewModel

	// chanel for communications with chat service
	_chatCh chan *Models.RawMqttMessage
}

// types instance
var _types = &types{}

// InitIoc initializes all services
func InitIoc() interfaces.IIoCProvider {
	dalCh := make(chan *models.RawMqttMessage)
	webSockCh := make(chan *models.RawMqttMessage)
	equipsCh := make(chan *models.RawMqttMessage)
	eventsCh := make(chan *models.RawMqttMessage)
	internalEventsCh := make(chan *models.MessageViewModel)
	chatCh := make(chan *models.RawMqttMessage)

	log := utils.LoggerNew()
	diagnosticService := DiagnosticServiceNew(log)
	authService := AuthServiceNew(log)
	topicStorage := utils.TopicStorageNew(log, "topics.json")
	settingsService := SettingsServiceNew(log, "settings.json")

	dalService := dal.DataLayerServiceNew(log, authService, settingsService, dalCh)
	equipsService := EquipsServiceNew(log, dalService, equipsCh, internalEventsCh)
	webSocketService := WebSocketServiceNew(log, _types, diagnosticService, settingsService, webSockCh)
	eventsService := EventsServiceNew(log, webSocketService, dalService, equipsService, webSockCh,
		eventsCh, internalEventsCh)
	serverStateService := ServerStateServiceNew(log, dalService)
	chatService := ChatServiceNew(log, webSocketService, dalService, equipsService, webSockCh, chatCh)
	mqttReceiverService := MqttReceiverServiceNew(log, _types, diagnosticService, webSocketService, dalService, equipsService, eventsService,
		topicStorage, dalCh, webSockCh, eventsCh)
	httpService := HTTPServiceNew(log, diagnosticService, settingsService, mqttReceiverService, webSocketService, dalService,
		equipsService, eventsService, chatService, authService, serverStateService)

	_types._log = log
	_types._diagnosticService = diagnosticService
	_types._authService = authService
	_types._mqttReceiverService = mqttReceiverService
	_types._webSocketService = webSocketService
	_types._dalService = dalService
	_types._equipsService = equipsService
	_types._httpService = httpService
	_types._topicStorage = topicStorage
	_types._settingsService = settingsService
	_types._eventsService = eventsService
	_types._chatService = chatService
	_types._dalCh = dalCh
	_types._webSockCh = webSockCh
	_types._equipsCh = equipsCh
	_types._eventsCh = eventsCh
	_types._internalEventsCh = internalEventsCh
	_types._chatCh = chatCh
	_types._serverStateService = serverStateService

	return _types
}

// GetLogger returns ILogger service
func (t *types) GetLogger() interfaces.ILogger {
	return t._log
}

// GetDiagnosticService returns IDiagnosticService service
func (t *types) GetDiagnosticService() interfaces.IDiagnosticService {
	return t._diagnosticService
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

// GetEventsService returns IEventsService service
func (t *types) GetEventsService() interfaces.IEventsService {
	return t._eventsService
}

// GetHTTPService returns IHttpService service
func (t *types) GetHTTPService() interfaces.IHttpService {
	return t._httpService
}

// GetWebSocket returns a new IWebSock instance
func (t *types) GetWebSocket() interfaces.IWebSock {
	return WebSockNew(t._log, t._diagnosticService, t._webSocketService)
}

// GetMqttClient returns a new IMqttClient instance
func (t *types) GetMqttClient() interfaces.IMqttClient {
	return MqttClientNew(t._log, t._diagnosticService, t._settingsService, t._mqttReceiverService, t._webSocketService,
		t._dalCh, t._webSockCh, t._equipsCh, t._eventsCh, t._chatCh)
}

// GetChatService returns IHttpService service
func (t *types) GetChatService() interfaces.IChatService {
	return t._chatService
}
