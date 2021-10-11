package bl

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	"ServerConsole/interfaces"
	"ServerConsole/models"
	"ServerConsole/utils"
)

// events service implementation type
type eventsService struct {
	// synchronization mutex
	_mtx sync.RWMutex

	//logger
	_log interfaces.ILogger

	// web socket service
	_webSocketService interfaces.IWebSocketService

	// DAL service
	_dalService interfaces.IDalService

	// equipment service
	_equipsService interfaces.IEquipsService

	// chanel for communications with websocket services
	_webSockCh chan *models.RawMqttMessage

	// chanel for communications with events service (outer events)
	_eventsCh chan *models.RawMqttMessage

	// chanel for communications with events service (internal events)
	_internalEventsCh chan *models.MessageViewModel
}

// EventsServiceNew creates an instance of equipsService
func EventsServiceNew(
	log interfaces.ILogger,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	webSockCh chan *models.RawMqttMessage,
	eventsCh chan *models.RawMqttMessage,
	internalEventsCh chan *models.MessageViewModel) interfaces.IEventsService {
	service := &eventsService{}

	service._log = log
	service._dalService = dalService
	service._webSocketService = webSocketService
	service._equipsService = equipsService
	service._eventsCh = eventsCh
	service._webSockCh = webSockCh
	service._internalEventsCh = internalEventsCh

	return service
}

// Starts the service
func (service *eventsService) Start() {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	dalService := service._dalService
	equipsService := service._equipsService
	webSocketService := service._webSocketService

	go func() {
		for d := range service._eventsCh {
			if strings.Contains(d.Topic, "/ARM/Software/msg") {
				viewmodel := models.SoftwareMessageViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)

				equipName := utils.GetEquipFromTopic(d.Topic)
				service.insertEvents(equipName, &viewmodel, false, nil)
			}
		}
	}()

	go func() {
		for msg := range service._internalEventsCh {
			equipName := msg.Level
			events := dalService.InsertEvents(equipName, "InternalEvent", []models.MessageViewModel{*msg}, nil)
			equipsService.SetLastSeen(equipName)
			go webSocketService.SendEvents(events)
		}
	}()
}

// InsertEvent inserts equipment connection state info into db
func (service *eventsService) InsertConnectEvent(equipName string, connected bool) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	dalService := service._dalService
	webSocketService := service._webSocketService

	go func() {
		msgCode := "подключен"
		if !connected {
			msgCode = "отключен"
		}
		msg := models.MessageViewModel{equipName, msgCode, ""}
		events := dalService.InsertEvents(equipName, "EquipConnected", []models.MessageViewModel{msg}, nil)
		service._equipsService.SetLastSeen(equipName)
		go webSocketService.SendEvents(events)
	}()
}

// GetEvents returns all events from db
func (service *eventsService) GetEvents(equipName string, startDate time.Time, endDate time.Time) []models.EventModel {
	dalService := service._dalService
	if equipName == "" {
		return dalService.GetEvents([]string{}, startDate, endDate)
	}

	equipsService := service._equipsService

	equipNames := append(equipsService.GetOldEquipNames(equipName), equipName)

	return dalService.GetEvents(equipNames, startDate, endDate)
}

func (service *eventsService) insertEvents(
	equipName string,
	viewmodel *models.SoftwareMessageViewModel,
	isOffline bool,
	msgDate *time.Time) {
	webSocketService := service._webSocketService
	equipsService := service._equipsService
	events := []models.EventModel{}
	typePostfix := ""
	if isOffline {
		typePostfix = "Offline"
	}
	if viewmodel.ErrorDescriptions != nil {
		events = service._dalService.InsertEvents(equipName, "ErrorDescriptions"+typePostfix, viewmodel.ErrorDescriptions, msgDate)
	}

	if viewmodel.AtlasErrorDescriptions != nil {
		events = append(events,
			service._dalService.InsertEvents(equipName, "AtlasErrorDescriptions"+typePostfix, viewmodel.AtlasErrorDescriptions, msgDate)...)
	}

	if viewmodel.HardwareErrorDescriptions != nil {
		events = append(events,
			service._dalService.InsertEvents(equipName, "HardwareErrorDescriptions"+typePostfix, viewmodel.HardwareErrorDescriptions, msgDate)...)
	}

	if viewmodel.OfflineMsg != nil && viewmodel.OfflineMsg.Message != nil {
		value := viewmodel.OfflineMsg.DateTime
		layout := time.RFC3339[:len(value)]
		date, err := time.Parse(layout, value)
		if err != nil {
			date = time.Now()
			service._log.Errorf("insertEvent error: %s %s", err, value)
		}

		service.insertEvents(equipName, viewmodel.OfflineMsg.Message, true, &date)
	}

	if viewmodel.SimpleMsgType != "" {
		msgCode := ""
		if viewmodel.SimpleMsgType == "AtlasExited" {
			msgCode = "Атлас выключен"
		} else if viewmodel.SimpleMsgType == "InstanceOnOffline" {
			msgCode = "подключен"
		}
		msg := models.MessageViewModel{equipName, msgCode, ""}
		events = append(events,
			service._dalService.InsertEvents(equipName, viewmodel.SimpleMsgType+typePostfix, []models.MessageViewModel{msg}, msgDate)...)
	}

	if viewmodel.AtlasUser.User != "" {
		msg := models.MessageViewModel{equipName, viewmodel.AtlasUser.User + " (" + viewmodel.AtlasUser.Role + ") вошел в Атлас", ""}
		events = append(events,
			service._dalService.InsertEvents(equipName, "AtlasUser"+typePostfix, []models.MessageViewModel{msg}, msgDate)...)
	}

	if len(events) > 0 {
		equipsService.SetLastSeen(equipName)
		go webSocketService.SendEvents(events)
	}
}
