package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	interfaces "../interfaces"
)

// EquipController describes equipment controller implementation type
type EquipController struct {
	_mqttReceiverService interfaces.IMqttReceiverService
	_webSocketService    interfaces.IWebSocketService

	// DAL service
	_dalService    interfaces.IDalService
	_equipsService interfaces.IEquipsService

	// http service
	_httpService   interfaces.IHttpService
}

// EquipControllerNew creates an instance of webSock
func EquipControllerNew(
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	httpService interfaces.IHttpService) *EquipController {
	service := &EquipController{}

	service._httpService = httpService
	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._equipsService = equipsService

	return service
}

// Handle handles incomming requests
func (service *EquipController) Handle() {
	mqttReceiverService := service._mqttReceiverService
	webSocketService := service._webSocketService
	equipsService := service._equipsService
	dalService := service._dalService

	// httpService := service._httpService
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
		sessionUID := sessionUids[0]

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

		log.Println("Url is: %s %s %s", sessionUID, activatedEquipInfo, deactivatedEquipInfo)
		if deactivatedEquipInfo != "" && deactivatedEquipInfo != activatedEquipInfo {
			mqttReceiverService.SendCommand(deactivatedEquipInfo, "deactivate")
		}
		webSocketService.Activate(sessionUID, activatedEquipInfo, deactivatedEquipInfo)
		mqttReceiverService.SendCommand(activatedEquipInfo, "activate")
	})

	http.HandleFunc("/equips/GetConnectedEquips", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
		equips := mqttReceiverService.GetConnectionNames()
		json.NewEncoder(w).Encode(equips)
	})

	http.HandleFunc("/equips/GetAllEquips", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
		equipInfos := equipsService.GetEquipInfos()
		json.NewEncoder(w).Encode(equipInfos)
	})

	http.HandleFunc("/equips/RunTeamViewer", func(w http.ResponseWriter, r *http.Request) {
		service.sendCommand(w, r, "runTV")
	})

	http.HandleFunc("/equips/RunTaskManager", func(w http.ResponseWriter, r *http.Request) {
		service.sendCommand(w, r, "runTaskMan")
	})

	http.HandleFunc("/equips/SendAtlasLogs", func(w http.ResponseWriter, r *http.Request) {
		service.sendCommand(w, r, "sendAtlasLogs")
	})

	http.HandleFunc("/equips/XilibLogsOn", func(w http.ResponseWriter, r *http.Request) {
		queryString := r.URL.Query()

		detailedXilibs, ok := queryString["detailedXilib"]
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

		service.sendCommand(w, r, "xilibLogsOn")
	})

	//(currType, equipName, startDate, endDate);
	http.HandleFunc("/equips/SearchEquip", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
		queryString := r.URL.Query()

		equipTypes, ok := queryString["currType"]
		if !ok || len(equipTypes[0]) < 1 {
			log.Println("Url Param 'currType' is missing")
			return
		}
		equipType := equipTypes[0]

		equipNames, ok := queryString["equipName"]
		if !ok {
			log.Println("Url Param 'equipName' is missing")
			return
		}
		equipName := equipNames[0]

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

		service.sendSearchResults(w, equipType, equipName, startDate, endDate)
	})

	http.HandleFunc("/equips/GetAllDBTableNames", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
		queryString := r.URL.Query()

		equipNames, ok := queryString["equipName"]
		if !ok {
			log.Println("Url Param 'equipName' is missing")
			return
		}
		equipName := equipNames[0]
		tables := dalService.GetAllTableNamesInfo(equipName)

		json.NewEncoder(w).Encode(tables)
	})

	http.HandleFunc("/equips/GetTableContent", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
		queryString := r.URL.Query()

		equipNames, ok := queryString["equipName"]
		if !ok {
			log.Println("Url Param 'equipName' is missing")
			return
		}
		equipName := equipNames[0]

		tableTypes, ok := queryString["tableType"]
		if !ok {
			log.Println("Url Param 'tableType' is missing")
			return
		}
		tableType := tableTypes[0]

		tableNames, ok := queryString["tableName"]
		if !ok {
			log.Println("Url Param 'tableName' is missing")
			return
		}
		tableName := tableNames[0]

		tables := dalService.GetTableContent(equipName, tableType, tableName)

		json.NewEncoder(w).Encode(tables)
	})

	http.HandleFunc("/equips/GetPermanentData", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
		queryString := r.URL.Query()

		equipTypes, ok := queryString["currType"]
		if !ok || len(equipTypes[0]) < 1 {
			log.Println("Url Param 'currType' is missing")
			return
		}
		equipType := equipTypes[0]

		equipNames, ok := queryString["equipName"]
		if !ok {
			log.Println("Url Param 'equipName' is missing")
			return
		}
		equipName := equipNames[0]

		log.Println("Url is: %s %s", equipType, equipName)

		service.sendPermanentSearchResults(w, equipType, equipName)
	})

	// address := Models.HttpServerAddress
	// fmt.Println("http server is listening... " + address)
	// http.ListenAndServe(address, nil)
}

func (service *EquipController) sendSearchResults(
	w http.ResponseWriter,
	equipType string,
	equipName string,
	startDate time.Time,
	endDate time.Time) {

	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, time.UTC)

	dalService := service._dalService
	if equipType == "SystemInfo" {
		sysInfo := dalService.GetSystemInfo(equipName, startDate, endDate)
		json.NewEncoder(w).Encode(sysInfo)
	} else if equipType == "OrganAutos" {
		organAutos := dalService.GetOrganAutoInfo(equipName, startDate, endDate)
		json.NewEncoder(w).Encode(organAutos)
	} else if equipType == "Generators" {
		genInfo := dalService.GetGeneratorInfo(equipName, startDate, endDate)
		json.NewEncoder(w).Encode(genInfo)
	} else if equipType == "Studies" {
		studies := dalService.GetStudiesInWork(equipName, startDate, endDate)
		json.NewEncoder(w).Encode(studies)
	} else if equipType == "Software" {
		swInfo := dalService.GetSoftwareInfo(equipName, startDate, endDate)
		json.NewEncoder(w).Encode(swInfo)
	} else if equipType == "Dicom" {
		dicomInfo := dalService.GetDicomInfo(equipName, startDate, endDate)
		json.NewEncoder(w).Encode(dicomInfo)
	} else if equipType == "Stands" {
		standInfo := dalService.GetStandInfo(equipName, startDate, endDate)
		json.NewEncoder(w).Encode(standInfo)
	}
}

func (service *EquipController) sendPermanentSearchResults(
	w http.ResponseWriter,
	equipType string,
	equipName string) {

	dalService := service._dalService
	if equipType == "SystemInfo" {
		// sysInfo := dalService.GetPermanentSystemInfo(equipName)
		sysInfo := dalService.GetDBSystemInfo(equipName)
		json.NewEncoder(w).Encode(sysInfo)
	} else if equipType == "Software" {
		// swInfo := dalService.GetPermanentSoftwareInfo(equipName)
		swInfo := dalService.GetDBSoftwareInfo(equipName)
		json.NewEncoder(w).Encode(swInfo)
	}
}

func (service *EquipController) sendCommand(w http.ResponseWriter, r *http.Request, command string) {
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
	service._mqttReceiverService.SendCommand(activatedEquipInfo, command)
}
