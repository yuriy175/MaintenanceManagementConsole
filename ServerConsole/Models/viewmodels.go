package models

import "time"

type GeneratorInfoViewModel struct {
	ID    float64
	State GeneratorInfoModel
}

type StandInfoViewModel struct {
	ID    float64
	State StandInfoModel
}

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
	Model    string
	CPU_Load string
}

type MotherboardInfoViewModel struct {
	Model string
}

type MemoryInfoViewModel struct {
	Memory_total_Gb string
	Memory_free_Gb  string
}

type NetworkInfoViewModel struct {
	NIC string
	IP  string
}

type VGAInfoViewModel struct {
	Card_Name      string
	Driver_Version string
	Memory_Gb      string
}

type MonitorInfoViewModel struct {
	Device_Name string `json:"Device_Name"`
	Width       string `json:"Width"`
	Height      string `json:"Height"`
}

type PhysicalDiskInfoViewModel struct {
	FriendlyName string
	MediaType    string
	Size_Gb      string
}

type PrinterInfoViewModel struct {
	Printer_Name string
	Port_Name    string
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
	DB_name        string
	DB_Status      string
	DB_compability string
}

type AtlasInfoViewModel struct {
	Atlas_Version  string
	Complex_type   string
	Language       string
	Multimonitor   string
	XiLibs_Version string
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
