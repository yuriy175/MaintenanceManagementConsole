package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	// "fmt"

	"ServerConsole/interfaces"
	"ServerConsole/models"
)

// AdminController describes admin controller implementation type
type AdminController struct {
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

	// authorization service
	_authService interfaces.IAuthService
}

// AdminControllerNew creates an instance of AdminController
func AdminControllerNew(
	log interfaces.ILogger,
	diagnosticService interfaces.IDiagnosticService,
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	authService interfaces.IAuthService) *AdminController {
	service := &AdminController{}

	service._log = log
	service._diagnosticService = diagnosticService
	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._authService = authService

	return service
}

// Handle handles incomming requests
/*func (service *AdminController) Handle() {
	dalService := service._dalService
	authService := service._authService
	log := service._log

	http.HandleFunc("/equips/GetAllUsers", func(w http.ResponseWriter, r *http.Request) {

		if CheckAdminAuthorization(authService, w, r) != nil {
			users := dalService.GetUsers()
			json.NewEncoder(w).Encode(users)
		}
	})

	http.HandleFunc("/equips/UpdateUser", func(w http.ResponseWriter, r *http.Request) {

		claims := CheckAdminAuthorization(authService, w, r)

		if claims == nil {
			return
		}

		defer r.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(r.Body)

		var userVM = &models.UserViewModel{}
		json.Unmarshal(bodyBytes, &userVM)

		user := dalService.UpdateUser(userVM)
		json.NewEncoder(w).Encode(user)

		log.Infof("User %s updated by %s", userVM.Login, claims.Login)
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
			service._log.Infof("login error: %s", userVM.Login)
			return
		}

		service._log.Infof("login success: %s", userVM.Login)
		token, userName := service._authService.CreateToken(user)
		json.NewEncoder(w).Encode(models.TokenWithUserViewModel{token, userName})
	})

	http.HandleFunc("/equips/GetServerLogs", func(w http.ResponseWriter, r *http.Request) {
		if CheckAdminAuthorization(authService, w, r) != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Encoding", "gzip")

			ok := log.WriteZipContent(w)
			if !ok {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	})
}*/

// GetServerLogs returns server logs
func (service *AdminController) GetServerLogs(w http.ResponseWriter, r *http.Request) {
	log := service._log
	authService := service._authService

	if CheckAdminAuthorization(authService, w, r) != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")

		ok := log.WriteZipContent(w)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

// Login logins a user
func (service *AdminController) Login(w http.ResponseWriter, r *http.Request) {
	log := service._log
	dalService := service._dalService

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
		log.Infof("login error: %s", userVM.Login)
		return
	}

	log.Infof("login success: %s", userVM.Login)
	token, userName := service._authService.CreateToken(user)
	json.NewEncoder(w).Encode(models.TokenWithUserViewModel{token, userName})
}

// GetAllUsers returns all users
func (service *AdminController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	dalService := service._dalService

	if CheckAdminAuthorization(authService, w, r) != nil {
		users := dalService.GetUsers()
		json.NewEncoder(w).Encode(users)
	}
}

// UpdateUser updates a user
func (service *AdminController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log := service._log
	authService := service._authService
	dalService := service._dalService

	claims := CheckAdminAuthorization(authService, w, r)

	if claims == nil {
		return
	}

	defer r.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(r.Body)

	var userVM = &models.UserViewModel{}
	json.Unmarshal(bodyBytes, &userVM)

	user := dalService.UpdateUser(userVM)
	json.NewEncoder(w).Encode(user)
	log.Infof("User %s updated by %s", userVM.Login, claims.Login)
}
