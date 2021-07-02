package bl

import (
	"encoding/json"
	"sync"

	"../interfaces"
	"../models"
	"../utils"
)

// chat service implementation type
type chatService struct {
	// synchronization mutex
	_mtx sync.RWMutex

	//logger
	_log interfaces.ILogger

	// web socket service
	_webSocketService interfaces.IWebSocketService

	// DAL service
	_dalService interfaces.IDalService

	// chanel for communications with chat services
	_chatCh chan *models.RawMqttMessage

	// chanel for communications with websocket services
	_webSockCh     chan *models.RawMqttMessage
}

// EventsServiceNew creates an instance of equipsService
func ChatServiceNew(
	log interfaces.ILogger,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	webSockCh chan *models.RawMqttMessage,
	chatCh chan *models.RawMqttMessage) interfaces.IChatService {
	service := &chatService{}

	service._log = log
	service._dalService = dalService
	service._webSocketService = webSocketService
	service._chatCh = chatCh
	service._webSockCh = webSockCh

	return service
}

// Starts the service
func (service *chatService) Start() {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	dalService := service._dalService
	// webSocketService := service._webSocketService
	go func() {
		for d := range service._chatCh {
			viewmodel := models.ChatViewModel{}
			json.Unmarshal([]byte(d.Data), &viewmodel)				
			note := dalService.InsertChatNote(utils.GetEquipFromTopic(d.Topic), "Chat", viewmodel.Message, viewmodel.User)

			data, _ := json.Marshal(note)
			rawMsg := models.RawMqttMessage{d.Topic, string(data)}
			service._webSockCh <- &rawMsg
		}
	}() 
}
