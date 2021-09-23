package mocks

import (
	"ServerConsole/interfaces"
	"ServerConsole/models"
	Models "ServerConsole/models"
	// "ServerConsole/tests/mocks"
	"ServerConsole/utils"
)

// IoC mock provider implementation type
type mockTypes struct {
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

	// init flag
	_inited bool
}

// mocks instance
var _mockTypes = &mockTypes{}

// InitIoc initializes all services
func InitMockIoc() interfaces.IIoCProvider {
	if !_mockTypes._inited {
		_mockTypes._dalCh = make(chan *models.RawMqttMessage)
		_mockTypes._webSockCh = make(chan *models.RawMqttMessage)
		_mockTypes._equipsCh = make(chan *models.RawMqttMessage)
		_mockTypes._eventsCh = make(chan *models.RawMqttMessage)
		_mockTypes._internalEventsCh = make(chan *models.MessageViewModel)
		_mockTypes._chatCh = make(chan *models.RawMqttMessage)

		_mockTypes._log = utils.LoggerNew()
		// _mockTypes._diagnosticService := bl.DiagnosticServiceNew(log)
		// _mockTypes._authService := bl.AuthServiceNew(log)
		// _mockTypes._topicStorage = utils.TopicStorageNew(_types._log)
		// _mockTypes._settingsService := bl.SettingsServiceNew(log)

		_mockTypes._dalService = DataLayerMockServiceNew()
		_mockTypes._equipsService = EquipsMockServiceNew(_mockTypes._log, _mockTypes._dalService, _mockTypes._equipsCh, _mockTypes._internalEventsCh)
		/*_mockTypes._webSocketService := bl.WebSocketServiceNew(log, _types, diagnosticService, settingsService, webSockCh)
		_mockTypes._eventsService := bl.EventsServiceNew(log, webSocketService, dalService, equipsService, webSockCh,
			eventsCh, internalEventsCh)
		_mockTypes._serverStateService := bl.ServerStateServiceNew(log, dalService)
		_mockTypes._chatService := bl.ChatServiceNew(log, webSocketService, dalService, equipsService, webSockCh, chatCh)
		_mockTypes._mqttReceiverService := bl.MqttReceiverServiceNew(log, _types, diagnosticService, webSocketService, dalService, equipsService, eventsService,
			topicStorage, dalCh, webSockCh, eventsCh)
		_mockTypes._httpService := bl.HTTPServiceNew(log, diagnosticService, settingsService, mqttReceiverService, webSocketService, dalService,
			equipsService, eventsService, chatService, authService, serverStateService)
*/

		_mockTypes._inited = true
	}

	return _mockTypes
}

// GetLogger returns ILogger service
func (t *mockTypes) GetLogger() interfaces.ILogger {
	return t._log
}

// GetDiagnosticService returns IDiagnosticService service
func (t *mockTypes) GetDiagnosticService() interfaces.IDiagnosticService {
	return t._diagnosticService
}

// GetMqttReceiverService returns IMqttReceiverService service
func (t *mockTypes) GetMqttReceiverService() interfaces.IMqttReceiverService {
	return t._mqttReceiverService
}

// GetWebSocketService returns IWebSocketService service
func (t *mockTypes) GetWebSocketService() interfaces.IWebSocketService {
	return t._webSocketService
}

// GetDalService returns IDalService service
func (t *mockTypes) GetDalService() interfaces.IDalService {
	return t._dalService
}

// GetEquipsService returns IEquipsService service
func (t *mockTypes) GetEquipsService() interfaces.IEquipsService {
	return t._equipsService
}

// GetEventsService returns IEventsService service
func (t *mockTypes) GetEventsService() interfaces.IEventsService {
	return t._eventsService
}

// GetHTTPService returns IHttpService service
func (t *mockTypes) GetHTTPService() interfaces.IHttpService {
	return t._httpService
}

// GetWebSocket returns a new IWebSock instance
func (t *mockTypes) GetWebSocket() interfaces.IWebSock {
	return nil // bl.WebSockNew(t._log, t._diagnosticService, t._webSocketService)
}

// GetMqttClient returns a new IMqttClient instance
func (t *mockTypes) GetMqttClient() interfaces.IMqttClient {
	return MqttClientMockNew()
}

// GetChatService returns IHttpService service
func (t *mockTypes) GetChatService() interfaces.IChatService {
	return t._chatService
}
