package BL

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../DAL"
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
		//queryString := r.URL.Query()
		defer r.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(r.Body)

		// Convert response body to string
		bodyString := string(bodyBytes)
		fmt.Println("API Response as String:\n" + bodyString)

		log.Println("Url Param 'detailedXilib' is missing")

		/*detailedXilibs, ok := queryString["detailedXilib"]
		if !ok || len(detailedXilibs[0]) < 1 {
			log.Println("Url Param 'detailedXilib' is missing")
			return
		}
		// detailedXilib := detailedXilibs[0]

		verboseXilibs, ok := queryString["verboseXilib"]
		if !ok || len(verboseXilibs[0]) < 1 {
			log.Println("Url Param 'verboseXilib' is missing")
			return
		}
		// verboseXilib := verboseXilibs[0]

		service.SendCommand(w, r, "xilibLogsOn")
		*/
	})
}
