package bl

import (
	"encoding/json"
	"strings"
	"sync"

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

	webSocketService := service._webSocketService

	go func() {
		for d := range service._eventsCh {
			if strings.Contains(d.Topic, "/ARM/Software/msg") {
				viewmodel := models.SoftwareMessageViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)

				equipName := utils.GetEquipFromTopic(d.Topic)
				events := []models.EventModel{}
				if viewmodel.ErrorDescriptions != nil {
					events = service._dalService.InsertEvents(equipName, "ErrorDescriptions", viewmodel.ErrorDescriptions)
				}

				if viewmodel.AtlasErrorDescriptions != nil {
					events = append(events,
						service._dalService.InsertEvents(equipName, "AtlasErrorDescriptions", viewmodel.AtlasErrorDescriptions)...)
				}

				if len(events) > 0 {
					go webSocketService.SendEvents(events)
				}
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
		msg := models.MessageViewModel{equipName, "connected"}
		events := dalService.InsertEvents(equipName, "EquipConnected", []models.MessageViewModel{msg})
		go webSocketService.SendEvents(events)
	}()
}
