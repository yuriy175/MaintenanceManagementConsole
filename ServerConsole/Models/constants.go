package Models

const (
	HttpServerAddress          = "localhost:8181"
	AppTasksNumber             = 2
	MongoDBConnectionString    = "mongodb://127.0.0.1"
	DBName                     = "test"
	EquipmentTableName         = "Equipment"
	DeviceConnectionsTableName = "DeviceConnections"
	StudyInWorkTableName       = "StudyInWork"
	HddDrivesInfoTableName     = "HddDrivesInfo"

	//RabbitMQConnectionString = "amqp://guest:guest@localhost:5672/"
	//"Server=mprom.ml;User=client1;Password=medtex"
	RabbitMQHost          = "mprom.ml" // "localhost"
	RabbitMQUser          = "client1"  // "guest"
	RabbitMQPassword      = "medtex"   // "guest"
	MQConnectionStateName = "HwConnectionStateArrived"
	MQInfoQueueName       = "SystemInfoQueue"

	MsgTypeInstanceOn               = "InstanceOn"
	MsgTypeInstanceOff              = "InstanceOff"
	MsgTypeStudyInWork              = "StudyInWork"
	MsgTypeNewImageCreated          = "NewImageCreated"
	MsgTypeHwConnectionStateArrived = "HwConnectionStateArrived"
	MsgTypeHddDrivesInfo            = "HddDrivesInfo"
)
