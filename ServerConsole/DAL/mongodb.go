package DAL

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

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
	GetSoftwareInfo(equipName string, startDate time.Time, endDate time.Time) []Models.SoftwareInfoModel
	GetDicomInfo(equipName string, startDate time.Time, endDate time.Time) []Models.DicomsInfoModel
	GetStandInfo(equipName string, startDate time.Time, endDate time.Time) []Models.StandInfoModel
}

type dalService struct {
	_dalCh chan *Models.RawMqttMessage
}

func DalServiceNew(
	dalCh chan *Models.RawMqttMessage) IDalService {
	service := &dalService{}

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
	dicomInfoCollection := db.C(Models.DicomInfoTableName)
	standInfoCollection := db.C(Models.StandInfoTableName)

	service.ensureIndeces(sysInfoCollection, []string{"equipname", "datetime"})
	service.ensureIndeces(sysVolatileInfoCollection, []string{"equipname", "datetime"})

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
			} else if strings.Contains(d.Topic, "/ARM/Software") {
				model := Models.SoftwareInfoModel{}
				json.Unmarshal([]byte(d.Data), &model)
				model.Id = bson.NewObjectId()
				model.DateTime = time.Now()
				model.EquipName = Utils.GetEquipFromTopic(d.Topic)

				softwareInfoCollection.Insert(model)
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

	// drivesCollection := session.DB(Models.DBName).C(Models.HddDrivesInfoTableName)
	sysInfoCollection := session.DB(Models.DBName).C(Models.SystemInfoTableName)
	sysVolatileInfoCollection := session.DB(Models.DBName).C(Models.SystemVolatileInfoTableName)

	query := service.getQuery(equipName, startDate, endDate)
	// // объект для сохранения результата
	sysInfo := []Models.SystemInfoModel{}
	sysInfoCollection.Find(query).Sort("-datetime").All(&sysInfo)

	sysVolatileInfo := []Models.SystemVolatileInfoModel{}
	sysVolatileInfoCollection.Find(query).Sort("-datetime").All(&sysVolatileInfo)

	return &Models.FullSystemInfoModel{sysInfo, sysVolatileInfo}
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

func (service *dalService) GetSoftwareInfo(equipName string, startDate time.Time, endDate time.Time) []Models.SoftwareInfoModel {
	session := service.createSession()
	defer session.Close()

	swInfoCollection := session.DB(Models.DBName).C(Models.SoftwareInfoTableName)

	// // критерий выборки
	query := service.getQuery(equipName, startDate, endDate)

	// // объект для сохранения результата
	swInfo := []Models.SoftwareInfoModel{}
	swInfoCollection.Find(query).Sort("-datetime").All(&swInfo)

	return swInfo
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

	createSystemInfo := func(paramName string, value string) Models.SystemInfoModel {
		model := Models.SystemInfoModel{}

		model.Id = bson.NewObjectId()
		model.DateTime = dateTime
		model.EquipName = equipName

		model.Parameter = paramName
		model.Value = value

		return model
	}

	volatileModel := Models.SystemVolatileInfoModel{}

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
	}

	volatileModel.Id = bson.NewObjectId()
	volatileModel.DateTime = dateTime
	volatileModel.EquipName = equipName
	//HDD                   []HDDVolatileInfoModel
	volatileModel.Processor_CPU_Load = viewmodel.Processor.CPU_Load
	volatileModel.Memory_Memory_free_Gb = viewmodel.Memory.Memory_free_Gb
	sysVolatileInfoCollection.Insert(volatileModel)
}

func (service *dalService) ensureIndeces(sysInfoCollection *mgo.Collection, keys []string) {
	idx := mgo.Index{
		Key: keys,
	}
	sysInfoCollection.EnsureIndex(idx)
}
