package BL

import (
	"fmt"

	"../Models"
)

var mqttConnections = map[string]bool{}

func MqttReceiverFactory(equipDalCh chan *Models.EquipmentMessage, message *Models.EquipmentMessage) {
	topic := fmt.Sprintf("%s/%s", message.EquipName, message.EquipNumber)

	if _, ok := mqttConnections[topic]; ok {
		fmt.Println(topic + " already exists")
		if message.MsgType == Models.MsgTypeInstanceOff {
			delete(mqttConnections, topic)
			fmt.Println(topic + " deleted")
		}

		return
	}

	if message.MsgType == Models.MsgTypeInstanceOn {
		mqttConnections[topic] = true
		go MqttReceiver(topic, equipDalCh)

		fmt.Println(topic + " created")
	}
}
