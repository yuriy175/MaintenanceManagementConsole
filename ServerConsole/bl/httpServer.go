package bl

import (
	"fmt"
	"net/http"

	"../controllers"
	"../interfaces"
)

// http service implementation type
type httpService struct {
	//logger
	_log interfaces.ILogger

	// authorization service
	_authService         interfaces.IAuthService
	_mqttReceiverService interfaces.IMqttReceiverService

	// web socket service
	_webSocketService    interfaces.IWebSocketService

	// DAL service
	_dalService      interfaces.IDalService

	// settings service
	_settingsService interfaces.ISettingsService

	// http server connection string
	_connectionString string 

	// equipment service
	_equipsService   interfaces.IEquipsService

	/// events service
	_eventsService interfaces.IEventsService

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
}

// HTTPServiceNew creates an instance of httpService
func HTTPServiceNew(
	log interfaces.ILogger,
	settingsService interfaces.ISettingsService,
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	eventsService interfaces.IEventsService,
	authService interfaces.IAuthService,
	serverStateService interfaces.IServerStateService) interfaces.IHttpService {
	service := &httpService{}

	service._log = log
	service._settingsService = settingsService
	service._connectionString = settingsService.GetHTTPServerConnectionString();
	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._equipsService = equipsService
	service._eventsService = eventsService
	service._authService = authService
	service._serverStateService = serverStateService

	service._equipController = controllers.EquipControllerNew(log, mqttReceiverService, webSocketService, dalService, equipsService, eventsService, service, authService)
	service._adminController = controllers.AdminControllerNew(log, mqttReceiverService, webSocketService, dalService, authService)
	service._chatController = controllers.ChatControllerNew(log, mqttReceiverService, webSocketService, dalService, service, authService)
	service._serverController = controllers.ServerControllerNew(log, service, serverStateService, authService)
	
	return service
}

// Starts the service
func (service *httpService) Start() {

	service._equipController.Handle()
	service._adminController.Handle()
	service._chatController.Handle()
	service._serverController.Handle()

	address := service._connectionString // models.HTTPServerAddress
	fmt.Println("http server is listening... " + address)
	http.ListenAndServe(address, nil)
}
