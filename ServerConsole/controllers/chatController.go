package controllers

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	
	interfaces "../interfaces"
)

// ChatController describes communication controller implementation type
type ChatController struct {
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

// ChatControllerNew creates an instance of ChatController
func ChatControllerNew(
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
func (service *ChatController) Handle() {
	mqttReceiverService := service._mqttReceiverService
	dalService := service._dalService
	log := service._log
	authService := service._authService
	///
	http.HandleFunc("/equips/GetCommunicationsData", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/equips/SendNewNote", func(w http.ResponseWriter, r *http.Request) {
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

		msgType := CheckQueryParameter(queryString, "msgType", w) 
		if msgType == ""{
			log.Error("Url Param 'msgType' is missing")
			return
		}

		id := CheckOptionalQueryParameter(queryString, "id", w) 

		/*message := CheckQueryParameter(queryString, "message", w) 
		if message == ""{
			log.Error("Url Param 'message' is missing")
			return
		}*/
		defer r.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		message := string (bodyBytes)

		if msgType == "Note"{
			note := dalService.UpsertChatNote(equipName, msgType, id, message, claims.Login)		
			json.NewEncoder(w).Encode(note)
		} else {
			if id == ""{
				mqttReceiverService.PublishChatNote(equipName, message, claims.Login)
			} else {
				dalService.UpsertChatNote(equipName, "Chat", id, message, claims.Login)
			}
		}
	})
	http.HandleFunc("/equips/DeleteNote", func(w http.ResponseWriter, r *http.Request) {
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

		msgType := CheckQueryParameter(queryString, "msgType", w) 
		if msgType == ""{
			log.Error("Url Param 'msgType' is missing")
			return
		}

		id := CheckQueryParameter(queryString, "id", w) 
		if id == ""{
			log.Error("Url Param 'id' is missing")
			return
		}

		dalService.DeleteChatNote(equipName, msgType, id)	
	})
}

