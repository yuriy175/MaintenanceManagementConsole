package DAL

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../Interfaces"
	"../Models"
	"../Utils"
)

type IDalService interface {
	Start()
	////
	GetStudiesInWork(equipName string, startDate time.Time, endDate time.Time) []Models.StudyInWorkModel
	GetSystemInfo(equipName string, startDate time.Time, endDate time.Time) *Models.FullSystemInfoModel
	GetOrganAutoInfo(equipName string, startDate time.Time, endDate time.Time) []Models.OrganAutoInfoModel
	GetGeneratorInfo(equipName string, startDate time.Time, endDate time.Time) []Models.GeneratorInfoModel
	GetSoftwareInfo(equipName string, startDate time.Time, endDate time.Time) *Models.FullSoftwareInfoModel
	GetDicomInfo(equipName string, startDate time.Time, endDate time.Time) []Models.DicomsInfoModel
	GetStandInfo(equipName string, startDate time.Time, endDate time.Time) []Models.StandInfoModel

	GetPermanentSystemInfo(equipName string) *Models.SystemInfoModel
	GetPermanentSoftwareInfo(equipName string) *Models.SoftwareInfoModel

	//user repository
	UpdateUser(user *Models.UserViewModel) *Models.UserModel
	GetUsers() []Models.UserModel
	GetUserByName(surname string, email string, password string) bool
}

type dalService struct {
	_dalCh       chan *Models.RawMqttMessage
	_authService Interfaces.IAuthService
}

func DalServiceNew(
	authService Interfaces.IAuthService,
	dalCh chan *Models.RawMqttMessage) IDalService {
	service := &dalService{}

	service._authService = authService
	service._dalCh = dalCh

	return service
}

func (service *dalService) Start() {
	quitCh := make(chan int)

	session := service.createSession()
	defer session.Close()

	db := session.DB(Models.DBName)
	// deviceCollection := db.C(Models.DeviceConnectionsTableName)
	studiesCollection := db.C(Models.StudyInWorkTableName)
	organAutoCollection := db.C(Models.OrganAutoTableName)
	genInfoCollection := db.C(Models.GeneratorInfoTableName)
	sysInfoCollection := db.C(Models.SystemInfoTableName)
	sysVolatileInfoCollection := db.C(Models.SystemVolatileInfoTableName)
	softwareInfoCollection := db.C(Models.SoftwareInfoTableName)
	softwareVolatileInfoCollection := db.C(Models.SoftwareVolatileInfoTableName)
	dicomInfoCollection := db.C(Models.DicomInfoTableName)
	standInfoCollection := db.C(Models.StandInfoTableName)

	service.ensureIndeces(sysInfoCollection, []string{"equipname", "datetime"})
	service.ensureIndeces(sysVolatileInfoCollection, []string{"equipname", "datetime"})

	service.ensureIndeces(softwareInfoCollection, []string{"equipname", "datetime"})
	service.ensureIndeces(softwareVolatileInfoCollection, []string{"equipname", "datetime"})

	go func() {
		for d := range service._dalCh {
			if strings.Contains(d.Topic, "/study") {
				model := Models.StudyInWorkModel{}
				json.Unmarshal([]byte(d.Data), &model)
				model.Id = bson.NewObjectId()
				model.DateTime = time.Now()
				model.EquipName = Utils.GetEquipFromTopic(d.Topic)
				model.State = 2

				studiesCollection.Insert(model)
			} else if strings.Contains(d.Topic, "/organauto") {
				model := Models.OrganAutoInfoModel{}
				json.Unmarshal([]byte(d.Data), &model)
				model.Id = bson.NewObjectId()
				model.DateTime = time.Now()
				model.EquipName = Utils.GetEquipFromTopic(d.Topic)

				organAutoCollection.Insert(model)
			} else if strings.Contains(d.Topic, "/generator/state") {
				viewmodel := Models.GeneratorInfoViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)
				viewmodel.State.Id = bson.NewObjectId()
				viewmodel.State.DateTime = time.Now()
				viewmodel.State.EquipName = Utils.GetEquipFromTopic(d.Topic)

				genInfoCollection.Insert(viewmodel.State)
			} else if strings.Contains(d.Topic, "/ARM/Hardware") {
				viewmodel := Models.SystemInfoViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)
				service.insertSystemInfo(&viewmodel, d.Topic, sysInfoCollection, sysVolatileInfoCollection)
			} else if strings.Contains(d.Topic, "/ARM/Software/Complex") {
				viewmodel := Models.SoftwareInfoViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)
				service.insertPermanentSoftwareInfo(&viewmodel, d.Topic, softwareInfoCollection)
			} else if strings.Contains(d.Topic, "/ARM/Software/msg") {
				viewmodel := Models.SoftwareInfoViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)
				service.insertPermanentSoftwareInfo(&viewmodel, d.Topic, softwareVolatileInfoCollection)
			} else if strings.Contains(d.Topic, "/dicom") {
				model := Models.DicomsInfoModel{}
				json.Unmarshal([]byte(d.Data), &model)
				model.Id = bson.NewObjectId()
				model.DateTime = time.Now()
				model.EquipName = Utils.GetEquipFromTopic(d.Topic)

				dicomInfoCollection.Insert(model)
			} else if strings.Contains(d.Topic, "/stand/state") {
				viewmodel := Models.StandInfoViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)
				viewmodel.State.Id = bson.NewObjectId()
				viewmodel.State.DateTime = time.Now()
				viewmodel.State.EquipName = Utils.GetEquipFromTopic(d.Topic)

				standInfoCollection.Insert(viewmodel.State)
			}
		}
	}() //deviceCollection)

	<-quitCh
	fmt.Println("DalWorker quitted")
}

