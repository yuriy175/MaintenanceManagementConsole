package BL

import (
	"fmt"
	"log"

	"../Models"
	"github.com/streadway/amqp"
)

func RabbitMqReceiver(mqttReceiverService *MqttReceiverService, equipDalCh chan *Models.RawMqttMessage, equipWebSockCh chan *Models.RawMqttMessage) {
	quitCh := make(chan int)

	//RabbitMQConnectionString = "amqp://guest:guest@localhost:5672/"
	rabbitMQConnectionString := fmt.Sprintf("amqp://%s:%s@%s:5672", Models.RabbitMQUser, Models.RabbitMQPassword, Models.RabbitMQHost)

	conn, err := amqp.Dial(rabbitMQConnectionString)
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

	if msgs == nil {

	}
	/*go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			content := Models.EquipmentMessage{}
			json.Unmarshal([]byte(d.Body), &content)

			fmt.Printf("%+v\n", content)

			if content.MsgType == Models.MsgTypeInstanceOn || content.MsgType == Models.MsgTypeInstanceOff {
				//rootTopic := fmt.Sprintf("%s/%s", content.EquipName, content.EquipNumber)
				//mqttReceiverService.UpdateMqtt(equipDalCh, rootTopic)
			} else {
				equipDalCh <- &content
			}

			d.Ack(false)
		}
	}()*/

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

	<-quitCh
	fmt.Println("DalWorker quitted")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
