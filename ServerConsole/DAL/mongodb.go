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

func DalWorker(equipDalCh chan *Models.RawMqttMessage) {
	quitCh := make(chan int)

	session := dalCreateSession()
	defer session.Close()

	db := session.DB(Models.DBName)
	// deviceCollection := db.C(Models.DeviceConnectionsTableName)
	studiesCollection := db.C(Models.StudyInWorkTableName)
	organAutoCollection := db.C(Models.OrganAutoTableName)
	genInfoCollection := db.C(Models.GeneratorInfoTableName)
	sysInfoCollection := db.C(Models.SystemInfoTableName)
	softwareInfoCollection := db.C(Models.SoftwareInfoTableName)

	go func() { 
		for d := range equipDalCh {
			/*if d.MsgType == Models.MsgTypeHwConnectionStateArrived {
				deviceId := d.Info["Id"].(float64)
				deviceName := d.Info["Name"].(string)
				deviceType := d.Info["Type"].(string)
				deviceConnection := d.Info["Connection"].(float64)

				model := &Models.DeviceConnectionModel{
					Id:               bson.NewObjectId(),
					DateTime:         time.Now(),
					EquipNumber:      d.EquipNumber,
					EquipName:        d.EquipName,
					EquipIP:          d.EquipIP,
					DeviceId:         deviceId,
					DeviceName:       deviceName,
					DeviceType:       deviceType,
					DeviceConnection: deviceConnection,
				}
				deviceCollection.Insert(model)
			} else */if strings.Contains(d.Topic, "/study") {
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
			}
		}
	}() //deviceCollection)

	<-quitCh
	fmt.Println("DalWorker quitted")
}

func GetDeviceConnections() []Models.DeviceConnectionModel {
	session := dalCreateSession()
	defer session.Close()

	deviceCollection := session.DB(Models.DBName).C(Models.DeviceConnectionsTableName)

	// критерий выборки
	query := bson.M{}
	// объект для сохранения результата
	devices := []Models.DeviceConnectionModel{}
	deviceCollection.Find(query).All(&devices)

	return devices
}

func GetStudiesInWork(startDate time.Time, endDate time.Time) []Models.StudyInWorkModel {
	session := dalCreateSession()
	defer session.Close()

	studiesCollection := session.DB(Models.DBName).C(Models.StudyInWorkTableName)

	// критерий выборки
	query := bson.M{
		"datetime": bson.M{
			"$gt": startDate,
			"$lt": endDate,
		},
	}
	// объект для сохранения результата
	studies := []Models.StudyInWorkModel{}
	err := studiesCollection.Find(query).All(&studies)
	if err != nil {
		fmt.Println(err)
	}

	/*studies := []Models.StudyInWorkModel{
		Models.StudyInWorkModel{
			EquipName:     "krt",
			State:         2,
			StudyId:       1,
			StudyDicomUid: "123",
			StudyName:     "нога",
		},
		Models.StudyInWorkModel{
			EquipName:     "krt",
			State:         2,
			StudyId:       2,
			StudyDicomUid: "124",
			StudyName:     "голова",
		},
	}*/

	return studies
}

func GetSystemInfo(startDate time.Time, endDate time.Time) []Models.SystemInfoModel {
	session := dalCreateSession()
	defer session.Close()

	// drivesCollection := session.DB(Models.DBName).C(Models.HddDrivesInfoTableName)
	sysInfoCollection := session.DB(Models.DBName).C(Models.SystemInfoTableName)

	// // критерий выборки
	query := bson.M{
		"datetime": bson.M{
			"$gt": startDate,
			"$lt": endDate,
		},
	}
	// // объект для сохранения результата
	sysInfo := []Models.SystemInfoModel{}
	sysInfoCollection.Find(query).All(&sysInfo)

	// sysInfo := []Models.SystemInfoModel{
	// 	Models.SystemInfoModel{
	// 		EquipName:     "krt",
	// 		State:         1,
	// 		CPULoad:       32,
	// 		TotalMemory:   16,
	// 		FreeMemory:    8,
	// 		HddName:       "C:/",
	// 		HddTotalSpace: 1000,
	// 		HddFreeSpace:  333,
	// 	},
	// 	Models.SystemInfoModel{
	// 		EquipName:     "krt",
	// 		State:         1,
	// 		HddName:       "D:/",
	// 		HddTotalSpace: 500,
	// 		HddFreeSpace:  443,
	// 	},
	// 	Models.SystemInfoModel{
	// 		EquipName: "krt",
	// 		CPULoad:   45,
	// 	},
	// }

	return sysInfo
}

