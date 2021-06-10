package interfaces

import (
	"../models"
)

// IWebSocketService describes web socket service interface
type IWebSocketService interface {
	// Start starts the service
	Start()

	// Activate activates a specified connection to equipment and deactivates the other
	Activate(sessionUID string, activatedEquipInfo string, deactivatedEquipInfo string)

	// UpdateWebClients notifies UI of a new equipment connection
	UpdateWebClients(state *models.EquipConnectionState)

	// HasActiveClients checks if there is an active connections
	HasActiveClients(topic string) bool

	// ClientClosed removes web socket connection
	ClientClosed(sessionUID string)

	// SendEvents sends events to all web connections
	SendEvents(events []models.EventModel) 
}
