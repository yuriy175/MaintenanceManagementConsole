package DAL

import (
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../Models"
)

func DalWorker(equipDalCh chan *Models.EquipmentMessage) {
	quitCh := make(chan int)

	session := dalCreateSession()
	defer session.Close()

	// db := session.DB(Models.DBName)
	// deviceCollection := db.C(Models.DeviceConnectionsTableName)
	// studiesCollection := db.C(Models.StudyInWorkTableName)

	go func() { //c *mgo.Collection) {
		for d := range equipDalCh {
			if d.MsgType == Models.MsgTypeHwConnectionStateArrived {
				/*deviceId := d.Info["Id"].(float64)
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
				deviceCollection.Insert(model)*/
			} else if d.MsgType == Models.MsgTypeStudyInWork {
				/*studyId := d.Info["StudyId"].(float64)
				studyDicomUid := d.Info["StudyDicomUid"].(string)
				studyName := d.Info["StudyName"].(string)

				model := &Models.StudyInWorkModel{
					Id:            bson.NewObjectId(),
					DateTime:      time.Now(),
					EquipNumber:   d.EquipNumber,
					EquipName:     d.EquipName,
					EquipIP:       d.EquipIP,
					StudyId:       studyId,
					StudyDicomUid: studyDicomUid,
					StudyName:     studyName,
				}
				studiesCollection.Insert(model)*/
			} else if d.MsgType == Models.MsgTypeHddDrivesInfo {
				var params []Models.HddDrivesInfoMessage
				json.Unmarshal(d.Info, &params)

				i := 0
				i++

				/*hddName := d.Info["HddName"].(string)
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
				studiesCollection.Insert(model)*/
			}
		}
	}() //deviceCollection)

	<-quitCh
	fmt.Println("DalWorker quitted")
}

func DalGetDeviceConnections() []Models.DeviceConnectionModel {
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

func DalGetStudiesInWork() []Models.StudyInWorkModel {
	session := dalCreateSession()
	defer session.Close()

	studiesCollection := session.DB(Models.DBName).C(Models.StudyInWorkTableName)

	// критерий выборки
	query := bson.M{}
	// объект для сохранения результата
	studies := []Models.StudyInWorkModel{}
	studiesCollection.Find(query).All(&studies)

	return studies
}

func DalGetHddDrivesInfo() []Models.HddDrivesInfoModel {
	session := dalCreateSession()
	defer session.Close()

	drivesCollection := session.DB(Models.DBName).C(Models.HddDrivesInfoTableName)

	// критерий выборки
	query := bson.M{}
	// объект для сохранения результата
	drives := []Models.HddDrivesInfoModel{}
	drivesCollection.Find(query).All(&drives)

	return drives
}

func DalGetOrganAutoInfo() []Models.OrganAutoInfoModel {
	session := dalCreateSession()
	defer session.Close()

	organAutos := []Models.OrganAutoInfoModel{
		Models.OrganAutoInfoModel{
			EquipName:    "krt",
			OrganAuto:    "нога",
			Projection:   "кривая",
			Direction:    "задняя",
			AgeGroupId:   4,
			Constitution: 2,
		},
		Models.OrganAutoInfoModel{
			EquipName:    "krt",
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

func dalCreateSession() *mgo.Session {
	session, err := mgo.Dial(Models.MongoDBConnectionString)
	if err != nil {
		panic(err)
	}

	return session
}
