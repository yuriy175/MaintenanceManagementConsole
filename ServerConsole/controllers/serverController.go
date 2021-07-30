package controllers

import (
	"encoding/json"
	"net/http"
	"io/ioutil"

	interfaces "../interfaces"
)

// ServerController describes server control controller implementation type
type ServerController struct {
	//logger
	_log interfaces.ILogger

	// mqtt receiver service
	_mqttReceiverService interfaces.IMqttReceiverService

	// web socket service
	_webSocketService    interfaces.IWebSocketService

	// DAL service
	_dalService    interfaces.IDalService

	// http service
	_httpService   interfaces.IHttpService

	// authorization service
	_authService interfaces.IAuthService
}

// ServerControllerNew creates an instance of ServerController
func ServerControllerNew(
	log interfaces.ILogger,
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	httpService interfaces.IHttpService,
	authService interfaces.IAuthService) *ChatController {
	service := &ChatController{}

	service._log = log
	service._httpService = httpService
	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._authService = authService

	return service
}

// Handle handles incomming requests
func (service *ServerController) Handle() {
	mqttReceiverService := service._mqttReceiverService
	dalService := service._dalService
	log := service._log
	authService := service._authService
	///
	http.HandleFunc("/equips/GetServerState", func(w http.ResponseWriter, r *http.Request) {
		claims := CheckUserAuthorization(authService, w, r) 
				
		if claims == nil{
			return
		}
		
		queryString := r.URL.Query()

		equipName := CheckQueryParameter(queryString, "equipName", w) 
		if equipName == ""{
			log.Error("Url Param 'equipName' is missing")
			return
		}
		
		notes := dalService.GetChatNotes(equipName)
		json.NewEncoder(w).Encode(notes)
	})
}

