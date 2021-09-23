package mocks

import (
	"ServerConsole/interfaces"
	"ServerConsole/models"
	"ServerConsole/bl"
)

// EquipsMockServiceNew creates an instance of equipsService
func EquipsMockServiceNew(
	log interfaces.ILogger,
	dalService interfaces.IDalService,
	equipCh chan *models.RawMqttMessage,
	internalEventsCh chan *models.MessageViewModel) interfaces.IEquipsService {
	return bl.EquipsServiceNew(log, dalService, equipCh, internalEventsCh)
}

