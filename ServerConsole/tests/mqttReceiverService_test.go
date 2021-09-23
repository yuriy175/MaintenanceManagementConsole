package tests

import (
	"ServerConsole/interfaces"
	"ServerConsole/tests/mocks"
)

var mqttReceiverServiceTest interfaces.IMqttReceiverService

func setUpMqttReceiverService() interfaces.IMqttReceiverService {

	if mqttReceiverServiceTest == nil {
		mqttReceiverServiceTest = mocks.InitMockIoc().GetMqttReceiverService() // EquipsServiceNew(nil, mocks.DataLayerMockServiceNew(), nil, nil)
	}

	return mqttReceiverServiceTest
}

