package BL

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type IWebSock interface {
	Create(w http.ResponseWriter, r *http.Request, uid string) IWebSock
	WriteMessage(data []byte) error
	IsValid() bool
}

func WebSockNew(webSocketService IWebSocketService) IWebSock {
	webSock := &webSock{}
	webSock._webSocketService = webSocketService

	return webSock
}

type webSock struct {
	_webSocketService IWebSocketService
	_uid              string
	_conn             *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (sock *webSock) Create(w http.ResponseWriter, r *http.Request, uid string) IWebSock {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panicf("create web socket: %s", err)
	}

	sock._uid = uid
	sock._conn = conn
	log.Println("web sock created Url %s", uid)

	// return sock

	/**/
	go func() {
		for {
			// Read message from browser
			_, msg, err := sock._conn.ReadMessage()
			if _, ok := err.(*websocket.CloseError); ok {
				sock._webSocketService.ClientClosed(sock._uid)
				sock.onClose()
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
		}
	}()

	return sock
}

func (ws *webSock) WriteMessage(data []byte) error {
	return ws._conn.WriteMessage(1, data)
}

func (ws *webSock) IsValid() bool {
	return ws._conn != nil
}

func (ws *webSock) onClose() {
	ws._conn.Close()
	ws._conn = nil
}
