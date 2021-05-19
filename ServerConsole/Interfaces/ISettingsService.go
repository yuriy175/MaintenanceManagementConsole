package interfaces

import (
	"../models"
)

// ISettingsService describes settings service interface
type ISettingsService interface {
	GetMongoDBSettings() *models.MongoDBSettingsModel
	GetRabbitMQSettings() *models.RabbitMQSettingsModel
}
