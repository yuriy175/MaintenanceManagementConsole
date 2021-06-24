package bl

import (
	"crypto/sha256"
	"github.com/dgrijalva/jwt-go"

	"../interfaces"
	"../models"
)

type userClaims struct {
	Login        string
	Surname 	 string
	Role         string
}

// authorization service implementation type
type authService struct {
	//logger
	_log interfaces.ILogger

	// jwt token secret
	_jwtSecret string
}

// AuthServiceNew creates an instance of authService
func AuthServiceNew(log interfaces.ILogger) interfaces.IAuthService {
	service := &authService{}
	service._log = log
	service._jwtSecret = "qweqrty1975!"

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

// CreateToken creates jwt token for user
func (service *authService) CreateToken(user *models.UserModel) string{
	claims := &userClaims{user.Login, user.Surname, user.Role}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, _ := token.SignedString([]byte(service._jwtSecret))
	return tokenString
}

func (c userClaims) Valid() error {
	return nil
}