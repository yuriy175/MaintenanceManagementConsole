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
}
