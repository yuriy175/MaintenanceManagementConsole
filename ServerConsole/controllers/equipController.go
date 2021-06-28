package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"strconv"

	interfaces "../interfaces"
)

// EquipController describes equipment controller implementation type
type EquipController struct {
	//logger
	_log interfaces.ILogger

	// mqtt receiver service
	_mqttReceiverService interfaces.IMqttReceiverService

	// web socket service
	_webSocketService    interfaces.IWebSocketService

	// DAL service
	_dalService    interfaces.IDalService

	// equipment service
	_equipsService interfaces.IEquipsService

	// http service
	_httpService   interfaces.IHttpService

	// authorization service
	_authService interfaces.IAuthService
}

// EquipControllerNew creates an instance of webSock
func EquipControllerNew(
	log interfaces.ILogger,
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	httpService interfaces.IHttpService,
	authService interfaces.IAuthService) *EquipController {
	service := &EquipController{}

	service._log = log
	service._httpService = httpService
	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._equipsService = equipsService
	service._authService = authService

	return service
}

// Handle handles incomming requests
func (service *EquipController) Handle() {
	mqttReceiverService := service._mqttReceiverService
	webSocketService := service._webSocketService
	equipsService := service._equipsService
	dalService := service._dalService
	log := service._log
	authService := service._authService

	// httpService := service._httpService
	http.HandleFunc("/equips/Activate", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		claims := CheckUserAuthorization(authService, w, r) 
						
		if claims == nil{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		queryString := r.URL.Query()

		sessionUids, ok := queryString["sessionUid"]

		if !ok || len(sessionUids[0]) < 1 {
			log.Error("Url Param 'sessionUid' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		sessionUID := sessionUids[0]

		activatedEquipInfos, ok := queryString["activatedEquipInfo"]

		if !ok || len(activatedEquipInfos[0]) < 1 {
			log.Error("Url Param 'activatedEquipInfo' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		activatedEquipInfo := activatedEquipInfos[0]

		deactivatedEquipInfos, ok := queryString["deactivatedEquipInfo"]

		if !ok {
			log.Error("Url Param 'deactivatedEquipInfo' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		deactivatedEquipInfo := ""
		if len(deactivatedEquipInfos[0]) > 0 {
			deactivatedEquipInfo = deactivatedEquipInfos[0]
		}

		log.Errorf("Url is: %s %s %s", sessionUID, activatedEquipInfo, deactivatedEquipInfo)
		if deactivatedEquipInfo != "" && deactivatedEquipInfo != activatedEquipInfo {
			mqttReceiverService.SendCommand(deactivatedEquipInfo, "deactivate")
		}
		webSocketService.Activate(sessionUID, activatedEquipInfo, deactivatedEquipInfo)
		mqttReceiverService.SendCommand(activatedEquipInfo, "activate")
	})

	http.HandleFunc("/equips/GetConnectedEquips", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		/*w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")*/
		claims := CheckUserAuthorization(authService, w, r) 
						
		if claims == nil{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		equips := mqttReceiverService.GetConnectionNames()
		json.NewEncoder(w).Encode(equips)
	})

	http.HandleFunc("/equips/GetAllEquips", func(w http.ResponseWriter, r *http.Request) {
		/*//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
*/
		claims := CheckUserAuthorization(authService, w, r) 
						
		if claims == nil{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		queryString := r.URL.Query()

		withDisableds, ok := queryString["withDisabled"]

		if !ok || len(withDisableds[0]) < 1 {
			log.Error("Url Param 'withDisabled' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		withDisabled, _ := strconv.ParseBool(withDisableds[0])

		equipInfos := equipsService.GetEquipInfos(withDisabled)
		json.NewEncoder(w).Encode(equipInfos)
	})

	http.HandleFunc("/equips/DisableEquipInfo", func(w http.ResponseWriter, r *http.Request) {
		/*//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
*/
		claims := CheckAdminAuthorization(authService, w, r) 
				
		if claims == nil{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		queryString := r.URL.Query()

		equipNames, ok := queryString["equipName"]
		if !ok {
			log.Error("Url Param 'equipName' is missing")
			return
		}
		equipName := equipNames[0]

		disableds, ok := queryString["disabled"]

		if !ok || len(disableds[0]) < 1 {
			log.Error("Url Param 'disabled' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		disabled, _ := strconv.ParseBool(disableds[0])

		equipsService.DisableEquipInfo(equipName, disabled)
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

	http.HandleFunc("/equips/UpdateDBInfo", func(w http.ResponseWriter, r *http.Request) {
		service.sendCommand(w, r, "updateDBInfo")
	})

	http.HandleFunc("/equips/XilibLogsOn", func(w http.ResponseWriter, r *http.Request) {		
		queryString := r.URL.Query()

		detailedXilibs, ok := queryString["detailedXilib"]
		if !ok || len(detailedXilibs[0]) < 1 {
			log.Error("Url Param 'detailedXilib' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// detailedXilib := detailedXilibs[0]

		verboseXilibs, ok := queryString["verboseXilib"]
		if !ok || len(verboseXilibs[0]) < 1 {
			log.Error("Url Param 'verboseXilib' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// verboseXilib := verboseXilibs[0]

		service.sendCommand(w, r, "xilibLogsOn")
	})

	//(currType, equipName, startDate, endDate);
	http.HandleFunc("/equips/SearchEquip", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		/*w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")*/
		claims := CheckAdminAuthorization(authService, w, r) 
				
		if claims == nil{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		
		queryString := r.URL.Query()

		equipTypes, ok := queryString["currType"]
		if !ok || len(equipTypes[0]) < 1 {
			log.Error("Url Param 'currType' is missing")
			return
		}
		equipType := equipTypes[0]

		equipNames, ok := queryString["equipName"]
		if !ok {
			log.Error("Url Param 'equipName' is missing")
			return
		}
		equipName := equipNames[0]

		startDates, ok := queryString["startDate"]
		if !ok || len(startDates[0]) < 1 {
			log.Error("Url Param 'startDate' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		endDates, ok := queryString["endDate"]
		if !ok {
			log.Error("Url Param 'endDate' is missing")
			w.WriteHeader(http.StatusBadRequest)
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

		// log.Println("Url is: %s %s %s", equipType, startDate, endDate)

		service.sendSearchResults(w, equipType, equipName, startDate, endDate)
	})

	http.HandleFunc("/equips/GetAllDBTableNames", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		/*w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")*/
		claims := CheckAdminAuthorization(authService, w, r) 
				
		if claims == nil{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		
		queryString := r.URL.Query()

		equipNames, ok := queryString["equipName"]
		if !ok {
			log.Error("Url Param 'equipName' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		equipName := equipNames[0]
		tables := dalService.GetAllTableNamesInfo(equipName)

		json.NewEncoder(w).Encode(tables)
	})

	http.HandleFunc("/equips/GetTableContent", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		/*w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")*/
		claims := CheckAdminAuthorization(authService, w, r) 
				
		if claims == nil{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		
		queryString := r.URL.Query()

		equipNames, ok := queryString["equipName"]
		if !ok {
			log.Error("Url Param 'equipName' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		equipName := equipNames[0]

		tableTypes, ok := queryString["tableType"]
		if !ok {
			log.Error("Url Param 'tableType' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tableType := tableTypes[0]

		tableNames, ok := queryString["tableName"]
		if !ok {
			log.Error("Url Param 'tableName' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tableName := tableNames[0]

		tables := dalService.GetTableContent(equipName, tableType, tableName)

		json.NewEncoder(w).Encode(tables)
	})

	http.HandleFunc("/equips/GetPermanentData", func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		/*w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")*/
		claims := CheckAdminAuthorization(authService, w, r) 
				
		if claims == nil{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		
		queryString := r.URL.Query()

		equipTypes, ok := queryString["currType"]
		if !ok || len(equipTypes[0]) < 1 {
			log.Error("Url Param 'currType' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		equipType := equipTypes[0]

		equipNames, ok := queryString["equipName"]
		if !ok {
			log.Error("Url Param 'equipName' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		equipName := equipNames[0]

		// log.Println("Url is: %s %s", equipType, equipName)

		service.sendPermanentSearchResults(w, equipType, equipName)
	})
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
	} else if equipType == "Events" {
		events := dalService.GetEvents(equipName, startDate, endDate)
		json.NewEncoder(w).Encode(events)
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
	} else if equipType == "LastSeen" {
		swInfo := dalService.GetLastSeenInfo(equipName)
		json.NewEncoder(w).Encode(swInfo)
	}
}

func (service *EquipController) sendCommand(w http.ResponseWriter, r *http.Request, command string) {
	//Allow CORS here By * or specific origin
	/*w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Content-Type", "application/json")*/
	log := service._log
	claims := CheckUserAuthorization(service._authService, w, r) 
						
	if claims == nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	queryString := r.URL.Query()
	activatedEquipInfos, ok := queryString["activatedEquipInfo"]

	if !ok || len(activatedEquipInfos[0]) < 1 {
		log.Error("Url Param 'activatedEquipInfo' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	activatedEquipInfo := activatedEquipInfos[0]
	service._mqttReceiverService.SendCommand(activatedEquipInfo, command)

	log.Infof("User %s sent command %s", claims.Login, command)
}
