package dal

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../interfaces"
	"../models"
	"../utils"
)

// DAL service implementation type
type dalService struct {
	// chanel for DAL communications
	_dalCh chan *models.RawMqttMessage

	// authorization service
	_authService interfaces.IAuthService

	// settings service
	_settingsService interfaces.ISettingsService

	// mongodb db name
	_dbName string

	// mongodb connection string
	_connectionString string

	// user repository
	_userRepository *userRepository

	// equipment db info repository
	_equipInfoRepository *equipInfoRepository

	// generator info repository
	_generatorRepository *generatorRepository

	// stand info repository
	_standRepository  *standRepository
	_dbInfoRepository *dbInfoRepository
}

// DalServiceNew creates an instance of dalService
func DalServiceNew(
	authService interfaces.IAuthService,
	settingsService interfaces.ISettingsService,
	dalCh chan *models.RawMqttMessage) interfaces.IDalService {
	service := &dalService{}

	service._authService = authService
	service._settingsService = settingsService
	service._dalCh = dalCh

	service._dbName = settingsService.GetMongoDBSettings().DBName
	service._connectionString = settingsService.GetMongoDBSettings().ConnectionString

	service._userRepository = UserRepositoryNew(service, service._dbName, authService)
	service._equipInfoRepository = EquipInfoRepositoryNew(service, service._dbName)
	service._generatorRepository = GeneratorRepositoryNew(service, service._dbName)
	service._standRepository = StandRepositoryNew(service, service._dbName)
	service._dbInfoRepository = DbInfoRepositoryNew(service, service._dbName)

	return service
}

// Start starts the service
func (service *dalService) Start() {
	quitCh := make(chan int)

	session := service.CreateSession()
	defer session.Close()

	db := session.DB(service._dbName)
	// deviceCollection := db.C(models.DeviceConnectionsTableName)
	studiesCollection := db.C(models.StudyInWorkTableName)
	organAutoCollection := db.C(models.OrganAutoTableName)
	// genInfoCollection := db.C(models.GeneratorInfoTableName)
	sysInfoCollection := db.C(models.SystemInfoTableName)
	sysVolatileInfoCollection := db.C(models.SystemVolatileInfoTableName)
	softwareInfoCollection := db.C(models.SoftwareInfoTableName)
	softwareVolatileInfoCollection := db.C(models.SoftwareVolatileInfoTableName)
	dicomInfoCollection := db.C(models.DicomInfoTableName)
	// standInfoCollection := db.C(models.StandInfoTableName)

	service.ensureIndeces(sysInfoCollection, []string{"equipname", "datetime"})
	service.ensureIndeces(sysVolatileInfoCollection, []string{"equipname", "datetime"})

	service.ensureIndeces(softwareInfoCollection, []string{"equipname", "datetime"})
	service.ensureIndeces(softwareVolatileInfoCollection, []string{"equipname", "datetime"})

	go func() {
		for d := range service._dalCh {
			if strings.Contains(d.Topic, "/study") {
				model := models.StudyInWorkModel{}
				json.Unmarshal([]byte(d.Data), &model)
				model.ID = bson.NewObjectId()
				model.DateTime = time.Now()
				model.EquipName = utils.GetEquipFromTopic(d.Topic)
				model.State = 2

				studiesCollection.Insert(model)
			} else if strings.Contains(d.Topic, "/organauto") {
				model := models.OrganAutoInfoModel{}
				json.Unmarshal([]byte(d.Data), &model)
				model.ID = bson.NewObjectId()
				model.DateTime = time.Now()
				model.EquipName = utils.GetEquipFromTopic(d.Topic)

				organAutoCollection.Insert(model)
			} else if strings.Contains(d.Topic, "/generator/state") {
				service._generatorRepository.InsertGeneratorInfo(utils.GetEquipFromTopic(d.Topic), d.Data)
			} else if strings.Contains(d.Topic, "/ARM/Hardware") {
				viewmodel := models.SystemInfoViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)
				service.insertSystemInfo(&viewmodel, d.Topic, sysInfoCollection, sysVolatileInfoCollection)
			} else if strings.Contains(d.Topic, "/ARM/Software/Complex") {
				viewmodel := models.SoftwareInfoViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)
				service.insertPermanentSoftwareInfo(&viewmodel, d.Topic, softwareInfoCollection)
			} else if strings.Contains(d.Topic, "/ARM/Software/msg") {
				viewmodel := models.SoftwareInfoViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)
				service.insertPermanentSoftwareInfo(&viewmodel, d.Topic, softwareVolatileInfoCollection)
			} else if strings.Contains(d.Topic, "/dicom") {
				model := models.DicomsInfoModel{}
				json.Unmarshal([]byte(d.Data), &model)
				model.ID = bson.NewObjectId()
				model.DateTime = time.Now()
				model.EquipName = utils.GetEquipFromTopic(d.Topic)

				dicomInfoCollection.Insert(model)
			} else if strings.Contains(d.Topic, "/stand/state") {
				service._standRepository.InsertStandInfo(utils.GetEquipFromTopic(d.Topic), d.Data)
			} else if strings.Contains(d.Topic, "/ARM/AllDBInfo") {
				service._dbInfoRepository.InsertDbInfoInfo(utils.GetEquipFromTopic(d.Topic), d.Data)
			}
		}
	}() //deviceCollection)

	<-quitCh
	fmt.Println("DalWorker quitted")
}

