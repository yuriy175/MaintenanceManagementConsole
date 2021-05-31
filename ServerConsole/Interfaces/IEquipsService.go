package interfaces

import (
	"../models"
)

// IEquipsService describes  equipment service interface
type IEquipsService interface {
	// Starts the service
	Start()

	// Inserts a new equipment
	InsertEquipInfo(equipName string, equipVM *models.EquipInfoViewModel) *models.EquipInfoModel

	// Checks if the equipment exists
	CheckEquipment(equipName string) bool

	// Returns all equipments
	GetEquipInfos() []models.EquipInfoModel
}
