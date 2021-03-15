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
	CPU_Load      float64
	TotalMemory   float64
	AvailableSize float64
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
	Name         string
	Projection   string
	Direction    string
	AgeId        float64
	Constitution float64
}

type ErrorDescription struct {
	Code string
	Description string
}

type GeneratorInfoModel struct {
	Id          bson.ObjectId `bson:"_id"`
	DateTime    time.Time
	EquipName   string
	State       float64
	Errors      []string
	Workstation float64
	HeatStatus  float64
	Mas         float64
	Kv          float64
	ErrorDescriptions []ErrorDescription
}

type SoftwareInfoModel struct {
	Id          bson.ObjectId `bson:"_id"`
	DateTime    time.Time
	EquipName   string
	SettingsDB bool
    ObservationsDB bool
    Version string
    XilibVersion string
	ErrorDescriptions []ErrorDescription
}

type DicomInfo struct {
	Name string
	IP string
	State       float64
}

type DicomsInfoModel struct {
	Id          bson.ObjectId `bson:"_id"`
	DateTime    time.Time
	EquipName   string
	PACS []DicomInfo
	WorkList []DicomInfo
}

type StandInfoModel struct {
	Id          bson.ObjectId `bson:"_id"`
	DateTime    time.Time
	EquipName   string
	State       float64
	Errors      []string
	RasterState float64
	Position_Current  float64
	ErrorDescriptions []ErrorDescription
}