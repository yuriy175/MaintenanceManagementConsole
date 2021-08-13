package bl

import (
	"encoding/json"
	"io/ioutil"

	"ServerConsole/interfaces"
	"ServerConsole/models"
)

// settings service implementation type
type settingsService struct {
	//logger
	_log interfaces.ILogger

	//rabbitMQ settings
	RabbitMQ models.RabbitMQSettingsModel `json:"RabbitMQ"`

	//mongodb settings
	MongoDB models.MongoDBSettingsModel `json:"MongoDB"`

	// http server connection string
	HTTPServerConnectionString string `json:"HTTPServerAddress"`

	// web socket server connection string
	WebSocketServerConnectionString string `json:"WSServerAddress"`
}

// SettingsServiceNew creates an instance of settingsService
func SettingsServiceNew(log interfaces.ILogger) interfaces.ISettingsService {
	data, err := ioutil.ReadFile("settings.json")
	var service settingsService // Models.RabbitMQSettingsModel //
	json.Unmarshal(data, &service)

	if err != nil {
		log.Error("failed reading data from settings file")
	}

	service._log = log
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

// GetHTTPServerConnectionString returns http server connection string
func (service *settingsService) GetHTTPServerConnectionString() string {
	return service.HTTPServerConnectionString
}

// GetWebSocketServerConnectionString returns web socket server connection string
func (service *settingsService) GetWebSocketServerConnectionString() string {
	return service.WebSocketServerConnectionString
}
