package BL

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"../Interfaces"
	"github.com/gorilla/websocket"
)

func WebSockNew(webSocketService Interfaces.IWebSocketService) Interfaces.IWebSock {
	webSock := &webSock{}
	webSock._webSocketService = webSocketService

	return webSock
}

type webSock struct {
	_webSocketService Interfaces.IWebSocketService
	_uid              string
	_conn             *websocket.Conn
	_mtx              sync.RWMutex
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (sock *webSock) Create(w http.ResponseWriter, r *http.Request, uid string) Interfaces.IWebSock {
	//timeoutDuration := 60 * time.Second
	timeoutDuration := 2 * time.Minute

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panicf("create web socket: %s", err)
	}

	sock._uid = uid
	sock._conn = conn
	log.Println("web sock created Url %s", uid)

	sock._conn.SetPongHandler(func(str string) error {
		log.Println("pong received", str)
		return nil
	})

	go func() {
		for {
			sock._conn.SetReadDeadline(time.Now().Add(timeoutDuration))

			// Read message from browser
			_, msg, err := sock._conn.ReadMessage()
			if _, ok := err.(*websocket.CloseError); ok {
				sock._webSocketService.ClientClosed(sock._uid)
				sock.onClose()
				return
			}

			if err != nil {
				fmt.Printf("websocket readmessage error met\n")
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
		}
	}()

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case _ = <-ticker.C:
				sock._mtx.Lock()
				if sock._conn == nil {
					return
				}

				err := sock._conn.WriteMessage(websocket.PingMessage, []byte{})
				if err != nil {
					log.Println("write:", err)
					return
				} else {
					log.Println("ping sent")
				}
				sock._mtx.Unlock()
				// case <-interrupt:
				// 	log.Println("interrupt")
				// 	// To cleanly close a connection, a client should send a close
				// 	// frame and wait for the server to close the connection.
				// 	err := c.WriteMessage(
				// 		websocket.CloseMessage,
				// 		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				// 	if err != nil {
				// 		log.Println("write close:", err)
				// 		return
				// 	}
				// 	select {
				// 	case <-done:
				// 	case <-time.After(time.Second):
				// 	}
				// 	c.Close()
				// 	return
			}
		}
	}()

	return sock
}

func (ws *webSock) WriteMessage(data []byte) error {
	ws._mtx.Lock()
	defer ws._mtx.Unlock()
	return ws._conn.WriteMessage(1, data)
}

func (ws *webSock) IsValid() bool {
	return ws._conn != nil
}

func (ws *webSock) onClose() {
	ws._conn.Close()
	ws._conn = nil
}
