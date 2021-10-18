package models

import (
	"encoding/json"
	"time"
)

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

// HddDriveInfoViewModel describes HDD info view model from equipment
type HddDriveInfoViewModel struct {
	Letter    string `json:"Letter"`
	TotalSize string `json:"TotalSize"`
	FreeSize  string `json:"FreeSize"`
}

// ProcessorInfoViewModel describes CPU info view model from equipment
type ProcessorInfoViewModel struct {
	Model   string
	CPULoad string
}

// MotherboardInfoViewModel describes motherboard info view model from equipment
type MotherboardInfoViewModel struct {
	Model string
}

// MemoryInfoViewModel describes memory info view model from equipment
type MemoryInfoViewModel struct {
	MemoryTotalGb string
	MemoryFreeGb  string
}

// NetworkInfoViewModel describes LAN adapters info view model from equipment
type NetworkInfoViewModel struct {
	NIC string
	IP  string
}

// VGAInfoViewModel describes videoadapters info view model from equipment
type VGAInfoViewModel struct {
	CardName      string
	DriverVersion string
	MemoryGb      string
}

// MonitorInfoViewModel describes monitors info view model from equipment
type MonitorInfoViewModel struct {
	DeviceName string `json:"Device_Name"`
	Width      string `json:"Width"`
	Height     string `json:"Height"`
}

// PhysicalDiskInfoViewModel describes HDD info view model from equipment
type PhysicalDiskInfoViewModel struct {
	FriendlyName string
	MediaType    string
	SizeGb       string
}

// PrinterInfoViewModel describes printers info view model from equipment
type PrinterInfoViewModel struct {
	PrinterName string
	PortName    string
}

// SystemInfoViewModel describes all system hardware view model from equipment
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

// SysInfoViewModel describes OS info view model from equipment
type SysInfoViewModel struct {
	OS          string
	Version     string
	BuildNumber string
}

// MSSQLInfoViewModel describes MSSQL info view model from equipment
type MSSQLInfoViewModel struct {
	SQL     string
	Version string
	Status  string
}

// DatabasesViewModel describes a database info view model from equipment
type DatabasesViewModel struct {
	DBName        string
	DBStatus      string
	DBCompability string
}

// AtlasInfoViewModel describes Atlas info view model from equipment
type AtlasInfoViewModel struct {
	AtlasVersion  string
	ComplexType   string
	Language      string
	Multimonitor  string
	XiLibsVersion string
}

// SoftwareInfoViewModel describes all software info view model from equipment
type SoftwareInfoViewModel struct {
	Sysinfo   SysInfoViewModel
	MSSQL     MSSQLInfoViewModel
	Databases []DatabasesViewModel
	Atlas     AtlasInfoViewModel
}

// MessageViewModel describes a message view model from equipment
type MessageViewModel struct {
	Level       string
	Code        string
	Description string
}

// AtlasUserViewModel describes a atlas user view model from equipment
type AtlasUserViewModel struct {
	User string
	Role string
}

// OfflineMsgViewModel describes a event offline view model from equipment
type OfflineMsgViewModel struct {
	Message  *SoftwareMessageViewModel
	DateTime string
}

// SoftwareMessageViewModel describes a software message view model from equipment
type SoftwareMessageViewModel struct {
	ErrorDescriptions         []MessageViewModel
	AtlasErrorDescriptions    []MessageViewModel
	HardwareErrorDescriptions []MessageViewModel
	AtlasUser                 AtlasUserViewModel
	OfflineMsg                *OfflineMsgViewModel
	SimpleMsgType             string
}

// UserViewModel describes user info view model to UI
type UserViewModel struct {
	ID       string
	Login    string
	Password string
	Surname  string
	Role     string
	Email    string
	Disabled bool
}

// EquipInfoViewModel describes hospital info view model from equipment
type EquipInfoViewModel struct {
	HospitalName      string
	HospitalAddress   string
	HospitalLongitude float32
	HospitalLatitude  float32
}

// DetailedEquipInfoViewModel describes hospital info view model to UI
type DetailedEquipInfoViewModel struct {
	RegisterDate      time.Time
	EquipName         string
	HospitalName      string
	HospitalAddress   string
	HospitalLongitude string
	HospitalLatitude  string
	IsActive          bool
	LastSeen          time.Time
}

// AllDBInfoViewModel describes all db info view model from equipment
type AllDBInfoViewModel struct {
	Hospital map[string]json.RawMessage //[]byte
	Software map[string]json.RawMessage //string
	System   map[string]json.RawMessage //string
	Atlas    map[string]json.RawMessage //string
}

// EventViewModel describes event view model to websocket
type EventsViewModel struct {
	Topic  string
	Events []EventModel
}

// ChatViewModel describes chat note view model
type ChatViewModel struct {
	Message    string
	User       string
	IsInternal bool
}

// TokenWithUserViewModel describes token and user info view model to UI
type TokenWithUserViewModel struct {
	Token   string
	Surname string
}

// StarpupInfoViewModel describes equipment start up info view model
type StartupInfoViewModel struct {
	StartupTime   *time.Time
	KernelPower41 *time.Time
}
