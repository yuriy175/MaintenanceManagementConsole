package models

import (
	"errors"
)

// UserClaims describes claim for authentification
type UserClaims struct {
	Login        string
	Surname 	 string
	Role         string
}

// Valid checks if claim is valid
func (c *UserClaims) Valid() error {
	if c.Login != "" && c.Role != ""{
		return nil
	}

	return  errors.New("wrong UserClaims")
}
