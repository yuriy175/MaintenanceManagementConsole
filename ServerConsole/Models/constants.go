package Models

const (
	HttpServerAddress = ":8181" // "localhost:8181"
	AppTasksNumber    = 2
	//MongoDBConnectionString       = "mongodb://127.0.0.1"
	//DBName                        = "test"
	EquipmentTableName            = "Equipment"
	DeviceConnectionsTableName    = "DeviceConnections"
	StudyInWorkTableName          = "StudyInWork"
	SystemInfoTableName           = "SystemInfo"
	SystemVolatileInfoTableName   = "SystemVolatileInfo"
	SoftwareInfoTableName         = "SoftwareInfo"
	SoftwareVolatileInfoTableName = "SoftwareVolatileInfo"
	DicomInfoTableName            = "DicomInfo"
	OrganAutoTableName            = "OrganAuto"
	GeneratorInfoTableName        = "GeneratorInfo"
	StandInfoTableName            = "StandInfo"
	DetectorInfoTableName         = "DetectorInfo"
	DosimeterInfoTableName        = "DosimeterInfo"

	RolesTableName = "Roles"
	UsersTableName = "Users"

	//RabbitMQConnectionString = "amqp://guest:guest@localhost:5672/"
	//"Server=mprom.ml;User=client1;Password=medtex"
	//RabbitMQHost          = "mprom.ml" // "localhost"
	//RabbitMQUser          = "client1"  // "guest"
	//RabbitMQPassword      = "medtex"   // "guest"
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
