package mocks

import (
	"ServerConsole/interfaces"
	"ServerConsole/models"
)

// mqtt client mock implementation type
type mqttClientMock struct {
	// main topic
	_topic string

	// is topic equipment
	_isEquipment bool
}

// MqttClientNew creates an instance of mqttClientMock
func MqttClientMockNew() interfaces.IMqttClient {
	client := &mqttClientMock{}
	return client
}

// Create initializes an instance of mqttClient
func (client *mqttClientMock) Create(
	rootTopic string,
	subTopics []string) interfaces.IMqttClient {
	client._topic = rootTopic
	client._isEquipment = rootTopic != models.CommonTopicPath &&
		rootTopic != models.BroadcastCommandsTopic &&
		rootTopic != models.CommonChatsPath &&
		rootTopic != models.CommonKeepAlive

	return client
}

// Disconnect disconnects the client
func (client *mqttClientMock) Disconnect() {
}

// SendCommand send command to a command topic
func (client *mqttClientMock) SendCommand(command string) {
}

// SendEquipCommand send command to another equip command topic
func (client *mqttClientMock) SendEquipCommand(equipment string, command string) {
}

// IsEquipTopic checks if root topic isn't common or broadcast
func (client *mqttClientMock) IsEquipTopic() bool {
	return client._isEquipment
}

// SendChatMessage send message to a chat topic
func (client *mqttClientMock) SendChatMessage(equipment string, user string, message string, isInternal bool) {
}

