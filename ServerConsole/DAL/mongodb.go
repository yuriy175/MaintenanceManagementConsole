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
	GetStudiesInWork(startDate time.Time, endDate time.Time) []Models.StudyInWorkModel
	GetSystemInfo(startDate time.Time, endDate time.Time) []Models.SystemInfoModel
	GetOrganAutoInfo(startDate time.Time, endDate time.Time) []Models.OrganAutoInfoModel
	GetGeneratorInfo(startDate time.Time, endDate time.Time) []Models.GeneratorInfoModel
	GetSoftwareInfo(startDate time.Time, endDate time.Time) []Models.SoftwareInfoModel
	GetDicomInfo(startDate time.Time, endDate time.Time) []Models.DicomsInfoModel
	GetStandInfo(startDate time.Time, endDate time.Time) []Models.StandInfoModel
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
	softwareInfoCollection := db.C(Models.SoftwareInfoTableName)
	dicomInfoCollection := db.C(Models.DicomInfoTableName)
	standInfoCollection := db.C(Models.StandInfoTableName)

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
				model := Models.SystemInfoModel{}
				json.Unmarshal([]byte(d.Data), &model)
				model.Id = bson.NewObjectId()
				model.DateTime = time.Now()
				model.EquipName = Utils.GetEquipFromTopic(d.Topic)

				sysInfoCollection.Insert(model)
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

func (service *dalService) GetStudiesInWork(startDate time.Time, endDate time.Time) []Models.StudyInWorkModel {
	session := service.createSession()
	defer session.Close()

	studiesCollection := session.DB(Models.DBName).C(Models.StudyInWorkTableName)

	query := service.getQuery(startDate, endDate)
	// объект для сохранения результата
	studies := []Models.StudyInWorkModel{}
	err := studiesCollection.Find(query).Sort("-datetime").All(&studies)
	if err != nil {
		fmt.Println(err)
	}

	return studies
}

func (service *dalService) GetSystemInfo(startDate time.Time, endDate time.Time) []Models.SystemInfoModel {
	session := service.createSession()
	defer session.Close()

	// drivesCollection := session.DB(Models.DBName).C(Models.HddDrivesInfoTableName)
	sysInfoCollection := session.DB(Models.DBName).C(Models.SystemInfoTableName)

	query := service.getQuery(startDate, endDate)
	// // объект для сохранения результата
	sysInfo := []Models.SystemInfoModel{}
	sysInfoCollection.Find(query).Sort("-datetime").All(&sysInfo)

	return sysInfo
}

func (service *dalService) GetOrganAutoInfo(startDate time.Time, endDate time.Time) []Models.OrganAutoInfoModel {
	session := service.createSession()
	defer session.Close()

	organAutoCollection := session.DB(Models.DBName).C(Models.OrganAutoTableName)

	query := service.getQuery(startDate, endDate)
	// объект для сохранения результата
	organAutos := []Models.OrganAutoInfoModel{}
	organAutoCollection.Find(query).Sort("-datetime").All(&organAutos)

	return organAutos
}

func (service *dalService) GetGeneratorInfo(startDate time.Time, endDate time.Time) []Models.GeneratorInfoModel {
	session := service.createSession()
	defer session.Close()

	// drivesCollection := session.DB(Models.DBName).C(Models.HddDrivesInfoTableName)
	genInfoCollection := session.DB(Models.DBName).C(Models.GeneratorInfoTableName)

	// // критерий выборки
	query := service.getQuery(startDate, endDate)
	// // объект для сохранения результата
	genInfo := []Models.GeneratorInfoModel{}
	genInfoCollection.Find(query).Sort("-datetime").All(&genInfo)

	return genInfo
}

func (service *dalService) GetSoftwareInfo(startDate time.Time, endDate time.Time) []Models.SoftwareInfoModel {
	session := service.createSession()
	defer session.Close()

	swInfoCollection := session.DB(Models.DBName).C(Models.SoftwareInfoTableName)

	// // критерий выборки
	query := service.getQuery(startDate, endDate)

	// // объект для сохранения результата
	swInfo := []Models.SoftwareInfoModel{}
	swInfoCollection.Find(query).Sort("-datetime").All(&swInfo)

	return swInfo
}

func (service *dalService) GetDicomInfo(startDate time.Time, endDate time.Time) []Models.DicomsInfoModel {
	session := service.createSession()
	defer session.Close()

	dicomInfoCollection := session.DB(Models.DBName).C(Models.DicomInfoTableName)

	// // критерий выборки
	query := service.getQuery(startDate, endDate)

	// // объект для сохранения результата
	dicomInfo := []Models.DicomsInfoModel{}
	dicomInfoCollection.Find(query).Sort("-datetime").All(&dicomInfo)

	return dicomInfo
}

func (service *dalService) GetStandInfo(startDate time.Time, endDate time.Time) []Models.StandInfoModel {
	session := service.createSession()
	defer session.Close()

	standInfoCollection := session.DB(Models.DBName).C(Models.StandInfoTableName)

	// // критерий выборки
	query := service.getQuery(startDate, endDate)

	// // объект для сохранения результата
	standInfo := []Models.StandInfoModel{}
	standInfoCollection.Find(query).Sort("-datetime").All(&standInfo)

	return standInfo
}

func (service *dalService) getQuery(startDate time.Time, endDate time.Time) bson.M {
	query := bson.M{
		"datetime": bson.M{
			"$gt": startDate,
			"$lt": endDate,
		},
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
