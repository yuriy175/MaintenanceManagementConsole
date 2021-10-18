package tests

import (
	"testing"
	"time"

	"ServerConsole/interfaces"
	"ServerConsole/models"
	"ServerConsole/tests/mocks"
)

func setUpMqttReceiverServiceAndEquipsService() (
	interfaces.IMqttReceiverService,
	interfaces.IEquipsService,
	chan *models.RawMqttMessage) {

	ioc := mocks.InitMockIoc().(*mocks.MockTypes)
	service, equipService, equipsChan := ioc.CreateMqttReceiverServiceAndEquipsService()
	equipService.Start()
	return service, equipService, equipsChan
	/*if mqttReceiverServiceTest == nil {
		mqttReceiverServiceTest = mocks.InitMockIoc().GetMqttReceiverService() // EquipsServiceNew(nil, mocks.DataLayerMockServiceNew(), nil, nil)
	}

	return mqttReceiverServiceTest*/
}

func TestUpdateMqttConnections_2Connected_1Rename(t *testing.T) {
	service, equipService, equipsChan := setUpMqttReceiverServiceAndEquipsService()

	equipName1 := "ABC/WKS_ABC1"
	state := &models.EquipConnectionState{
		Name:      equipName1,
		Connected: true,
	}
	service.UpdateMqttConnections(state)
	service.SetKeepAliveReceived(equipName1)

	msg := &models.RawMqttMessage{
		Topic:   equipName1 + "/hospital",
		Data:    "",
		Arrival: time.Now(),
	}

	equipsChan <- msg

	equipName2 := "ABC/WKS_ABC2"
	state = &models.EquipConnectionState{
		Name:      equipName2,
		Connected: true,
	}
	service.UpdateMqttConnections(state)
	service.SetKeepAliveReceived(equipName2)

	msg = &models.RawMqttMessage{
		Topic:   equipName2 + "/hospital",
		Data:    "",
		Arrival: time.Now(),
	}

	equipsChan <- msg
	time.Sleep(1 * time.Second)

	equips := equipService.GetEquipInfos(false)
	if len(equips) != 2 {
		t.Errorf(`GetEquipInfos before rename = %v `, len(equips))
		return
	}

	connectionNames := service.GetConnectionNames()
	if len(connectionNames) != 2 {
		t.Errorf(`GetConnectionNames before rename = %v`, len(connectionNames))
		return
	}

	equipName22 := "ABC/WKS2_ABC2"
	msg = &models.RawMqttMessage{
		Topic:   equipName22 + "/hospital",
		Data:    "",
		Arrival: time.Now(),
	}

	equipsChan <- msg
	time.Sleep(1 * time.Second)

	equips = equipService.GetEquipInfos(false)
	if len(equips) != 2 {
		t.Errorf(`GetEquipInfos after rename = %v `, len(equips))
		return
	}

	connectionNames = service.GetConnectionNames()
	if len(connectionNames) != 2 {
		t.Errorf(`GetConnectionNames after rename = %v`, len(connectionNames))
		return
	}
}

