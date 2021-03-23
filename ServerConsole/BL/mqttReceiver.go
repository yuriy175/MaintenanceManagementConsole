package BL

import (
	"encoding/json"
	"fmt"

	"../Models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type IMqttClient interface {
	Create(rootTopic string, subTopics []string) IMqttClient
	Disconnect()
	SendCommand(command string)
	IsEquipTopic() bool
}

type mqttClient struct {
	_mqttReceiverService IMqttReceiverService
	_webSocketService    IWebSocketService
	_dalCh               chan *Models.RawMqttMessage
	_webSockCh           chan *Models.RawMqttMessage

	_client      mqtt.Client
	_topic       string
	_isEquipment bool
}

func MqttClientNew(
	mqttReceiverService IMqttReceiverService,
	webSocketService IWebSocketService,
	dalCh chan *Models.RawMqttMessage,
	webSockCh chan *Models.RawMqttMessage) IMqttClient {
	client := &mqttClient{}
	client._mqttReceiverService = mqttReceiverService
	client._webSocketService = webSocketService
	client._dalCh = dalCh
	client._webSockCh = webSockCh

	return client
}

func (client *mqttClient) Create(
	rootTopic string,
	subTopics []string) IMqttClient {
	//quitCh := make(chan int)

	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected " + rootTopic)
	}

	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %s %v", rootTopic, err)
	}

	var messagePubHandler mqtt.MessageHandler = func(c mqtt.Client, msg mqtt.Message) {
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

			state := &Models.EquipConnectionState{newRootTopic, data == "on"}
			go client._mqttReceiverService.UpdateMqttConnections(state)
			go client._webSocketService.UpdateWebClients(state)
		} else {
			//content := Models.EquipmentMessage{}
			//json.Unmarshal([]byte(payload), &content)
			// equipDalCh <- &content
			rawMsg := Models.RawMqttMessage{topic, string(payload)}
			//json.Unmarshal([]byte(payload), &content)
			client._dalCh <- &rawMsg
			client._webSockCh <- &rawMsg
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

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client._client = c
	client._topic = rootTopic
	client._isEquipment = rootTopic != Models.CommonTopicPath && rootTopic != Models.BroadcastCommandsTopic

	go func() {
		var topics = map[string]byte{}
		topics[rootTopic] = 0
		for _, value := range subTopics {
			topics[rootTopic+value] = 0
		}

		token := c.SubscribeMultiple(topics, nil) // callback MessageHandler)
		token.Wait()
		fmt.Printf("Subscribed to topic: %s", rootTopic)
	}()

	return client
}

func (client *mqttClient) Disconnect() {
	client._client.Disconnect(0)
}

func (client *mqttClient) SendCommand(command string) {
	commandTopic := client._topic + "/command"
	client._client.Publish(commandTopic, 0, false, command)
	fmt.Println("Sent command " + commandTopic + " " + command)
}

func (client *mqttClient) IsEquipTopic() bool {
	return client._isEquipment
}
