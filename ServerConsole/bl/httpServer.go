package bl

import (
	"fmt"
	"net/http"

	"../controllers"
	"../interfaces"
	"../models"
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
	_equipsService   interfaces.IEquipsService
	_equipController *controllers.EquipController
	_adminController *controllers.AdminController
}

// HTTPServiceNew creates an instance of httpService
func HTTPServiceNew(
	log interfaces.ILogger,
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	authService interfaces.IAuthService) interfaces.IHttpService {
	service := &httpService{}

	service._log = log
	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._equipsService = equipsService
	service._authService = authService

	service._equipController = controllers.EquipControllerNew(log, mqttReceiverService, webSocketService, dalService, equipsService, service)
	service._adminController = controllers.AdminControllerNew(log, mqttReceiverService, webSocketService, dalService)

	return service
}

// Starts the service
func (service *httpService) Start() {
	// mqttReceiverService := service._mqttReceiverService
	// webSocketService := service._webSocketService
	// dalService := service._dalService

	service._equipController.Handle()
	service._adminController.Handle()

	address := models.HTTPServerAddress
	fmt.Println("http server is listening... " + address)
	http.ListenAndServe(address, nil)
}