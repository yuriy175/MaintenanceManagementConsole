package BL

import (
	"encoding/json"
	"strings"
	"sync"

	"../Interfaces"
	"../Models"
	"../Utils"
)

type equipsService struct {
	_mtx        sync.RWMutex
	_dalService Interfaces.IDalService
	_equips     map[string]Models.EquipInfoModel
	_equipCh    chan *Models.RawMqttMessage
}

func EquipsServiceNew(
	dalService Interfaces.IDalService,
	equipCh chan *Models.RawMqttMessage) Interfaces.IEquipsService {
	service := &equipsService{}

	service._dalService = dalService
	service._equipCh = equipCh
	service._equips = map[string]Models.EquipInfoModel{}

	return service
}

func (service *equipsService) Start() {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equipInfos := service._dalService.GetEquipInfos()
	service._equips = make(map[string]Models.EquipInfoModel, len(equipInfos))

	for _, equip := range equipInfos {
		service._equips[equip.EquipName] = equip
	}

	go func() {
		for d := range service._equipCh {
			if strings.Contains(d.Topic, "/hospital") {
				viewmodel := Models.EquipInfoViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)

				equipName := Utils.GetEquipFromTopic(d.Topic)

				service._mtx.Lock()
				if _, ok := service._equips[equipName]; !ok {
					service._dalService.InsertEquipInfo(equipName, &viewmodel)
				}
				service._mtx.Unlock()
			}
		}
	}()
}

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

func (service *equipsService) InsertEquipInfo(equipName string, equipVM *Models.EquipInfoViewModel) *Models.EquipInfoModel {
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

func (service *equipsService) GetEquipInfos() []Models.EquipInfoModel {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equips := service._equips
	v := make([]Models.EquipInfoModel, 0, len(equips))

	for _, value := range equips {
		v = append(v, value)
	}

	return v
}
