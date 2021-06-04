package dal

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"../interfaces"
	"../models"
)

// StandRepository describes stand repository implementation type
type StandRepository struct {
	//logger
	_log interfaces.ILogger

	// DAL service
	_dalService *dalService

	// mongodb db name
	_dbName string
}

// StandRepositoryNew creates an instance of standRepository
func StandRepositoryNew(
	log interfaces.ILogger,
	dalService *dalService,
	dbName string) *StandRepository {
	repository := &StandRepository{}

	repository._log = log
	repository._dalService = dalService
	repository._dbName = dbName

	return repository
}

// InsertStandInfo inserts stand info into db
func (repository *StandRepository) InsertStandInfo(equipName string, data string) *models.RawDeviceInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	genInfoCollection := session.DB(repository._dbName).C(models.StandInfoTableName)

	model := models.RawDeviceInfoModel{}

	model.ID = bson.NewObjectId()
	model.DateTime = time.Now()
	model.EquipName = equipName
	model.Data = data

	error := genInfoCollection.Insert(model)
	if error != nil {
		fmt.Println("error InsertStandInfo")
	}

	return &model
}

// GetStandInfo returns stand info from db by equipment name and within the specified dates
func (repository *StandRepository) GetStandInfo(equipName string, startDate time.Time, endDate time.Time) []models.RawDeviceInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	standInfoCollection := session.DB(repository._dbName).C(models.StandInfoTableName)

	query := repository._dalService.GetQuery(equipName, startDate, endDate)

	standInfo := []models.RawDeviceInfoModel{}
	standInfoCollection.Find(query).Sort("-datetime").All(&standInfo)

	return standInfo
}
