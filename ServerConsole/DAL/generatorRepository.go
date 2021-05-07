package DAL

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"../Models"
)

type generatorRepository struct {
	_dalService *dalService
	_dbName     string
}

func GeneratorRepositoryNew(
	dalService *dalService,
	dbName string) *generatorRepository {
	repository := &generatorRepository{}

	repository._dalService = dalService
	repository._dbName = dbName

	return repository
}

func (repository *generatorRepository) InsertGeneratorInfo(equipName string, data string) *Models.RawDeviceInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	genInfoCollection := session.DB(repository._dbName).C(Models.GeneratorInfoTableName)

	model := Models.RawDeviceInfoModel{}

	model.Id = bson.NewObjectId()
	model.DateTime = time.Now()
	model.EquipName = equipName
	model.Data = data

	error := genInfoCollection.Insert(model)
	if error != nil {
		fmt.Println("error InsertGeneratorInfo")
	}

	return &model
}

func (repository *generatorRepository) GetGeneratorInfo(equipName string, startDate time.Time, endDate time.Time) []Models.RawDeviceInfoModel {
	//[]Models.GeneratorInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	genInfoCollection := session.DB(repository._dbName).C(Models.GeneratorInfoTableName)

	// // критерий выборки
	query := repository._dalService.GetQuery(equipName, startDate, endDate)
	// // объект для сохранения результата
	genInfo := []Models.RawDeviceInfoModel{} // GeneratorInfoModel{}
	genInfoCollection.Find(query).Sort("-datetime").All(&genInfo)

	return genInfo
}
