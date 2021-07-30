package bl

import (
	"../interfaces"
	"../models"
)

// server state service implementation type
type serverStateService struct {
	//logger
	_log interfaces.ILogger
	
	// DAL service
	_dalService interfaces.IDalService
}

// ServerStateServiceNew creates an instance of serverStateService
func ServerStateServiceNew(
	log interfaces.ILogger,
	dalService interfaces.IDalService) interfaces.IServerStateService {
	service := &serverStateService{}

	service._log = log
	service._dalService = dalService
	return service
}

// GetState returns server state
func (service *serverStateService) GetState() *models.ServerState{
	state := &models.ServerState{}

	dbState := service._dalService.GetState()
	state.DBUsedSize = dbState["totalSize"].(float64)
	state.DiskTotalSpace = dbState["fsTotalSize"].(float64)
	state.DiskUsedSpace = dbState["fsUsedSize"].(float64)
	
	return state
}
