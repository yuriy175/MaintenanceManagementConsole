package interfaces

import (
	"../models"
)

// IEquipsService describes  equipment service interface
type IEquipsService interface {
	// Starts the service
	Start()

	// InsertEquipInfo inserts a new equipment
	InsertEquipInfo(equipName string, equipVM *models.EquipInfoViewModel) *models.EquipInfoModel

	// CheckEquipment checks if the equipment exists
	CheckEquipment(equipName string) bool

	// GetEquipInfos returns all equipments
	GetEquipInfos(withDisabled bool) []models.EquipInfoModel

	// DisableEquipInfo disables an equipment
	DisableEquipInfo(equipName string, disabled bool) 

	// GetFullInfo returns full equipment permanent info
	GetFullInfo(equipName string)*models.FullEquipInfoModel

	// GetOldEquipNames returns out of date equipment names
	GetOldEquipNames(equipName string) []string 
}