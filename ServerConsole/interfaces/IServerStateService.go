package interfaces

import (
	"ServerConsole/models"
)

// IServerStateService describes server state interface
type IServerStateService interface {
	// GetState returns server state
	GetState() *models.ServerState
}
