package dal

import (
	"time"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"

	"ServerConsole/interfaces"
	"ServerConsole/models"
)

// UserRepository describes user repository implementation type
type UserRepository struct {
	//logger
	_log interfaces.ILogger

	// DAL service
	_dalService *dalService

	// mongodb db name
	_dbName string

	// authorization service
	_authService interfaces.IAuthService
}

// UserRepositoryNew creates an instance of UserRepository
func UserRepositoryNew(
	log interfaces.ILogger,
	dalService *dalService,
	dbName string,
	authService interfaces.IAuthService) *UserRepository {
	repository := &UserRepository{}

	repository._log = log
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
		model.ID = bson.ObjectIdHex(userVM.ID)
		newBson := bson.D{
			{"disabled", model.Disabled},
			{"login", model.Login},
			{"role", model.Role},
			{"surname", model.Surname},
			{"email", model.Email},
		}

		if userVM.Password != "" {
			sum := repository._authService.GetSum(userVM.Password)
			newBson = append(newBson, bson.DocElem{"passwordhash", sum})
		}

		err := userCollection.Update(
			bson.M{"_id": model.ID},
			bson.D{
				{"$set", newBson}})
		if err != nil {
			repository._log.Errorf("UpdateUser error %v", err)
		}

		if model.Disabled {
			repository.EnsureAdmin()
		}
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
	if login == "" && email == "" {
		return nil
	}

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

func (repository *UserRepository) EnsureAdmin() {
	session := repository._dalService.CreateSession()
	defer session.Close()

	userCollection := session.DB(repository._dbName).C(models.UsersTableName)
	query := bson.M{}

	//критерий выборки
	query = bson.M{"role": models.AdminRoleName, "disabled": false}

	// // объект для сохранения результата
	user := models.UserModel{}
	userCollection.Find(query).One(&user)

	if user.ID != "" {
		//admin exists
		return
	}

	query = bson.M{"login": models.DefaultAdminName}
	userCollection.RemoveAll(query)

	userVM := models.UserViewModel{}

	userVM.Login = models.DefaultAdminName
	userVM.Surname = models.DefaultAdminName
	userVM.Role = models.AdminRoleName
	userVM.Disabled = false
	userVM.Password = "medtex"

	repository.UpdateUser(&userVM)
}
