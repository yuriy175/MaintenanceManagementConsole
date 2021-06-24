package dal

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"../interfaces"
	"../models"
)

// EquipInfoRepository describes equipment info repository implementation type
type EquipInfoRepository struct {
	//logger
	_log interfaces.ILogger

	// DAL service
	_dalService *dalService

	// mongodb db name
	_dbName string
}

// EquipInfoRepositoryNew creates an instance of EquipInfoRepository
func EquipInfoRepositoryNew(
	log interfaces.ILogger,
	dalService *dalService,
	dbName string) *EquipInfoRepository {
	repository := &EquipInfoRepository{}

	repository._log = log
	repository._dalService = dalService
	repository._dbName = dbName

	return repository
}

// InsertEquipInfo inserts equipment info into db
func (repository *EquipInfoRepository) InsertEquipInfo(equipName string, equipVM *models.EquipInfoViewModel) *models.EquipInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipInfoCollection := session.DB(repository._dbName).C(models.EquipmentTableName)

	model := models.EquipInfoModel{}

	model.ID = bson.NewObjectId()
	model.RegisterDate = time.Now()
	model.EquipName = equipName

	model.HospitalName = equipVM.HospitalName
	model.HospitalAddress = equipVM.HospitalAddress
	model.HospitalLongitude = equipVM.HospitalLongitude
	model.HospitalLatitude = equipVM.HospitalLatitude

	equipInfoCollection.Insert(model)

	return &model
}

// GetEquipInfos returns all equipment infos from db
func (repository *EquipInfoRepository) GetEquipInfos() []models.EquipInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipInfoCollection := session.DB(repository._dbName).C(models.EquipmentTableName)

	// // критерий выборки
	query := bson.M{}

	// // объект для сохранения результата
	equipInfos := []models.EquipInfoModel{}
	// equipInfoCollection.Find(query).Sort("-registerdate").All(&equipInfos)
	equipInfoCollection.Find(query).Sort("-equipname").All(&equipInfos)

	return equipInfos
}

// CheckEquipment checks if the equipment exists in db
func (repository *EquipInfoRepository) CheckEquipment(equipName string) bool {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipCollection := session.DB(repository._dbName).C(models.EquipmentTableName)

	// // критерий выборки
	query := bson.M{"equipname": equipName}

	// // объект для сохранения результата
	equip := models.EquipInfoModel{}
	equipCollection.Find(query).One(&equip)

	return equip.ID != ""
}

// DisableEquipInfo disables an equipment
func (repository *EquipInfoRepository) DisableEquipInfo(equipName string, disabled bool) {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipCollection := session.DB(repository._dbName).C(models.EquipmentTableName)	
	equipCollection.Update(
		bson.M{"equipname": equipName},
		bson.D{
			{"$set", bson.D{{"disabled", disabled}}}})
}