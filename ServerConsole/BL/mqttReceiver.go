package BL

import (
	"encoding/json"
	"fmt"

	"../Models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type MqttClient struct {
	Client mqtt.Client
	Topic  string
}

func CreateMqttClient(rootTopic string, subTopics []string, equipDalCh chan *Models.EquipmentMessage) *MqttClient {
	//quitCh := make(chan int)

	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected")
	}

	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %v", err)
	}

	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		payload := msg.Payload()
		fmt.Printf("Received message: %s from topic: %s\n", payload, msg.Topic())
		content := Models.EquipmentMessage{}
		json.Unmarshal([]byte(payload), &content)
		equipDalCh <- &content
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

	return &MqttClient{client, rootTopic}
}

func (client *MqttClient) Disconnect() {
	client.Client.Disconnect(0)
}

func (client *MqttClient) SendCommand(command string) {
	commandTopic := client.Topic + "/command"
	client.Client.Publish(commandTopic, 0, false, command)
	fmt.Println(commandTopic + " " + command)
}
