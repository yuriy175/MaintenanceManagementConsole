package mocks

import (
	"time"
	"ServerConsole/interfaces"
)

// diagnostic mock service implementation type
type diagnosticMockService struct {
}

// DiagnosticMockServiceNew creates an instance of diagnosticMockService
func DiagnosticMockServiceNew() interfaces.IDiagnosticService {
	service := &diagnosticMockService{}
	return service
}

// IncCount increments specified counter
func (service *diagnosticMockService) IncCount(counterName string) {
}

// SetDuration sets specified duration
func (service *diagnosticMockService) SetDuration(counterName string, duration time.Duration) {
}
