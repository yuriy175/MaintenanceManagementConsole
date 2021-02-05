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

	deviceCollection := session.DB(Models.DBName).C(Models.DeviceConnectionsTableName)

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

	deviceCollection := session.DB(Models.DBName).C(Models.EquipmentTableName)

	// критерий выборки
	query := bson.M{}
	// объект для сохранения результата
	devices := []Models.DeviceConnectionModel{}
	deviceCollection.Find(query).All(&devices)

	return devices
}
