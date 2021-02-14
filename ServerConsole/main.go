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
	equipDalCh := make(chan *Models.EquipmentMessage)

	go DAL.DalWorker(equipDalCh)
	// go BL.RabbitMqReceiver(equipDalCh)
	go BL.MqttReceiver(equipDalCh)
	go BL.HttpServer()

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
