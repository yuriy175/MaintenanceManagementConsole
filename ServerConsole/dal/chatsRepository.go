package dal

import (
	"time"
	// "fmt"

	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../interfaces"
	"../models"
)

// ChatsRepository describes chat notes repository implementation type
type ChatsRepository struct {
	//logger
	_log interfaces.ILogger

	// DAL service
	_dalService *dalService

	// mongodb db name
	_dbName string
}

// ChatsRepositoryNew creates an instance of ChatsRepository
func ChatsRepositoryNew(
	log interfaces.ILogger,
	dalService *dalService,
	dbName string) *ChatsRepository {
	repository := &ChatsRepository{}

	repository._log = log
	repository._dalService = dalService
	repository._dbName = dbName

	return repository
}

// GetChatNotes returns all chat notes from db
func (repository *ChatsRepository) GetChatNotes(equipNames []string) []models.ChatModel {
	service := repository._dalService
	session := service.CreateSession()
	defer session.Close()

	chatsCollection := session.DB(repository._dbName).C(models.ChatsTableName)

	// query := bson.M{"equipname": equipName, "hidden": false}
	query := bson.M{"$and": []bson.M{
		bson.M{"equipname": bson.M{"$in": equipNames}},
		bson.M{"hidden": false}}}

	// // объект для сохранения результата
	chats := []models.ChatModel{}
	chatsCollection.Find(query).Sort("-datetime").All(&chats)

	return chats
}

// UpsertChatNote upserts a new chat note into db
func (repository *ChatsRepository) UpsertChatNote(equipName string, msgType string, id string, message string, 
	userLogin string, isInternal bool) *models.ChatModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	chatsCollection := session.DB(repository._dbName).C(models.ChatsTableName)

	model := models.ChatModel{}
	model.Hidden = false
	model.EquipName = equipName
	model.Type = msgType	
	model.User = userLogin
	model.Message = message
	model.IsInternal = isInternal

	if id == ""{
		model.ID = bson.NewObjectId()
		model.DateTime = time.Now()	

		error := chatsCollection.Insert(model)
		if error != nil {
			repository._log.Errorf("error InsertChatNote")
		}
	} else {
		model.ID = bson.ObjectIdHex(id) 
		newBson := bson.D{
			{"message", model.Message},
		}

		err := chatsCollection.Update(
			bson.M{"_id": model.ID }, 
			bson.D{
				{"$set", newBson}})
		if err != nil{
			repository._log.Errorf("UpdateUser error %v", err)
		}
	}

	return &model
}

// DeleteChatNote hides a chat note in db
func (repository *ChatsRepository) DeleteChatNote(equipName string, msgType string, id string) {
	session := repository._dalService.CreateSession()
	defer session.Close()

	chatsCollection := session.DB(repository._dbName).C(models.ChatsTableName)

	if id != ""{
		hexId := bson.ObjectIdHex(id) 
		newBson := bson.D{
			{"hidden", true},
		}

		err := chatsCollection.Update(
			bson.M{"_id": hexId }, 
			bson.D{
				{"$set", newBson}})
		if err != nil{
			repository._log.Errorf("DeleteChatNote error %v", err)
		}
	}
}

