package BL

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../Models"
)

func HttpServer(mqttReceiverService *MqttReceiverService) {

	/*http.HandleFunc("/devices/", func(w http.ResponseWriter, r *http.Request) {
		devices := DAL.DalGetDeviceConnections()
		for _, device := range devices {
			fmt.Fprintf(w, "time %s device : %d name : %s type : %s connection : %d\n",
				device.DateTime.Format(time.RFC3339), device.DeviceId, device.DeviceName, device.DeviceType, device.DeviceConnection)
		}

		fmt.Fprint(w, "Index Page")
	})

	http.HandleFunc("/studies/", func(w http.ResponseWriter, r *http.Request) {
		devices := DAL.DalGetStudiesInWork()
		for _, study := range devices {
			fmt.Fprintf(w, "time %s study : %d dicom : %s name : %s\n",
				study.DateTime.Format(time.RFC3339), study.StudyId, study.StudyDicomUid, study.StudyName)
		}

		fmt.Fprint(w, "Index Page")
	})

	http.HandleFunc("/commands/", func(w http.ResponseWriter, r *http.Request) {
		mqttReceiverService.SendCommand("KRT/HOMEPC", "runTV")

		fmt.Fprint(w, "Index Page")
	})*/
	//

	http.HandleFunc("/equips/Activate", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		equipInfos, ok := r.URL.Query()["equipInfo"]

		if !ok || len(equipInfos[0]) < 1 {
			log.Println("Url Param 'equipInfo' is missing")
			return
		}
		equipInfo := equipInfos[0]

		mqttReceiverService.SendCommand(equipInfo, "activate")
		log.Println("Url Param 'key' is: " + string(equipInfo))
	})

	http.HandleFunc("/equips/GetAllEquips", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
		equips := mqttReceiverService.GetConnectionNames()
		json.NewEncoder(w).Encode(equips)
	})

	address := Models.HttpServerAddress
	fmt.Println("http server is listening... " + address)
	http.ListenAndServe(address, nil)
}