func GetOrganAutoInfo(startDate time.Time, endDate time.Time) []Models.OrganAutoInfoModel {
	session := dalCreateSession()
	defer session.Close()

	/*organAutos := []Models.OrganAutoInfoModel{
		Models.OrganAutoInfoModel{
			EquipName:    "krt",
			State:        0,
			OrganAuto:    "нога",
			Projection:   "кривая",
			Direction:    "задняя",
			AgeGroupId:   4,
			Constitution: 2,
		},
		Models.OrganAutoInfoModel{
			EquipName:    "krt",
			State:        1,
			OrganAuto:    "рука",
			Projection:   "левая",
			Direction:    "задняя",
			AgeGroupId:   4,
			Constitution: 2,
		},
	}*/
	// drivesCollection := session.DB(Models.DBName).C(Models.HddDrivesInfoTableName)
	organAutoCollection := session.DB(Models.DBName).C(Models.OrganAutoTableName)

	// критерий выборки
	query := bson.M{
		"datetime": bson.M{
			"$gt": startDate,
			"$lt": endDate,
		},
	}
	// объект для сохранения результата
	organAutos := []Models.OrganAutoInfoModel{}
	organAutoCollection.Find(query).All(&organAutos)

	return organAutos
}

func GetGeneratorInfo(startDate time.Time, endDate time.Time) []Models.GeneratorInfoModel {
	session := dalCreateSession()
	defer session.Close()

	// drivesCollection := session.DB(Models.DBName).C(Models.HddDrivesInfoTableName)
	genInfoCollection := session.DB(Models.DBName).C(Models.GeneratorInfoTableName)

	// // критерий выборки
	query := bson.M{
		"datetime": bson.M{
			"$gt": startDate,
			"$lt": endDate,
		},
	}
	// // объект для сохранения результата
	genInfo := []Models.GeneratorInfoModel{}
	genInfoCollection.Find(query).All(&genInfo)

	// genInfo := []Models.GeneratorInfoModel{
	// 	Models.GeneratorInfoModel{
	// 		EquipName:   "krt",
	// 		State:       0,
	// 		Errors:      "все умерло",
	// 		Workstation: 1,
	// 		Heat:        1,
	// 		Current:     10,
	// 		Voltage:     66,
	// 	},
	// 	Models.GeneratorInfoModel{
	// 		EquipName:   "krt",
	// 		State:       1,
	// 		Workstation: 2,
	// 		Heat:        2,
	// 		Current:     5,
	// 		Voltage:     106,
	// 	},
	// 	Models.GeneratorInfoModel{
	// 		EquipName: "krt",
	// 		Voltage:   107,
	// 	},
	// }

	return genInfo
}

func GetSoftwareInfo(startDate time.Time, endDate time.Time) []Models.SoftwareInfoModel {
	session := dalCreateSession()
	defer session.Close()

	swInfoCollection := session.DB(Models.DBName).C(Models.SoftwareInfoTableName)

	// // критерий выборки
	query := GetQuery(startDate, endDate)

	// // объект для сохранения результата
	swInfo := []Models.SoftwareInfoModel{}
	swInfoCollection.Find(query).All(&swInfo)

	return swInfo
}

func GetQuery(startDate time.Time, endDate time.Time) bson.M{
	query := bson.M{
		"datetime": bson.M{
			"$gt": startDate,
			"$lt": endDate,
		},
	}

	return query
}

func dalCreateSession() *mgo.Session {
	session, err := mgo.Dial(Models.MongoDBConnectionString)
	if err != nil {
		panic(err)
	}

	return session
}
