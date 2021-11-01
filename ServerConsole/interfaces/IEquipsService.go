package interfaces

import (
	"ServerConsole/models"
)

// IEquipsService describes  equipment service interface
type IEquipsService interface {
	// Starts the service
	Start()

	// InsertEquipInfo inserts a new equipment
	InsertEquipInfo(equipName string, equipVM *models.EquipInfoViewModel) *models.DetailedEquipInfoModel

	// CheckEquipment checks if the equipment exists
	CheckEquipment(equipName string) bool

	// GetEquipInfos returns all equipments
	GetEquipInfos(withDisabled bool) []models.DetailedEquipInfoModel

	// DisableEquipInfo disables an equipment
	DisableEquipInfo(equipName string, disabled bool)

	// GetFullInfo returns full equipment permanent info
	GetFullInfo(equipName string) *models.FullEquipInfoModel

	// GetOldEquipNames returns out of date equipment names
	GetOldEquipNames(equipName string) []string

	// SetLastSeen sets last seen event time
	SetLastSeen(equipName string)

	// SetActivate sets whether equipment is active
	SetActivate(equipName string, isOn bool)

	// UpdateEquipmentDetails updates an equipment details
	UpdateEquipmentDetails(equipName string, equipVM *models.EquipDetailsViewModel)
}
