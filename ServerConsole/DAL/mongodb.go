package DAL

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../Models"
)

func DalWorker(devicesCh chan *Models.DeviceConnection) {
	quitCh := make(chan int)

	session, err := mgo.Dial(Models.MongoDBConnectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	deviceCollection := session.DB(Models.DBName).C(Models.DevicesTableName)

	go func(c *mgo.Collection) {
		for d := range devicesCh {
			model := Models.NewDeviceConnectionModel(d)
			err = c.Insert(model)
		}
	}(deviceCollection)

	<-quitCh
	fmt.Println("DalWorker quitted")
}

func DalGetDeviceConnections() []Models.DeviceConnectionModel {
	session, err := mgo.Dial(Models.MongoDBConnectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	deviceCollection := session.DB(Models.DBName).C(Models.DevicesTableName)

	// критерий выборки
	query := bson.M{}
	// объект для сохранения результата
	devices := []Models.DeviceConnectionModel{}
	deviceCollection.Find(query).All(&devices)

	return devices
}
