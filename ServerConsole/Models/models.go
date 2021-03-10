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
	EquipName     string
	StudyId       float64
	StudyDicomUid string
	StudyName     string
	State         float64
}

type SystemInfoModel struct {
	Id            bson.ObjectId `bson:"_id"`
	DateTime      time.Time
	EquipName     string
	State         float64
	CPULoad       float64
	TotalMemory   float64
	FreeMemory    float64
	HddName       string
	HddTotalSpace float64
	HddFreeSpace  float64
}

type RawMqttMessage struct {
	Topic string
	Data  string
}

type OrganAutoInfoModel struct {
	Id           bson.ObjectId `bson:"_id"`
	DateTime     time.Time
	EquipName    string
	State        float64
	OrganAuto    string
	Projection   string
	Direction    string
	AgeGroupId   float64
	Constitution float64
}

type GeneratorInfoModel struct {
	Id          bson.ObjectId `bson:"_id"`
	DateTime    time.Time
	EquipName   string
	State       float64
	Errors      string
	Workstation float64
	Heat        float64
	Current     float64
	Voltage     float64
}
