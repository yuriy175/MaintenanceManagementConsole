package interfaces

import (
	"time"
)

// IDiagnosticService describes diagnostic service interface
type IDiagnosticService interface {
	// IncCount increments specified counter
	IncCount(counterName string) 

	// SetDuration sets specified duration
	SetDuration(counterName string, duration time.Duration)
}
