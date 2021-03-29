package Models

/*type DeviceConnection struct {
	DeviceId   int    `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Connection int    `json:"connection"`
}

type EquipmentMessage struct {
	EquipNumber string `json:"Number"`
	EquipName   string `json:"Name"`
	EquipIP     string `json:"ipAddress"`
	MsgType     string `json:"msgType"`
	//Info        map[string]interface{} `json:"info"` // Rest of the fields should go here.
	Info json.RawMessage `json:"info"`
	//Info string `json:"info"` // Rest of the fields should go here.
	//Info string `json:"-"` // Rest of the fields should go here.
}
*/

type GeneratorInfoViewModel struct {
	Id    float64
	State GeneratorInfoModel
}

type StandInfoViewModel struct {
	Id    float64
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
	Monitor_Name   string `json:"Device_Name"`
	Monitor_Width  string `json:"Width"`
	Monitor_Height string `json:"Height"`
}

type SystemInfoViewModel struct {
	HDD         []HddDriveInfoViewModel
	Processor   ProcessorInfoViewModel
	Motherboard MotherboardInfoViewModel
	Memory      MemoryInfoViewModel
	Network     []NetworkInfoViewModel
	VGA         []VGAInfoViewModel
	Monitor     []MonitorInfoViewModel
}
