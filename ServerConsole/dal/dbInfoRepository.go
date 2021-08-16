package dal

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"ServerConsole/interfaces"
	"ServerConsole/models"
)

// DbInfoRepository describes equipment db info repository implementation type
type DbInfoRepository struct {
	//logger
	_log interfaces.ILogger

	// DAL service
	_dalService *dalService

	//mongodb db name
	_dbName string
}

// DbInfoRepositoryNew creates an instance of dbInfoRepository
func DbInfoRepositoryNew(
	log interfaces.ILogger,
	dalService *dalService,
	dbName string) *DbInfoRepository {
	repository := &DbInfoRepository{}

	repository._log = log
	repository._dalService = dalService
	repository._dbName = dbName

	return repository
}

// InsertDbInfoInfo inserts equipment db info into db
func (repository *DbInfoRepository) InsertDbInfoInfo(equipName string, data string) *models.AllDBInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	allInfoCollection := session.DB(repository._dbName).C(models.AllDBInfoTableName)

	model := models.AllDBInfoModel{}

	// // критерий выборки
	/*query := bson.M{"equipname": equipName}
	allInfoCollection.Find(query).One(&model)

	if model.ID != "" {
		return &model
	}*/

	// viewmodel := models.AllDBInfoViewModel{}
	json.Unmarshal([]byte(data), &model)

	model.ID = bson.NewObjectId()
	model.DateTime = time.Now()
	model.EquipName = equipName
	// model.Data = data

	error := allInfoCollection.Insert(model)
	if error != nil {
		fmt.Println("error InsertDbInfoInfo")
	}

	return &model
}

// GetAllTableNamesInfo returns all table names from db for equipment name
func (repository *DbInfoRepository) GetAllTableNamesInfo(equipName string) *models.AllDBTablesModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	allInfoCollection := session.DB(repository._dbName).C(models.AllDBInfoTableName)

	query := bson.M{"equipname": equipName, "hidden": false}

	allInfo := models.AllDBInfoModel{}
	allInfoCollection.Find(query).Sort("datetime").One(&allInfo)

	tablesModel := models.AllDBTablesModel{}
	tablesModel.EquipName = allInfo.EquipName

	for k := range allInfo.Hospital {
		tablesModel.Hospital = append(tablesModel.Hospital, k)
	}

	for k := range allInfo.Software {
		tablesModel.Software = append(tablesModel.Software, k)
	}

	for k := range allInfo.System {
		tablesModel.System = append(tablesModel.System, k)
	}

	for k := range allInfo.Atlas {
		tablesModel.Atlas = append(tablesModel.Atlas, k)
	}

	return &tablesModel
}

// GetTableContent returns the content of the specified table from equipment
func (repository *DbInfoRepository) GetTableContent(equipName string, tableType string, tableName string) []string {
	session := repository._dalService.CreateSession()
	defer session.Close()

	allInfoCollection := session.DB(repository._dbName).C(models.AllDBInfoTableName)

	query := bson.M{"equipname": equipName, "hidden": false}

	allInfos := []models.AllDBInfoModel{}
	allInfoCollection.Find(query).Sort("datetime").All(&allInfos)

	tablesModel := models.AllDBTablesModel{}
	tablesModel.EquipName = equipName

	contents := []string{}
	if tableType == "Hospital" {
		for _, allInfo := range allInfos {
			for k, v := range allInfo.Hospital {
				if k == tableName {
					content := string(v)
					contents = append(contents, content)
					break
				}
			}
		}
	} else if tableType == "System" {
		for _, allInfo := range allInfos {
			for k, v := range allInfo.System {
				if k == tableName {
					content := string(v)
					contents = append(contents, content)
					break
				}
			}
		}
	} else if tableType == "Software" {
		for _, allInfo := range allInfos {
			for k, v := range allInfo.Software {
				if k == tableName {
					content := string(v)
					contents = append(contents, content)
					break
				}
			}
		}
	} else if tableType == "Atlas" {
		for _, allInfo := range allInfos {
			for k, v := range allInfo.Atlas {
				if k == tableName {
					content := string(v)
					contents = append(contents, content)
					break
				}
			}
		}
	}

	return contents
}

// GetDBSystemInfo returns permanent system info from db
func (repository *DbInfoRepository) GetDBSystemInfo(equipName string) []map[string]json.RawMessage {
	session := repository._dalService.CreateSession()
	defer session.Close()

	allInfoCollection := session.DB(repository._dbName).C(models.AllDBInfoTableName)

	// query := bson.M{"equipname": equipName, "hidden": false}
	// query := bson.M{"equipname": equipName}
	query := bson.M{"$and": []bson.M{
		bson.M{"equipname": equipName},
		bson.M{"hidden": false}}}

	maps := []map[string]json.RawMessage{}
	allInfos := []models.AllDBInfoModel{}
	allInfoCollection.Find(query).Sort("datetime").All(&allInfos)

	for _, allInfo := range allInfos {
		maps = append(maps, allInfo.System)
	}

	return maps
}

// GetDBSoftwareInfo returns permanent software info from db
func (repository *DbInfoRepository) GetDBSoftwareInfo(equipName string) *models.DBSoftwareInfoModel {
	swInfo := models.DBSoftwareInfoModel{}

	session := repository._dalService.CreateSession()
	defer session.Close()

	allInfoCollection := session.DB(repository._dbName).C(models.AllDBInfoTableName)

	query := bson.M{"equipname": equipName, "hidden": false}

	allInfos := []models.AllDBInfoModel{}
	allInfoCollection.Find(query).Sort("datetime").All(&allInfos)

	for _, allInfo := range allInfos {
		swInfo.Software = append(swInfo.Software, allInfo.Software)
		swInfo.Atlas = append(swInfo.Atlas, allInfo.Atlas)
	}

	return &swInfo
}

// DisableAllDBInfo disables all db info
func (repository *DbInfoRepository) DisableAllDBInfo(equipName string) {
	session := repository._dalService.CreateSession()
	defer session.Close()

	allInfoCollection := session.DB(repository._dbName).C(models.AllDBInfoTableName)

	_, err := allInfoCollection.UpdateAll(
		bson.M{"equipname": equipName},
		bson.D{
			{"$set", bson.D{{"hidden", true}}}})

	if err != nil {
		repository._log.Errorf("DisableAllDBInfo error %v", err)
	}
}
