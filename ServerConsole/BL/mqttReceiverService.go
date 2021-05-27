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
	// synchronization mutex
	_mtx              sync.RWMutex
	_ioCProvider      interfaces.IIoCProvider
	_webSocketService interfaces.IWebSocketService

	// DAL service
	_dalService    interfaces.IDalService
	_equipsService interfaces.IEquipsService
	_dalCh         chan *models.RawMqttMessage
	_webSockCh     chan *models.RawMqttMessage

	// mqtt connections map
	// key - topic
	// value - mqtt client
	_mqttConnections map[string]interfaces.IMqttClient
	_topicStorage    interfaces.ITopicStorage
}

// MqttReceiverServiceNew creates an instance of mqttReceiverService
func MqttReceiverServiceNew(
	ioCProvider interfaces.IIoCProvider,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	topicStorage interfaces.ITopicStorage,
	dalCh chan *models.RawMqttMessage,
	webSockCh chan *models.RawMqttMessage) interfaces.IMqttReceiverService {
	service := &mqttReceiverService{}

	service._ioCProvider = ioCProvider
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._equipsService = equipsService
	service._topicStorage = topicStorage
	service._dalCh = dalCh
	service._webSockCh = webSockCh
	service._mqttConnections = map[string]interfaces.IMqttClient{}

	return service
}

//UpdateMqttConnections updates mqtt connections map for an equipment connection state
func (service *mqttReceiverService) UpdateMqttConnections(state *models.EquipConnectionState) {
	rootTopic := state.Name
	isOff := !state.Connected
	topics := service._topicStorage.GetTopics()
	mqttConnections := service._mqttConnections
	ioCProvider := service._ioCProvider
	// dalService := service._dalService
	equipsService := service._equipsService

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
		if !equipsService.CheckEquipment(rootTopic) {
			go service.SendCommand(rootTopic, "equipInfo")
		}
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
