package bl

import (
	"crypto/sha256"

	"github.com/dgrijalva/jwt-go"

	"ServerConsole/interfaces"
	"ServerConsole/models"
)

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
func (service *authService) CreateToken(user *models.UserModel) (string, string) {
	claims := &models.UserClaims{user.Login, user.Role}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, _ := token.SignedString([]byte(service._jwtSecret))
	return tokenString, user.Surname
}

func (service *authService) VerifyToken(tokenString string) *models.UserClaims {
	claims := &models.UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(service._jwtSecret), nil
	})

	if token == nil || err != nil {
		return nil
	}

	return claims
}
