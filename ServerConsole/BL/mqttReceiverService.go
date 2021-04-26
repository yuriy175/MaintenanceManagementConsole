package BL

import (
	"fmt"
	"sync"

	"../Interfaces"
	"../Models"
)

type mqttReceiverService struct {
	_mtx              sync.RWMutex
	_ioCProvider      Interfaces.IIoCProvider
	_webSocketService Interfaces.IWebSocketService
	_dalCh            chan *Models.RawMqttMessage
	_webSockCh        chan *Models.RawMqttMessage
	_mqttConnections  map[string]Interfaces.IMqttClient
	_topicStorage     Interfaces.ITopicStorage
}

//var mqttConnections = map[string]*MqttClient{}

func MqttReceiverServiceNew(
	ioCProvider Interfaces.IIoCProvider,
	webSocketService Interfaces.IWebSocketService,
	topicStorage Interfaces.ITopicStorage,
	dalCh chan *Models.RawMqttMessage,
	webSockCh chan *Models.RawMqttMessage) Interfaces.IMqttReceiverService {
	service := &mqttReceiverService{}

	service._ioCProvider = ioCProvider
	service._webSocketService = webSocketService
	service._topicStorage = topicStorage
	service._dalCh = dalCh
	service._webSockCh = webSockCh
	service._mqttConnections = map[string]Interfaces.IMqttClient{}

	return service
}

func (service *mqttReceiverService) UpdateMqttConnections(state *Models.EquipConnectionState) {
	rootTopic := state.Name
	isOff := !state.Connected
	topics := service._topicStorage.GetTopics()
	mqttConnections := service._mqttConnections
	ioCProvider := service._ioCProvider

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
	}

	fmt.Println(rootTopic + " created")
}

func (service *mqttReceiverService) CreateCommonConnections() {
	mqttConnections := service._mqttConnections
	ioCProvider := service._ioCProvider

	service._mtx.Lock()
	defer service._mtx.Unlock()
	mqttConnections[Models.CommonTopicPath] = ioCProvider.GetMqttClient().Create(Models.CommonTopicPath, []string{})
	mqttConnections[Models.BroadcastCommandsTopic] = ioCProvider.GetMqttClient().Create(Models.BroadcastCommandsTopic, []string{})

	return
}

func (service *mqttReceiverService) SendCommand(equipment string, command string) {
	fmt.Printf("SendCommand from topic: %s %s\n", equipment, command)

	service._mtx.Lock()
	defer service._mtx.Unlock()

	if client, ok := service._mqttConnections[equipment]; ok {
		go client.SendCommand(command)
	}

	return
}

func (service *mqttReceiverService) SendBroadcastCommand(command string) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	if client, ok := service._mqttConnections[Models.BroadcastCommandsTopic]; ok {
		go client.SendCommand(command)
	}

	return
}

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
