package interfaces

import (
	"../models"
)

// IWebSocketService describes web socket service interface
type IWebSocketService interface {
	Start()
	Activate(sessionUID string, activatedEquipInfo string, deactivatedEquipInfo string)
	UpdateWebClients(state *models.EquipConnectionState)
	HasActiveClients(topic string) bool
	ClientClosed(sessionUID string)
}
