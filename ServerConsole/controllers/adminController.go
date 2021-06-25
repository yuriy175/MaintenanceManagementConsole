package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../interfaces"
	"../models"
)

// AdminController describes admin controller implementation type
type AdminController struct {
	//logger
	_log interfaces.ILogger

	// mqtt receiver service
	_mqttReceiverService interfaces.IMqttReceiverService

	// web socket service
	_webSocketService    interfaces.IWebSocketService

	// DAL service
	_dalService interfaces.IDalService

	// authorization service
	_authService interfaces.IAuthService
}

// AdminControllerNew creates an instance of webSock
func AdminControllerNew(
	log interfaces.ILogger,
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	authService interfaces.IAuthService) *AdminController {
	service := &AdminController{}

	service._log = log
	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._authService = authService

	return service
}

// Handle handles incomming requests
func (service *AdminController) Handle() {
	//mqttReceiverService := service._mqttReceiverService
	//webSocketService := service._webSocketService
	dalService := service._dalService
	authService := service._authService

	http.HandleFunc("/equips/GetAllUsers", func(w http.ResponseWriter, r *http.Request) {

		/*if CheckOptionsAndSetCORSMethod(w, r){
			return
		}

		tokenString := CheckAuthorization(w, r)
		if tokenString == ""{
			return;
		}
		
		service._authService.VerifyToken(tokenString) */
		if CheckAdminAuthorization(authService, w, r) != nil{
			users := dalService.GetUsers()
			json.NewEncoder(w).Encode(users)
		}
	})

	http.HandleFunc("/equips/UpdateUser", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(r.Body)

		var userVM = &models.UserViewModel{}
		json.Unmarshal(bodyBytes, &userVM)

		user := dalService.UpdateUser(userVM)
		json.NewEncoder(w).Encode(user)
	})

	http.HandleFunc("/equips/Login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(r.Body)

		var userVM = &models.UserViewModel{}
		json.Unmarshal(bodyBytes, &userVM)

		user := dalService.GetUserByName(userVM.Login, userVM.Email, userVM.Password)
		if user == nil {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			service._log.Infof("login error: %s", userVM.Login);
			return
		}

		service._log.Infof("login success: %s", userVM.Login);
		token := service._authService.CreateToken(user)
		json.NewEncoder(w).Encode(token)
	})
}
