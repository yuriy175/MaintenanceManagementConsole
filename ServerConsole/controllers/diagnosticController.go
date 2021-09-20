package controllers

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"ServerConsole/interfaces"
)

// DiagnosticController describes diagnostic controller implementation type
type DiagnosticController struct {
	//logger
	_log interfaces.ILogger

	// authorization service
	_authService interfaces.IAuthService

	// metrics handler
	_handler http.Handler
}

// DiagnosticController creates an instance of DiagnosticController
func DiagnosticControllerNew(
	log interfaces.ILogger,
	authService interfaces.IAuthService) *DiagnosticController {
	service := &DiagnosticController{}

	service._log = log
	service._authService = authService
	service._handler = promhttp.Handler()

	return service
}

// Handle handles incomming requests
/*func (service *DiagnosticController) Handle() {
	authService := service._authService

	handler := promhttp.Handler()
	http.HandleFunc("/equips/metrics", func(w http.ResponseWriter, r *http.Request) {
		if CheckAdminAuthorization(authService, w, r) != nil {
			handler.ServeHTTP(w, r)
		}
	})
}*/

// GetServerMetrics returns server performance metrics
func (service *DiagnosticController) GetServerMetrics(w http.ResponseWriter, r *http.Request) {
	authService := service._authService
	if CheckAdminAuthorization(authService, w, r) != nil {
		service._handler.ServeHTTP(w, r)
	}
}
