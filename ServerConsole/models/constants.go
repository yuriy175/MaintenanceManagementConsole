package models

const (
	// HTTPServerAddress is a connecting string for a local http server
	// HTTPServerAddress = ":8181" // "localhost:8181"

	// EquipmentTableName is Equipment table name
	EquipmentTableName = "Equipment"

	// RenamedEquipmentTableName is RenamedEquipment table name
	RenamedEquipmentTableName = "RenamedEquipment"

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

	// StandInfoTableName is StandInfo table name
	StandInfoTableName = "StandInfo"

	// DetectorInfoTableName is DetectorInfo table name
	DetectorInfoTableName = "DetectorInfo"

	// DosimeterInfoTableName is DosimeterInfoT table name
	DosimeterInfoTableName = "DosimeterInfo"

	// AllDBInfoTableName is AllDBInfo table name
	AllDBInfoTableName = "AllDBInfo"

	// RolesTableName is Roles table name
	RolesTableName = "Roles"

	// UsersTableName is Users table name
	UsersTableName = "Users"

	// EventsTableName is Events table name
	EventsTableName = "Events"

	// ChatsTableName is Chats table name
	ChatsTableName = "Chats"

	// EquipInfoTableName is Equipment info table name
	EquipInfoTableName = "EquipInfo"

	// MQInfoQueueName is rabbit mq queue name
	MQInfoQueueName = "SystemInfoQueue"

	// CommonTopicPath is common mqtt topic
	CommonTopicPath = "Subscribe"

	// CommonChatsPath is common mqtt topic for chats
	CommonChatsPath = "Chats"

	// BroadcastCommandsTopic is broadcasting command mqtt topic
	BroadcastCommandsTopic = "Broadcast"

	// WebSocketQueryString is websocket url query subpath
	WebSocketQueryString = "/websocket"

	// EventsTopicPath is common events topic
	EventsTopicPath = "Events"

	// AdminRoleName is administrator role name
	AdminRoleName = "Administrator"

	// DefaultAdminName is default administrator account name
	DefaultAdminName = "sa"

	// KeepAliveCheckPeriod is KeepAlive messages check period, sec
	KeepAliveCheckPeriod = 10

	// CommonChat is mqtt topic for common chat
	CommonChat = "CommonChat"

	// CommonKeepAlive is mqtt topic for keepalive messages
	CommonKeepAlive = "CommonKeepAlive"
)
