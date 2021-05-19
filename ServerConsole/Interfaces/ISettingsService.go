package interfaces

import (
	"../models"
)

// settings service interface
type ISettingsService interface {
	GetMongoDBSettings() *models.MongoDBSettingsModel
	GetRabbitMQSettings() *models.RabbitMQSettingsModel
}
