package mocks

import (
	"ServerConsole/interfaces"
	"ServerConsole/models"
	"ServerConsole/bl"
)

// WebSocketMockServiceNew creates an instance of webSocketService
func WebSocketMockServiceNew(
	log interfaces.ILogger,
	ioCProvider interfaces.IIoCProvider,
	diagnosticService interfaces.IDiagnosticService,
	settingsService interfaces.ISettingsService,
	webSockCh chan *models.RawMqttMessage) interfaces.IWebSocketService {
	return bl.WebSocketServiceNew(
		log,
		ioCProvider,
		diagnosticService,
		settingsService,
		webSockCh)
}
