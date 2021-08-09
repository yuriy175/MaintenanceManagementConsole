package bl

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"../interfaces"
	"../models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

// mqtt client implementation type
type mqttClient struct {
	//logger
	_log interfaces.ILogger

	// diagnostic service
	_diagnosticService interfaces.IDiagnosticService

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

	// chanel for communications with chat services
	_chatCh chan *models.RawMqttMessage

	// mqtt client
	_client      mqtt.Client

	// main topic
	_topic       string

	// is topic equipment
	_isEquipment bool

	// last alive message time
	_lastAliveMessage time.Time
}

// MqttClientNew creates an instance of mqttClient
func MqttClientNew(
	log interfaces.ILogger,
	diagnosticService interfaces.IDiagnosticService,
	settingsService interfaces.ISettingsService,
	mqttReceiverService interfaces.IMqttReceiverService,
	webSocketService interfaces.IWebSocketService,
	dalCh chan *models.RawMqttMessage,
	webSockCh chan *models.RawMqttMessage,
	equipsCh chan *models.RawMqttMessage,
	eventsCh chan *models.RawMqttMessage,
	chatCh chan *models.RawMqttMessage) interfaces.IMqttClient {
	client := &mqttClient{}
	
	client._log = log
	client._diagnosticService = diagnosticService
	client._settingsService = settingsService
	client._mqttReceiverService = mqttReceiverService
	client._webSocketService = webSocketService
	client._dalCh = dalCh
	client._webSockCh = webSockCh
	client._equipsCh = equipsCh
	client._eventsCh = eventsCh
	client._chatCh = chatCh

	client._lastAliveMessage = time.Now()

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
			actualRootTopic := rootTopic
			if rootTopic == models.CommonChatsPath {
				actualRootTopic = rootTopic + "/#"
			}

			topics[actualRootTopic] = 0
			for _, value := range subTopics {
				topics[rootTopic+value] = 0
			}

			token := client.SubscribeMultiple(topics, nil) // callback MessageHandler)
			token.Wait()
			fmt.Printf("Subscribed to topic: %s", actualRootTopic)
		}()
	}

	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %s %v", rootTopic, err)
	}

	var messagePubHandler mqtt.MessageHandler = func(c mqtt.Client, msg mqtt.Message) {
		payload := msg.Payload()
		topic := msg.Topic()
		if strings.Contains(topic, "HOME") {
			// client._equipsCh <- &rawMsg
			fmt.Printf("Received message: %s from topic: %s\n", payload, topic)
			t := 1
			t = t + 1
		}

		// 
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
		} else if strings.HasPrefix(topic, models.CommonChatsPath) { // strings.HasPrefix(topic, models.CommonChatsPath) {
			fmt.Printf("Received chat message: %s from topic: %s\n", payload, topic)
			charTopic := topic[len(models.CommonChatsPath)+1:len(topic)]
			if charTopic != models.CommonChat{
				charTopic = charTopic + "/chat"
			}
			rawMsg := models.RawMqttMessage{charTopic, string(payload), time.Now()}
			client._chatCh <- &rawMsg
		} else {
			rawMsg := models.RawMqttMessage{topic, string(payload), time.Now()}

			if strings.Contains(topic, "/keepalive") {
				client._lastAliveMessage = time.Now()
				return
			}

			/*if strings.Contains(topic, "/chat") {
				client._chatCh <- &rawMsg
				return
			}*/

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
	client._isEquipment = rootTopic != models.CommonTopicPath && 
						rootTopic != models.BroadcastCommandsTopic &&
						rootTopic != models.CommonChatsPath

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

// SendChatMessage send message to a chat topic
func (client *mqttClient) SendChatMessage(equipment string, user string, message string, isInternal bool) {
	// chatTopic := client._topic + "/chat"
	
	chatTopic := models.CommonChatsPath + "/" + equipment
	fmt.Printf("PublishChatNote from topic: %s %s %s\n", chatTopic, message, user)
	
	viewmodel := &models.ChatViewModel{message, user, isInternal}
    data, _ := json.Marshal(viewmodel)

	client._client.Publish(chatTopic, 0, false, string(data))
	fmt.Println("Sent chat " + chatTopic)
}

// GetLastAliveMessage returns the client is last alive message time
func (client *mqttClient) GetLastAliveMessage() time.Time {
	return client._lastAliveMessage
}