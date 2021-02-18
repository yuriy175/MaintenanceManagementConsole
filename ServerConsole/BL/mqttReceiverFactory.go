package BL

import (
	"fmt"
	"time"

	"../Models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttConnections = map[string]mqtt.Client{}

func MqttReceiverFactory(equipDalCh chan *Models.EquipmentMessage, message *Models.EquipmentMessage) {
	topicStorage := &TopicStorage{}
	topics := topicStorage.getTopics()

	rootTopic := fmt.Sprintf("%s/%s", message.EquipName, message.EquipNumber)

	if _, ok := mqttConnections[rootTopic]; ok {
		fmt.Println(rootTopic + " already exists")
		if message.MsgType == Models.MsgTypeInstanceOff {
			mqttConnections[rootTopic].Disconnect(0)
			delete(mqttConnections, rootTopic)
			fmt.Println(rootTopic + " deleted")
		}

		return
	}

	if message.MsgType == Models.MsgTypeInstanceOn {
		//go MqttReceiver(topic, equipDalCh)
		go func() {
			mqttConnections[rootTopic] = MqttReceiver(rootTopic, topics, equipDalCh)
			time.Sleep(time.Second * 5)
			text := "epona hren"
			commandTopic := rootTopic + "/command"
			mqttConnections[rootTopic].Publish(commandTopic, 0, false, text)
			fmt.Println(commandTopic + " " + text)
		}()

		fmt.Println(rootTopic + " created")
	}
}
