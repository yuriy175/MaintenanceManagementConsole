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

	// authService := service._authService
	serverController := service._serverController
	diagnosticController := service._diagnosticController
	chatController := service._chatController
	adminController := service._adminController
	equipController := service._equipController

	// service._equipController.Handle()
	// service._adminController.Handle()
	// service._chatController.Handle()
	// service._serverController.Handle()
	// service._diagnosticController.Handle()

	http.HandleFunc("/equips/Activate", equipController.ActivateEquip)
	http.HandleFunc("/equips/GetConnectedEquips", equipController.GetConnectedEquips)
	http.HandleFunc("/equips/GetAllEquips", equipController.GetAllEquips)
	http.HandleFunc("/equips/DisableEquipInfo", equipController.DisableEquipInfo)
	http.HandleFunc("/equips/RunTeamViewer", equipController.RunTeamViewer)
	http.HandleFunc("/equips/RunTaskManager", equipController.RunTaskManager)
	http.HandleFunc("/equips/SendAtlasLogs", equipController.SendAtlasLogs)
	http.HandleFunc("/equips/UpdateDBInfo", equipController.UpdateDBInfo)
	http.HandleFunc("/equips/RecreateDBInfo", equipController.RecreateDBInfo)
	http.HandleFunc("/equips/XilibLogsOn", equipController.XilibLogsOn)
	http.HandleFunc("/equips/SetEquipLogsOn", equipController.SetEquipLogsOn)
	http.HandleFunc("/equips/SearchEquip", equipController.SearchEquip)
	http.HandleFunc("/equips/GetAllDBTableNames", equipController.GetAllDBTableNames)
	http.HandleFunc("/equips/GetTableContent", equipController.GetTableContent)
	http.HandleFunc("/equips/GetPermanentData", equipController.GetPermanentData)
	http.HandleFunc("/equips/UpdateEquipDetails", equipController.UpdateEquipDetails)
	http.HandleFunc("/equips/GetEquipInfo", equipController.GetEquipCardInfo)
	http.HandleFunc("/equips/UpdateEquipInfo", equipController.UpdateEquipCardInfo)

	http.HandleFunc("/equips/GetAllUsers", adminController.GetAllUsers)
	http.HandleFunc("/equips/UpdateUser", adminController.UpdateUser)
	http.HandleFunc("/equips/Login", adminController.Login)
	http.HandleFunc("/equips/GetServerLogs", adminController.GetServerLogs)

	http.HandleFunc("/equips/GetCommunicationsData", chatController.GetCommunicationsData)
	http.HandleFunc("/equips/SendNewNote", chatController.UpsertNote)
	http.HandleFunc("/equips/DeleteNote", chatController.DeleteNote)

	http.HandleFunc("/equips/GetServerState", serverController.GetServerState)

	http.HandleFunc("/equips/metrics", diagnosticController.GetServerMetrics)

	address := service._connectionString // models.HTTPServerAddress
	fmt.Println("http server is listening... " + address)
	http.ListenAndServe(address, nil)
}
