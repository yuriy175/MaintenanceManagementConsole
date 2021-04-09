package BL

import (
	"crypto/sha256"

	"../Interfaces"
)

type authService struct {
}

func AuthServiceNew() Interfaces.IAuthService {
	service := &authService{}

	return service
}

func (service *authService) GetSum(value string) [32]byte {
	sum := sha256.Sum256([]byte(value))

	return sum
}

func (service *authService) CheckSum(value string, hash [32]byte) bool {
	sum := sha256.Sum256([]byte(value))
	ok := hash == sum

	return ok
}
