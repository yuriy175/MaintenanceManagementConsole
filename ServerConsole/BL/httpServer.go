package BL

import (
	"fmt"
	"net/http"

	"../Controllers"
	"../Interfaces"
	"../Models"
)

type httpService struct {
	_authService         Interfaces.IAuthService
	_mqttReceiverService Interfaces.IMqttReceiverService
	_webSocketService    Interfaces.IWebSocketService
	_dalService          Interfaces.IDalService
	_equipController     *Controllers.EquipController
	_adminController     *Controllers.AdminController
}

func HttpServiceNew(
	mqttReceiverService Interfaces.IMqttReceiverService,
	webSocketService Interfaces.IWebSocketService,
	dalService Interfaces.IDalService,
	authService Interfaces.IAuthService) Interfaces.IHttpService {
	service := &httpService{}

	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._authService = authService

	service._equipController = Controllers.EquipControllerNew(mqttReceiverService, webSocketService, dalService, service)
	service._adminController = Controllers.AdminControllerNew(mqttReceiverService, webSocketService, dalService)

	return service
}

func (service *httpService) Start() {
	// mqttReceiverService := service._mqttReceiverService
	// webSocketService := service._webSocketService
	// dalService := service._dalService

	service._equipController.Handle()
	service._adminController.Handle()

	address := Models.HttpServerAddress
	fmt.Println("http server is listening... " + address)
	http.ListenAndServe(address, nil)
}
