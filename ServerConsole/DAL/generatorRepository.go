package dal

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"../models"
)

// generator info repository implementation type
type generatorRepository struct {
	// DAL service
	_dalService *dalService

	// mongodb db name
	_dbName string
}

// GeneratorRepositoryNew creates an instance of generatorRepository
func GeneratorRepositoryNew(
	dalService *dalService,
	dbName string) *generatorRepository {
	repository := &generatorRepository{}

	repository._dalService = dalService
	repository._dbName = dbName

	return repository
}

// InsertGeneratorInfo inserts generator info into db
func (repository *generatorRepository) InsertGeneratorInfo(equipName string, data string) *models.RawDeviceInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	genInfoCollection := session.DB(repository._dbName).C(models.GeneratorInfoTableName)

	model := models.RawDeviceInfoModel{}

	model.ID = bson.NewObjectId()
	model.DateTime = time.Now()
	model.EquipName = equipName
	model.Data = data

	error := genInfoCollection.Insert(model)
	if error != nil {
		fmt.Println("error InsertGeneratorInfo")
	}

	return &model
}

// GetGeneratorInfo returns generator info from db by equipment name and within the specified dates
func (repository *generatorRepository) GetGeneratorInfo(equipName string, startDate time.Time, endDate time.Time) []models.RawDeviceInfoModel {
	//[]Models.GeneratorInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	genInfoCollection := session.DB(repository._dbName).C(models.GeneratorInfoTableName)

	// // критерий выборки
	query := repository._dalService.GetQuery(equipName, startDate, endDate)
	// // объект для сохранения результата
	genInfo := []models.RawDeviceInfoModel{} // GeneratorInfoModel{}
	genInfoCollection.Find(query).Sort("-datetime").All(&genInfo)

	return genInfo
}
