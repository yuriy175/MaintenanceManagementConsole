package interfaces

import (
	"../models"
)

// web socket service interface
type IWebSocketService interface {
	Start()
	Activate(sessionUid string, activatedEquipInfo string, deactivatedEquipInfo string)
	UpdateWebClients(state *models.EquipConnectionState)
	HasActiveClients(topic string) bool
	ClientClosed(sessionUid string)
}
