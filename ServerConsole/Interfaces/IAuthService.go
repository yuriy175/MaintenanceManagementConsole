package Interfaces

type IAuthService interface {
	GetSum(value string) [32]byte
	CheckSum(value string, hash [32]byte) bool
}
