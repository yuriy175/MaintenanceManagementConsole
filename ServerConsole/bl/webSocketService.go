package bl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"../interfaces"
	"../models"
	"../utils"
)

// web socket service implementation type
type webSocketService struct {
	//logger
	_log interfaces.ILogger

	// synchronization mutex
	_mtx sync.RWMutex

	// IoC provider
	_ioCProvider interfaces.IIoCProvider

	// settings service
	_settingsService interfaces.ISettingsService

	// web socket server connection string
	_connectionString string 

	// chanel for communications with websocket services
	_webSockCh chan *models.RawMqttMessage

	// map of active websocket connections
	// keys - sessionUids
	_webSocketConnections map[string]interfaces.IWebSock

	// map of active equipment topics connections
	// keys - main equipment topics
	// values - slice of session uids
	_topicConnections map[string][]string
}

// WebSocketServiceNew creates an instance of webSocketService
func WebSocketServiceNew(
	log interfaces.ILogger,
	ioCProvider interfaces.IIoCProvider,
	settingsService interfaces.ISettingsService,
	webSockCh chan *models.RawMqttMessage) interfaces.IWebSocketService {
	service := &webSocketService{}

	service._log = log
	service._ioCProvider = ioCProvider
	service._settingsService = settingsService
	service._connectionString = settingsService.GetWebSocketServerConnectionString();
	service._webSockCh = webSockCh
	service._webSocketConnections = map[string]interfaces.IWebSock{}
	service._topicConnections = map[string][]string{}

	return service
}

// Start starts the service
func (service *webSocketService) Start() {
	http.HandleFunc(models.WebSocketQueryString, func(w http.ResponseWriter, r *http.Request) {
		uids, ok := r.URL.Query()["uid"]

		if !ok || len(uids[0]) < 1 {
			log.Println("Url Param 'uid' is missing")
			return
		}
		uid := uids[0]

		service._mtx.Lock()
		defer service._mtx.Unlock()

		service._webSocketConnections[uid] = service._ioCProvider.GetWebSocket().Create(w, r, uid)

		fmt.Printf("created websocket uid: %s \n", uid)
	})

	go func() {
		for d := range service._webSockCh {

			//find equipment name of a new message
			//topicParts := strings.Split(d.Topic, "/")
			activatedEquipInfo := utils.GetEquipFromTopic(d.Topic) //strings.Join([]string{topicParts[0], topicParts[1]}, "/")

			b, _ := json.Marshal(d)
			service.writeMessageToSocket(activatedEquipInfo, b)
			/*service._mtx.Lock()

			//find all sessions activated this equipment
			if sessionUids, ok := service._topicConnections[activatedEquipInfo]; ok {
				for _, uid := range sessionUids {
					v := service._webSocketConnections[uid]
					if v == nil || !v.IsValid() {
						log.Println(" no connection for  %s", uid)
						service.removeFromTopicMap(activatedEquipInfo, uid)
					} else if err = v.WriteMessage(b); err != nil {
						//log.Println("send message error for  %s", uid)
					}
				}
			}

			service._mtx.Unlock()
			*/
		}
	}()

	http.ListenAndServe(service._connectionString, nil) //":8080", nil)
}

// Activate activates a specified connection to equipment and deactivates the other
func (service *webSocketService) Activate(sessionUID string, activatedEquipInfo string, deactivatedEquipInfo string) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	service.removeFromTopicMap(deactivatedEquipInfo, sessionUID)
	topicConnections := service._topicConnections

	topicConnections[activatedEquipInfo] = append(topicConnections[activatedEquipInfo], sessionUID)

	fmt.Printf("websocket Activate %s\n", activatedEquipInfo)

	return
}

// UpdateWebClients notifies UI of a new equipment connection
func (service *webSocketService) UpdateWebClients(state *models.EquipConnectionState) {
	stateVM := &models.EquipConnectionStateViewModel{models.CommonTopicPath, *state}

	service._mtx.Lock()
	defer service._mtx.Unlock()

	for _, ws := range service._webSocketConnections {
		if ws.IsValid() {
			b, _ := json.Marshal(stateVM)
			ws.WriteMessage(b)
		}
	}
}

// SendEvents sends events to all web connections
func (service *webSocketService) SendEvents(events []models.EventModel) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	eventsVm := &models.EventsViewModel{models.EventsTopicPath, events}
	for _, ws := range service._webSocketConnections {
		if ws.IsValid() {
			b, _ := json.Marshal(eventsVm)
			ws.WriteMessage(b)
		}
	}
}

// HasActiveClients checks if there is an active connections
func (service *webSocketService) HasActiveClients(topic string) bool {
	_, ok := service._topicConnections[topic]
	return ok
}

// ClientClosed removes web socket connection
func (service *webSocketService) ClientClosed(uid string) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	delete(service._webSocketConnections, uid)
	fmt.Printf("removed websocket uid: %s \n", uid)
}

func (service *webSocketService) removeFromTopicMap(equipInfo string, uid string) {
	if equipInfo == "" || uid == "" {
		return
	}

	if sessionUids, ok := service._topicConnections[equipInfo]; ok {
		ind := -1
		for i, v := range sessionUids {
			if v == uid {
				ind = i
				break
			}
		}

		if ind < 0 {
			return
		}

		topicConnection := service._topicConnections[equipInfo]
		// service._topicConnections[equipInfo] = append(
		// 	service._topicConnections[equipInfo][:ind], service._topicConnections[equipInfo][ind+1:]...)
		service._topicConnections[equipInfo] = append(
			topicConnection[:ind], topicConnection[ind+1:]...)
		if len(service._topicConnections[equipInfo]) == 0 {
			delete(service._topicConnections, equipInfo)
		}

		log.Println("removed absent connection for  %s", uid)
	}
}

func (service *webSocketService)writeMessageToSocket(activatedEquipInfo string, data []byte) {
	service._mtx.Lock()
	defer service._mtx.Unlock()
	
	if activatedEquipInfo == models.CommonChat{
		for _, ws := range service._webSocketConnections {
			if ws.IsValid() {
				ws.WriteMessage(data)
			}
		}
		return
	}

	//find all sessions activated this equipment
	if sessionUids, ok := service._topicConnections[activatedEquipInfo]; ok {
		for _, uid := range sessionUids {
			v := service._webSocketConnections[uid]
			if v == nil || !v.IsValid() {
				log.Println(" no connection for  %s", uid)
				service.removeFromTopicMap(activatedEquipInfo, uid)
			} else if err := v.WriteMessage(data); err != nil {
				//log.Println("send message error for  %s", uid)
			}
		}
	}

	// service._mtx.Unlock()
}