func (service *dalService) GetStudiesInWork(equipName string, startDate time.Time, endDate time.Time) []models.StudyInWorkModel {
	session := service.CreateSession()
	defer session.Close()

	studiesCollection := session.DB(service._dbName).C(models.StudyInWorkTableName)

	query := service.GetQuery(equipName, startDate, endDate)
	// объект для сохранения результата
	studies := []models.StudyInWorkModel{}
	err := studiesCollection.Find(query).Sort("-datetime").All(&studies)
	if err != nil {
		fmt.Println(err)
	}

	return studies
}

func (service *dalService) GetSystemInfo(equipName string, startDate time.Time, endDate time.Time) *models.FullSystemInfoModel {
	session := service.CreateSession()
	defer session.Close()

	sysVolatileInfoCollection := session.DB(service._dbName).C(models.SystemVolatileInfoTableName)

	query := service.GetQuery(equipName, startDate, endDate)
	// // объект для сохранения результата
	sysInfo := service.GetPermanentSystemInfo(equipName)
	sysInfos := []models.SystemInfoModel{*sysInfo}

	sysVolatileInfo := []models.SystemVolatileInfoModel{}
	sysVolatileInfoCollection.Find(query).Sort("-datetime").All(&sysVolatileInfo)

	return &models.FullSystemInfoModel{sysInfos, sysVolatileInfo}
}

func (service *dalService) GetOrganAutoInfo(equipName string, startDate time.Time, endDate time.Time) []models.OrganAutoInfoModel {
	session := service.CreateSession()
	defer session.Close()

	organAutoCollection := session.DB(service._dbName).C(models.OrganAutoTableName)

	query := service.GetQuery(equipName, startDate, endDate)
	// объект для сохранения результата
	organAutos := []models.OrganAutoInfoModel{}
	organAutoCollection.Find(query).Sort("-datetime").All(&organAutos)

	return organAutos
}

func (service *dalService) GetGeneratorInfo(equipName string, startDate time.Time, endDate time.Time) []models.RawDeviceInfoModel {
	return service._generatorRepository.GetGeneratorInfo(equipName, startDate, endDate)
}

func (service *dalService) GetSoftwareInfo(equipName string, startDate time.Time, endDate time.Time) *models.FullSoftwareInfoModel {
	session := service.CreateSession()
	defer session.Close()

	swVolatileInfoCollection := session.DB(service._dbName).C(models.SystemVolatileInfoTableName)

	query := service.GetQuery(equipName, startDate, endDate)
	// // объект для сохранения результата
	swInfo := service.GetPermanentSoftwareInfo(equipName)
	swInfos := []models.SoftwareInfoModel{*swInfo}

	swVolatileInfo := []models.SoftwareVolatileInfoModel{}
	swVolatileInfoCollection.Find(query).Sort("-datetime").All(&swVolatileInfo)

	return &models.FullSoftwareInfoModel{swInfos, swVolatileInfo}
}

func (service *dalService) GetDicomInfo(equipName string, startDate time.Time, endDate time.Time) []models.DicomsInfoModel {
	session := service.CreateSession()
	defer session.Close()

	dicomInfoCollection := session.DB(service._dbName).C(models.DicomInfoTableName)

	// // критерий выборки
	query := service.GetQuery(equipName, startDate, endDate)

	// // объект для сохранения результата
	dicomInfo := []models.DicomsInfoModel{}
	dicomInfoCollection.Find(query).Sort("-datetime").All(&dicomInfo)

	return dicomInfo
}

func (service *dalService) GetStandInfo(equipName string, startDate time.Time, endDate time.Time) []models.RawDeviceInfoModel {
	return service._standRepository.GetStandInfo(equipName, startDate, endDate)
}

