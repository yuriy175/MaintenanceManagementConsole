package models

import "time"

// GeneratorInfoViewModel describes generator info view model from equipment
type GeneratorInfoViewModel struct {
	ID    float64
	State GeneratorInfoModel
}

// StandInfoViewModel describes stand info view model from equipment
type StandInfoViewModel struct {
	ID    float64
	State StandInfoModel
}

// EquipConnectionStateViewModel describes connection state view model to websocket
type EquipConnectionStateViewModel struct {
	Topic string
	State EquipConnectionState
}

type HddDriveInfoViewModel struct {
	Letter    string `json:"Letter"`
	TotalSize string `json:"TotalSize"`
	FreeSize  string `json:"FreeSize"`
}

type ProcessorInfoViewModel struct {
	Model   string
	CPULoad string
}

type MotherboardInfoViewModel struct {
	Model string
}

type MemoryInfoViewModel struct {
	MemoryTotalGb string
	MemoryFreeGb  string
}

type NetworkInfoViewModel struct {
	NIC string
	IP  string
}

type VGAInfoViewModel struct {
	CardName      string
	DriverVersion string
	MemoryGb      string
}

type MonitorInfoViewModel struct {
	DeviceName string `json:"Device_Name"`
	Width      string `json:"Width"`
	Height     string `json:"Height"`
}

type PhysicalDiskInfoViewModel struct {
	FriendlyName string
	MediaType    string
	SizeGb       string
}

type PrinterInfoViewModel struct {
	PrinterName string
	PortName    string
}

type SystemInfoViewModel struct {
	HDD           []HddDriveInfoViewModel
	PhysicalDisks []PhysicalDiskInfoViewModel
	Processor     ProcessorInfoViewModel
	Motherboard   MotherboardInfoViewModel
	Memory        MemoryInfoViewModel
	Network       []NetworkInfoViewModel
	VGA           []VGAInfoViewModel
	Monitor       []MonitorInfoViewModel
	Printer       []PrinterInfoViewModel
}

type SysInfoViewModel struct {
	OS           string
	Version      string
	Build_Number string
}

type MSSQLInfoViewModel struct {
	SQL     string
	Version string
	Status  string
}

type DatabasesViewModel struct {
	DBName        string
	DBStatus      string
	DBCompability string
}

type AtlasInfoViewModel struct {
	AtlasVersion  string
	ComplexType   string
	Language      string
	Multimonitor  string
	XiLibsVersion string
}

type SoftwareInfoViewModel struct {
	Sysinfo   SysInfoViewModel
	MSSQL     MSSQLInfoViewModel
	Databases []DatabasesViewModel
	Atlas     AtlasInfoViewModel
}

type MessageViewModel struct {
	Code        string
	Description string
}

type SoftwareMessageViewModel struct {
	ErrorDescriptions      []MessageViewModel
	AtlasErrorDescriptions []MessageViewModel
}

type UserViewModel struct {
	ID       string
	Login    string
	Password string
	Surname  string
	Role     string
	Email    string
	Disabled bool
}

type EquipInfoViewModel struct {
	HospitalName      string
	HospitalAddress   string
	HospitalLongitude string
	HospitalLatitude  string
}

type DetailedEquipInfoViewModel struct {
	RegisterDate      time.Time
	EquipName         string
	HospitalName      string
	HospitalAddress   string
	HospitalLongitude string
	HospitalLatitude  string
	IsActive          bool
}
