package controllers

import (
	"encoding/json"
	"net/http"

	interfaces "ServerConsole/interfaces"
)

// ServerController describes server control controller implementation type
type ServerController struct {
	//logger
	_log interfaces.ILogger

	// diagnostic service
	_diagnosticService interfaces.IDiagnosticService

	// http service
	_httpService interfaces.IHttpService

	// server state service
	_serverStateService interfaces.IServerStateService

	// authorization service
	_authService interfaces.IAuthService
}

// ServerControllerNew creates an instance of ServerController
func ServerControllerNew(
	log interfaces.ILogger,
	diagnosticService interfaces.IDiagnosticService,
	httpService interfaces.IHttpService,
	serverStateService interfaces.IServerStateService,
	authService interfaces.IAuthService) *ServerController {
	service := &ServerController{}

	service._log = log
	service._diagnosticService = diagnosticService
	service._httpService = httpService
	service._serverStateService = serverStateService
	service._authService = authService

	return service
}

// Handle handles incomming requests
func (service *ServerController) Handle() {
	serverStateService := service._serverStateService
	authService := service._authService
	http.HandleFunc("/equips/GetServerState", func(w http.ResponseWriter, r *http.Request) {
		claims := CheckUserAuthorization(authService, w, r)

		if claims == nil {
			return
		}

		state := serverStateService.GetState()
		json.NewEncoder(w).Encode(state)
	})
}
