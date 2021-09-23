package mocks

import (
	"ServerConsole/interfaces"
	"ServerConsole/models"
	"ServerConsole/bl"
)

// EventsMockServiceNew creates an instance of equipsService
func EventsMockServiceNew(
	log interfaces.ILogger,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	webSockCh chan *models.RawMqttMessage,
	eventsCh chan *models.RawMqttMessage,
	internalEventsCh chan *models.MessageViewModel) interfaces.IEventsService {
		return bl.EventsServiceNew(
			log,
			webSocketService,
			dalService,
			equipsService,
			webSockCh,
			eventsCh,
			internalEventsCh)
}
