package interfaces

import (
	"ServerConsole/models"
)

// IAuthService describes auth service interface
type IAuthService interface {
	// GetSum calculates SHA256 hash
	GetSum(value string) [32]byte

	// CheckSum checks SHA256 hash
	CheckSum(value string, hash [32]byte) bool

	// CreateToken creates jwt token for user
	CreateToken(user *models.UserModel) (string, string)

	// VerifyToken verifies token
	VerifyToken(tokenString string) *models.UserClaims
}
