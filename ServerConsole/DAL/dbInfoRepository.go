package dal

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"../models"
)

// equipment db info repository implementation type
type dbInfoRepository struct {
	_dalService *dalService
	_dbName     string
}

// DbInfoRepositoryNew creates an instance of dbInfoRepository
func DbInfoRepositoryNew(
	dalService *dalService,
	dbName string) *dbInfoRepository {
	repository := &dbInfoRepository{}

	repository._dalService = dalService
	repository._dbName = dbName

	return repository
}

// InsertDbInfoInfo inserts equipment db info into db
func (repository *dbInfoRepository) InsertDbInfoInfo(equipName string, data string) *models.AllDBInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	allInfoCollection := session.DB(repository._dbName).C(models.AllDBInfoTableName)

	model := models.AllDBInfoModel{}

	// // критерий выборки
	query := bson.M{"equipname": equipName}
	allInfoCollection.Find(query).One(&model)

	if model.ID != "" {
		return &model
	}

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
func (repository *dbInfoRepository) GetAllTableNamesInfo(equipName string) *models.AllDBTablesModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	allInfoCollection := session.DB(repository._dbName).C(models.AllDBInfoTableName)

	query := bson.M{"equipname": equipName}

	allInfo := models.AllDBInfoModel{}
	allInfoCollection.Find(query).Sort("datetime").One(&allInfo)

	tablesModel := models.AllDBTablesModel{}
	tablesModel.EquipName = allInfo.EquipName

	for k, _ := range allInfo.Hospital {
		tablesModel.Hospital = append(tablesModel.Hospital, k)
	}

	for k, _ := range allInfo.Software {
		tablesModel.Software = append(tablesModel.Software, k)
	}

	for k, _ := range allInfo.System {
		tablesModel.System = append(tablesModel.System, k)
	}

	for k, _ := range allInfo.Atlas {
		tablesModel.Atlas = append(tablesModel.Atlas, k)
	}

	return &tablesModel
}

// GetTableContent returns the content of the specified table from equipment
func (repository *dbInfoRepository) GetTableContent(equipName string, tableType string, tableName string) string {
	session := repository._dalService.CreateSession()
	defer session.Close()

	allInfoCollection := session.DB(repository._dbName).C(models.AllDBInfoTableName)

	query := bson.M{"equipname": equipName}

	allInfo := models.AllDBInfoModel{}
	allInfoCollection.Find(query).Sort("datetime").One(&allInfo)

	tablesModel := models.AllDBTablesModel{}
	tablesModel.EquipName = allInfo.EquipName

	if tableType == "Hospital" {
		for k, v := range allInfo.Hospital {
			if k == tableName {
				content := string(v)
				return content
			}
		}
	}

	if tableType == "System" {
		for k, v := range allInfo.System {
			if k == tableName {
				content := string(v)
				return content
			}
		}
	}

	if tableType == "Software" {
		for k, v := range allInfo.Software {
			if k == tableName {
				content := string(v)
				return content
			}
		}
	}

	if tableType == "Atlas" {
		for k, v := range allInfo.Atlas {
			if k == tableName {
				content := string(v)
				return content
			}
		}
	}

	return ""
}

func (repository *dbInfoRepository) GetDBSystemInfo(equipName string) map[string]json.RawMessage {
	session := repository._dalService.CreateSession()
	defer session.Close()

	allInfoCollection := session.DB(repository._dbName).C(models.AllDBInfoTableName)

	query := bson.M{"equipname": equipName}

	allInfo := models.AllDBInfoModel{}
	allInfoCollection.Find(query).Sort("-datetime").One(&allInfo)

	return allInfo.System
}

func (repository *dbInfoRepository) GetDBSoftwareInfo(equipName string) *models.DBSoftwareInfoModel {
	swInfo := models.DBSoftwareInfoModel{}

	return &swInfo
}
