package interfaces

import (
	"encoding/json"
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

	// obsolete
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
	
	// DisableEquipInfo disables an equipment
	DisableEquipInfo(equipName string, disabled bool) 

	//all db repository
	GetAllTableNamesInfo(equipName string) *models.AllDBTablesModel
	GetTableContent(equipName string, tableType string, tableName string) []string

	// GetDBSystemInfo returns permanent system info from db
	GetDBSystemInfo(equipName string) []map[string]json.RawMessage

	// GetDBSoftwareInfo returns permanent software info from db
	GetDBSoftwareInfo(equipName string) *models.DBSoftwareInfoModel

	// GetLastSeenInfo returns last event datetime from db
	GetLastSeenInfo(equipName string) time.Time

	// events repository
	// GetEvents returns all events from db
	GetEvents(equipName string, startDate time.Time, endDate time.Time) []models.EventModel

	// InsertEvent inserts events into db
	InsertEvents(equipName string, msgType string, msgVms []models.MessageViewModel, msgDate *time.Time) []models.EventModel

	// chats repository
	// GetChatNotes returns all chat notes from db
	GetChatNotes(equipName string) []models.ChatModel

	// UpsertChatNote upserts a new chat note into db
	UpsertChatNote(equipName string, msgType string, id string, message string, userLogin string) *models.ChatModel 

	// DeleteChatNote hides a chat note in db
	DeleteChatNote(equipName string, msgType string, id string)
}
