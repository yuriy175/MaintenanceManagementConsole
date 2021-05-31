package bl

import (
	"encoding/json"
	"strings"
	"sync"

	"../interfaces"
	"../models"
	"../utils"
)

// equipment service implementation type
type equipsService struct {
	// synchronization mutex
	_mtx        sync.RWMutex

	// DAL service
	_dalService interfaces.IDalService
	_equips     map[string]models.EquipInfoModel
	_equipCh    chan *models.RawMqttMessage
}

// EquipsServiceNew creates an instance of equipsService
func EquipsServiceNew(
	dalService interfaces.IDalService,
	equipCh chan *models.RawMqttMessage) interfaces.IEquipsService {
	service := &equipsService{}

	service._dalService = dalService
	service._equipCh = equipCh
	service._equips = map[string]models.EquipInfoModel{}

	return service
}

// Starts the service
func (service *equipsService) Start() {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equipInfos := service._dalService.GetEquipInfos()
	service._equips = make(map[string]models.EquipInfoModel, len(equipInfos))

	for _, equip := range equipInfos {
		service._equips[equip.EquipName] = equip
	}

	go func() {
		for d := range service._equipCh {
			if strings.Contains(d.Topic, "/hospital") {
				viewmodel := models.EquipInfoViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)

				equipName := utils.GetEquipFromTopic(d.Topic)

				service._mtx.Lock()
				if _, ok := service._equips[equipName]; !ok {
					equip := service._dalService.InsertEquipInfo(equipName, &viewmodel)
					service._equips[equip.EquipName] = *equip
				}
				service._mtx.Unlock()
			}
		}
	}()
}

// Checks if the equipment exists
func (service *equipsService) CheckEquipment(equipName string) bool {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equips := service._equips
	if _, ok := equips[equipName]; ok {
		return true
	}

	ok := service._dalService.CheckEquipment(equipName)
	return ok
}

// Inserts a new equipment
func (service *equipsService) InsertEquipInfo(equipName string, equipVM *models.EquipInfoViewModel) *models.EquipInfoModel {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equips := service._equips
	if equip, ok := equips[equipName]; ok {
		return &equip
	}

	equip := service._dalService.InsertEquipInfo(equipName, equipVM)
	equips[equipName] = *equip
	return equip
}

// Returns all equipments
func (service *equipsService) GetEquipInfos() []models.EquipInfoModel {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equips := service._equips
	v := make([]models.EquipInfoModel, 0, len(equips))

	for _, value := range equips {
		v = append(v, value)
	}

	return v
}