func (service *dalService) GetStudiesInWork(equipName string, startDate time.Time, endDate time.Time) []Models.StudyInWorkModel {
	session := service.createSession()
	defer session.Close()

	studiesCollection := session.DB(Models.DBName).C(Models.StudyInWorkTableName)

	query := service.getQuery(equipName, startDate, endDate)
	// объект для сохранения результата
	studies := []Models.StudyInWorkModel{}
	err := studiesCollection.Find(query).Sort("-datetime").All(&studies)
	if err != nil {
		fmt.Println(err)
	}

	return studies
}

func (service *dalService) GetSystemInfo(equipName string, startDate time.Time, endDate time.Time) *Models.FullSystemInfoModel {
	session := service.createSession()
	defer session.Close()

	sysVolatileInfoCollection := session.DB(Models.DBName).C(Models.SystemVolatileInfoTableName)

	query := service.getQuery(equipName, startDate, endDate)
	// // объект для сохранения результата
	sysInfo := service.GetPermanentSystemInfo(equipName)
	sysInfos := []Models.SystemInfoModel{*sysInfo}

	sysVolatileInfo := []Models.SystemVolatileInfoModel{}
	sysVolatileInfoCollection.Find(query).Sort("-datetime").All(&sysVolatileInfo)

	return &Models.FullSystemInfoModel{sysInfos, sysVolatileInfo}
}

func (service *dalService) GetOrganAutoInfo(equipName string, startDate time.Time, endDate time.Time) []Models.OrganAutoInfoModel {
	session := service.createSession()
	defer session.Close()

	organAutoCollection := session.DB(Models.DBName).C(Models.OrganAutoTableName)

	query := service.getQuery(equipName, startDate, endDate)
	// объект для сохранения результата
	organAutos := []Models.OrganAutoInfoModel{}
	organAutoCollection.Find(query).Sort("-datetime").All(&organAutos)

	return organAutos
}

func (service *dalService) GetGeneratorInfo(equipName string, startDate time.Time, endDate time.Time) []Models.GeneratorInfoModel {
	session := service.createSession()
	defer session.Close()

	// drivesCollection := session.DB(Models.DBName).C(Models.HddDrivesInfoTableName)
	genInfoCollection := session.DB(Models.DBName).C(Models.GeneratorInfoTableName)

	// // критерий выборки
	query := service.getQuery(equipName, startDate, endDate)
	// // объект для сохранения результата
	genInfo := []Models.GeneratorInfoModel{}
	genInfoCollection.Find(query).Sort("-datetime").All(&genInfo)

	return genInfo
}

func (service *dalService) GetSoftwareInfo(equipName string, startDate time.Time, endDate time.Time) *Models.FullSoftwareInfoModel {
	session := service.createSession()
	defer session.Close()

	swVolatileInfoCollection := session.DB(Models.DBName).C(Models.SystemVolatileInfoTableName)

	query := service.getQuery(equipName, startDate, endDate)
	// // объект для сохранения результата
	swInfo := service.GetPermanentSoftwareInfo(equipName)
	swInfos := []Models.SoftwareInfoModel{*swInfo}

	swVolatileInfo := []Models.SoftwareVolatileInfoModel{}
	swVolatileInfoCollection.Find(query).Sort("-datetime").All(&swVolatileInfo)

	return &Models.FullSoftwareInfoModel{swInfos, swVolatileInfo}
}

