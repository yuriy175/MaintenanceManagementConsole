package interfaces

import (
	"time"

	"../models"
)

// IDalService describes DAL service interface
type IDalService interface {
	Start()
	////
	GetStudiesInWork(equipName string, startDate time.Time, endDate time.Time) []models.StudyInWorkModel
	GetSystemInfo(equipName string, startDate time.Time, endDate time.Time) *models.FullSystemInfoModel
	GetOrganAutoInfo(equipName string, startDate time.Time, endDate time.Time) []models.OrganAutoInfoModel
	GetGeneratorInfo(equipName string, startDate time.Time, endDate time.Time) []models.RawDeviceInfoModel // GeneratorInfoModel
	GetSoftwareInfo(equipName string, startDate time.Time, endDate time.Time) *models.FullSoftwareInfoModel
	GetDicomInfo(equipName string, startDate time.Time, endDate time.Time) []models.DicomsInfoModel
	GetStandInfo(equipName string, startDate time.Time, endDate time.Time) []models.RawDeviceInfoModel

	GetPermanentSystemInfo(equipName string) *models.SystemInfoModel
	GetPermanentSoftwareInfo(equipName string) *models.SoftwareInfoModel

	//user repository
	UpdateUser(user *models.UserViewModel) *models.UserModel
	GetUsers() []models.UserModel
	GetUserByName(surname string, email string, password string) *models.UserModel

	//equip info repository
	CheckEquipment(equipName string) bool
	InsertEquipInfo(equipName string, equipVM *models.EquipInfoViewModel) *models.EquipInfoModel
	GetEquipInfos() []models.EquipInfoModel

	//all db repository
	GetAllTableNamesInfo(equipName string) *models.AllDBTablesModel
	GetTableContent(equipName string, tableType string, tableName string) string
}
