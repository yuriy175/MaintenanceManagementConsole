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
	_authService         interfaces.IAuthService
	_mqttReceiverService interfaces.IMqttReceiverService
	_webSocketService    interfaces.IWebSocketService
	_dalService          interfaces.IDalService
	_equipsService       interfaces.IEquipsService
	_equipController     *controllers.EquipController
	_adminController     *controllers.AdminController
}

// Creates an instance of httpService
func HttpServiceNew(
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	authService interfaces.IAuthService) interfaces.IHttpService {
	service := &httpService{}

	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._equipsService = equipsService
	service._authService = authService

	service._equipController = controllers.EquipControllerNew(mqttReceiverService, webSocketService, dalService, equipsService, service)
	service._adminController = controllers.AdminControllerNew(mqttReceiverService, webSocketService, dalService)

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