func (service *dalService) GetDicomInfo(equipName string, startDate time.Time, endDate time.Time) []Models.DicomsInfoModel {
	session := service.createSession()
	defer session.Close()

	dicomInfoCollection := session.DB(Models.DBName).C(Models.DicomInfoTableName)

	// // критерий выборки
	query := service.getQuery(equipName, startDate, endDate)

	// // объект для сохранения результата
	dicomInfo := []Models.DicomsInfoModel{}
	dicomInfoCollection.Find(query).Sort("-datetime").All(&dicomInfo)

	return dicomInfo
}

func (service *dalService) GetStandInfo(equipName string, startDate time.Time, endDate time.Time) []Models.StandInfoModel {
	session := service.createSession()
	defer session.Close()

	standInfoCollection := session.DB(Models.DBName).C(Models.StandInfoTableName)

	// // критерий выборки
	query := service.getQuery(equipName, startDate, endDate)

	// // объект для сохранения результата
	standInfo := []Models.StandInfoModel{}
	standInfoCollection.Find(query).Sort("-datetime").All(&standInfo)

	return standInfo
}

func (service *dalService) GetPermanentSystemInfo(equipName string) *Models.SystemInfoModel {
	session := service.createSession()
	defer session.Close()

	sysInfoCollection := session.DB(Models.DBName).C(Models.SystemInfoTableName)

	sysInfo := Models.SystemInfoModel{}
	sysQuery := bson.M{"equipname": equipName}
	sysInfoCollection.Find(sysQuery).Sort("-datetime").One(&sysInfo)

	return &sysInfo
}

func (service *dalService) GetPermanentSoftwareInfo(equipName string) *Models.SoftwareInfoModel {
	session := service.createSession()
	defer session.Close()

	softwareInfoCollection := session.DB(Models.DBName).C(Models.SoftwareInfoTableName)

	softwareInfo := Models.SoftwareInfoModel{}
	sysQuery := bson.M{"equipname": equipName}
	softwareInfoCollection.Find(sysQuery).Sort("-datetime").One(&softwareInfo)

	return &softwareInfo
}

