package tests

import (
	"time"
	"testing"
	"strconv"

	"ServerConsole/interfaces"
	"ServerConsole/tests/mocks"
	"ServerConsole/models"
)

const(
	NumberOfMqttConnections = 100
)

var mqttReceiverServiceTest interfaces.IMqttReceiverService

func setUpMqttReceiverService() interfaces.IMqttReceiverService {

	if mqttReceiverServiceTest == nil {
		mqttReceiverServiceTest = mocks.InitMockIoc().GetMqttReceiverService() // EquipsServiceNew(nil, mocks.DataLayerMockServiceNew(), nil, nil)
	}

	return mqttReceiverServiceTest
}

func TestUpdateMqttConnections_100ConnectionCreated(t *testing.T) {
	service := setUpMqttReceiverService()
	equipName := "ABC/WKS_ABC"
	for i := 0; i < NumberOfMqttConnections; i++{
		state := &models.EquipConnectionState{
			Name: equipName + strconv.Itoa(i),
			Connected: true,
		}
        service.UpdateMqttConnections(state)
    }
	
	connectionNames := service.GetConnectionNames()
	if len(connectionNames) != NumberOfMqttConnections{
		t.Errorf(`UpdateMqttConnections created number = %v`, len(connectionNames))
	}
}

func TestUpdateMqttConnections_100ConnectionCreatedAndTimedOut(t *testing.T) {
	service := setUpMqttReceiverService()
	equipName := "ABC/WKS_ABC"
	for i := 0; i < NumberOfMqttConnections; i++{
		state := &models.EquipConnectionState{
			Name: equipName + strconv.Itoa(i),
			Connected: true,
		}
        service.UpdateMqttConnections(state)
    }
	
	connectionNames := service.GetConnectionNames()
	if len(connectionNames) != NumberOfMqttConnections{
		t.Errorf(`UpdateMqttConnections created number = %v`, len(connectionNames))
	}

	for i := 0; i < NumberOfMqttConnections; i++{
        service.SetKeepAliveReceived(equipName + strconv.Itoa(i))
    }
	time.Sleep((models.KeepAliveCheckPeriod * 2 + 1) * time.Second)

	connectionNames = service.GetConnectionNames()
	if len(connectionNames) != 0{
		t.Errorf(`UpdateMqttConnections after timeout number = %v`, len(connectionNames))
	}
}

func TestUpdateMqttConnections_1Connected_TimedOut_Reconnected_Active(t *testing.T) {
	service := setUpMqttReceiverService()
	equipService := setUpEquipsService()
	equipsChan := mocks.InitMockIoc().GetEquipsChan()

	equipName := "ABC/WKS_ABC"
	state := &models.EquipConnectionState{
			Name: equipName,
			Connected: true,
	}
    service.UpdateMqttConnections(state)
	service.SetKeepAliveReceived(equipName)	

	msg := &models.RawMqttMessage{
		Topic: equipName + "/hospital",
		Data: "",
		Arrival: time.Now(),
	}

	equipsChan <- msg
	time.Sleep(1 * time.Second)

	equips := equipService.GetEquipInfos(false)

	if len(equips) != 1{
		t.Errorf(`GetEquipInfos before timeout number = %v `, len(equips))
		return
	}

	if !equips[0].IsActive{
		t.Errorf(`GetEquipInfos before timeout active = %v`, equips[0].IsActive)
		return
	}

	time.Sleep((models.KeepAliveCheckPeriod * 2 + 1) * time.Second)

	connectionNames := service.GetConnectionNames()
	if len(connectionNames) != 0{
		t.Errorf(`GetConnectionNames after timeout number = %v`, len(connectionNames))
		return
	}

	equips = equipService.GetEquipInfos(false)
	if equips[0].IsActive{
		t.Errorf(`GetEquipInfos after timeout active = %v`, equips[0].IsActive)
		return
	}

	service.UpdateMqttConnections(state)

	time.Sleep(1 * time.Second)
	equips = equipService.GetEquipInfos(false)

	if len(equips) != 1{
		t.Errorf(`GetEquipInfos after timeout number = %v `, len(equips))
		return
	}

	if !equips[0].IsActive{
		t.Errorf(`GetEquipInfos after timeout active = %v`, equips[0].IsActive)
		return
	}
}

