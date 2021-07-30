package models

import (
	"time"

	"encoding/json"

	"gopkg.in/mgo.v2/bson"
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

// OrganAutoInfoModel describes organ auto info model from equipment
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

// ErrorDescription describes error message from equipment
type ErrorDescription struct {
	Code        string
	Description string
}

// GeneratorInfoModel describes generator info model from equipment
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

// DicomInfo describes dicom connection model from equipment
type DicomInfo struct {
	Name  string
	IP    string
	State float64
}

// DicomsInfoModel describes all dicom connections model from equipment
type DicomsInfoModel struct {
	ID        bson.ObjectId `bson:"_id"`
	DateTime  time.Time
	EquipName string
	PACS      []DicomInfo
	WorkList  []DicomInfo
}

// StandInfoModel describes stand info model from equipment
type StandInfoModel struct {
	ID                bson.ObjectId `bson:"_id"`
	DateTime          time.Time
	EquipName         string
	State             float64
	Errors            []string
	RasterState       float64
	PositionCurrent   float64
	ErrorDescriptions []ErrorDescription
}

// EquipConnectionState describes connection state model to websocket
type EquipConnectionState struct {
	Name      string
	Connected bool
}

// HDDVolatileInfoModel describes HDD volatile info model from equipment
type HDDVolatileInfoModel struct {
	Letter   string
	FreeSize string
}

// SystemVolatileInfoModel describes system volatile hardware info model from equipment
type SystemVolatileInfoModel struct {
	ID                 bson.ObjectId `bson:"_id"`
	DateTime           time.Time
	EquipName          string
	HDD                []HDDVolatileInfoModel
	ProcessorCPULoad   string
	MemoryMemoryFreeGb string
}

// HddDriveInfoModel describes HDD info model from equipment
type HddDriveInfoModel struct {
	Letter    string
	TotalSize string
}

// ProcessorInfoModel describes CPU info model from equipment
type ProcessorInfoModel struct {
	Model string
}

// MotherboardInfoModel describes motherboard info model from equipment
type MotherboardInfoModel struct {
	Model string
}

// MemoryInfoModel describes memory info model from equipment
type MemoryInfoModel struct {
	MemoryTotalGb string
}

// NetworkInfoModel describes LAN adapters info model from equipment
type NetworkInfoModel struct {
	NIC string
	IP  string
}

// VGAInfoModel describes videoadapters info model from equipment
type VGAInfoModel struct {
	CardName      string
	DriverVersion string
	MemoryGb      string
}

// MonitorInfoModel describes monitors info model from equipment
type MonitorInfoModel struct {
	DeviceName string
	Width      string
	Height     string
}

// PhysicalDiskInfoModel describes HDD info model from equipment
type PhysicalDiskInfoModel struct {
	FriendlyName string
	MediaType    string
	SizeGb       string
}

// PrinterInfoModel describes printers info model from equipment
type PrinterInfoModel struct {
	PrinterName string
	PortName    string
}

// SystemInfoModel describes system permanent hardware info model from equipment
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

// FullSystemInfoModel describes full system hardware model from equipment
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

// RabbitMQSettingsModel describes rabbitMQ settings model
type RabbitMQSettingsModel struct {
	Host     string
	User     string
	Password string
}

// MongoDBSettingsModel describes mongodb settings model
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
	Disabled     bool
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
	Hospital  map[string]json.RawMessage //[]byte
	Software  map[string]json.RawMessage //string
	System    map[string]json.RawMessage //string
	Atlas     map[string]json.RawMessage //string
}

// AllDBTablesModel describes all db info from equipment info DB model
type AllDBTablesModel struct {
	EquipName string
	Hospital  []string
	Software  []string
	System    []string
	Atlas     []string
}

// DBSoftwareInfoModel describes both general and atlas software permanent info DB model
type DBSoftwareInfoModel struct {
	Software []map[string]json.RawMessage
	Atlas    []map[string]json.RawMessage
}

// EventModel describes event DB model
type EventModel struct {
	ID               bson.ObjectId `bson:"_id"`
	DateTime         time.Time
	EquipName        string
	Type        string
	Title        string
	Description string
	Details string
}

// ChatModel describes chat notes DB model
type ChatModel struct {
	ID               bson.ObjectId `bson:"_id"`
	DateTime         time.Time
	EquipName        string
	Type        string
	Message        string
	User string
	Hidden bool
	IsInternal bool 
}

// FullEquipInfoModel describes full equipment permanent info view model to ui
type FullEquipInfoModel struct {
	SoftwareInfo   DBSoftwareInfoModel
	SystemInfo []map[string]json.RawMessage
	LastSeen time.Time
	LocationInfo EquipInfoModel
}

// ServerState describes server state model
type ServerState struct {	
	DBVolume        string
	DiskTotalSpace string
	DiskFreeSpace string
}