package main

import (
	"fmt"
	"log"

	"./BL"
	"./DAL"
	"./Models"
)

func main() {

	intCh := make(chan int)
	devicesDalCh := make(chan *Models.DeviceConnection)

	go DAL.DalWorker(devicesDalCh)
	go BL.RabbitMqReceiver(devicesDalCh)

	/*conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//
	err = ch.ExchangeDeclare(
		"HwConnectionStateArrived", // name
		"topic",                    // type
		false,                      // durable
		false,                      // auto-deleted
		false,                      // internal
		false,                      // no-wait
		nil,                        // arguments
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

	// if len(os.Args) < 2 {
	// 	log.Printf("Usage: %s [binding_key]...", os.Args[0])
	// 	os.Exit(0)
	// }
	//for _, s := range os.Args[1:] {
	log.Printf("Binding queue %s to exchange %s with routing key %s",
		q.Name, "logs_topic", "#")
	err = ch.QueueBind(
		q.Name,                     // queue name
		"#",                        // routing key
		"HwConnectionStateArrived", // exchange
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

	forever := make(chan bool)

	go func() { //c *mgo.Collection) {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			content := Models.DeviceConnection{}
			json.Unmarshal([]byte(d.Body), &content)

			fmt.Printf("%+v\n", content)

			devicesDalCh <- &content
		}
	}() //productCollection)

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

	// go a(5, intCh)
	// a(6)
	//<-intCh
	*/
	fmt.Println("Hello Go")
	<-intCh
	fmt.Println("app quitted")
}

func a(val int, intCh chan int) {
	fmt.Println(val)
	intCh <- val
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
