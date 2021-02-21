package BL

import (
	"fmt"

	"../Models"
)

type MqttReceiverService struct {
}

var mqttConnections = map[string]*MqttClient{}

func (*MqttReceiverService) UpdateMqtt(equipDalCh chan *Models.EquipmentMessage, message *Models.EquipmentMessage) {
	topicStorage := &TopicStorage{}
	topics := topicStorage.getTopics()

	rootTopic := fmt.Sprintf("%s/%s", message.EquipName, message.EquipNumber)

	if client, ok := mqttConnections[rootTopic]; ok {
		fmt.Println(rootTopic + " already exists")
		if message.MsgType == Models.MsgTypeInstanceOff {
			go client.Disconnect()
			delete(mqttConnections, rootTopic)
			fmt.Println(rootTopic + " deleted")
		}

		return
	}

	if message.MsgType == Models.MsgTypeInstanceOn {
		//go MqttReceiver(topic, equipDalCh)
		go func() {
			mqttConnections[rootTopic] = CreateMqttClient(rootTopic, topics, equipDalCh)
		}()

		fmt.Println(rootTopic + " created")
	}
}

func (*MqttReceiverService) SendCommand(equipment string, command string) {
	if client, ok := mqttConnections[equipment]; ok {
		go client.SendCommand(command)
	}

	return
}
