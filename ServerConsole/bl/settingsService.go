package bl

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"../interfaces"
	"../models"
)

// settings service implementation type
type settingsService struct {
	//rabbitMQ settings
	RabbitMQ models.RabbitMQSettingsModel `json:"RabbitMQ"`

	//mongodb settings
	MongoDB  models.MongoDBSettingsModel  `json:"MongoDB"`
}

// SettingsServiceNew creates an instance of settingsService
func SettingsServiceNew() interfaces.ISettingsService {
	data, err := ioutil.ReadFile("settings.json")
	var service settingsService // Models.RabbitMQSettingsModel //
	json.Unmarshal(data, &service)

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	return &service
}

// GetRabbitMQSettings returns rabbitMQ settings
func (service *settingsService) GetRabbitMQSettings() *models.RabbitMQSettingsModel {
	return &service.RabbitMQ
}

// GetMongoDBSettings returns mongodb settings
func (service *settingsService) GetMongoDBSettings() *models.MongoDBSettingsModel {
	return &service.MongoDB
}
