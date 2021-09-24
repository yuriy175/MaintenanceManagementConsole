package mocks

import (
	"ServerConsole/bl"
	"ServerConsole/interfaces"
	"ServerConsole/models"
)

// MqttReceiverMockServiceNew creates an instance of mqttReceiverService
func MqttReceiverMockServiceNew(
	log interfaces.ILogger,
	outWriter interfaces.IOutputWriter,
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
		outWriter,
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
