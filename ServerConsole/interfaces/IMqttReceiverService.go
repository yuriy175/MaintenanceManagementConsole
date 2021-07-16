package interfaces

import (
	"../models"
)

// IMqttReceiverService describes mqtt receiver service interface
type IMqttReceiverService interface {
	//UpdateMqttConnections updates mqtt connections map for an equipment connection state
	UpdateMqttConnections(state *models.EquipConnectionState)

	// CreateCommonConnections reates common mqtt connections
	CreateCommonConnections()

	// SendCommand sends a command to equipment via mqtt
	SendCommand(equipment string, command string)

	// SendCommand sends a broadcast command to equipments via mqtt
	SendBroadcastCommand(command string)

	// GetConnectionNames returns connected equipment names
	GetConnectionNames() []string

	// PublishChatNote sends a chat note to equipment via mqtt
	PublishChatNote(equipment string, message string, user string) 

	// Activate activates a specified connection to equipment and deactivates the other
	Activate(activatedEquipInfo string, deactivatedEquipInfo string)
}
