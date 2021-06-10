package bl

import (
	"encoding/json"
	"fmt"
	"strings"

	"../interfaces"
	"../models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

// mqtt client implementation type
type mqttClient struct {
	//logger
	_log interfaces.ILogger

	_settingsService interfaces.ISettingsService

	// mqtt receiver service
	_mqttReceiverService interfaces.IMqttReceiverService

	// web socket service
	_webSocketService    interfaces.IWebSocketService

	// chanel for DAL communications
	_dalCh               chan *models.RawMqttMessage

	// chanel for communications with websocket services
	_webSockCh           chan *models.RawMqttMessage
	
	// chanel for communications with equipment service
	_equipsCh chan *models.RawMqttMessage

	// chanel for communications with events service
	_eventsCh chan *models.RawMqttMessage

	// mqtt client
	_client      mqtt.Client

	// main topic
	_topic       string

	// is topic equipment
	_isEquipment bool
}

// MqttClientNew creates an instance of mqttClient
func MqttClientNew(
	log interfaces.ILogger,
	settingsService interfaces.ISettingsService,
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalCh chan *models.RawMqttMessage,
	webSockCh chan *models.RawMqttMessage,
	equipsCh chan *models.RawMqttMessage,
	eventsCh chan *models.RawMqttMessage) interfaces.IMqttClient {
	client := &mqttClient{}
	
	client._log = log
	client._settingsService = settingsService
	client._mqttReceiverService = mqttReceiverService
	client._webSocketService = webSocketService
	client._dalCh = dalCh
	client._webSockCh = webSockCh
	client._equipsCh = equipsCh
	client._eventsCh = eventsCh

	return client
}

// Create initializes an instance of mqttClient
func (client *mqttClient) Create(
	rootTopic string,
	subTopics []string) interfaces.IMqttClient {
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
		if topic == models.CommonTopicPath {
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
			state := &models.EquipConnectionState{newRootTopic, isOn}

			go client._mqttReceiverService.UpdateMqttConnections(state)
			go client._webSocketService.UpdateWebClients(state)
		} else {
			rawMsg := models.RawMqttMessage{topic, string(payload)}

			client._dalCh <- &rawMsg
			client._webSockCh <- &rawMsg
			client._eventsCh <- &rawMsg
			if strings.Contains(topic, "/hospital") {
				client._equipsCh <- &rawMsg
			}
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
	client._isEquipment = rootTopic != models.CommonTopicPath && rootTopic != models.BroadcastCommandsTopic

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

// Disconnect disconnects the client
func (client *mqttClient) Disconnect() {
	client._client.Disconnect(0)
}

// SendCommand send command to a command topic
func (client *mqttClient) SendCommand(command string) {
	commandTopic := client._topic + "/command"
	client._client.Publish(commandTopic, 0, false, command)
	fmt.Println("Sent command " + commandTopic + " " + command)
}

// IsEquipTopic checks if root topic isn't common or broadcast
func (client *mqttClient) IsEquipTopic() bool {
	return client._isEquipment
}
