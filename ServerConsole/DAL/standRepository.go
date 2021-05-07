package DAL

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"../Models"
)

type standRepository struct {
	_dalService *dalService
	_dbName     string
}

func StandRepositoryNew(
	dalService *dalService,
	dbName string) *standRepository {
	repository := &standRepository{}

	repository._dalService = dalService
	repository._dbName = dbName

	return repository
}

func (repository *standRepository) InsertStandInfo(equipName string, data string) *Models.RawDeviceInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	genInfoCollection := session.DB(repository._dbName).C(Models.StandInfoTableName)

	model := Models.RawDeviceInfoModel{}

	model.Id = bson.NewObjectId()
	model.DateTime = time.Now()
	model.EquipName = equipName
	model.Data = data

	error := genInfoCollection.Insert(model)
	if error != nil {
		fmt.Println("error InsertStandInfo")
	}

	return &model
}

func (repository *standRepository) GetStandInfo(equipName string, startDate time.Time, endDate time.Time) []Models.RawDeviceInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	standInfoCollection := session.DB(repository._dbName).C(Models.StandInfoTableName)

	query := repository._dalService.GetQuery(equipName, startDate, endDate)

	standInfo := []Models.RawDeviceInfoModel{}
	standInfoCollection.Find(query).Sort("-datetime").All(&standInfo)

	return standInfo
}
