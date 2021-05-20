package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type DeviceConnectionModel struct {
	ID               bson.ObjectId `bson:"_id"`
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
	ID            bson.ObjectId `bson:"_id"`
	DateTime      time.Time
	EquipName     string
	StudyId       float64
	StudyDicomUid string
	StudyName     string
	State         float64
}

/*
type SystemInfoModel struct {
	Id        bson.ObjectId `bson:"_id"`
	DateTime  time.Time
	EquipName string
		State         float64
		CPU_Load      float64
		TotalMemory   float64
		AvailableSize float64
		HddName       string
		HddTotalSpace float64
		HddFreeSpace  float64
}
*/

type RawMqttMessage struct {
	Topic string
	Data  string
}

type OrganAutoInfoModel struct {
	ID           bson.ObjectId `bson:"_id"`
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
	Code        string
	Description string
}

type GeneratorInfoModel struct {
	ID                bson.ObjectId `bson:"_id"`
	DateTime          time.Time
	EquipName         string
	State             float64
	Errors            []string
	Workstation       float64
	HeatStatus        float64
	Mas               float64
	Kv                float64
	ErrorDescriptions []ErrorDescription
}

// type SoftwareInfoModel struct {
// 	Id                bson.ObjectId `bson:"_id"`
// 	DateTime          time.Time
// 	EquipName         string
// 	SettingsDB        bool
// 	ObservationsDB    bool
// 	Version           string
// 	XilibVersion      string
// 	ErrorDescriptions []ErrorDescription
// }

type DicomInfo struct {
	Name  string
	IP    string
	State float64
}

type DicomsInfoModel struct {
	ID        bson.ObjectId `bson:"_id"`
	DateTime  time.Time
	EquipName string
	PACS      []DicomInfo
	WorkList  []DicomInfo
}

type StandInfoModel struct {
	ID                bson.ObjectId `bson:"_id"`
	DateTime          time.Time
	EquipName         string
	State             float64
	Errors            []string
	RasterState       float64
	Position_Current  float64
	ErrorDescriptions []ErrorDescription
}

type EquipConnectionState struct {
	Name      string
	Connected bool
}

type HDDVolatileInfoModel struct {
	Letter   string
	FreeSize string
}

type SystemVolatileInfoModel struct {
	ID                 bson.ObjectId `bson:"_id"`
	DateTime           time.Time
	EquipName          string
	HDD                []HDDVolatileInfoModel
	ProcessorCPULoad   string
	MemoryMemoryFreeGb string
}

////
type HddDriveInfoModel struct {
	Letter    string
	TotalSize string
}

type ProcessorInfoModel struct {
	Model string
}

type MotherboardInfoModel struct {
	Model string
}

type MemoryInfoModel struct {
	MemoryTotalGb string
}

type NetworkInfoModel struct {
	NIC string
	IP  string
}

type VGAInfoModel struct {
	CardName      string
	DriverVersion string
	MemoryGb      string
}

type MonitorInfoModel struct {
	DeviceName string
	Width      string
	Height     string
}

type PhysicalDiskInfoModel struct {
	FriendlyName string
	MediaType    string
	SizeGb       string
}

type PrinterInfoModel struct {
	PrinterName string
	Port_Name   string
}

////

type SystemInfoModel struct {
	ID        bson.ObjectId `bson:"_id"`
	DateTime  time.Time
	EquipName string
	//Parameter string
	//Value     string
	HDD           []HddDriveInfoModel
	PhysicalDisks []PhysicalDiskInfoModel
	Processor     ProcessorInfoModel
	Motherboard   MotherboardInfoModel
	Memory        MemoryInfoModel
	Network       []NetworkInfoModel
	VGA           []VGAInfoModel
	Monitor       []MonitorInfoModel
	Printer       []PrinterInfoModel
}

type FullSystemInfoModel struct {
	PermanentInfo []SystemInfoModel
	VolatileInfo  []SystemVolatileInfoModel
}

type SysInfoModel struct {
	OS           string
	Version      string
	Build_Number string
}

type MSSQLInfoModel struct {
	SQL     string
	Version string
	Status  string
}

type DatabasesModel struct {
	DBName        string
	DBStatus      string
	DBCompability string
}

type AtlasInfoModel struct {
	AtlasVersion  string
	ComplexType   string
	Language      string
	Multimonitor  string
	XiLibsVersion string
}

type SoftwareVolatileInfoModel struct {
	ID               bson.ObjectId `bson:"_id"`
	DateTime         time.Time
	EquipName        string
	ErrorCode        string
	ErrorDescription string
}

type SoftwareInfoModel struct {
	ID        bson.ObjectId `bson:"_id"`
	DateTime  time.Time
	EquipName string
	Sysinfo   SysInfoModel
	MSSQL     MSSQLInfoModel
	Databases []DatabasesModel
	Atlas     AtlasInfoModel
}

type FullSoftwareInfoModel struct {
	PermanentInfo []SoftwareInfoModel
	VolatileInfo  []SoftwareVolatileInfoModel
}

type UserModel struct {
	ID           bson.ObjectId `bson:"_id"`
	DateTime     time.Time
	Login        string
	PasswordHash [32]byte
	Surname      string
	Role         string
	Email        string
	Disabled     bool
}

type RabbitMQSettingsModel struct {
	Host     string
	User     string
	Password string
}

type MongoDBSettingsModel struct {
	ConnectionString string
	DBName           string
}

type EquipInfoModel struct {
	ID                bson.ObjectId `bson:"_id"`
	RegisterDate      time.Time
	EquipName         string
	HospitalName      string
	HospitalAddress   string
	HospitalLongitude string
	HospitalLatitude  string
}

type RawDeviceInfoModel struct {
	ID        bson.ObjectId `bson:"_id"`
	DateTime  time.Time
	EquipName string
	Data      string
}
