package Interfaces

type IMqttClient interface {
	Create(rootTopic string, subTopics []string) IMqttClient
	Disconnect()
	SendCommand(command string)
	IsEquipTopic() bool
}
