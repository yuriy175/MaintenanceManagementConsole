package main

import (
	"fmt"
	"log"

	"./BL"
)

func main() {

	intCh := make(chan int)

	ioc := BL.InitIoc()
	mqttReceiverService := ioc.GetMqttReceiverService()
	webSocketService := ioc.GetWebSocketService()
	dalService := ioc.GetDalService()
	httpService := ioc.GetHttpService()
	mqttReceiverService.CreateCommonConnections()

	go dalService.Start()
	// go BL.RabbitMqReceiver(mqttReceiverService, equipDalCh, equipWebSockCh)

	go webSocketService.Start()
	go httpService.Start()

	mqttReceiverService.SendBroadcastCommand("reconnect")

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
