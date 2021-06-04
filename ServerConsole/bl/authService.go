package bl

import (
	"crypto/sha256"

	"../interfaces"
)

// authorization service implementation type
type authService struct {
	//logger
	_log interfaces.ILogger
}

// AuthServiceNew creates an instance of authService
func AuthServiceNew(log interfaces.ILogger) interfaces.IAuthService {
	service := &authService{}
	service._log = log

	return service
}

// GetSum calculates SHA256 hash
func (service *authService) GetSum(value string) [32]byte {
	sum := sha256.Sum256([]byte(value))

	return sum
}

// CheckSum checks SHA256 hash
func (service *authService) CheckSum(value string, hash [32]byte) bool {
	sum := sha256.Sum256([]byte(value))
	ok := hash == sum

	return ok
}
