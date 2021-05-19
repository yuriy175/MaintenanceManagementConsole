package interfaces

// IMqttClient describes mqtt client interface
type IMqttClient interface {
	Create(rootTopic string, subTopics []string) IMqttClient
	Disconnect()
	SendCommand(command string)
	IsEquipTopic() bool
}
