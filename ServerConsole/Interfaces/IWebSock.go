package interfaces

import (
	"net/http"
)

// IWebSock describes web socket interface
type IWebSock interface {
	Create(w http.ResponseWriter, r *http.Request, uid string) IWebSock
	WriteMessage(data []byte) error
	IsValid() bool
}
