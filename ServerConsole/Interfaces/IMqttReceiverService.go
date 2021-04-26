package Interfaces

import (
	"../Models"
)

type IMqttReceiverService interface {
	///
	UpdateMqttConnections(state *Models.EquipConnectionState)
	CreateCommonConnections()
	SendCommand(equipment string, command string)
	SendBroadcastCommand(command string)
	GetConnectionNames() []string
	///
}
