package bl

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	"../interfaces"
	"../models"
	"../utils"
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

	// chanel for communications with websocket services
	_webSockCh chan *models.RawMqttMessage

	// chanel for communications with events service
	_eventsCh chan *models.RawMqttMessage
}

// EventsServiceNew creates an instance of equipsService
func EventsServiceNew(
	log interfaces.ILogger,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	webSockCh chan *models.RawMqttMessage,
	eventsCh chan *models.RawMqttMessage) interfaces.IEventsService {
	service := &eventsService{}

	service._log = log
	service._dalService = dalService
	service._webSocketService = webSocketService
	service._eventsCh = eventsCh
	service._webSockCh = webSockCh

	return service
}

// Starts the service
func (service *eventsService) Start() {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	go func() {
		for d := range service._eventsCh {
			if strings.Contains(d.Topic, "/ARM/Software/msg") {
				viewmodel := models.SoftwareMessageViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)

				equipName := utils.GetEquipFromTopic(d.Topic)
				service.insertEvents(equipName, &viewmodel, false, nil)
				/*events := []models.EventModel{}
				if viewmodel.ErrorDescriptions != nil {
					events = service._dalService.InsertEvents(equipName, "ErrorDescriptions", viewmodel.ErrorDescriptions)
				}

				if viewmodel.AtlasErrorDescriptions != nil {
					events = append(events,
						service._dalService.InsertEvents(equipName, "AtlasErrorDescriptions", viewmodel.AtlasErrorDescriptions)...)
				}

				if viewmodel.SimpleMsgType != "" {
					msgCode := ""
					if viewmodel.SimpleMsgType == "AtlasExited"{
						msgCode = "Атлас выключен"
					} 
					msg := models.MessageViewModel{equipName, msgCode, ""}
					events = append(events,
						service._dalService.InsertEvents(equipName, viewmodel.SimpleMsgType, []models.MessageViewModel{msg})...)
				}

				if viewmodel.AtlasUser.User != "" {
					msg := models.MessageViewModel{equipName,  viewmodel.AtlasUser.User + " (" + viewmodel.AtlasUser.Role + ") вошел в Атлас", ""}
					events = append(events,
						service._dalService.InsertEvents(equipName, "AtlasUser", []models.MessageViewModel{msg})...)
				}

				if len(events) > 0 {
					go webSocketService.SendEvents(events)
				}*/
			}
		}
	}() //deviceCollection)
}

// InsertEvent inserts equipment connection state info into db
func (service *eventsService) InsertConnectEvent(equipName string) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	dalService := service._dalService
	webSocketService := service._webSocketService

	go func() {
		msg := models.MessageViewModel{equipName, "подключен", ""}
		events := dalService.InsertEvents(equipName, "EquipConnected", []models.MessageViewModel{msg}, nil)
		go webSocketService.SendEvents(events)
	}()
}

func (service *eventsService) insertEvents(
	equipName string, 
	viewmodel *models.SoftwareMessageViewModel, 
	isOffline bool, 
	msgDate *time.Time) {
	webSocketService := service._webSocketService
	events := []models.EventModel{}
	typePostfix := ""
	if isOffline{
		typePostfix = "Offline"
	}
	if viewmodel.ErrorDescriptions != nil {
		events = service._dalService.InsertEvents(equipName, "ErrorDescriptions" + typePostfix, viewmodel.ErrorDescriptions, msgDate)
	}

	if viewmodel.AtlasErrorDescriptions != nil {
		events = append(events,
			service._dalService.InsertEvents(equipName, "AtlasErrorDescriptions" + typePostfix, viewmodel.AtlasErrorDescriptions, msgDate)...)
	}

	if viewmodel.OfflineMsg != nil && viewmodel.OfflineMsg.Message != nil {
		value := viewmodel.OfflineMsg.DateTime
		layout := time.RFC3339[:len(value)]
		date, err := time.Parse(layout, value)
		if err != nil{
			date = time.Now()
			service._log.Errorf("insertEvent error: %s %s", err, value);
		} 
		
		service.insertEvents(equipName, viewmodel.OfflineMsg.Message, true, &date)
	}

	if viewmodel.SimpleMsgType != "" {
		msgCode := ""
		if viewmodel.SimpleMsgType == "AtlasExited"{
			msgCode = "Атлас выключен"
		} else if viewmodel.SimpleMsgType == "InstanceOnOffline"{
			msgCode = "подключен"
		}
		msg := models.MessageViewModel{equipName, msgCode, ""}
		events = append(events,
			service._dalService.InsertEvents(equipName, viewmodel.SimpleMsgType + typePostfix, []models.MessageViewModel{msg}, msgDate)...)
	}

	if viewmodel.AtlasUser.User != "" {
		msg := models.MessageViewModel{equipName,  viewmodel.AtlasUser.User + " (" + viewmodel.AtlasUser.Role + ") вошел в Атлас", ""}
		events = append(events,
			service._dalService.InsertEvents(equipName, "AtlasUser" + typePostfix, []models.MessageViewModel{msg}, msgDate)...)
	}

	if len(events) > 0 {
		go webSocketService.SendEvents(events)
	}
}