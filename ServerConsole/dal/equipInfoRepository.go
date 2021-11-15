package dal

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"ServerConsole/interfaces"
	"ServerConsole/models"
	"ServerConsole/utils"
)

// EquipInfoRepository describes equipment info repository implementation type
type EquipInfoRepository struct {
	//logger
	_log interfaces.ILogger

	// DAL service
	_dalService *dalService

	// mongodb db name
	_dbName string
}

// EquipInfoRepositoryNew creates an instance of EquipInfoRepository
func EquipInfoRepositoryNew(
	log interfaces.ILogger,
	dalService *dalService,
	dbName string) *EquipInfoRepository {
	repository := &EquipInfoRepository{}

	repository._log = log
	repository._dalService = dalService
	repository._dbName = dbName

	return repository
}

// InsertEquipInfo inserts equipment info into db
func (repository *EquipInfoRepository) InsertEquipInfo(equipName string, equipVM *models.EquipInfoViewModel) *models.EquipInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipInfoCollection := session.DB(repository._dbName).C(models.EquipmentTableName)

	model := models.EquipInfoModel{}

	model.ID = bson.NewObjectId()
	model.RegisterDate = time.Now()
	model.EquipName = equipName

	model.HospitalName = equipVM.HospitalName
	model.HospitalAddress = equipVM.HospitalAddress
	model.HospitalLongitude = equipVM.HospitalLongitude
	model.HospitalLatitude = equipVM.HospitalLatitude

	equipInfoCollection.Insert(model)

	return &model
}

// GetEquipInfos returns all equipment infos from db
func (repository *EquipInfoRepository) GetEquipInfos() []models.EquipInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipInfoCollection := session.DB(repository._dbName).C(models.EquipmentTableName)

	// // критерий выборки
	query := bson.M{}

	// // объект для сохранения результата
	equipInfos := []models.EquipInfoModel{}
	equipInfoCollection.Find(query).Sort("-equipname").All(&equipInfos)

	/*equipInfoCollection.Update(
	bson.M{"equipname": "KRT/HOMEPC89_MYHOMEHDD"},
	bson.D{
		{"$set", bson.D{
			{"hospitalhame", "House where I have to work"},
			{"hospitaladdress", "Гагарина 32"},
			{"hospitallongitude", 30.3378224},
			{"hospitallatitude", 59.8923184},
		}}})*/

	return equipInfos
}

// GetOldEquipInfos returns all renamed equipment infos from db
func (repository *EquipInfoRepository) GetOldEquipInfos() []models.RenamedEquipInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	renamedInfoCollection := session.DB(repository._dbName).C(models.RenamedEquipmentTableName)

	// // критерий выборки
	query := bson.M{}

	// // объект для сохранения результата
	equipInfos := []models.RenamedEquipInfoModel{}
	renamedInfoCollection.Find(query).Sort("-hddnumber").All(&equipInfos)

	return equipInfos
}

// UpdateEquipmentInfo updates equipment info in db
func (repository *EquipInfoRepository) UpdateEquipmentInfo(equip *models.DetailedEquipInfoModel) {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipCollection := session.DB(repository._dbName).C(models.EquipmentTableName)

	equipCollection.Update(
		bson.M{"equipname": equip.EquipName},
		bson.D{
			{"$set", bson.D{
				{"hospitalhame", equip.HospitalName},
				{"hospitaladdress", equip.HospitalAddress},
				{"hospitallongitude", equip.HospitalLongitude},
				{"hospitallatitude", equip.HospitalLatitude},
			}}})
}

// UpdateEquipmentDetails updates equipment details in db
func (repository *EquipInfoRepository) UpdateEquipmentDetails(equipVM *models.EquipDetailsViewModel) {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipCollection := session.DB(repository._dbName).C(models.EquipmentTableName)

	// IsManuallySet
	equipCollection.Update(
		bson.M{"equipname": equipVM.EquipName},
		bson.D{
			{"$set", bson.D{
				{"hospitallongitude", equipVM.HospitalLongitude},
				{"hospitallatitude", equipVM.HospitalLatitude},
				{"equipalias", equipVM.EquipAlias},
				{"hospitalhame", equipVM.HospitalName},
				{"hospitaladdress", equipVM.HospitalAddress},
				{"ismanuallyset", true},
				{"hospitalzones", equipVM.HospitalZones},
			}}})
}

// RenameEquip appends equipment to renamedequipment
func (repository *EquipInfoRepository) RenameEquip(oldEquipName string) bool {
	session := repository._dalService.CreateSession()
	defer session.Close()

	renamedInfoCollection := session.DB(repository._dbName).C(models.RenamedEquipmentTableName)
	equipCollection := session.DB(repository._dbName).C(models.EquipmentTableName)

	equipCollection.Update(
		bson.M{"equipname": oldEquipName},
		bson.D{
			{"$set", bson.D{{"renamed", true}}}})

	hddNumber := utils.GetHddNumberFromEquip(oldEquipName)

	query := bson.M{"hddnumber": hddNumber}
	renamedEquip := models.RenamedEquipInfoModel{}
	renamedInfoCollection.Find(query).One(&renamedEquip)

	if renamedEquip.ID != "" {
		oldEquipNames := append(renamedEquip.OldEquipNames, oldEquipName)
		renamedInfoCollection.Update(
			bson.M{"hddnumber": hddNumber},
			bson.D{
				{"$set", bson.D{{"oldequipnames", oldEquipNames}}}})
	} else {
		renamedEquip.ID = bson.NewObjectId()
		renamedEquip.HddNumber = hddNumber
		renamedEquip.OldEquipNames = []string{oldEquipName}

		renamedInfoCollection.Insert(renamedEquip)
	}

	return true
}

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

// DisableEquipInfo disables an equipment
func (repository *EquipInfoRepository) DisableEquipInfo(equipName string, disabled bool) {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipCollection := session.DB(repository._dbName).C(models.EquipmentTableName)
	equipCollection.Update(
		bson.M{"equipname": equipName},
		bson.D{
			{"$set", bson.D{{"disabled", disabled}}}})
}

///

// GetOldEquipInfos returns all renamed equipment infos from db
func (repository *EquipInfoRepository) GetEquipCardInfo(equipName string) models.EquipCardInfoModel {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipCollection := session.DB(repository._dbName).C(models.EquipInfoTableName)

	// // критерий выборки
	query := bson.M{"equipname": equipName}

	// // объект для сохранения результата
	equipInfo := models.EquipCardInfoModel{}
	equipCollection.Find(query).One(&equipInfo)

	return equipInfo
}

// UpdateEquipCardInfo updates equipment info in db
func (repository *EquipInfoRepository) UpdateEquipCardInfo(equip *models.EquipCardInfoModel) {
	session := repository._dalService.CreateSession()
	defer session.Close()

	equipCollection := session.DB(repository._dbName).C(models.EquipInfoTableName)

	equipCollection.Update(
		bson.M{"equipname": equip.EquipName},
		bson.D{
			{"$set", bson.D{
				{"serialnum", equip.SerialNum},
				{"model", equip.Model},
				{"agreement", equip.Agreement},
				{"contactinfo", equip.ContactInfo},
				{"reparinfo", equip.ReparInfo},
				{"manufacturingdate", equip.ManufacturingDate},
				{"montagedate", equip.MontageDate},
			}}})
}
