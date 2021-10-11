package bl

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	"ServerConsole/interfaces"
	"ServerConsole/models"
	"ServerConsole/utils"
)

// equipment service implementation type
type equipsService struct {
	// synchronization mutex
	_mtx sync.RWMutex

	//logger
	_log interfaces.ILogger

	// DAL service
	_dalService interfaces.IDalService
	_equips     map[string]models.DetailedEquipInfoModel
	_equipCh    chan *models.RawMqttMessage

	// chanel for communications with events service (internal events)
	_internalEventsCh chan *models.MessageViewModel

	_renamedEquips map[string][]string
}

// EquipsServiceNew creates an instance of equipsService
func EquipsServiceNew(
	log interfaces.ILogger,
	dalService interfaces.IDalService,
	equipCh chan *models.RawMqttMessage,
	internalEventsCh chan *models.MessageViewModel) interfaces.IEquipsService {
	service := &equipsService{}

	service._log = log
	service._dalService = dalService
	service._equipCh = equipCh
	service._internalEventsCh = internalEventsCh
	service._equips = map[string]models.DetailedEquipInfoModel{}

	return service
}

// Starts the service
func (service *equipsService) Start() {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	/*equipInfos := service._dalService.GetEquipInfos()
	service._equips = make(map[string]models.EquipInfoModel, len(equipInfos))

	for _, equip := range equipInfos {
		service._equips[equip.EquipName] = equip
	}*/
	service.initEquipInfos()

	go func() {
		for d := range service._equipCh {
			if strings.Contains(d.Topic, "/hospital") {
				viewmodel := models.EquipInfoViewModel{}
				json.Unmarshal([]byte(d.Data), &viewmodel)

				equipName := utils.GetEquipFromTopic(d.Topic)

				existingEquip, ok := service._equips[equipName]
				if !ok {
					service.checkIfEquipmentRenamed(equipName)
					equip := service._dalService.InsertEquipInfo(equipName, &viewmodel)

					service._mtx.Lock()
					now := time.Now()
					detailedEquip := models.DetailedEquipInfoModel{*equip, true, &now}
					service._equips[equip.EquipName] = detailedEquip

					service._mtx.Unlock()
				} else {
					if existingEquip.HospitalName != viewmodel.HospitalName ||
						existingEquip.HospitalAddress != viewmodel.HospitalAddress ||
						(viewmodel.HospitalLongitude != 0 && existingEquip.HospitalLongitude != viewmodel.HospitalLongitude) ||
						(viewmodel.HospitalLatitude != 0 && existingEquip.HospitalLatitude != viewmodel.HospitalLatitude) {
						go service.updateEquipmentInfo(equipName, viewmodel)
					}
				}
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

	if !ok {
		go service.checkIfEquipmentRenamed(equipName)
	}

	return ok
}

// Inserts a new equipment
func (service *equipsService) InsertEquipInfo(equipName string, equipVM *models.EquipInfoViewModel) *models.DetailedEquipInfoModel {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equips := service._equips
	if equip, ok := equips[equipName]; ok {
		return &equip
	}

	equip := service._dalService.InsertEquipInfo(equipName, equipVM)
	// equips[equipName] = *equip
	now := time.Now()
	detailedEquip := models.DetailedEquipInfoModel{*equip, true, &now}
	equips[equip.EquipName] = detailedEquip

	return &detailedEquip
}

// DisableEquipInfo disables an equipment
func (service *equipsService) DisableEquipInfo(equipName string, disabled bool) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equips := service._equips
	if equip, ok := equips[equipName]; ok {
		equip.Disabled = disabled
		equips[equipName] = equip
	}

	go service._dalService.DisableEquipInfo(equipName, disabled)
}

// GetEquipInfos returns all equipments
func (service *equipsService) GetEquipInfos(withDisabled bool) []models.DetailedEquipInfoModel {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equips := service._equips
	v := make([]models.DetailedEquipInfoModel, 0, len(equips))

	for _, value := range equips {
		if !value.Renamed && (withDisabled || (!withDisabled && !value.Disabled)) {
			v = append(v, value)
		}
	}

	return v
}

// GetOldEquipNames returns out of date equipment names
func (service *equipsService) GetOldEquipNames(equipName string) []string {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	renamedEquips := service._renamedEquips
	hddNumber := utils.GetHddNumberFromEquip(equipName)
	if equip, ok := renamedEquips[hddNumber]; ok {
		return equip
	}

	return []string{}
}

// SetLastSeen sets last seen event time
func (service *equipsService) SetLastSeen(equipName string) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equips := service._equips
	if equip, ok := equips[equipName]; ok {
		now := time.Now()
		equip.LastSeen = &now
		equips[equipName] = equip
	}
}

// SetActivate sets whether equipment is active
func (service *equipsService) SetActivate(equipName string, isOn bool) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equips := service._equips
	if equip, ok := equips[equipName]; ok {
		equip.IsActive = isOn
		equips[equipName] = equip
	}
}

