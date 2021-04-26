package Interfaces

import (
	"../Models"
)

type IWebSocketService interface {
	Start()
	Activate(sessionUid string, activatedEquipInfo string, deactivatedEquipInfo string)
	UpdateWebClients(state *Models.EquipConnectionState)
	HasActiveClients(topic string) bool
	ClientClosed(sessionUid string)
}
