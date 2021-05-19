package main

import (
	"fmt"
	"log"

	"./bl"
)

func main() {

	intCh := make(chan int)

	ioc := bl.InitIoc()
	mqttReceiverService := ioc.GetMqttReceiverService()
	webSocketService := ioc.GetWebSocketService()
	dalService := ioc.GetDalService()
	equipsService := ioc.GetEquipsService()
	httpService := ioc.GetHttpService()
	mqttReceiverService.CreateCommonConnections()

	go dalService.Start()
	go equipsService.Start()
	// go BL.RabbitMqReceiver(mqttReceiverService, equipDalCh, equipWebSockCh)

	go webSocketService.Start()
	go httpService.Start()

	go mqttReceiverService.SendBroadcastCommand("reconnect")

	fmt.Println("App Go started")
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
