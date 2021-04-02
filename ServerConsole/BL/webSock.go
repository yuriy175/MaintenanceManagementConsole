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

func WebSockNew() IWebSock {
	webSock := &webSock{}
	return webSock
}

type webSock struct {
	Uid  string
	Conn *websocket.Conn
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

	sock.Uid = uid
	sock.Conn = conn
	log.Println("web sock created Url %s", uid)

	// return sock

	/**/
	go func() {
		for {
			// Read message from browser
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
	
			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))	
		}
	}()

	return sock
}

func (ws *webSock) WriteMessage(data []byte) error {
	return ws.Conn.WriteMessage(1, data)
}

func (ws *webSock) IsValid() bool {
	return ws.Conn != nil
}
