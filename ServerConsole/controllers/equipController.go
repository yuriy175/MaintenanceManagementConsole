package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	interfaces "ServerConsole/interfaces"
)

// EquipController describes equipment controller implementation type
type EquipController struct {
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

	// equipment service
	_equipsService interfaces.IEquipsService

	// events service
	_eventsService interfaces.IEventsService

	// http service
	_httpService interfaces.IHttpService

	// authorization service
	_authService interfaces.IAuthService
}

// EquipControllerNew creates an instance of webSock
func EquipControllerNew(
	log interfaces.ILogger,
	diagnosticService interfaces.IDiagnosticService,
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	eventsService interfaces.IEventsService,
	httpService interfaces.IHttpService,
	authService interfaces.IAuthService) *EquipController {
	service := &EquipController{}

	service._log = log
	service._diagnosticService = diagnosticService
	service._httpService = httpService
	service._mqttReceiverService = mqttReceiverService
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._equipsService = equipsService
	service._eventsService = eventsService
	service._authService = authService

	return service
}

// Handle handles incomming requests
/*func (service *EquipController) Handle() {
	mqttReceiverService := service._mqttReceiverService
	webSocketService := service._webSocketService
	equipsService := service._equipsService
	dalService := service._dalService
	log := service._log
	authService := service._authService
	diagnosticService := service._diagnosticService

	// httpService := service._httpService
	http.HandleFunc("/equips/Activate", func(w http.ResponseWriter, r *http.Request) {
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
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

		go func() {
			webSocketService.Activate(sessionUID, activatedEquipInfo, deactivatedEquipInfo)
			mqttReceiverService.Activate(activatedEquipInfo, deactivatedEquipInfo)
		}()
	})

	http.HandleFunc("/equips/GetConnectedEquips", func(w http.ResponseWriter, r *http.Request) {
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
			return
		}
		equips := mqttReceiverService.GetConnectionNames()
		json.NewEncoder(w).Encode(equips)
	})

	http.HandleFunc("/equips/GetAllEquips", func(w http.ResponseWriter, r *http.Request) {
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
			return
		}
		methodName := "/equips/GetAllEquips"
		diagnosticService.IncCount(methodName)
		start := time.Now()

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

		diagnosticService.SetDuration(methodName, time.Since(start))
	})

	http.HandleFunc("/equips/DisableEquipInfo", func(w http.ResponseWriter, r *http.Request) {
		claims := CheckAdminAuthorization(authService, w, r)

		if claims == nil {
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
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
			return
		}

		service.sendCommand(w, r, "updateDBInfo")
	})

	http.HandleFunc("/equips/RecreateDBInfo", func(w http.ResponseWriter, r *http.Request) {
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
			return
		}

		queryString := r.URL.Query()

		equipInfo := CheckQueryParameter(queryString, "activatedEquipInfo", w)
		if equipInfo == "" {
			log.Error("Url Param 'equipInfo' is missing")
			return
		}

		go dalService.DisableAllDBInfo(equipInfo)
		service.sendCommand(w, r, "recreateDBInfo")
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

	http.HandleFunc("/equips/SetEquipLogsOn", func(w http.ResponseWriter, r *http.Request) {
		log := service._log
		claims := CheckUserAuthorization(service._authService, w, r)

		if claims == nil {
			return
		}

		queryString := r.URL.Query()
		equipName := CheckQueryParameter(queryString, "activatedEquipInfo", w)
		if equipName == "" {
			log.Error("Url Param 'activatedEquipInfo' is missing")
			return
		}

		hardwareType := CheckQueryParameter(queryString, "hardwareType", w)
		if hardwareType == "" {
			log.Error("Url Param 'hardwareType' is missing")
			return
		}

		value := CheckQueryParameter(queryString, "value", w)
		if value == "" {
			log.Error("Url Param 'value' is missing")
			return
		}

		service._mqttReceiverService.SendCommand(equipName,
			"equipLogsOn"+"?hardwareType="+hardwareType+
				"&value="+value)
	})

	//(currType, equipName, startDate, endDate);
	http.HandleFunc("/equips/SearchEquip", func(w http.ResponseWriter, r *http.Request) {
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
			return
		}

		start := time.Now()
		queryString := r.URL.Query()

		equipTypes, ok := queryString["currType"]
		if !ok || len(equipTypes[0]) < 1 {
			log.Error("Url Param 'currType' is missing")
			return
		}
		equipType := equipTypes[0]

		methodName := "/equips/SearchEquip_" + equipType
		diagnosticService.IncCount(methodName)

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

		diagnosticService.SetDuration(methodName, time.Since(start))
	})

	http.HandleFunc("/equips/GetAllDBTableNames", func(w http.ResponseWriter, r *http.Request) {
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
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
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
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
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
			return
		}

		queryString := r.URL.Query()

		equipType := CheckQueryParameter(queryString, "currType", w)
		if equipType == "" {
			log.Error("Url Param 'equipType' is missing")
			return
		}

		equipName := CheckQueryParameter(queryString, "equipName", w)
		if equipName == "" {
			log.Error("Url Param 'equipName' is missing")
			return
		}

		service.sendPermanentSearchResults(w, equipType, equipName)
	})
}*/

