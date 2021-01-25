package BL

import (
	"fmt"
	"log"

	"encoding/json"

	"github.com/streadway/amqp"

	"../Models"
)

func RabbitMqReceiver(devicesDalCh chan *Models.DeviceConnection) {
	quitCh := make(chan int)

	conn, err := amqp.Dial(Models.RabbitMQConnectionString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//
	err = ch.ExchangeDeclare(
		Models.MQConnectionStateName, // name
		"topic",                      // type
		false,                        // durable
		false,                        // auto-deleted
		false,                        // internal
		false,                        // no-wait
		nil,                          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	log.Printf("Binding queue %s to exchange %s with routing key %s",
		q.Name, "logs_topic", "#")
	err = ch.QueueBind(
		q.Name,                       // queue name
		"#",                          // routing key
		Models.MQConnectionStateName, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")
	//}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			content := Models.DeviceConnection{}
			json.Unmarshal([]byte(d.Body), &content)

			fmt.Printf("%+v\n", content)

			devicesDalCh <- &content
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

	<-quitCh
	fmt.Println("DalWorker quitted")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
