package interfaces

import (
	"net/http"
)

// IWebSock describes web socket interface
type IWebSock interface {
	// Create initializes an instance of webSock
	Create(w http.ResponseWriter, r *http.Request, uid string) IWebSock

	// WriteMessage write a message to the web socket
	WriteMessage(data []byte) error

	// IsValid checks if the web socket is valid
	IsValid() bool
}