func (service *dalService) getQuery(equipName string, startDate time.Time, endDate time.Time) bson.M {
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

func (service *dalService) createSession() *mgo.Session {
	session, err := mgo.Dial(Models.MongoDBConnectionString)
	if err != nil {
		panic(err)
	}

	return session
}

func (service *dalService) insertSystemInfo(
	viewmodel *Models.SystemInfoViewModel,
	topic string,
	sysInfoCollection *mgo.Collection,
	sysVolatileInfoCollection *mgo.Collection) {
	dateTime := time.Now()
	equipName := Utils.GetEquipFromTopic(topic)

	model := Models.SystemInfoModel{}
	volatileModel := Models.SystemVolatileInfoModel{}

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
	hdds := []Models.HddDriveInfoModel{}
	for _, value := range viewmodel.HDD {
		hdds = append(hdds, Models.HddDriveInfoModel{
			Letter:    value.Letter,
			TotalSize: value.TotalSize,
		})
	}

	phyDisks := []Models.PhysicalDiskInfoModel{}
	for _, value := range viewmodel.PhysicalDisks {
		phyDisks = append(phyDisks, Models.PhysicalDiskInfoModel{
			FriendlyName: value.FriendlyName,
			MediaType:    value.MediaType,
			Size_Gb:      value.Size_Gb,
		})
	}

	nets := []Models.NetworkInfoModel{}
	for _, value := range viewmodel.Network {
		nets = append(nets, Models.NetworkInfoModel{
			NIC: value.NIC,
			IP:  value.IP,
		})
	}

	vgas := []Models.VGAInfoModel{}
	for _, value := range viewmodel.VGA {
		vgas = append(vgas, Models.VGAInfoModel{
			Card_Name:      value.Card_Name,
			Driver_Version: value.Driver_Version,
			Memory_Gb:      value.Memory_Gb,
		})
	}

	monitors := []Models.MonitorInfoModel{}
	for _, value := range viewmodel.Monitor {
		monitors = append(monitors, Models.MonitorInfoModel{
			Device_Name: value.Device_Name,
			Width:       value.Width,
			Height:      value.Height,
		})
	}

	model.Id = bson.NewObjectId()
	model.DateTime = dateTime
	model.EquipName = equipName
	model.HDD = hdds
	model.PhysicalDisks = phyDisks
	model.Processor = Models.ProcessorInfoModel{Model: viewmodel.Processor.Model}
	model.Motherboard = Models.MotherboardInfoModel{Model: viewmodel.Motherboard.Model}
	model.Memory = Models.MemoryInfoModel{Memory_total_Gb: viewmodel.Memory.Memory_total_Gb}
	model.Network = nets
	model.VGA = vgas
	model.Monitor = monitors
	//model.Printer           []PrinterInfoModel
	sysInfoCollection.Insert(model)

	volatileModel.Id = bson.NewObjectId()
	volatileModel.DateTime = dateTime
	volatileModel.EquipName = equipName
	//HDD                   []HDDVolatileInfoModel
	volatileModel.Processor_CPU_Load = viewmodel.Processor.CPU_Load
	volatileModel.Memory_Memory_free_Gb = viewmodel.Memory.Memory_free_Gb
	sysVolatileInfoCollection.Insert(volatileModel)
}

////
func (service *dalService) insertPermanentSoftwareInfo(
	viewmodel *Models.SoftwareInfoViewModel,
	topic string,
	swInfoCollection *mgo.Collection) {
	dateTime := time.Now()
	equipName := Utils.GetEquipFromTopic(topic)

	model := Models.SoftwareInfoModel{}

	dbs := []Models.DatabasesModel{}
	for _, value := range viewmodel.Databases {
		dbs = append(dbs, Models.DatabasesModel{
			DB_name:        value.DB_name,
			DB_Status:      value.DB_Status,
			DB_compability: value.DB_compability,
		})
	}

	model.Id = bson.NewObjectId()
	model.DateTime = dateTime
	model.EquipName = equipName
	model.Databases = dbs

	model.Sysinfo = Models.SysInfoModel{
		OS:           viewmodel.Sysinfo.OS,
		Version:      viewmodel.Sysinfo.Version,
		Build_Number: viewmodel.Sysinfo.Build_Number,
	}
	model.MSSQL = Models.MSSQLInfoModel{
		SQL:     viewmodel.MSSQL.SQL,
		Version: viewmodel.MSSQL.Version,
		Status:  viewmodel.MSSQL.Status,
	}
	model.Atlas = Models.AtlasInfoModel{
		Atlas_Version:  viewmodel.Atlas.Atlas_Version,
		Complex_type:   viewmodel.Atlas.Complex_type,
		Language:       viewmodel.Atlas.Language,
		Multimonitor:   viewmodel.Atlas.Multimonitor,
		XiLibs_Version: viewmodel.Atlas.XiLibs_Version,
	}

	swInfoCollection.Insert(model)
}

func (service *dalService) insertVolatileSoftwareInfo(
	viewmodel *Models.SoftwareMessageViewModel,
	topic string,
	swVolatileInfoCollection *mgo.Collection) {
	dateTime := time.Now()
	equipName := Utils.GetEquipFromTopic(topic)
	volatileModel := Models.SoftwareVolatileInfoModel{}
	volatileModel.Id = bson.NewObjectId()
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

func (service *dalService) UpdateUser(userVM *Models.UserViewModel) *Models.UserModel {
	session := service.createSession()
	defer session.Close()

	userCollection := session.DB(Models.DBName).C(Models.UsersTableName)

	sum := service._authService.GetSum(userVM.Password)

	model := Models.UserModel{}
	model.Id = bson.NewObjectId()
	model.DateTime = time.Now()

	model.Login = userVM.Login
	model.PasswordHash = sum
	model.Surname = userVM.Surname
	model.Role = userVM.Role
	model.Email = userVM.Email

	userCollection.Insert(model)

	return &model
}

func (service *dalService) GetUsers() []Models.UserModel {
	session := service.createSession()
	defer session.Close()

	userCollection := session.DB(Models.DBName).C(Models.UsersTableName)

	// // критерий выборки
	query := bson.M{}

	// // объект для сохранения результата
	users := []Models.UserModel{}
	userCollection.Find(query).All(&users)

	return users
}

func (service *dalService) GetUserByName(login string, email string, password string) bool {
	session := service.createSession()
	defer session.Close()

	userCollection := session.DB(Models.DBName).C(Models.UsersTableName)

	// // критерий выборки
	query := bson.M{"login": login}

	// // объект для сохранения результата
	user := Models.UserModel{}
	userCollection.Find(query).One(&user)

	ok := service._authService.CheckSum(password, user.PasswordHash)

	return ok
}
