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
func (repository *ChatsRepository) GetChatNotes(equipName string) []models.ChatModel {
	service := repository._dalService
	session := service.CreateSession()
	defer session.Close()

	chatsCollection := session.DB(repository._dbName).C(models.ChatsTableName)

	// // критерий выборки
	query := bson.M{"equipname": equipName}

	// // объект для сохранения результата
	chats := []models.ChatModel{}
	chatsCollection.Find(query).Sort("-datetime").All(&chats)

	return chats
}

// InsertChatNote inserts a new chat note into db
func (repository *ChatsRepository) InsertChatNote(equipName string, msgType string, message string, userLogin string) *models.ChatModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	chatsCollection := session.DB(repository._dbName).C(models.ChatsTableName)

	model := models.ChatModel{}

	model.ID = bson.NewObjectId()
	model.DateTime = time.Now()
	model.EquipName = equipName
	model.Type = msgType
	model.User = userLogin
	model.Message= message

	error := chatsCollection.Insert(model)
	if error != nil {
		repository._log.Errorf("error InsertChatNote")
	}

	return &model
}

