package interfaces

import (
	"time"

	"../models"
)

// IEventsService describes events service interface
type IEventsService interface {
	// Starts the service
	Start()

	// InsertEvent inserts equipment connection state info into db
	InsertConnectEvent(equipName string)
	
	// GetEvents returns all events from db
	GetEvents(equipName string, startDate time.Time, endDate time.Time) []models.EventModel
}
