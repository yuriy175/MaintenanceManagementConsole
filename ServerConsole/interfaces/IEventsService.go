package interfaces

import (
	"../models"
)

// IEventsService describes events service interface
type IEventsService interface {
	// Starts the service
	Start()

	// InsertEvent inserts equipment connection state info into db
	InsertConnectEvent(equipName string)*models.EventModel 
}
