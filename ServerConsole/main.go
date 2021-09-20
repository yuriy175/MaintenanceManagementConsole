package main

import (
	"fmt"

	"ServerConsole/bl"
)

func main() {

	fmt.Println("App Go started")

	intCh := make(chan int)

	ioc := bl.InitIoc()
	log := ioc.GetLogger()
	mqttReceiverService := ioc.GetMqttReceiverService()
	webSocketService := ioc.GetWebSocketService()
	dalService := ioc.GetDalService()
	equipsService := ioc.GetEquipsService()
	httpService := ioc.GetHTTPService()
	eventsService := ioc.GetEventsService()
	chatService := ioc.GetChatService()
	mqttReceiverService.CreateCommonConnections()

	go dalService.Start()
	go equipsService.Start()
	go eventsService.Start()
	// go BL.RabbitMqReceiver(mqttReceiverService, equipDalCh, equipWebSockCh)

	go webSocketService.Start()
	go httpService.Start()
	go chatService.Start()

	go mqttReceiverService.SendBroadcastCommand("reconnect")

	log.Info("App Go started")
	<-intCh
	log.Info("app quitted")
}

/*func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}*/
