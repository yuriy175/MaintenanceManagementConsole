package BL

import (
	"fmt"

	"../Models"
)

type MqttReceiverService struct {
}

var mqttConnections = map[string]*MqttClient{}

func (service *MqttReceiverService) UpdateMqtt(rootTopic string, equipDalCh chan *Models.EquipmentMessage) {
	topicStorage := &TopicStorage{}
	topics := topicStorage.getTopics()

	/*if client, ok := mqttConnections[rootTopic]; ok {
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
	}*/
	go func() {
		mqttConnections[rootTopic] = CreateMqttClient(rootTopic, topics, equipDalCh, service)
	}()

	fmt.Println(rootTopic + " created")
}

func (service *MqttReceiverService) CreateCommonConnection(equipDalCh chan *Models.EquipmentMessage) {
	mqttConnections[Models.CommonTopicPath] = CreateMqttClient(Models.CommonTopicPath, []string{}, equipDalCh, service)
	return
}

func (*MqttReceiverService) SendCommand(equipment string, command string) {
	if client, ok := mqttConnections[equipment]; ok {
		go client.SendCommand(command)
	}

	return
}

func (*MqttReceiverService) GetConnectionNames() []string {
	// return []string{"first", "second"}
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
