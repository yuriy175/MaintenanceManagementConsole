package BL

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"../DAL"
	"../Models"
)

func HttpServer(mqttReceiverService *MqttReceiverService, webSocketService *WebSocketService) {

	/*http.HandleFunc("/devices/", func(w http.ResponseWriter, r *http.Request) {
		devices := DAL.DalGetDeviceConnections()
		for _, device := range devices {
			fmt.Fprintf(w, "time %s device : %d name : %s type : %s connection : %d\n",
				device.DateTime.Format(time.RFC3339), device.DeviceId, device.DeviceName, device.DeviceType, device.DeviceConnection)
		}

		fmt.Fprint(w, "Index Page")
	})*/
	//

	http.HandleFunc("/equips/Activate", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		queryString := r.URL.Query()

		sessionUids, ok := queryString["sessionUid"]

		if !ok || len(sessionUids[0]) < 1 {
			log.Println("Url Param 'sessionUid' is missing")
			return
		}
		sessionUid := sessionUids[0]

		activatedEquipInfos, ok := queryString["activatedEquipInfo"]

		if !ok || len(activatedEquipInfos[0]) < 1 {
			log.Println("Url Param 'activatedEquipInfo' is missing")
			return
		}
		activatedEquipInfo := activatedEquipInfos[0]

		deactivatedEquipInfos, ok := queryString["deactivatedEquipInfo"]

		if !ok {
			log.Println("Url Param 'deactivatedEquipInfo' is missing")
			return
		}
		deactivatedEquipInfo := ""
		if len(deactivatedEquipInfos[0]) > 0 {
			deactivatedEquipInfo = deactivatedEquipInfos[0]
		}

		log.Println("Url is: %s %s %s", sessionUid, activatedEquipInfo, deactivatedEquipInfo)
		if deactivatedEquipInfo != "" && deactivatedEquipInfo != activatedEquipInfo {
			mqttReceiverService.SendCommand(deactivatedEquipInfo, "deactivate")
		}
		webSocketService.Activate(sessionUid, activatedEquipInfo, deactivatedEquipInfo)
		mqttReceiverService.SendCommand(activatedEquipInfo, "activate")
	})

	http.HandleFunc("/equips/GetAllEquips", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
		equips := mqttReceiverService.GetConnectionNames()
		json.NewEncoder(w).Encode(equips)
	})

	http.HandleFunc("/equips/RunTeamViewer", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")

		queryString := r.URL.Query()
		activatedEquipInfos, ok := queryString["activatedEquipInfo"]

		if !ok || len(activatedEquipInfos[0]) < 1 {
			log.Println("Url Param 'activatedEquipInfo' is missing")
			return
		}
		activatedEquipInfo := activatedEquipInfos[0]
		mqttReceiverService.SendCommand(activatedEquipInfo, "runTV")
	})

	//(currEquip, startDate, endDate);
	http.HandleFunc("/equips/SearchEquip", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
		queryString := r.URL.Query()

		equipTypes, ok := queryString["equipType"]
		if !ok || len(equipTypes[0]) < 1 {
			log.Println("Url Param 'equipType' is missing")
			return
		}
		equipType := equipTypes[0]

		startDates, ok := queryString["startDate"]
		if !ok || len(startDates[0]) < 1 {
			log.Println("Url Param 'startDate' is missing")
			return
		}

		endDates, ok := queryString["endDate"]
		if !ok {
			log.Println("Url Param 'endDate' is missing")
			return
		}

		startDate, err := time.Parse("2006-01-02", startDates[0])
		if err != nil {
			fmt.Println(err)
		}

		endDate, err2 := time.Parse("2006-01-02", endDates[0])
		if err2 != nil {
			fmt.Println(err)
		}

		log.Println("Url is: %s %s %s", equipType, startDate, endDate)

		SendSearchResults(w, equipType, startDate, endDate)
	})

	address := Models.HttpServerAddress
	fmt.Println("http server is listening... " + address)
	http.ListenAndServe(address, nil)
}

func SendSearchResults(w http.ResponseWriter, equipType string, startDate time.Time, endDate time.Time) {

	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, time.UTC)

	if equipType == "SystemInfo" {
		sysInfo := DAL.GetSystemInfo(startDate, endDate)
		json.NewEncoder(w).Encode(sysInfo)
	} else if equipType == "OrganAutos" {
		organAutos := DAL.GetOrganAutoInfo(startDate, endDate)
		json.NewEncoder(w).Encode(organAutos)
	} else if equipType == "Generators" {
		genInfo := DAL.GetGeneratorInfo(startDate, endDate)
		json.NewEncoder(w).Encode(genInfo)
	} else if equipType == "Studies" {
		studies := DAL.GetStudiesInWork(startDate, endDate)
		json.NewEncoder(w).Encode(studies)
	} else if equipType == "Software" {
		swInfo := DAL.GetSoftwareInfo(startDate, endDate)
		json.NewEncoder(w).Encode(swInfo)
	} else if equipType == "Dicom" {
		dicomInfo := DAL.GetDicomInfo(startDate, endDate)
		json.NewEncoder(w).Encode(dicomInfo)
	} 
}
