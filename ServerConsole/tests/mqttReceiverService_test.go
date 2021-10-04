package tests

import (
	"strconv"
	"testing"
	"time"

	"ServerConsole/interfaces"
	"ServerConsole/models"
	"ServerConsole/tests/mocks"
)

const (
	NumberOfMqttConnections = 100
)

var mqttReceiverServiceTest interfaces.IMqttReceiverService

func setUpMqttReceiverService() interfaces.IMqttReceiverService {

	ioc := mocks.InitMockIoc().(*mocks.MockTypes)
	return ioc.CreateMqttReceiverService()
	/*if mqttReceiverServiceTest == nil {
		mqttReceiverServiceTest = mocks.InitMockIoc().GetMqttReceiverService() // EquipsServiceNew(nil, mocks.DataLayerMockServiceNew(), nil, nil)
	}

	return mqttReceiverServiceTest*/
}

func TestUpdateMqttConnections_100ConnectionCreated(t *testing.T) {
	service := setUpMqttReceiverService()
	initialConnectionNamesNumber := len(service.GetConnectionNames())
	equipName := "ABC/WKS_ABC"
	for i := 0; i < NumberOfMqttConnections; i++ {
		state := &models.EquipConnectionState{
			Name:      equipName + strconv.Itoa(i),
			Connected: true,
		}
		service.UpdateMqttConnections(state)
	}

	connectionNames := service.GetConnectionNames()
	connectionNamesNumber := NumberOfMqttConnections + initialConnectionNamesNumber
	if len(connectionNames) != connectionNamesNumber {
		t.Errorf(`UpdateMqttConnections created number = %v`, connectionNamesNumber)
	}
}

func TestUpdateMqttConnections_100ConnectionCreatedAndTimedOut(t *testing.T) {
	service := setUpMqttReceiverService()
	initialConnectionNamesNumber := len(service.GetConnectionNames())

	equipName := "ABC/WKS_ABC"
	for i := 0; i < NumberOfMqttConnections; i++ {
		state := &models.EquipConnectionState{
			Name:      equipName + strconv.Itoa(i),
			Connected: true,
		}
		service.UpdateMqttConnections(state)
	}

	connectionNames := service.GetConnectionNames()
	connectionNamesNumber := NumberOfMqttConnections + initialConnectionNamesNumber
	if len(connectionNames) != connectionNamesNumber {
		t.Errorf(`UpdateMqttConnections created number = %v`, connectionNamesNumber)
		return
	}

	for i := 0; i < NumberOfMqttConnections; i++ {
		service.SetKeepAliveReceived(equipName + strconv.Itoa(i))
	}
	time.Sleep((models.KeepAliveCheckPeriod*2 + 1) * time.Second)

	connectionNames = service.GetConnectionNames()
	if len(connectionNames) != 0 {
		t.Errorf(`UpdateMqttConnections after timeout number = %v`, len(connectionNames))
	}
}
