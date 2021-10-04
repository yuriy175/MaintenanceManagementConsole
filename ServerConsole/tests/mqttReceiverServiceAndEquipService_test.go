package tests

import (
	"testing"
	"time"

	"ServerConsole/models"
	"ServerConsole/tests/mocks"
)

func TestUpdateMqttConnections_1Connected_TimedOut_Reconnected_Active(t *testing.T) {
	service := setUpMqttReceiverService()
	equipService := setUpEquipsService()
	equipsChan := mocks.InitMockIoc().GetEquipsChan()

	initEquipsNumber := len(equipService.GetEquipInfos(false))

	equipName := "ABC/WKS_ABC"
	state := &models.EquipConnectionState{
		Name:      equipName,
		Connected: true,
	}
	service.UpdateMqttConnections(state)
	service.SetKeepAliveReceived(equipName)

	msg := &models.RawMqttMessage{
		Topic:   equipName + "/hospital",
		Data:    "",
		Arrival: time.Now(),
	}

	equipsChan <- msg
	time.Sleep(1 * time.Second)

	equips := equipService.GetEquipInfos(false)
	equipsNumber := initEquipsNumber + 1
	activeEquip := equips[initEquipsNumber]
	if len(equips) != equipsNumber {
		t.Errorf(`GetEquipInfos before timeout number = %v `, equipsNumber)
		return
	}

	if !activeEquip.IsActive {
		t.Errorf(`GetEquipInfos before timeout active = %v`, activeEquip.IsActive)
		return
	}

	time.Sleep((models.KeepAliveCheckPeriod*2 + 1) * time.Second)

	connectionNames := service.GetConnectionNames()
	if len(connectionNames) != 0 {
		t.Errorf(`GetConnectionNames after timeout number = %v`, len(connectionNames))
		return
	}

	equips = equipService.GetEquipInfos(false)
	activeEquip = equips[initEquipsNumber]
	if activeEquip.IsActive {
		t.Errorf(`GetEquipInfos after timeout active = %v`, activeEquip.IsActive)
		return
	}

	service.UpdateMqttConnections(state)

	time.Sleep(1 * time.Second)
	equips = equipService.GetEquipInfos(false)
	activeEquip = equips[initEquipsNumber]
	if len(equips) != equipsNumber {
		t.Errorf(`GetEquipInfos after timeout number = %v `, equipsNumber)
		return
	}

	if !activeEquip.IsActive {
		t.Errorf(`GetEquipInfos after second timeout active = %v`, activeEquip.IsActive)
		return
	}
}
