package mocks

import (
	"encoding/json"
	"time"

	"ServerConsole/interfaces"
	"ServerConsole/models"
)

// DAL service mock type
type dalServiceMock struct {
}

// DataLayerMockServiceNew creates an instance of dalServiceMock
func DataLayerMockServiceNew() interfaces.IDalService {
	service := &dalServiceMock{}
	return service
}

// Start starts the service
func (service *dalServiceMock) Start() {
}

func (service *dalServiceMock) GetStudiesInWork(equipName string, startDate time.Time, endDate time.Time) []models.StudyInWorkModel {
	return nil
}

func (service *dalServiceMock) GetSystemInfo(equipName string, startDate time.Time, endDate time.Time) *models.FullSystemInfoModel {
	return nil
}

func (service *dalServiceMock) GetOrganAutoInfo(equipName string, startDate time.Time, endDate time.Time) []models.OrganAutoInfoModel {
	return nil
}

func (service *dalServiceMock) GetGeneratorInfo(equipName string, startDate time.Time, endDate time.Time) []models.RawDeviceInfoModel {
	return nil
}

func (service *dalServiceMock) GetSoftwareInfo(equipName string, startDate time.Time, endDate time.Time) *models.FullSoftwareInfoModel {
	return nil
}

func (service *dalServiceMock) GetDicomInfo(equipName string, startDate time.Time, endDate time.Time) []models.DicomsInfoModel {
	return nil
}

func (service *dalServiceMock) GetStandInfo(equipName string, startDate time.Time, endDate time.Time) []models.RawDeviceInfoModel {
	return nil
}

func (service *dalServiceMock) GetPermanentSystemInfo(equipName string) *models.SystemInfoModel {
	return nil
}

func (service *dalServiceMock) GetPermanentSoftwareInfo(equipName string) *models.SoftwareInfoModel {
	return nil
}

func (service *dalServiceMock) UpdateUser(userVM *models.UserViewModel) *models.UserModel {
	return nil
}

func (service *dalServiceMock) GetUsers() []models.UserModel {
	return nil
}

func (service *dalServiceMock) GetUserByName(login string, email string, password string) *models.UserModel {
	return nil
}

func (service *dalServiceMock) CheckEquipment(equipName string) bool {
	return true
}

// GetOldEquipInfos returns all existing equipment infos
func (service *dalServiceMock) GetEquipInfos() []models.EquipInfoModel {
	return nil
}

// GetOldEquipInfos returns all renamed equipment infos
func (service *dalServiceMock) GetOldEquipInfos() []models.RenamedEquipInfoModel {
	return nil
}

// RenameEquip appends equipment to renamedequipment
func (service *dalServiceMock) RenameEquip(oldEquipName string) bool {
	return true
}

// UpdateEquipmentInfo updates equipment info in db
func (service *dalServiceMock) UpdateEquipmentInfo(equip *models.DetailedEquipInfoModel) {
}

func (service *dalServiceMock) InsertEquipInfo(equipName string, equipVM *models.EquipInfoViewModel) *models.EquipInfoModel {
	return &models.EquipInfoModel{
		EquipName:         equipName,
		HospitalName:      equipVM.HospitalName,
		HospitalAddress:   equipVM.HospitalAddress,
		HospitalLongitude: equipVM.HospitalLongitude,
		HospitalLatitude:  equipVM.HospitalLatitude,
		Disabled:          false,
		Renamed:           false,
	}
}

func (service *dalServiceMock) GetAllTableNamesInfo(equipName string) *models.AllDBTablesModel {
	return nil
}

func (service *dalServiceMock) GetTableContent(equipName string, tableType string, tableName string) []string {
	return nil
}

// GetDBSystemInfo returns permanent system info from db
func (service *dalServiceMock) GetDBSystemInfo(equipName string) []map[string]json.RawMessage {
	return nil
}

// GetDBSoftwareInfo returns permanent software info from db
func (service *dalServiceMock) GetDBSoftwareInfo(equipName string) *models.DBSoftwareInfoModel {
	return nil
}

// GetLastSeenInfo returns last event datetime from db
func (service *dalServiceMock) GetLastSeenInfo(equipName string) time.Time {
	return time.Now()
}

// DisableAllDBInfo disables all db info
func (service *dalServiceMock) DisableAllDBInfo(equipName string) {
}

// GetEvents returns all events from db
func (service *dalServiceMock) GetEvents(equipNames []string, startDate time.Time, endDate time.Time) []models.EventModel {
	return nil
}

// InsertEvent inserts events into db
func (service *dalServiceMock) InsertEvents(equipName string, msgType string, msgVms []models.MessageViewModel, msgDate *time.Time) []models.EventModel {
	return nil
}

// DisableEquipInfo disables an equipment
func (service *dalServiceMock) DisableEquipInfo(equipName string, disabled bool) {
}

// GetChatNotes returns all chat notes from db
func (service *dalServiceMock) GetChatNotes(equipNames []string) []models.ChatModel {
	return nil
}

// UpsertChatNote upserts a new chat note into db
func (service *dalServiceMock) UpsertChatNote(equipName string, msgType string, id string, message string,
	userLogin string, isInternal bool) *models.ChatModel {
	return nil
}

// DeleteChatNote hides a chat note in db
func (service *dalServiceMock) DeleteChatNote(equipName string, msgType string, id string) {
}

// GetState returns db server state
func (service *dalServiceMock) GetState() map[string]interface{} {
	return nil
}