func (service *dalService) GetPermanentSystemInfo(equipName string) *models.SystemInfoModel {
	session := service.CreateSession()
	defer session.Close()

	sysInfoCollection := session.DB(service._dbName).C(models.SystemInfoTableName)

	sysInfo := models.SystemInfoModel{}
	sysQuery := bson.M{"equipname": equipName}
	sysInfoCollection.Find(sysQuery).Sort("-datetime").One(&sysInfo)

	return &sysInfo
}

func (service *dalService) GetPermanentSoftwareInfo(equipName string) *models.SoftwareInfoModel {
	session := service.CreateSession()
	defer session.Close()

	softwareInfoCollection := session.DB(service._dbName).C(models.SoftwareInfoTableName)

	softwareInfo := models.SoftwareInfoModel{}
	sysQuery := bson.M{"equipname": equipName}
	softwareInfoCollection.Find(sysQuery).Sort("-datetime").One(&softwareInfo)

	return &softwareInfo
}

func (service *dalService) GetQuery(equipName string, startDate time.Time, endDate time.Time) bson.M {
	var query bson.M
	query = bson.M{
		"datetime": bson.M{
			"$gt": startDate,
			"$lt": endDate,
		},
	}
	if equipName != "" {
		query = bson.M{"$and": []bson.M{
			bson.M{"equipname": equipName},
			query}}
	}

	return query
}

func (service *dalService) CreateSession() *mgo.Session {
	session, err := mgo.Dial(service._connectionString)
	if err != nil {
		panic(err)
	}

	return session
}

func (service *dalService) insertSystemInfo(
	viewmodel *models.SystemInfoViewModel,
	topic string,
	sysInfoCollection *mgo.Collection,
	sysVolatileInfoCollection *mgo.Collection) {
	dateTime := time.Now()
	equipName := utils.GetEquipFromTopic(topic)

	model := models.SystemInfoModel{}
	volatileModel := models.SystemVolatileInfoModel{}

	/*createSystemInfo := func(paramName string, value string) Models.SystemInfoModel {
		model := Models.SystemInfoModel{}

		model.Id = bson.NewObjectId()
		model.DateTime = dateTime
		model.EquipName = equipName

		model.Parameter = paramName
		model.Value = value

		return model
	}

	bulk := sysInfoCollection.Bulk()
	//users := make([]interface{}, count)
	infos := []interface{}{
		createSystemInfo("Memory_Model_Memory_total_Gb", viewmodel.Memory.Memory_total_Gb),
		createSystemInfo("Processor_Model", viewmodel.Processor.Model),
		createSystemInfo("Motherboard_Model", viewmodel.Motherboard.Model),
	}
	bulk.Insert(infos...)
	_, bulkErr := bulk.Run()
	if bulkErr != nil {
		fmt.Println("bulkErr error! ")
	}*/
	hdds := []models.HddDriveInfoModel{}
	for _, value := range viewmodel.HDD {
		hdds = append(hdds, models.HddDriveInfoModel{
			Letter:    value.Letter,
			TotalSize: value.TotalSize,
		})
	}

	phyDisks := []models.PhysicalDiskInfoModel{}
	for _, value := range viewmodel.PhysicalDisks {
		phyDisks = append(phyDisks, models.PhysicalDiskInfoModel{
			FriendlyName: value.FriendlyName,
			MediaType:    value.MediaType,
			SizeGb:       value.SizeGb,
		})
	}

	nets := []models.NetworkInfoModel{}
	for _, value := range viewmodel.Network {
		nets = append(nets, models.NetworkInfoModel{
			NIC: value.NIC,
			IP:  value.IP,
		})
	}

	vgas := []models.VGAInfoModel{}
	for _, value := range viewmodel.VGA {
		vgas = append(vgas, models.VGAInfoModel{
			CardName:      value.CardName,
			DriverVersion: value.DriverVersion,
			MemoryGb:      value.MemoryGb,
		})
	}

	monitors := []models.MonitorInfoModel{}
	for _, value := range viewmodel.Monitor {
		monitors = append(monitors, models.MonitorInfoModel{
			DeviceName: value.DeviceName,
			Width:      value.Width,
			Height:     value.Height,
		})
	}

	model.ID = bson.NewObjectId()
	model.DateTime = dateTime
	model.EquipName = equipName
	model.HDD = hdds
	model.PhysicalDisks = phyDisks
	model.Processor = models.ProcessorInfoModel{Model: viewmodel.Processor.Model}
	model.Motherboard = models.MotherboardInfoModel{Model: viewmodel.Motherboard.Model}
	model.Memory = models.MemoryInfoModel{MemoryTotalGb: viewmodel.Memory.MemoryTotalGb}
	model.Network = nets
	model.VGA = vgas
	model.Monitor = monitors
	//model.Printer           []PrinterInfoModel
	sysInfoCollection.Insert(model)

	volatileModel.ID = bson.NewObjectId()
	volatileModel.DateTime = dateTime
	volatileModel.EquipName = equipName
	//HDD                   []HDDVolatileInfoModel
	volatileModel.ProcessorCPULoad = viewmodel.Processor.CPULoad
	volatileModel.MemoryMemoryFreeGb = viewmodel.Memory.MemoryFreeGb
	sysVolatileInfoCollection.Insert(volatileModel)
}

