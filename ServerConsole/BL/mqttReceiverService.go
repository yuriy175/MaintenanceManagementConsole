package BL

import (
	"fmt"

	"../Models"
)

type IMqttReceiverService interface {
	///
	UpdateMqtt(state *Models.EquipConnectionState)
	CreateCommonConnections()
	SendCommand(equipment string, command string)
	SendBroadcastCommand(command string)
	GetConnectionNames() []string
	///
}

type mqttReceiverService struct {
	_ioCProvider      IIoCProvider
	_webSocketService IWebSocketService
	_dalCh            chan *Models.RawMqttMessage
	_webSockCh        chan *Models.RawMqttMessage
	_mqttConnections  map[string]IMqttClient
}

//var mqttConnections = map[string]*MqttClient{}

func MqttReceiverServiceNew(
	ioCProvider IIoCProvider,
	webSocketService IWebSocketService,
	dalCh chan *Models.RawMqttMessage,
	webSockCh chan *Models.RawMqttMessage) IMqttReceiverService {
	service := &mqttReceiverService{}

	service._ioCProvider = ioCProvider
	service._webSocketService = webSocketService
	service._dalCh = dalCh
	service._webSockCh = webSockCh
	service._mqttConnections = map[string]IMqttClient{}

	return service
}

func (service *mqttReceiverService) UpdateMqtt(state *Models.EquipConnectionState) {
	rootTopic := state.Name
	isOff := !state.Connected
	topicStorage := &TopicStorage{}
	topics := topicStorage.getTopics()
	mqttConnections := service._mqttConnections
	ioCProvider := service._ioCProvider

	if client, ok := mqttConnections[rootTopic]; ok {
		fmt.Println(rootTopic + " already exists")
		if isOff {
			go client.Disconnect()
			delete(mqttConnections, rootTopic)
			fmt.Println(rootTopic + " deleted")
		}

		return
	}

	if !isOff {
		go func() {
			mqttConnections[rootTopic] = ioCProvider.GetMqttClient().Create(rootTopic, topics)
		}()

		fmt.Println(rootTopic + " created")
	}
}

func (service *mqttReceiverService) CreateCommonConnections() {
	mqttConnections := service._mqttConnections
	ioCProvider := service._ioCProvider

	mqttConnections[Models.CommonTopicPath] = ioCProvider.GetMqttClient().Create(Models.CommonTopicPath, []string{})
	mqttConnections[Models.BroadcastCommandsTopic] = ioCProvider.GetMqttClient().Create(Models.BroadcastCommandsTopic, []string{})
	return
}

func (service *mqttReceiverService) SendCommand(equipment string, command string) {
	if client, ok := service._mqttConnections[equipment]; ok {
		go client.SendCommand(command)
	}

	return
}

func (service *mqttReceiverService) SendBroadcastCommand(command string) {
	if client, ok := service._mqttConnections[Models.BroadcastCommandsTopic]; ok {
		go client.SendCommand(command)
	}

	return
}

func (service *mqttReceiverService) GetConnectionNames() []string {
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
