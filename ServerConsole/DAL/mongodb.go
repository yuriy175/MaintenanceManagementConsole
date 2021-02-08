package DAL

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../Models"
)

func DalWorker(equipDalCh chan *Models.EquipmentMessage) {
	quitCh := make(chan int)

	session, err := mgo.Dial(Models.MongoDBConnectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	db := session.DB(Models.DBName)
	deviceCollection := db.C(Models.DeviceConnectionsTableName)
	studiesCollection := db.C(Models.StudyInWorkTableName)

	go func() { //c *mgo.Collection) {
		for d := range equipDalCh {
			if d.MsgType == Models.MsgTypeHwConnectionStateArrived {
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
				err = deviceCollection.Insert(model)
			} else if d.MsgType == Models.MsgTypeStudyInWork {
				studyId := d.Info["StudyId"].(float64)
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
				err = studiesCollection.Insert(model)
			}
		}
	}() //deviceCollection)

	<-quitCh
	fmt.Println("DalWorker quitted")
}

func DalGetDeviceConnections() []Models.DeviceConnectionModel {
	session, err := mgo.Dial(Models.MongoDBConnectionString)
	if err != nil {
		panic(err)
	}
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
	session, err := mgo.Dial(Models.MongoDBConnectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	studiesCollection := session.DB(Models.DBName).C(Models.StudyInWorkTableName)

	// критерий выборки
	query := bson.M{}
	// объект для сохранения результата
	studies := []Models.StudyInWorkModel{}
	studiesCollection.Find(query).All(&studies)

	return studies
}
