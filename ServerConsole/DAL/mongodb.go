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

	go func() { //c *mgo.Collection) {
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
				/*studyId := d.Info["StudyId"].(float64)
				studyDicomUid := d.Info["StudyDicomUid"].(string)
				studyName := d.Info["StudyName"].(string)
				*/
				// model := &Models.StudyInWorkModel{
				// 	Id:            bson.NewObjectId(),
				// 	DateTime:      time.Now(),
				// 	EquipName:     d.EquipName,
				// 	StudyId:       studyId,
				// 	StudyDicomUid: studyDicomUid,
				// 	StudyName:     studyName,
				// 	State:         2,
				// }
				model := Models.StudyInWorkModel{}
				json.Unmarshal([]byte(d.Data), &model)
				model.Id = bson.NewObjectId()
				model.DateTime = time.Now()
				model.EquipName = Utils.GetEquipFromTopic(d.Topic)
				model.State = 2
				studiesCollection.Insert(model)
			} /*else if d.MsgType == Models.MsgTypeHddDrivesInfo {
				var params []Models.HddDrivesInfoMessage
				json.Unmarshal(d.Info, &params)

				i := 0
				i++

				hddName := d.Info["HddName"].(string)
				hddTotalSpace := d.Info["TotalSize"].(float64)
				hddFreeSpace := d.Info["FreeSize"].(float64)

				model := &Models.HddDrivesInfoModel{
					Id:            bson.NewObjectId(),
					DateTime:      time.Now(),
					EquipNumber:   d.EquipNumber,
					EquipName:     d.EquipName,
					EquipIP:       d.EquipIP,
					HddName:       hddName,
					HddTotalSpace: hddTotalSpace,
					HddFreeSpace:  hddFreeSpace,
				}
				studiesCollection.Insert(model)
			}*/
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

func GetStudiesInWork() []Models.StudyInWorkModel {
	session := dalCreateSession()
	defer session.Close()

	studiesCollection := session.DB(Models.DBName).C(Models.StudyInWorkTableName)

	// критерий выборки
	query := bson.M{}
	// объект для сохранения результата
	studies := []Models.StudyInWorkModel{}
	studiesCollection.Find(query).All(&studies)

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

func GetSystemInfo() []Models.SystemInfoModel {
	session := dalCreateSession()
	defer session.Close()

	// drivesCollection := session.DB(Models.DBName).C(Models.HddDrivesInfoTableName)

	// // критерий выборки
	// query := bson.M{}
	// // объект для сохранения результата
	// drives := []Models.HddDrivesInfoModel{}
	// drivesCollection.Find(query).All(&drives)

	sysInfo := []Models.SystemInfoModel{
		Models.SystemInfoModel{
			EquipName:     "krt",
			State:         1,
			CPULoad:       32,
			TotalMemory:   16,
			FreeMemory:    8,
			HddName:       "C:/",
			HddTotalSpace: 1000,
			HddFreeSpace:  333,
		},
		Models.SystemInfoModel{
			EquipName:     "krt",
			State:         1,
			HddName:       "D:/",
			HddTotalSpace: 500,
			HddFreeSpace:  443,
		},
		Models.SystemInfoModel{
			EquipName: "krt",
			CPULoad:   45,
		},
	}

	return sysInfo
}

func GetOrganAutoInfo() []Models.OrganAutoInfoModel {
	session := dalCreateSession()
	defer session.Close()

	organAutos := []Models.OrganAutoInfoModel{
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
	}
	// drivesCollection := session.DB(Models.DBName).C(Models.HddDrivesInfoTableName)

	// // критерий выборки
	// query := bson.M{}
	// // объект для сохранения результата
	// drives := []Models.HddDrivesInfoModel{}
	// drivesCollection.Find(query).All(&drives)

	return organAutos
}

func GetGeneratorInfo() []Models.GeneratorInfoModel {
	session := dalCreateSession()
	defer session.Close()

	// drivesCollection := session.DB(Models.DBName).C(Models.HddDrivesInfoTableName)

	// // критерий выборки
	// query := bson.M{}
	// // объект для сохранения результата
	// drives := []Models.HddDrivesInfoModel{}
	// drivesCollection.Find(query).All(&drives)

	genInfo := []Models.GeneratorInfoModel{
		Models.GeneratorInfoModel{
			EquipName:   "krt",
			State:       0,
			Errors:      "все умерло",
			Workstation: 1,
			Heat:        1,
			Current:     10,
			Voltage:     66,
		},
		Models.GeneratorInfoModel{
			EquipName:   "krt",
			State:       1,
			Workstation: 2,
			Heat:        2,
			Current:     5,
			Voltage:     106,
		},
		Models.GeneratorInfoModel{
			EquipName: "krt",
			Voltage:   107,
		},
	}

	return genInfo
}

func dalCreateSession() *mgo.Session {
	session, err := mgo.Dial(Models.MongoDBConnectionString)
	if err != nil {
		panic(err)
	}

	return session
}
