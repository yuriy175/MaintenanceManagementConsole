package BL

import (
	"encoding/json"
	"fmt"
	"log"

	"../Models"
	"github.com/streadway/amqp"
)

func RabbitMqReceiver(equipDalCh chan *Models.EquipmentMessage) {
	quitCh := make(chan int)

	conn, err := amqp.Dial(Models.RabbitMQConnectionString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		Models.MQInfoQueueName, // name
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

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			content := Models.EquipmentMessage{}
			json.Unmarshal([]byte(d.Body), &content)

			fmt.Printf("%+v\n", content)

			equipDalCh <- &content
			d.Ack(false)
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
