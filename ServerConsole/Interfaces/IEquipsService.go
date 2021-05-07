package Interfaces

import (
	"../Models"
)

type IEquipsService interface {
	Start()

	InsertEquipInfo(equipName string, equipVM *Models.EquipInfoViewModel) *Models.EquipInfoModel
	CheckEquipment(equipName string) bool
	GetEquipInfos() []Models.EquipInfoModel
}
