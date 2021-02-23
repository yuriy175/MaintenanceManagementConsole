package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type DeviceConnectionModel struct {
	Id               bson.ObjectId `bson:"_id"`
	DateTime         time.Time
	EquipNumber      string
	EquipName        string
	EquipIP          string
	DeviceId         float64
	DeviceName       string
	DeviceType       string
	DeviceConnection float64
}

type StudyInWorkModel struct {
	Id            bson.ObjectId `bson:"_id"`
	DateTime      time.Time
	EquipNumber   string
	EquipName     string
	EquipIP       string
	StudyId       float64
	StudyDicomUid string
	StudyName     string
}

type HddDrivesInfoModel struct {
	Id            bson.ObjectId `bson:"_id"`
	DateTime      time.Time
	EquipNumber   string
	EquipName     string
	EquipIP       string
	HddName       string
	HddTotalSpace float64
	HddFreeSpace  float64
}
