package bl

import (
	"fmt"
	"sync"

	"../interfaces"
	"../models"
	Models "../models"
)

// mqtt receiver service implementation type
type mqttReceiverService struct {
	//logger
	_log interfaces.ILogger

	// synchronization mutex
	_mtx              sync.RWMutex

	// IoC provider
	_ioCProvider      interfaces.IIoCProvider

	// web socket service
	_webSocketService interfaces.IWebSocketService

	// DAL service
	_dalService    interfaces.IDalService

	// equipment service
	_equipsService interfaces.IEquipsService

	// events service
	_eventsService interfaces.IEventsService

	// chanel for DAL communications
	_dalCh         chan *models.RawMqttMessage

	// chanel for communications with websocket services
	_webSockCh     chan *models.RawMqttMessage

	// chanel for communications with events services
	_eventsCh     chan *models.RawMqttMessage

	// mqtt connections map
	// key - topic
	// value - mqtt client
	_mqttConnections map[string]interfaces.IMqttClient
	_topicStorage    interfaces.ITopicStorage

	// topics : server may communicate with a client
	_supportedTopics []string
}

// MqttReceiverServiceNew creates an instance of mqttReceiverService
func MqttReceiverServiceNew(
	log interfaces.ILogger,
	ioCProvider interfaces.IIoCProvider,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	eventsService interfaces.IEventsService,
	topicStorage interfaces.ITopicStorage,
	dalCh chan *models.RawMqttMessage,
	webSockCh chan *models.RawMqttMessage,
	eventsCh  chan *models.RawMqttMessage) interfaces.IMqttReceiverService {
	service := &mqttReceiverService{}

	service._log = log
	service._ioCProvider = ioCProvider
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._equipsService = equipsService
	service._eventsService = eventsService
	service._topicStorage = topicStorage
	service._dalCh = dalCh
	service._webSockCh = webSockCh
	service._eventsCh = eventsCh
	service._mqttConnections = map[string]interfaces.IMqttClient{}

	service._supportedTopics = topicStorage.GetTopics()

	return service
}

//UpdateMqttConnections updates mqtt connections map for an equipment connection state
func (service *mqttReceiverService) UpdateMqttConnections(state *models.EquipConnectionState) {
	rootTopic := state.Name
	isOff := !state.Connected
	topics := service._supportedTopics
	mqttConnections := service._mqttConnections
	ioCProvider := service._ioCProvider
	// dalService := service._dalService
	// equipsService := service._equipsService
	eventsService := service._eventsService

	fmt.Printf("UpdateMqttConnections from topic: %s\n", rootTopic)

	service._mtx.Lock()
	defer service._mtx.Unlock()

	fmt.Printf("UpdateMqttConnections unlocked")

	if client, ok := mqttConnections[rootTopic]; ok {
		fmt.Println(rootTopic + " already exists")
		if isOff {
			go client.Disconnect()
			delete(mqttConnections, rootTopic)
			fmt.Println(rootTopic + " deleted")
		}

		// if the topic is observed by any client -> send activate command
		if service._webSocketService.HasActiveClients(rootTopic) {
			go service.SendCommand(rootTopic, "activate")
		}

		return
	}

	if !isOff {
		mqttConnections[rootTopic] = ioCProvider.GetMqttClient().Create(rootTopic, topics)
		// if !equipsService.CheckEquipment(rootTopic) {
		go service.SendCommand(rootTopic, "serverReady")
		// }

		go eventsService.InsertConnectEvent(rootTopic)
	}

	fmt.Println(rootTopic + " created")
}

// CreateCommonConnections reates common mqtt connections
func (service *mqttReceiverService) CreateCommonConnections() {
	mqttConnections := service._mqttConnections
	ioCProvider := service._ioCProvider

	service._mtx.Lock()
	defer service._mtx.Unlock()
	mqttConnections[Models.CommonTopicPath] = ioCProvider.GetMqttClient().Create(models.CommonTopicPath, []string{})
	mqttConnections[Models.BroadcastCommandsTopic] = ioCProvider.GetMqttClient().Create(models.BroadcastCommandsTopic, []string{})

	return
}

// SendCommand sends a command to equipment via mqtt
func (service *mqttReceiverService) SendCommand(equipment string, command string) {
	fmt.Printf("SendCommand from topic: %s %s\n", equipment, command)

	service._mtx.Lock()
	defer service._mtx.Unlock()

	if client, ok := service._mqttConnections[equipment]; ok {
		go client.SendCommand(command)
	}

	return
}

// PublishChatNote sends a chat note to equipment via mqtt
func (service *mqttReceiverService) PublishChatNote(equipment string, message string, user string) {
	fmt.Printf("PublishChatNote from topic: %s %s %s\n", equipment, message, user)

	service._mtx.Lock()
	defer service._mtx.Unlock()

	// we may have no connection to this client
	topics := service._supportedTopics
	mqttConnections := service._mqttConnections
	ioCProvider := service._ioCProvider

	if _, ok := mqttConnections[equipment]; !ok {
		mqttConnections[equipment] = ioCProvider.GetMqttClient().Create(equipment, topics)
	}	

	if client, ok := mqttConnections[equipment]; ok {
		go client.SendChatMessage(user, message)
	}

	return
}

// SendCommand sends a broadcast command to equipments via mqtt
func (service *mqttReceiverService) SendBroadcastCommand(command string) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	if client, ok := service._mqttConnections[models.BroadcastCommandsTopic]; ok {
		go client.SendCommand(command)
	}

	return
}

// GetConnectionNames returns connected equipment names
func (service *mqttReceiverService) GetConnectionNames() []string {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	mqttConnections := service._mqttConnections

	keys := make([]string, len(mqttConnections))

	i := 0
	for k, d := range mqttConnections {
		if d.IsEquipTopic() {
			keys[i] = k
			i++
		}
	}

	return keys
}
