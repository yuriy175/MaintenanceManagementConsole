package DAL

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"../Models"
)

type equipInfoRepository struct {
	_dalService *dalService
	_dbName     string
}

func EquipInfoRepositoryNew(
	dalService *dalService,
	dbName string) *equipInfoRepository {
	repository := &equipInfoRepository{}

	repository._dalService = dalService
	repository._dbName = dbName

	return repository
}

func (repository *equipInfoRepository) InsertEquipInfo(equipName string, equipVM *Models.EquipInfoViewModel) *Models.EquipInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipInfoCollection := session.DB(repository._dbName).C(Models.EquipmentTableName)

	model := Models.EquipInfoModel{}

	model.Id = bson.NewObjectId()
	model.RegisterDate = time.Now()
	model.EquipName = equipName

	model.HospitalName = equipVM.HospitalName
	model.HospitalAddress = equipVM.HospitalAddress
	model.HospitalLongitude = equipVM.HospitalLongitude
	model.HospitalLatitude = equipVM.HospitalLatitude

	equipInfoCollection.Insert(model)

	return &model
}

func (repository *equipInfoRepository) GetEquipInfos() []Models.EquipInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipInfoCollection := session.DB(repository._dbName).C(Models.EquipmentTableName)

	// // критерий выборки
	query := bson.M{}

	// // объект для сохранения результата
	equipInfos := []Models.EquipInfoModel{}
	equipInfoCollection.Find(query).Sort("-registerdate").All(&equipInfos)

	return equipInfos
}

func (repository *equipInfoRepository) CheckEquipment(equipName string) bool {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipCollection := session.DB(repository._dbName).C(Models.EquipmentTableName)

	// // критерий выборки
	query := bson.M{"equipname": equipName}

	// // объект для сохранения результата
	equip := Models.EquipInfoModel{}
	equipCollection.Find(query).One(&equip)

	return equip.Id != ""
}
