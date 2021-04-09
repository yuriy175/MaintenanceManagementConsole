package BL

import (
	"fmt"
	"net/http"

	"../DAL"
	"../Interfaces"
	"../Models"
)

type IHttpService interface {
	Start()
}

type httpService struct {
	_authService         Interfaces.IAuthService
	_mqttReceiverService IMqttReceiverService
	_webSocketService    IWebSocketService
	_dalService          DAL.IDalService
	_equipController     *EquipController
	_adminController     *AdminController
}

func HttpServiceNew(
	mqttReceiverService IMqttReceiverService,
	webSocketService IWebSocketService,
	dalService DAL.IDalService,
	authService Interfaces.IAuthService) IHttpService {
	service := &httpService{}

	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._authService = authService

	service._equipController = EquipControllerNew(mqttReceiverService, webSocketService, dalService, service)
	service._adminController = AdminControllerNew(mqttReceiverService, webSocketService, dalService)

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
