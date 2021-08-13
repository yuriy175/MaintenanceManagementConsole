package bl

import (
	"fmt"
	"log"

	"ServerConsole/interfaces"
	"ServerConsole/models"

	"github.com/streadway/amqp"
)

//RabbitMqReceiver creates rabbit mq receiver
func RabbitMqReceiver(
	settingsService interfaces.ISettingsService,
	mqttReceiverService *interfaces.IMqttReceiverService,
	equipDalCh chan *models.RawMqttMessage,
	equipWebSockCh chan *models.RawMqttMessage) {
	quitCh := make(chan int)

	//RabbitMQConnectionString = "amqp://guest:guest@localhost:5672/"
	rabbitMQSettings := settingsService.GetRabbitMQSettings()
	rabbitMQConnectionString := fmt.Sprintf("amqp://%s:%s@%s:5672", rabbitMQSettings.User, rabbitMQSettings.Password, rabbitMQSettings.Host)

	conn, err := amqp.Dial(rabbitMQConnectionString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		models.MQInfoQueueName, // name
		false,                  // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	if msgs == nil {

	}

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

	<-quitCh
	fmt.Println("DalWorker quitted")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