func (service *EquipController) sendSearchResults(
	w http.ResponseWriter,
	equipType string,
	equipName string,
	startDate time.Time,
	endDate time.Time) {

	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, time.UTC)

	dalService := service._dalService
	eventsService := service._eventsService
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
		events := eventsService.GetEvents(equipName, startDate, endDate)
		json.NewEncoder(w).Encode(events)
	}
}

func (service *EquipController) sendPermanentSearchResults(
	w http.ResponseWriter,
	equipType string,
	equipName string) {

	dalService := service._dalService
	equipsService := service._equipsService

	if equipType == "FullInfo" {
		fullInfo := equipsService.GetFullInfo(equipName)
		json.NewEncoder(w).Encode(fullInfo)
	} else if equipType == "SystemInfo" {
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
	log := service._log
	claims := CheckUserAuthorization(service._authService, w, r)

	if claims == nil {
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

// ActivateEquip activates an equipment
func (service *EquipController) ActivateEquip(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	mqttReceiverService := service._mqttReceiverService
	webSocketService := service._webSocketService
	log := service._log
	if CheckUserAuthorization(authService, w, r) != nil {
		queryString := r.URL.Query()

		sessionUID := CheckQueryParameter(queryString, "sessionUid", w)
		if sessionUID == "" {
			log.Error("Url Param 'sessionUid' is missing")
			return
		}

		activatedEquipInfo := CheckQueryParameter(queryString, "activatedEquipInfo", w)
		if activatedEquipInfo == "" {
			log.Error("Url Param 'activatedEquipInfo' is missing")
			return
		}

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

		go func() {
			webSocketService.Activate(sessionUID, activatedEquipInfo, deactivatedEquipInfo)
			mqttReceiverService.Activate(activatedEquipInfo, deactivatedEquipInfo)
		}()
	}
}

// GetConnectedEquips returns connected equipments
func (service *EquipController) GetConnectedEquips(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	mqttReceiverService := service._mqttReceiverService
	if CheckUserAuthorization(authService, w, r) != nil {
		equips := mqttReceiverService.GetConnectionNames()
		json.NewEncoder(w).Encode(equips)
	}
}

// GetConnectedEquips returns connected equipments
func (service *EquipController) GetAllEquips(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	diagnosticService := service._diagnosticService
	equipsService := service._equipsService
	log := service._log
	if CheckUserAuthorization(authService, w, r) != nil {
		methodName := "/equips/GetAllEquips"
		diagnosticService.IncCount(methodName)
		start := time.Now()

		queryString := r.URL.Query()

		withDisabledParam := CheckQueryParameter(queryString, "withDisabled", w)
		if withDisabledParam == "" {
			log.Error("Url Param 'withDisabled' is missing")
			return
		}

		withDisabled, _ := strconv.ParseBool(withDisabledParam)
		equipInfos := equipsService.GetEquipInfos(withDisabled)
		json.NewEncoder(w).Encode(equipInfos)

		diagnosticService.SetDuration(methodName, time.Since(start))
	}
}

// DisableEquipInfo disables equipment
func (service *EquipController) DisableEquipInfo(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	equipsService := service._equipsService
	log := service._log
	if CheckAdminAuthorization(authService, w, r) != nil {
		queryString := r.URL.Query()

		equipName := CheckQueryParameter(queryString, "equipName", w)
		if equipName == "" {
			log.Error("Url Param 'equipName' is missing")
			return
		}

		disableds, ok := queryString["disabled"]

		if !ok || len(disableds[0]) < 1 {
			log.Error("Url Param 'disabled' is missing")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		disabled, _ := strconv.ParseBool(disableds[0])

		equipsService.DisableEquipInfo(equipName, disabled)
	}
}

// RunTeamViewer sends RunTeamViewer command
func (service *EquipController) RunTeamViewer(w http.ResponseWriter, r *http.Request) {
	service.sendCommand(w, r, "runTV")
}

// RunTaskManager sends RunTaskManager command
func (service *EquipController) RunTaskManager(w http.ResponseWriter, r *http.Request) {
	service.sendCommand(w, r, "runTaskMan")
}

// SendAtlasLogs sends SendAtlasLogs command
func (service *EquipController) SendAtlasLogs(w http.ResponseWriter, r *http.Request) {
	service.sendCommand(w, r, "sendAtlasLogs")
}

// UpdateDBInfo sends UpdateDBInfo command
func (service *EquipController) UpdateDBInfo(w http.ResponseWriter, r *http.Request) {
	service.sendCommand(w, r, "updateDBInfo")
}

// RecreateDBInfo sends RecreateDBInfo command
func (service *EquipController) RecreateDBInfo(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	dalService := service._dalService
	log := service._log
	if CheckUserAuthorization(authService, w, r) != nil {
		queryString := r.URL.Query()

		equipInfo := CheckQueryParameter(queryString, "activatedEquipInfo", w)
		if equipInfo == "" {
			log.Error("Url Param 'equipInfo' is missing")
			return
		}

		go dalService.DisableAllDBInfo(equipInfo)
		service.sendCommand(w, r, "recreateDBInfo")
	}
}

// XilibLogsOn sends XilibLogsOn command
func (service *EquipController) XilibLogsOn(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	log := service._log
	if CheckUserAuthorization(authService, w, r) != nil {
		queryString := r.URL.Query()

		detailedXilib := CheckQueryParameter(queryString, "detailedXilib", w)
		if detailedXilib == "" {
			log.Error("Url Param 'detailedXilib' is missing")
			return
		}

		verboseXilib := CheckQueryParameter(queryString, "verboseXilib", w)
		if verboseXilib == "" {
			log.Error("Url Param 'verboseXilib' is missing")
			return
		}

		service.sendCommand(w, r, "xilibLogsOn")
	}
}

// SetEquipLogsOn sends equipLogsOn command
func (service *EquipController) SetEquipLogsOn(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	log := service._log
	if CheckUserAuthorization(authService, w, r) != nil {
		queryString := r.URL.Query()
		equipName := CheckQueryParameter(queryString, "activatedEquipInfo", w)
		if equipName == "" {
			log.Error("Url Param 'activatedEquipInfo' is missing")
			return
		}

		hardwareType := CheckQueryParameter(queryString, "hardwareType", w)
		if hardwareType == "" {
			log.Error("Url Param 'hardwareType' is missing")
			return
		}

		value := CheckQueryParameter(queryString, "value", w)
		if value == "" {
			log.Error("Url Param 'value' is missing")
			return
		}

		service._mqttReceiverService.SendCommand(equipName,
			"equipLogsOn"+"?hardwareType="+hardwareType+
				"&value="+value)
	}
}

// SearchEquip returns search results for equipment
func (service *EquipController) SearchEquip(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	diagnosticService := service._diagnosticService
	log := service._log
	if CheckUserAuthorization(authService, w, r) != nil {
		start := time.Now()
		queryString := r.URL.Query()

		equipType := CheckQueryParameter(queryString, "currType", w)
		if equipType == "" {
			log.Error("Url Param 'currType' is missing")
			return
		}

		methodName := "/equips/SearchEquip_" + equipType
		diagnosticService.IncCount(methodName)

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
			log.Errorf("Url Param 'startDate' time.Parse %v", err)
		}

		endDate, err := time.Parse("2006-01-02", endDates[0])
		if err != nil {
			log.Errorf("Url Param 'endDate' time.Parse %v", err)
		}

		service.sendSearchResults(w, equipType, equipName, startDate, endDate)

		diagnosticService.SetDuration(methodName, time.Since(start))
	}
}

// GetAllDBTableNames returns all db table names
func (service *EquipController) GetAllDBTableNames(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	dalService := service._dalService
	log := service._log
	if CheckUserAuthorization(authService, w, r) != nil {
		queryString := r.URL.Query()

		equipName := CheckQueryParameter(queryString, "equipName", w)
		if equipName == "" {
			log.Error("Url Param 'equipName' is missing")
			return
		}
		tables := dalService.GetAllTableNamesInfo(equipName)

		json.NewEncoder(w).Encode(tables)
	}
}

// GetTableContent returns a db table content
func (service *EquipController) GetTableContent(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	dalService := service._dalService
	log := service._log
	if CheckUserAuthorization(authService, w, r) != nil {
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
	}
}

// GetPermanentData returns equipment permanent data
func (service *EquipController) GetPermanentData(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	log := service._log
	if CheckUserAuthorization(authService, w, r) != nil {
		queryString := r.URL.Query()

		equipType := CheckQueryParameter(queryString, "currType", w)
		if equipType == "" {
			log.Error("Url Param 'equipType' is missing")
			return
		}

		equipName := CheckQueryParameter(queryString, "equipName", w)
		if equipName == "" {
			log.Error("Url Param 'equipName' is missing")
			return
		}

		service.sendPermanentSearchResults(w, equipType, equipName)
	}
}
