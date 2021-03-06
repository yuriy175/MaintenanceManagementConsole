package interfaces

// IMqttClient describes mqtt client interface
type IMqttClient interface {
	// Create initializes an instance of mqttClient
	Create(rootTopic string, subTopics []string) IMqttClient

	// Disconnect disconnects the client
	Disconnect()

	// SendCommand send command to a command topic
	SendCommand(command string)

	// IsEquipTopic checks if root topic isn't common or broadcast
	IsEquipTopic() bool

	// SendChatMessage send message to a chat topic
	SendChatMessage(user string, message string) 
}
