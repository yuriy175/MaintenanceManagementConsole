package mocks

import (
	"ServerConsole/interfaces"
	"ServerConsole/models"
	Models "ServerConsole/models"

	// "ServerConsole/tests/mocks"
	"ServerConsole/bl"
	"ServerConsole/utils"
)

// IoC mock provider implementation type
type MockTypes struct {
	//logger
	_log interfaces.ILogger

	// output writer
	_outWriter interfaces.IOutputWriter

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
var _mockTypes = &MockTypes{}

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
		_mockTypes._outWriter = OutputMockWriterNew()
		_mockTypes._diagnosticService = DiagnosticMockServiceNew()
		// _mockTypes._authService := bl.AuthServiceNew(log)
		_mockTypes._topicStorage = utils.TopicStorageNew(_mockTypes._log, "../topics.json")
		_mockTypes._settingsService = bl.SettingsServiceNew(_mockTypes._log, "../settings.json")

		_mockTypes._dalService = DataLayerMockServiceNew()
		_mockTypes._equipsService = EquipsMockServiceNew(_mockTypes._log, _mockTypes._dalService, _mockTypes._equipsCh, _mockTypes._internalEventsCh)
		_mockTypes._webSocketService = WebSocketMockServiceNew(_mockTypes._log, _mockTypes, _mockTypes._diagnosticService,
			_mockTypes._settingsService, _mockTypes._webSockCh)
		_mockTypes._eventsService = EventsMockServiceNew(_mockTypes._log, _mockTypes._webSocketService,
			_mockTypes._dalService, _mockTypes._equipsService, _mockTypes._webSockCh, _mockTypes._eventsCh, _mockTypes._internalEventsCh)
		//_mockTypes._serverStateService := bl.ServerStateServiceNew(log, dalService)
		//_mockTypes._chatService := bl.ChatServiceNew(log, webSocketService, dalService, equipsService, webSockCh, chatCh)

		_mockTypes._mqttReceiverService = MqttReceiverMockServiceNew(
			_mockTypes._log, _mockTypes._outWriter,
			_mockTypes, _mockTypes._diagnosticService,
			_mockTypes._webSocketService,
			_mockTypes._dalService, _mockTypes._equipsService,
			_mockTypes._eventsService, _mockTypes._topicStorage,
			_mockTypes._dalCh, _mockTypes._webSockCh, _mockTypes._eventsCh)
		// _mockTypes._httpService := bl.HTTPServiceNew(log, diagnosticService, settingsService, mqttReceiverService, webSocketService, dalService,
		//	equipsService, eventsService, chatService, authService, serverStateService)

		_mockTypes._inited = true
	}

	return _mockTypes
}

// GetLogger returns ILogger service
func (t *MockTypes) GetLogger() interfaces.ILogger {
	return t._log
}

// GetDiagnosticService returns IDiagnosticService service
func (t *MockTypes) GetDiagnosticService() interfaces.IDiagnosticService {
	return t._diagnosticService
}

// GetMqttReceiverService returns IMqttReceiverService service
func (t *MockTypes) GetMqttReceiverService() interfaces.IMqttReceiverService {
	return t._mqttReceiverService
}

// CreateMqttReceiverService returns IMqttReceiverService service
func (t *MockTypes) CreateMqttReceiverService() interfaces.IMqttReceiverService {
	mqttReceiverService := MqttReceiverMockServiceNew(
		_mockTypes._log, _mockTypes._outWriter,
		_mockTypes, _mockTypes._diagnosticService,
		_mockTypes._webSocketService,
		_mockTypes._dalService,
		// _mockTypes._equipsService,
		t.CreateEquipsService(),
		_mockTypes._eventsService, _mockTypes._topicStorage,
		_mockTypes._dalCh, _mockTypes._webSockCh, _mockTypes._eventsCh)

	return mqttReceiverService
}

// CreateMqttReceiverServiceAndEquipsService returns IMqttReceiverService and IEquipsService services
func (t *MockTypes) CreateMqttReceiverServiceAndEquipsService() (
	interfaces.IMqttReceiverService,
	interfaces.IEquipsService,
	chan *models.RawMqttMessage) {
	equipsCh := make(chan *models.RawMqttMessage)
	equipsService := EquipsMockServiceNew(_mockTypes._log, _mockTypes._dalService, equipsCh, _mockTypes._internalEventsCh)
	mqttReceiverService := MqttReceiverMockServiceNew(
		_mockTypes._log, _mockTypes._outWriter,
		_mockTypes, _mockTypes._diagnosticService,
		_mockTypes._webSocketService,
		_mockTypes._dalService,
		equipsService,
		_mockTypes._eventsService, _mockTypes._topicStorage,
		_mockTypes._dalCh, _mockTypes._webSockCh, _mockTypes._eventsCh)

	return mqttReceiverService, equipsService, equipsCh
}

// GetWebSocketService returns IWebSocketService service
func (t *MockTypes) GetWebSocketService() interfaces.IWebSocketService {
	return t._webSocketService
}

// GetDalService returns IDalService service
func (t *MockTypes) GetDalService() interfaces.IDalService {
	return t._dalService
}

// GetEquipsService returns IEquipsService service
func (t *MockTypes) GetEquipsService() interfaces.IEquipsService {
	return t._equipsService
}

// CreateEquipsService returns IEquipsService service
func (t *MockTypes) CreateEquipsService() interfaces.IEquipsService {
	equipsService := EquipsMockServiceNew(_mockTypes._log, _mockTypes._dalService, _mockTypes._equipsCh, _mockTypes._internalEventsCh)
	return equipsService
}

// GetEventsService returns IEventsService service
func (t *MockTypes) GetEventsService() interfaces.IEventsService {
	return t._eventsService
}

// GetHTTPService returns IHttpService service
func (t *MockTypes) GetHTTPService() interfaces.IHttpService {
	return t._httpService
}

// GetWebSocket returns a new IWebSock instance
func (t *MockTypes) GetWebSocket() interfaces.IWebSock {
	return WebSockMockNew()
}

// GetMqttClient returns a new IMqttClient instance
func (t *MockTypes) GetMqttClient() interfaces.IMqttClient {
	return MqttClientMockNew()
}

// GetChatService returns IHttpService service
func (t *MockTypes) GetChatService() interfaces.IChatService {
	return t._chatService
}

// GetEquipsChan returns equips channel
func (t *MockTypes) GetEquipsChan() chan *Models.RawMqttMessage {
	return t._equipsCh
}
