package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	interfaces "ServerConsole/interfaces"
)

// ChatController describes communication controller implementation type
type ChatController struct {
	//logger
	_log interfaces.ILogger

	// diagnostic service
	_diagnosticService interfaces.IDiagnosticService

	// mqtt receiver service
	_mqttReceiverService interfaces.IMqttReceiverService

	// web socket service
	_webSocketService interfaces.IWebSocketService

	// DAL service
	_dalService interfaces.IDalService

	// chat service
	_chatService interfaces.IChatService

	// http service
	_httpService interfaces.IHttpService

	// authorization service
	_authService interfaces.IAuthService
}

// ChatControllerNew creates an instance of ChatController
func ChatControllerNew(
	log interfaces.ILogger,
	diagnosticService interfaces.IDiagnosticService,
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	chatService interfaces.IChatService,
	httpService interfaces.IHttpService,
	authService interfaces.IAuthService) *ChatController {
	service := &ChatController{}

	service._log = log
	service._diagnosticService = diagnosticService
	service._httpService = httpService
	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._chatService = chatService
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
	chatService := service._chatService
	diagnosticService := service._diagnosticService
	///
	http.HandleFunc("/equips/GetCommunicationsData", func(w http.ResponseWriter, r *http.Request) {
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
			return
		}

		start := time.Now()
		methodName := "/equips/GetCommunicationsData"
		diagnosticService.IncCount(methodName)

		queryString := r.URL.Query()

		equipName := CheckQueryParameter(queryString, "equipName", w)
		if equipName == "" {
			log.Error("Url Param 'equipName' is missing")
			return
		}

		notes := chatService.GetChatNotes(equipName)
		json.NewEncoder(w).Encode(notes)

		diagnosticService.SetDuration(methodName, time.Since(start))
	})

	http.HandleFunc("/equips/SendNewNote", func(w http.ResponseWriter, r *http.Request) {
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
			return
		}

		queryString := r.URL.Query()

		equipName := CheckQueryParameter(queryString, "equipName", w)
		if equipName == "" {
			log.Error("Url Param 'equipName' is missing")
			return
		}

		msgType := CheckQueryParameter(queryString, "msgType", w)
		if msgType == "" {
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
		message := string(bodyBytes)

		if msgType == "Note" {
			note := dalService.UpsertChatNote(equipName, msgType, id, message, claims.Login, true)
			json.NewEncoder(w).Encode(note)
		} else {
			if id == "" {
				mqttReceiverService.PublishChatNote(equipName, message, claims.Login, true)
			} else {
				note := dalService.UpsertChatNote(equipName, "Chat", id, message, claims.Login, true)
				json.NewEncoder(w).Encode(note)
			}
		}
	})
	http.HandleFunc("/equips/DeleteNote", func(w http.ResponseWriter, r *http.Request) {
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
			return
		}

		queryString := r.URL.Query()

		equipName := CheckQueryParameter(queryString, "equipName", w)
		if equipName == "" {
			log.Error("Url Param 'equipName' is missing")
			return
		}

		msgType := CheckQueryParameter(queryString, "msgType", w)
		if msgType == "" {
			log.Error("Url Param 'msgType' is missing")
			return
		}

		id := CheckQueryParameter(queryString, "id", w)
		if id == "" {
			log.Error("Url Param 'id' is missing")
			return
		}

		dalService.DeleteChatNote(equipName, msgType, id)
	})
}
