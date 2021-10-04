package tests

import (
	"testing"

	"ServerConsole/interfaces"
	"ServerConsole/models"
	"ServerConsole/tests/mocks"
)

var equipsServiceTest interfaces.IEquipsService

func setUpEquipsService() interfaces.IEquipsService {

	ioc := mocks.InitMockIoc().(*mocks.MockTypes)
	return ioc.CreateEquipsService()
	/*if equipsServiceTest == nil {
		equipsServiceTest = mocks.InitMockIoc().GetEquipsService() // EquipsServiceNew(nil, mocks.DataLayerMockServiceNew(), nil, nil)
		equipsServiceTest.Start()
	}

	return equipsServiceTest*/
}

func TestInsertEquipInfo(t *testing.T) {
	service := setUpEquipsService()
	equipName := "ABC/WKS_123"
	hospitalName := "HospitalName"
	model := service.InsertEquipInfo(equipName,
		&models.EquipInfoViewModel{HospitalName: "HospitalName",
			HospitalAddress:   "HospitalAddress",
			HospitalLongitude: "12",
			HospitalLatitude:  "34"})

	if !model.IsActive || model.EquipName != equipName || model.HospitalName != hospitalName {
		t.Errorf(`InsertEquipInfo= %v %v %v`, model.IsActive, model.EquipName, model.HospitalName)
	}
}

/*func TestGetEquipInfos(t *testing.T) {
	service := setUp()
}*/

func TestCheckEquipment(t *testing.T) {
	service := setUpEquipsService()
	equipName := "ABC/WKS_123"

	if ok := service.CheckEquipment(equipName); !ok {
		t.Errorf(`TestCheckEquipment= %v`, equipName)
	}
}
