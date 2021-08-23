package interfaces

// IMqttClient describes mqtt client interface
type IMqttClient interface {
	// Create initializes an instance of mqttClient
	Create(rootTopic string, subTopics []string) IMqttClient

	// Disconnect disconnects the client
	Disconnect()

	// SendCommand send command to a command topic
	SendCommand(command string)

	// SendEquipCommand send command to another equip command topic
	SendEquipCommand(equipment string, command string)

	// IsEquipTopic checks if root topic isn't common or broadcast
	IsEquipTopic() bool

	// SendChatMessage send message to a chat topic
	SendChatMessage(equipment string, user string, message string, isInternal bool)

	// GetLastAliveMessage returns the client is last alive message time
	// GetLastAliveMessage() time.Time
}
