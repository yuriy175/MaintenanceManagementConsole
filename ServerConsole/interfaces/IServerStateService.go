package interfaces
import (
	"../models"
)

// IServerStateService describes server state interface
type IServerStateService interface {
    // GetState returns server state
	GetState() *models.ServerState
}
