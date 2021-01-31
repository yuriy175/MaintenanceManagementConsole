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
