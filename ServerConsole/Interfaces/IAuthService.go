package interfaces

// IAuthService describes auth service interface
type IAuthService interface {
   // GetSum calculates SHA256 hash
	GetSum(value string) [32]byte

	// CheckSum checks SHA256 hash
	CheckSum(value string, hash [32]byte) bool
}
