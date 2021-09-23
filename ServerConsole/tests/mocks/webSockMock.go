package mocks

import (
	"net/http"

	"ServerConsole/interfaces"
)

// web socket client mock implementation type
type webSockMock struct {
	// web socket uid
	_uid string
}

// WebSockMockNew creates an instance of webSockMock
func WebSockMockNew() interfaces.IWebSock {
	webSock := &webSockMock{}

	return webSock
}

// Create initializes an instance of webSockMock
func (sock *webSockMock) Create(w http.ResponseWriter, r *http.Request, uid string) interfaces.IWebSock {
	sock._uid = uid

	return sock
}

// WriteMessage write a message to the web socket
func (sock *webSockMock) WriteMessage(data []byte) error {
	return nil
}

// IsValid checks if the web socket is valid
func (sock *webSockMock) IsValid() bool {
	return true
}

func (sock *webSockMock) onClose() {
}