// GetFullInfo returns full equipment permanent info
func (service *equipsService) GetFullInfo(equipName string) *models.FullEquipInfoModel {
	var wg sync.WaitGroup
	model := &models.FullEquipInfoModel{}
	dalService := service._dalService

	wg.Add(4)
	go func() {
		defer wg.Done()
		model.SystemInfo = dalService.GetDBSystemInfo(equipName)
	}()

	go func() {
		defer wg.Done()
		model.SoftwareInfo = *dalService.GetDBSoftwareInfo(equipName)
	}()

	go func() {
		defer wg.Done()
		model.LastSeen = dalService.GetLastSeenInfo(equipName)
	}()

	go func() {
		defer wg.Done()
		if equip, ok := service._equips[equipName]; ok {
			model.LocationInfo = equip
		}
	}()

	wg.Wait()

	return model
}

func (service *equipsService) checkIfEquipmentRenamed(equipName string) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equips := service._equips
	dalService := service._dalService
	hddNumber := utils.GetHddNumberFromEquip(equipName)

	anyRenamed := ""
	for oldEquipName, oldEquip := range equips {
		if !oldEquip.Renamed && hddNumber == utils.GetHddNumberFromEquip(oldEquipName) {
			dalService.RenameEquip(oldEquipName)
			anyRenamed = oldEquip.EquipName
		}
	}

	if anyRenamed != "" {
		msg := models.MessageViewModel{equipName, "переименован из " + anyRenamed, ""}

		service._internalEventsCh <- &msg
		// service.initEquipInfos()
		if equip, ok := equips[anyRenamed]; ok {
			equip.EquipName = equipName
			delete(equips, anyRenamed)
			equips[equipName] = equip
		}

		service.initRenamedEquipInfos()
	}
}

func (service *equipsService) initEquipInfos() {
	equipInfos := service._dalService.GetEquipInfos()
	service._equips = make(map[string]models.DetailedEquipInfoModel, len(equipInfos))

	for _, equip := range equipInfos {
		detailedEquip := models.DetailedEquipInfoModel{equip, false, nil}
		service._equips[equip.EquipName] = detailedEquip
	}

	service.initRenamedEquipInfos()
}

func (service *equipsService) initRenamedEquipInfos() {
	renamedInfos := service._dalService.GetOldEquipInfos()
	service._renamedEquips = make(map[string][]string, len(renamedInfos))

	for _, equip := range renamedInfos {
		service._renamedEquips[equip.HddNumber] = equip.OldEquipNames
	}
}

func (service *equipsService) updateEquipmentInfo(equipName string, viewmodel models.EquipInfoViewModel) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	equips := service._equips
	if equip, ok := equips[equipName]; ok {
		equip.HospitalName = viewmodel.HospitalName
		equip.HospitalAddress = viewmodel.HospitalAddress
		if viewmodel.HospitalLongitude != 0 {
			equip.HospitalLongitude = viewmodel.HospitalLongitude
		}
		if viewmodel.HospitalLatitude != 0 {
			equip.HospitalLatitude = viewmodel.HospitalLatitude
		}
		equips[equipName] = equip

		go service._dalService.UpdateEquipmentInfo(&equip)
	}
}
