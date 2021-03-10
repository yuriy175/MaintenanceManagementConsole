package BL

import (
	"encoding/json"
	"fmt"

	"../Models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type MqttClient struct {
	Client      mqtt.Client
	Topic       string
	IsEquipment bool
}

func CreateMqttClient(
	rootTopic string,
	subTopics []string,
	equipDalCh chan *Models.RawMqttMessage,
	equipWebSockCh chan *Models.RawMqttMessage,
	mqttReceiverService *MqttReceiverService) *MqttClient {
	//quitCh := make(chan int)

	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected " + rootTopic)
	}

	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %s %v", rootTopic, err)
	}

	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		payload := msg.Payload()
		topic := msg.Topic()
		fmt.Printf("Received message: %s from topic: %s\n", payload, topic)
		if topic == Models.CommonTopicPath {
			var content = map[string]string{}
			json.Unmarshal([]byte(payload), &content)

			newRootTopic := ""
			data := ""
			for k, d := range content {
				newRootTopic = k
				data = d
			}

			mqttReceiverService.UpdateMqtt(newRootTopic, data == "off", equipDalCh, equipWebSockCh)
		} else {
			//content := Models.EquipmentMessage{}
			//json.Unmarshal([]byte(payload), &content)
			// equipDalCh <- &content
			rawMsg := Models.RawMqttMessage{topic, string(payload)}
			//json.Unmarshal([]byte(payload), &content)
			equipDalCh <- &rawMsg
			equipWebSockCh <- &rawMsg
		}
	}

	var broker = Models.RabbitMQHost
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(uuid.New().String())
	opts.SetUsername(Models.RabbitMQUser)
	opts.SetPassword(Models.RabbitMQPassword)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go func() {
		var topics = map[string]byte{}
		topics[rootTopic] = 0
		for _, value := range subTopics {
			topics[rootTopic+value] = 0
		}

		token := client.SubscribeMultiple(topics, nil) // callback MessageHandler)
		token.Wait()
		fmt.Printf("Subscribed to topic: %s", rootTopic)
	}()

	return &MqttClient{client, rootTopic, IsEquipTopic(rootTopic)}
}

func (client *MqttClient) Disconnect() {
	client.Client.Disconnect(0)
}

func (client *MqttClient) SendCommand(command string) {
	commandTopic := client.Topic + "/command"
	client.Client.Publish(commandTopic, 0, false, command)
	fmt.Println("Sent command " + commandTopic + " " + command)
}

func IsEquipTopic(rootTopic string) bool {
	return rootTopic != Models.CommonTopicPath && rootTopic != Models.BroadcastCommandsTopic
}
