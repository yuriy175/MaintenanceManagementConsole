package interfaces

import (
	"../models"
)

// equipment service interface
type IEquipsService interface {
	Start()

	InsertEquipInfo(equipName string, equipVM *models.EquipInfoViewModel) *models.EquipInfoModel
	CheckEquipment(equipName string) bool
	GetEquipInfos() []models.EquipInfoModel
}
