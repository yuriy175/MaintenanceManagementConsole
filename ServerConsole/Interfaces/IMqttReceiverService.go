package interfaces

import (
	"../models"
)

// mqtt receiver service interface
type IMqttReceiverService interface {
	///
	UpdateMqttConnections(state *models.EquipConnectionState)
	CreateCommonConnections()
	SendCommand(equipment string, command string)
	SendBroadcastCommand(command string)
	GetConnectionNames() []string
	///
}
