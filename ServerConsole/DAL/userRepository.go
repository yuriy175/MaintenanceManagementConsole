package dal

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"../interfaces"
	"../models"
)

type userRepository struct {
	_dalService  *dalService
	_dbName      string
	_authService interfaces.IAuthService
}

func UserRepositoryNew(
	dalService *dalService,
	dbName string,
	authService interfaces.IAuthService) *userRepository {
	repository := &userRepository{}

	repository._dalService = dalService
	repository._dbName = dbName
	repository._authService = authService

	return repository
}

func (repository *userRepository) UpdateUser(userVM *models.UserViewModel) *models.UserModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	userCollection := session.DB(repository._dbName).C(models.UsersTableName)

	model := models.UserModel{}

	model.Login = userVM.Login
	model.Surname = userVM.Surname
	model.Role = userVM.Role
	model.Email = userVM.Email
	model.Disabled = userVM.Disabled

	if userVM.Id == "" {
		sum := repository._authService.GetSum(userVM.Password)

		model.Id = bson.NewObjectId()
		model.DateTime = time.Now()
		model.PasswordHash = sum

		userCollection.Insert(model)
	} else {
		userCollection.Update(
			bson.M{"login": model.Login},
			bson.D{
				{"$set", bson.D{{"disabled", model.Disabled}}}})
	}

	return &model
}

func (repository *userRepository) GetUsers() []models.UserModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	userCollection := session.DB(repository._dbName).C(models.UsersTableName)

	// // критерий выборки
	query := bson.M{}

	// // объект для сохранения результата
	users := []models.UserModel{}
	userCollection.Find(query).Sort("-datetime").All(&users)

	return users
}

func (repository *userRepository) GetUserByName(login string, email string, password string) *models.UserModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	userCollection := session.DB(repository._dbName).C(models.UsersTableName)

	// // критерий выборки
	query := bson.M{"login": login}
	if login == "" {
		query = bson.M{"email": email}
	}

	// // объект для сохранения результата
	user := models.UserModel{}
	userCollection.Find(query).One(&user)

	if ok := repository._authService.CheckSum(password, user.PasswordHash); ok && !user.Disabled {
		return &user
	}

	return nil
}
