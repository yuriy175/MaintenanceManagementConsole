package dal

import (
	"time"
	// "fmt"

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

// InsertEvent inserts events into db
func (repository *EventsRepository) InsertEvents(equipName string, msgType string, msgVms []models.MessageViewModel, msgDate *time.Time) []models.EventModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	eventsCollection := session.DB(repository._dbName).C(models.EventsTableName)
	events := []models.EventModel{}

	for _, msg := range msgVms {
		dbEvent := repository.insertEvent(eventsCollection, equipName, msgType, &msg, msgDate)
		events = append(events, *dbEvent)
	}

	return events
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
func (repository *EventsRepository) InsertConnectEvent(equipName string) *models.EventModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	eventsCollection := session.DB(repository._dbName).C(models.EventsTableName)

	msg := models.MessageViewModel{equipName, "подключен", ""}
	model := repository.insertEvent(eventsCollection, equipName, "EquipConnected", &msg, nil)

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
	vm *models.MessageViewModel,
	msgDate *time.Time) *models.EventModel {
	model := models.EventModel{}

	model.ID = bson.NewObjectId()
	if msgDate == nil{
		model.DateTime = time.Now()
	} else {
		model.DateTime = *msgDate
	}
	model.EquipName = equipName

	model.Type = msgType
	model.Title = vm.Level
	model.Description = vm.Code
	model.Details = vm.Description

	eventsCollection.Insert(model)

	return &model
}
