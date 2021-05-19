package dal

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"../models"
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

func (repository *standRepository) InsertStandInfo(equipName string, data string) *models.RawDeviceInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	genInfoCollection := session.DB(repository._dbName).C(models.StandInfoTableName)

	model := models.RawDeviceInfoModel{}

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

func (repository *standRepository) GetStandInfo(equipName string, startDate time.Time, endDate time.Time) []models.RawDeviceInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	standInfoCollection := session.DB(repository._dbName).C(models.StandInfoTableName)

	query := repository._dalService.GetQuery(equipName, startDate, endDate)

	standInfo := []models.RawDeviceInfoModel{}
	standInfoCollection.Find(query).Sort("-datetime").All(&standInfo)

	return standInfo
}
