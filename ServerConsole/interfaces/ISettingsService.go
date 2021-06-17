package interfaces

import (
	"../models"
)

// ISettingsService describes settings service interface
type ISettingsService interface {
	// GetRabbitMQSettings returns rabbitMQ settings
	GetMongoDBSettings() *models.MongoDBSettingsModel

	// GetMongoDBSettings returns mongodb settings
	GetRabbitMQSettings() *models.RabbitMQSettingsModel

	// GetHTTPServerConnectionString returns http server connection string
	GetHTTPServerConnectionString() string

	// GetWebSocketServerConnectionString returns web socket server connection string
	GetWebSocketServerConnectionString() string
}
