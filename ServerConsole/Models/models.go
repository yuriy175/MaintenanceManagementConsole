package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

// StudyInWorkModel describes study in work model
type StudyInWorkModel struct {
	ID            bson.ObjectId `bson:"_id"`
	DateTime      time.Time
	EquipName     string
	StudyID       float64
	StudyDicomUID string
	StudyName     string
	State         float64
}

// RawMqttMessage describes a raw mqtt message from equipment
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
	AgeID        float64
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

// SysInfoModel describes OS info model from equipment
type SysInfoModel struct {
	OS          string
	Version     string
	BuildNumber string
}

// MSSQLInfoModel describes general MSSQL info model from equipment
type MSSQLInfoModel struct {
	SQL     string
	Version string
	Status  string
}

// DatabasesModel describes a database info model from equipment
type DatabasesModel struct {
	DBName        string
	DBStatus      string
	DBCompability string
}

// AtlasInfoModel describes Atlas info model from equipment
type AtlasInfoModel struct {
	AtlasVersion  string
	ComplexType   string
	Language      string
	Multimonitor  string
	XiLibsVersion string
}

// SoftwareVolatileInfoModel describes software volatile info DB model
type SoftwareVolatileInfoModel struct {
	ID               bson.ObjectId `bson:"_id"`
	DateTime         time.Time
	EquipName        string
	ErrorCode        string
	ErrorDescription string
}

// SoftwareInfoModel describes software permanent info DB model
type SoftwareInfoModel struct {
	ID        bson.ObjectId `bson:"_id"`
	DateTime  time.Time
	EquipName string
	Sysinfo   SysInfoModel
	MSSQL     MSSQLInfoModel
	Databases []DatabasesModel
	Atlas     AtlasInfoModel
}

// FullSoftwareInfoModel describes full software info DB model
type FullSoftwareInfoModel struct {
	PermanentInfo []SoftwareInfoModel
	VolatileInfo  []SoftwareVolatileInfoModel
}

// UserModel describes user info DB model
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

// EquipInfoModel describes hospital info DB model
type EquipInfoModel struct {
	ID                bson.ObjectId `bson:"_id"`
	RegisterDate      time.Time
	EquipName         string
	HospitalName      string
	HospitalAddress   string
	HospitalLongitude string
	HospitalLatitude  string
}

// RawDeviceInfoModel describes raw device data from equipment info DB model
type RawDeviceInfoModel struct {
	ID        bson.ObjectId `bson:"_id"`
	DateTime  time.Time
	EquipName string
	Data      string
}

// AllDBInfoModel describes all db info from equipment info DB model
type AllDBInfoModel struct {
	ID        bson.ObjectId `bson:"_id"`
	DateTime  time.Time
	EquipName string
	Hospital      map[string]json.RawMessage//[]byte
	Software   map[string]json.RawMessage//string
	System map[string]json.RawMessage//string
	Atlas  map[string]json.RawMessage//string
}

// AllDBInfoModel describes all db info from equipment info DB model
type AllDBTablesModel struct {
	EquipName string
	Hospital      []string
	Software   []string
	System []string
	Atlas  []string
}