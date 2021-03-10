package Models

import "encoding/json"

type DeviceConnection struct {
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

type HddDrivesInfoMessage struct {
	HddName       string  `json:"Letter"`
	HddTotalSpace float64 `json:"TotalSize"`
	HddFreeSpace  float64 `json:"FreeSize"`
}

type GeneratorInfoViewModel struct {
	Id    float64
	State GeneratorInfoModel
}
