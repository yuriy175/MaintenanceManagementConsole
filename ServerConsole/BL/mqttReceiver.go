package BL

import (
	"encoding/json"
	"fmt"
	"log"

	"../Models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

func MqttReceiver(topic string, equipDalCh chan *Models.EquipmentMessage) {
	quitCh := make(chan int)

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
		// topic := "topic/test"
		token := client.Subscribe(topic, 1, nil)
		token.Wait()
		fmt.Printf("Subscribed to topic: %s", topic)
	}()
	// sub(client)
	// publish(client)

	// client.Disconnect(250000)

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

	<-quitCh
	fmt.Println("DalWorker quitted")
}

/*func publish(client mqtt.Client) {
	num := 5
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d qwerty", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second * 5)
	}
}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}*/
