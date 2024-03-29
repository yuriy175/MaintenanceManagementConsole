package interfaces

import (
	"ServerConsole/models"
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
	PublishChatNote(equipment string, message string, user string, isInternal bool)

	// Activate activates a specified connection to equipment and deactivates the other
	Activate(activatedEquipInfo string, deactivatedEquipInfo string)

	//ReconnectMqttConnectionIfAbsent sends reconnect command  if connection is absent in connections map
	// ReconnectMqttConnectionIfAbsent(equipment string)

	// SetKeepAliveReceived sets keepalive message from equipment
	SetKeepAliveReceived(equipment string)
}
