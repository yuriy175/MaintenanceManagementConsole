package BL

import (
	"encoding/json"
	"fmt"

	"../Interfaces"
	"../Models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type mqttClient struct {
	_settingsService     Interfaces.ISettingsService
	_mqttReceiverService Interfaces.IMqttReceiverService
	_webSocketService    Interfaces.IWebSocketService
	_dalCh               chan *Models.RawMqttMessage
	_webSockCh           chan *Models.RawMqttMessage

	_client      mqtt.Client
	_topic       string
	_isEquipment bool
}

func MqttClientNew(
	settingsService Interfaces.ISettingsService,
	mqttReceiverService Interfaces.IMqttReceiverService,
	webSocketService Interfaces.IWebSocketService,
	dalCh chan *Models.RawMqttMessage,
	webSockCh chan *Models.RawMqttMessage) Interfaces.IMqttClient {
	client := &mqttClient{}
	client._settingsService = settingsService
	client._mqttReceiverService = mqttReceiverService
	client._webSocketService = webSocketService
	client._dalCh = dalCh
	client._webSockCh = webSockCh

	return client
}

func (client *mqttClient) Create(
	rootTopic string,
	subTopics []string) Interfaces.IMqttClient {
	//quitCh := make(chan int)

	var reconnectingHandler mqtt.ReconnectHandler = func(client mqtt.Client, optins *mqtt.ClientOptions) {
		fmt.Println("Reconnecting " + rootTopic)
	}

	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected " + rootTopic)
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
	}

	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %s %v", rootTopic, err)
	}

	var messagePubHandler mqtt.MessageHandler = func(c mqtt.Client, msg mqtt.Message) {
		payload := msg.Payload()
		topic := msg.Topic()
		//fmt.Printf("Received message: %s from topic: %s\n", payload, topic)
		if topic == Models.CommonTopicPath {
			fmt.Printf("Received message: %s from topic: %s\n", payload, topic)

			var content = map[string]string{}
			json.Unmarshal([]byte(payload), &content)

			newRootTopic := ""
			data := ""
			for k, d := range content {
				newRootTopic = k
				data = d
			}

			isOn := data == "on"
			state := &Models.EquipConnectionState{newRootTopic, isOn}

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

	rabbitMQSettings := client._settingsService.GetRabbitMQSettings()

	var broker = rabbitMQSettings.Host
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(uuid.New().String())
	opts.SetUsername(rabbitMQSettings.User)
	opts.SetPassword(rabbitMQSettings.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.OnReconnecting = reconnectingHandler

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client._client = c
	client._topic = rootTopic
	client._isEquipment = rootTopic != Models.CommonTopicPath && rootTopic != Models.BroadcastCommandsTopic

	// go func() {
	// 	var topics = map[string]byte{}
	// 	topics[rootTopic] = 0
	// 	for _, value := range subTopics {
	// 		topics[rootTopic+value] = 0
	// 	}

	// 	token := c.SubscribeMultiple(topics, nil) // callback MessageHandler)
	// 	token.Wait()
	// 	fmt.Printf("Subscribed to topic: %s", rootTopic)
	// }()

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
