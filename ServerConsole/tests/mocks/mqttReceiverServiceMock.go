package mocks

import (
	"ServerConsole/interfaces"
	"ServerConsole/models"
	"ServerConsole/bl"
)

// MqttReceiverMockServiceNew creates an instance of mqttReceiverService
func MqttReceiverMockServiceNew(
	log interfaces.ILogger,
	ioCProvider interfaces.IIoCProvider,
	diagnosticService interfaces.IDiagnosticService,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	eventsService interfaces.IEventsService,
	topicStorage interfaces.ITopicStorage,
	dalCh chan *models.RawMqttMessage,
	webSockCh chan *models.RawMqttMessage,
	eventsCh chan *models.RawMqttMessage) interfaces.IMqttReceiverService {
		return bl.MqttReceiverServiceNew(log,
			ioCProvider,
			diagnosticService,
			webSocketService,
			dalService,
			equipsService,
			eventsService,
			topicStorage,
			dalCh,
			webSockCh,
			eventsCh)
}
