package Interfaces

import (
	"../Models"
)

type ISettingsService interface {
	GetMongoDBSettings() *Models.MongoDBSettingsModel
	GetRabbitMQSettings() *Models.RabbitMQSettingsModel
}
