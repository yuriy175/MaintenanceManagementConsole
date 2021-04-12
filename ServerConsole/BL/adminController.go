package BL

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../DAL"
	"../Models"
)

type AdminController struct {
	_mqttReceiverService IMqttReceiverService
	_webSocketService    IWebSocketService
	_dalService          DAL.IDalService
}

func AdminControllerNew(
	mqttReceiverService IMqttReceiverService, webSocketService IWebSocketService, dalService DAL.IDalService) *AdminController {
	service := &AdminController{}

	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService

	return service
}

func (service *AdminController) Handle() {
	//mqttReceiverService := service._mqttReceiverService
	//webSocketService := service._webSocketService
	dalService := service._dalService

	http.HandleFunc("/equips/GetAllUsers", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
		//roles := dalService.GetRoles()
		users := dalService.GetUsers()
		json.NewEncoder(w).Encode(users)
	})

	http.HandleFunc("/equips/UpdateUser", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(r.Body)

		var userVM = &Models.UserViewModel{}
		json.Unmarshal(bodyBytes, &userVM)

		user := dalService.UpdateUser(userVM)
		json.NewEncoder(w).Encode(user)
	})

	http.HandleFunc("/equips/Login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(r.Body)

		var userVM = &Models.UserViewModel{}
		json.Unmarshal(bodyBytes, &userVM)

		user := dalService.GetUserByName(userVM.Login, userVM.Email, userVM.Password)
		if user == nil {
			http.Error(w, "Not authorized", 401)
			return
		}
		json.NewEncoder(w).Encode(user)
	})
}
