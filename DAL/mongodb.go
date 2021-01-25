package DAL

import (
	"fmt"

	"gopkg.in/mgo.v2"

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