func TestUpdateMqttConnections_1Connected_TimedOut_Reconnected_Active(t *testing.T) {
	service, equipService, equipsChan := setUpMqttReceiverServiceAndEquipsService()

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

func TestUpdateMqttConnections_2Connected_1Rename_1BackRenamed(t *testing.T) {
	service, equipService, equipsChan := setUpMqttReceiverServiceAndEquipsService()

	equipName1 := "ABC/WKS_ABC1"
	state := &models.EquipConnectionState{
		Name:      equipName1,
		Connected: true,
	}
	service.UpdateMqttConnections(state)
	service.SetKeepAliveReceived(equipName1)

	msg := &models.RawMqttMessage{
		Topic:   equipName1 + "/hospital",
		Data:    "",
		Arrival: time.Now(),
	}

	equipsChan <- msg

	equipName2 := "ABC/WKS_ABC2"
	state = &models.EquipConnectionState{
		Name:      equipName2,
		Connected: true,
	}
	service.UpdateMqttConnections(state)
	service.SetKeepAliveReceived(equipName2)

	msg = &models.RawMqttMessage{
		Topic:   equipName2 + "/hospital",
		Data:    "",
		Arrival: time.Now(),
	}

	equipsChan <- msg
	time.Sleep(1 * time.Second)

	equips := equipService.GetEquipInfos(false)
	if len(equips) != 2 {
		t.Errorf(`GetEquipInfos before rename = %v `, len(equips))
		return
	}

	connectionNames := service.GetConnectionNames()
	if len(connectionNames) != 2 {
		t.Errorf(`GetConnectionNames before rename = %v`, len(connectionNames))
		return
	}

	equipName22 := "ABC/WKS2_ABC2"
	msg = &models.RawMqttMessage{
		Topic:   equipName22 + "/hospital",
		Data:    "",
		Arrival: time.Now(),
	}

	equipsChan <- msg
	time.Sleep(1 * time.Second)

	service.SetKeepAliveReceived(equipName1)
	// service.SetKeepAliveReceived(equipName22)

	equips = equipService.GetEquipInfos(false)
	if len(equips) != 2 {
		t.Errorf(`GetEquipInfos after rename = %v `, len(equips))
		return
	}

	connectionNames = service.GetConnectionNames()
	if len(connectionNames) != 2 {
		t.Errorf(`GetConnectionNames after rename = %v`, len(connectionNames))
		return
	}

	equipNameB2 := "ABC/WKS_ABC2"
	msg = &models.RawMqttMessage{
		Topic:   equipNameB2 + "/hospital",
		Data:    "",
		Arrival: time.Now(),
	}

	service.SetKeepAliveReceived(equipName1)
	service.SetKeepAliveReceived(equipNameB2)
	equipsChan <- msg
	time.Sleep(1 * time.Second)

	equips = equipService.GetEquipInfos(false)
	if len(equips) != 2 {
		t.Errorf(`GetEquipInfos after back rename = %v `, len(equips))
		return
	}

	connectionNames = service.GetConnectionNames()
	if len(connectionNames) != 2 {
		t.Errorf(`GetConnectionNames after back rename = %v`, len(connectionNames))
		return
	}
}

func TestUpdateMqttConnections_1Connected_Off_Rename_Off_BackRenamed(t *testing.T) {
	service, equipService, equipsChan := setUpMqttReceiverServiceAndEquipsService()

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
	if len(equips) != 1 {
		t.Errorf(`GetEquipInfos before rename = %v `, len(equips))
		return
	}

	connectionNames := service.GetConnectionNames()
	if len(connectionNames) != 1 {
		t.Errorf(`GetConnectionNames before rename = %v`, len(connectionNames))
		return
	}

	// time.Sleep((models.KeepAliveCheckPeriod*2 + 1) * time.Second)
	state = &models.EquipConnectionState{
		Name:      equipName,
		Connected: false,
	}
	service.UpdateMqttConnections(state)
	time.Sleep(1 * time.Second)

	equips = equipService.GetEquipInfos(false)
	activeEquip := equips[0]
	if activeEquip.IsActive {
		t.Errorf(`GetEquipInfos before rename  after timeout active = %v`, activeEquip.IsActive)
		return
	}

	connectionNames = service.GetConnectionNames()
	if len(connectionNames) != 0 {
		t.Errorf(`GetConnectionNames before rename after off = %v`, len(connectionNames))
		return
	}

	equipName = "ABC/WKS2_ABC"
	state = &models.EquipConnectionState{
		Name:      equipName,
		Connected: true,
	}
	service.UpdateMqttConnections(state)
	service.SetKeepAliveReceived(equipName)

	msg = &models.RawMqttMessage{
		Topic:   equipName + "/hospital",
		Data:    "",
		Arrival: time.Now(),
	}

	equipsChan <- msg
	time.Sleep(1 * time.Second)

	equips = equipService.GetEquipInfos(false)
	if len(equips) != 1 {
		t.Errorf(`GetEquipInfos after rename = %v `, len(equips))
		return
	}

	connectionNames = service.GetConnectionNames()
	if len(connectionNames) != 1 {
		t.Errorf(`GetConnectionNames after rename = %v`, len(connectionNames))
		return
	}

	// time.Sleep((models.KeepAliveCheckPeriod*2 + 1) * time.Second)
	state = &models.EquipConnectionState{
		Name:      equipName,
		Connected: false,
	}
	service.UpdateMqttConnections(state)
	time.Sleep(1 * time.Second)

	equips = equipService.GetEquipInfos(false)
	activeEquip = equips[0]
	if activeEquip.IsActive {
		t.Errorf(`GetEquipInfos after rename after timeout active = %v`, activeEquip.IsActive)
		return
	}

	connectionNames = service.GetConnectionNames()
	if len(connectionNames) != 0 {
		t.Errorf(`GetConnectionNames after rename after off = %v`, len(connectionNames))
		return
	}

	equipName = "ABC/WKS_ABC"
	state = &models.EquipConnectionState{
		Name:      equipName,
		Connected: true,
	}
	service.UpdateMqttConnections(state)
	service.SetKeepAliveReceived(equipName)

	msg = &models.RawMqttMessage{
		Topic:   equipName + "/hospital",
		Data:    "",
		Arrival: time.Now(),
	}

	equipsChan <- msg
	time.Sleep(1 * time.Second)

	equips = equipService.GetEquipInfos(false)
	activeEquip = equips[0]
	if !activeEquip.IsActive || activeEquip.EquipName != equipName {
		t.Errorf(`GetEquipInfos after back rename after timeout active = %v name = %v`,
			activeEquip.IsActive, activeEquip.EquipName)
		return
	}

	connectionNames = service.GetConnectionNames()
	if len(connectionNames) != 1 {
		t.Errorf(`GetConnectionNames after back rename = %v`, len(connectionNames))
		return
	}
}
