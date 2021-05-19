package interfaces

// mqtt client interface
type IMqttClient interface {
	Create(rootTopic string, subTopics []string) IMqttClient
	Disconnect()
	SendCommand(command string)
	IsEquipTopic() bool
}
