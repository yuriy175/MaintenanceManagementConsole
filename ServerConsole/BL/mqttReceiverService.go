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
	_webSocketService IWebSocketService
	_dalCh            chan *Models.RawMqttMessage
	_webSockCh        chan *Models.RawMqttMessage
	_mqttConnections  map[string]*MqttClient
}

//var mqttConnections = map[string]*MqttClient{}

func MqttReceiverServiceNew(
	webSocketService IWebSocketService,
	dalCh chan *Models.RawMqttMessage,
	webSockCh chan *Models.RawMqttMessage) IMqttReceiverService {
	service := &mqttReceiverService{}

	service._webSocketService = webSocketService
	service._dalCh = dalCh
	service._webSockCh = webSockCh
	service._mqttConnections = map[string]*MqttClient{}

	return service
}

func (service *mqttReceiverService) UpdateMqtt(state *Models.EquipConnectionState) {
	rootTopic := state.Name
	isOff := !state.Connected
	topicStorage := &TopicStorage{}
	topics := topicStorage.getTopics()
	mqttConnections := service._mqttConnections

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
			mqttConnections[rootTopic] = CreateMqttClient(rootTopic, topics, equipDalCh, equipWebSockCh, service, webSocketService)
		}()

		fmt.Println(rootTopic + " created")
	}
}

func (service *mqttReceiverService) CreateCommonConnections() {
	mqttConnections := service._mqttConnections

	mqttConnections[Models.CommonTopicPath] = CreateMqttClient(Models.CommonTopicPath, []string{}, equipDalCh, equipWebSockCh, service, webSocketService)
	mqttConnections[Models.BroadcastCommandsTopic] = CreateMqttClient(Models.BroadcastCommandsTopic, []string{}, equipDalCh, equipWebSockCh, service, webSocketService)
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
		if d.IsEquipment {
			keys[i] = k
			i++
		}
	}

	return keys
}
