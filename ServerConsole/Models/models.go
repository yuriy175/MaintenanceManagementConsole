package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type DeviceConnectionModel struct {
	Id         bson.ObjectId `bson:"_id"`
	DateTime   time.Time
	DeviceId   int
	Name       string
	Type       string
	Connection int
}

func NewDeviceConnectionModel(vm *DeviceConnection) *DeviceConnectionModel {
	return &DeviceConnectionModel{
		Id:         bson.NewObjectId(),
		DateTime:   time.Now(),
		DeviceId:   vm.DeviceId,
		Name:       vm.Name,
		Type:       vm.Type,
		Connection: vm.Connection}
}
