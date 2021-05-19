package dal

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"../models"
)

// equipment info repository implementation type
type equipInfoRepository struct {
	_dalService *dalService
	_dbName     string
}

// EquipInfoRepositoryNew creates an instance of equipInfoRepository
func EquipInfoRepositoryNew(
	dalService *dalService,
	dbName string) *equipInfoRepository {
	repository := &equipInfoRepository{}

	repository._dalService = dalService
	repository._dbName = dbName

	return repository
}

// InsertEquipInfo inserts equipment info into db
func (repository *equipInfoRepository) InsertEquipInfo(equipName string, equipVM *models.EquipInfoViewModel) *models.EquipInfoModel {
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

func (repository *equipInfoRepository) GetEquipInfos() []models.EquipInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipInfoCollection := session.DB(repository._dbName).C(models.EquipmentTableName)

	// // критерий выборки
	query := bson.M{}

	// // объект для сохранения результата
	equipInfos := []models.EquipInfoModel{}
	equipInfoCollection.Find(query).Sort("-registerdate").All(&equipInfos)

	return equipInfos
}

func (repository *equipInfoRepository) CheckEquipment(equipName string) bool {
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
