package dal

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"../interfaces"
	"../models"
)

// UserRepository describes user repository implementation type
type UserRepository struct {
	// DAL service
	_dalService *dalService

	// mongodb db name
	_dbName string

	// authorization service
	_authService interfaces.IAuthService
}

// UserRepositoryNew creates an instance of UserRepository
func UserRepositoryNew(
	dalService *dalService,
	dbName string,
	authService interfaces.IAuthService) *UserRepository {
	repository := &UserRepository{}

	repository._dalService = dalService
	repository._dbName = dbName
	repository._authService = authService

	return repository
}

// UpdateUser upserts user info into db
func (repository *UserRepository) UpdateUser(userVM *models.UserViewModel) *models.UserModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	userCollection := session.DB(repository._dbName).C(models.UsersTableName)

	model := models.UserModel{}

	model.Login = userVM.Login
	model.Surname = userVM.Surname
	model.Role = userVM.Role
	model.Email = userVM.Email
	model.Disabled = userVM.Disabled

	if userVM.ID == "" {
		sum := repository._authService.GetSum(userVM.Password)

		model.ID = bson.NewObjectId()
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

// GetUsers returns all users from db
func (repository *UserRepository) GetUsers() []models.UserModel {
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

// GetUserByName returns a valid user by login or email or nil
func (repository *UserRepository) GetUserByName(login string, email string, password string) *models.UserModel {
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
