package Interfaces

import (
	"net/http"
)

type IWebSock interface {
	Create(w http.ResponseWriter, r *http.Request, uid string) IWebSock
	WriteMessage(data []byte) error
	IsValid() bool
}
