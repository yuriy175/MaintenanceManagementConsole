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
	equipWebSockCh := make(chan *Models.RawMqttMessage)

	mqttReceiverService := &BL.MqttReceiverService{}
	webSocketService := &BL.WebSocketService{}
	mqttReceiverService.CreateCommonConnection(equipDalCh, equipWebSockCh)

	go DAL.DalWorker(equipDalCh)
	go BL.RabbitMqReceiver(mqttReceiverService, equipDalCh, equipWebSockCh)
	// go BL.MqttReceiver(equipDalCh)

	go BL.WebServer(equipWebSockCh)
	go BL.HttpServer(mqttReceiverService, webSocketService)

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
