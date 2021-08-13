package bl

import (
	"fmt"
	"net/http"

	"ServerConsole/controllers"
	"ServerConsole/interfaces"
)

// http service implementation type
type httpService struct {
	//logger
	_log interfaces.ILogger

	// authorization service
	_authService interfaces.IAuthService

	// diagnostic service
	_diagnosticService interfaces.IDiagnosticService

	// mqtt receiver service
	_mqttReceiverService interfaces.IMqttReceiverService

	// web socket service
	_webSocketService interfaces.IWebSocketService

	// DAL service
	_dalService interfaces.IDalService

	// settings service
	_settingsService interfaces.ISettingsService

	// http server connection string
	_connectionString string

	// equipment service
	_equipsService interfaces.IEquipsService

	/// events service
	_eventsService interfaces.IEventsService

	// chat service
	_chatService interfaces.IChatService

	// server state service
	_serverStateService interfaces.IServerStateService

	//equipment http controller
	_equipController *controllers.EquipController

	// admin http controller
	_adminController *controllers.AdminController

	// communication http controller
	_chatController *controllers.ChatController

	// server control http controller
	_serverController *controllers.ServerController

	// diagnostic http controller
	_diagnosticController *controllers.DiagnosticController
}

// HTTPServiceNew creates an instance of httpService
func HTTPServiceNew(
	log interfaces.ILogger,
	diagnosticService interfaces.IDiagnosticService,
	settingsService interfaces.ISettingsService,
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	eventsService interfaces.IEventsService,
	chatService interfaces.IChatService,
	authService interfaces.IAuthService,
	serverStateService interfaces.IServerStateService) interfaces.IHttpService {
	service := &httpService{}

	service._log = log
	service._diagnosticService = diagnosticService
	service._settingsService = settingsService
	service._connectionString = settingsService.GetHTTPServerConnectionString()
	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._equipsService = equipsService
	service._eventsService = eventsService
	service._chatService = chatService
	service._authService = authService
	service._serverStateService = serverStateService

	service._equipController = controllers.EquipControllerNew(log, diagnosticService, mqttReceiverService, webSocketService, dalService, equipsService, eventsService, service, authService)
	service._adminController = controllers.AdminControllerNew(log, diagnosticService, mqttReceiverService, webSocketService, dalService, authService)
	service._chatController = controllers.ChatControllerNew(log, diagnosticService, mqttReceiverService, webSocketService, dalService, chatService, service, authService)
	service._serverController = controllers.ServerControllerNew(log, diagnosticService, service, serverStateService, authService)
	service._diagnosticController = controllers.DiagnosticControllerNew(log, authService)

	return service
}

// Starts the service
func (service *httpService) Start() {

	service._equipController.Handle()
	service._adminController.Handle()
	service._chatController.Handle()
	service._serverController.Handle()
	service._diagnosticController.Handle()

	address := service._connectionString // models.HTTPServerAddress
	fmt.Println("http server is listening... " + address)
	http.ListenAndServe(address, nil)
}
