package BL

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"../Interfaces"
	"../Models"
)

type settingsService struct {
	RabbitMQ Models.RabbitMQSettingsModel `json:"RabbitMQ"`
	MongoDB  Models.MongoDBSettingsModel  `json:"MongoDB"`
}

func SettingsServiceNew() Interfaces.ISettingsService {
	data, err := ioutil.ReadFile("settings.json")
	var service settingsService // Models.RabbitMQSettingsModel //
	json.Unmarshal(data, &service)

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	return &service
}

func (service *settingsService) GetRabbitMQSettings() *Models.RabbitMQSettingsModel {
	return &service.RabbitMQ
}

func (service *settingsService) GetMongoDBSettings() *Models.MongoDBSettingsModel {
	return &service.MongoDB
}
