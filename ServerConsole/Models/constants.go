package models

const (
	// HTTPServerAddress is a connecting string for a local http server
	HTTPServerAddress = ":8181" // "localhost:8181"

	// EquipmentTableName is Equipment table name
	EquipmentTableName = "Equipment"

	// DeviceConnectionsTableName is DeviceConnections table name
	DeviceConnectionsTableName = "DeviceConnections"

	// StudyInWorkTableName is StudyInWork table name
	StudyInWorkTableName = "StudyInWork"

	// SystemInfoTableName is SystemInfo table name
	SystemInfoTableName = "SystemInfo"

	// SystemVolatileInfoTableName is SystemVolatileInfo table name
	SystemVolatileInfoTableName = "SystemVolatileInfo"

	// SoftwareInfoTableName is SoftwareInfo table name
	SoftwareInfoTableName = "SoftwareInfo"

	// SoftwareVolatileInfoTableName is SoftwareVolatileInfo table name
	SoftwareVolatileInfoTableName = "SoftwareVolatileInfo"

	// DicomInfoTableName is DicomInfo table name
	DicomInfoTableName = "DicomInfo"

	// OrganAutoTableName is OrganAuto table name
	OrganAutoTableName = "OrganAuto"

	// GeneratorInfoTableName is GeneratorInfo table name
	GeneratorInfoTableName = "GeneratorInfo"
	StandInfoTableName     = "StandInfo"
	DetectorInfoTableName  = "DetectorInfo"
	DosimeterInfoTableName = "DosimeterInfo"

	RolesTableName = "Roles"
	UsersTableName = "Users"

	MQConnectionStateName = "HwConnectionStateArrived"
	MQInfoQueueName       = "SystemInfoQueue"

	MsgTypeInstanceOn               = "InstanceOn"
	MsgTypeInstanceOff              = "InstanceOff"
	MsgTypeStudyInWork              = "StudyInWork"
	MsgTypeNewImageCreated          = "NewImageCreated"
	MsgTypeHwConnectionStateArrived = "HwConnectionStateArrived"
	MsgTypeHddDrivesInfo            = "HddDrivesInfo"

	CommonTopicPath        = "Subscribe"
	BroadcastCommandsTopic = "Broadcast"
	WebSocketQueryString   = "/websocket"
)
