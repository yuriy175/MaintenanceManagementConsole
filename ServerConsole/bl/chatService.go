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

	// equipment service
	_equipsService   interfaces.IEquipsService

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
	equipsService interfaces.IEquipsService,
	webSockCh chan *models.RawMqttMessage,
	chatCh chan *models.RawMqttMessage) interfaces.IChatService {
	service := &chatService{}

	service._log = log
	service._dalService = dalService
	service._webSocketService = webSocketService
	service._equipsService = equipsService
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
			note := dalService.UpsertChatNote(utils.GetEquipFromTopic(d.Topic), "Chat", "", 
				viewmodel.Message, viewmodel.User, viewmodel.IsInternal)

			data, _ := json.Marshal(note)
			rawMsg := models.RawMqttMessage{d.Topic, string(data)}
			service._webSockCh <- &rawMsg
		}
	}() 
}
