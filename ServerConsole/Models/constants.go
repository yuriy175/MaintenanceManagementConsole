package Models

const (
	AppTasksNumber             = 2
	MongoDBConnectionString    = "mongodb://127.0.0.1"
	DBName                     = "test"
	EquipmentTableName         = "Equipment"
	DeviceConnectionsTableName = "DeviceConnections"

	RabbitMQConnectionString = "amqp://guest:guest@localhost:5672/"
	MQConnectionStateName    = "HwConnectionStateArrived"
	MQInfoQueueName          = "SystemInfoQueue"

	MsgTypeStudyInWork              = "StudyInWork"
	MsgTypeNewImageCreated          = "NewImageCreated"
	MsgTypeHwConnectionStateArrived = "HwConnectionStateArrived"
)