////
func (service *dalService) insertPermanentSoftwareInfo(
	viewmodel *models.SoftwareInfoViewModel,
	topic string,
	swInfoCollection *mgo.Collection) {
	dateTime := time.Now()
	equipName := utils.GetEquipFromTopic(topic)

	model := models.SoftwareInfoModel{}

	dbs := []models.DatabasesModel{}
	for _, value := range viewmodel.Databases {
		dbs = append(dbs, models.DatabasesModel{
			DBName:        value.DBName,
			DBStatus:      value.DBStatus,
			DBCompability: value.DBCompability,
		})
	}

	model.ID = bson.NewObjectId()
	model.DateTime = dateTime
	model.EquipName = equipName
	model.Databases = dbs

	model.Sysinfo = models.SysInfoModel{
		OS:          viewmodel.Sysinfo.OS,
		Version:     viewmodel.Sysinfo.Version,
		BuildNumber: viewmodel.Sysinfo.BuildNumber,
	}
	model.MSSQL = models.MSSQLInfoModel{
		SQL:     viewmodel.MSSQL.SQL,
		Version: viewmodel.MSSQL.Version,
		Status:  viewmodel.MSSQL.Status,
	}
	model.Atlas = models.AtlasInfoModel{
		AtlasVersion:  viewmodel.Atlas.AtlasVersion,
		ComplexType:   viewmodel.Atlas.ComplexType,
		Language:      viewmodel.Atlas.Language,
		Multimonitor:  viewmodel.Atlas.Multimonitor,
		XiLibsVersion: viewmodel.Atlas.XiLibsVersion,
	}

	swInfoCollection.Insert(model)
}

func (service *dalService) insertVolatileSoftwareInfo(
	viewmodel *models.SoftwareMessageViewModel,
	topic string,
	swVolatileInfoCollection *mgo.Collection) {
	dateTime := time.Now()
	equipName := utils.GetEquipFromTopic(topic)
	volatileModel := models.SoftwareVolatileInfoModel{}
	volatileModel.ID = bson.NewObjectId()
	volatileModel.DateTime = dateTime
	volatileModel.EquipName = equipName
	//volatileModel.ErrorCode = viewmodel.ErrorCode
	//volatileModel.ErrorDescription = viewmodel.ErrorDescription
	swVolatileInfoCollection.Insert(volatileModel)
}

////

func (service *dalService) ensureIndeces(sysInfoCollection *mgo.Collection, keys []string) {
	idx := mgo.Index{
		Key: keys,
	}
	sysInfoCollection.EnsureIndex(idx)
}

func (service *dalService) UpdateUser(userVM *models.UserViewModel) *models.UserModel {
	return service._userRepository.UpdateUser(userVM)
}

func (service *dalService) GetUsers() []models.UserModel {
	return service._userRepository.GetUsers()
}

func (service *dalService) GetUserByName(login string, email string, password string) *models.UserModel {
	return service._userRepository.GetUserByName(login, email, password)
}

func (service *dalService) CheckEquipment(equipName string) bool {
	return service._equipInfoRepository.CheckEquipment(equipName)
}

func (service *dalService) GetEquipInfos() []models.EquipInfoModel {
	return service._equipInfoRepository.GetEquipInfos()
}

func (service *dalService) InsertEquipInfo(equipName string, equipVM *models.EquipInfoViewModel) *models.EquipInfoModel {
	return service._equipInfoRepository.InsertEquipInfo(equipName, equipVM)
}

func (service *dalService) GetAllTableNamesInfo(equipName string) *models.AllDBTablesModel {
	return service._dbInfoRepository.GetAllTableNamesInfo(equipName)
}

func (service *dalService) GetTableContent(equipName string, tableType string, tableName string) string {
	return service._dbInfoRepository.GetTableContent(equipName, tableType, tableName)
}

func (service *dalService) GetDBSystemInfo(equipName string) map[string]json.RawMessage {
	return service._dbInfoRepository.GetDBSystemInfo(equipName)
}

func (service *dalService) GetDBSoftwareInfo(equipName string) *models.DBSoftwareInfoModel {
	return service._dbInfoRepository.GetDBSoftwareInfo(equipName)
}
