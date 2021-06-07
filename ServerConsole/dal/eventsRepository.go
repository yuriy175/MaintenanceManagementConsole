package dal

import (
	"time"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../interfaces"
	"../models"
)

// EventsRepository describes events repository implementation type
type EventsRepository struct {
	//logger
	_log interfaces.ILogger

	// DAL service
	_dalService *dalService

	// mongodb db name
	_dbName string
}

// EventsRepositoryNew creates an instance of EventsRepository
func EventsRepositoryNew(
	log interfaces.ILogger,
	dalService *dalService,
	dbName string) *EventsRepository {
	repository := &EventsRepository{}

	repository._log = log
	repository._dalService = dalService
	repository._dbName = dbName

	return repository
}

// InsertEvent inserts equipment info into db
func (repository *EventsRepository) InsertEvent(equipName string, data string) {
	session := repository._dalService.CreateSession()
	defer session.Close()

	eventsCollection := session.DB(repository._dbName).C(models.EventsTableName)

	viewmodel := models.SoftwareMessageViewModel{}
	json.Unmarshal([]byte(data), &viewmodel)

	for _, msg := range viewmodel.ErrorDescriptions {
		repository.insertEvent(eventsCollection, equipName, "ErrorDescriptions", &msg)
	}

	for _, msg := range viewmodel.AtlasErrorDescriptions {
		repository.insertEvent(eventsCollection, equipName, "AtlasErrorDescriptions", &msg)
	}
}


// GetEvents returns all events from db
func (repository *EventsRepository) GetEvents(equipName string, startDate time.Time, endDate time.Time) []models.EventModel {
	service := repository._dalService
	session := service.CreateSession()
	defer session.Close()

	eventsCollection := session.DB(repository._dbName).C(models.EventsTableName)

	// // критерий выборки
	query := service.GetQuery(equipName, startDate, endDate)

	// // объект для сохранения результата
	events := []models.EventModel{}
	eventsCollection.Find(query).Sort("-datetime").All(&events)

	return events
}

// InsertEvent inserts equipment connection state info into db
func (repository *EventsRepository) InsertConnectEvent(equipName string)*models.EventModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	eventsCollection := session.DB(repository._dbName).C(models.EventsTableName)

	msg := models.MessageViewModel {equipName, "connected"}
	model := repository.insertEvent(eventsCollection, equipName, "EquipConnected", &msg)

	return model
}

/*
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
*/

func (repository *EventsRepository) insertEvent(
	eventsCollection *mgo.Collection, 
	equipName string, 
	msgType string, 
	vm *models.MessageViewModel) *models.EventModel {	
	model := models.EventModel{}

	model.ID = bson.NewObjectId()
	model.DateTime = time.Now()
	model.EquipName = equipName

	model.Type = msgType
	model.Title = vm.Code
	model.Description = vm.Description

	eventsCollection.Insert(model)

	return &model
}
