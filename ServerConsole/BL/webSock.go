package bl

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"../interfaces"
	"github.com/gorilla/websocket"
)

// web socket client implementation type
type webSock struct {
	_webSocketService interfaces.IWebSocketService
	_uid              string
	_conn             *websocket.Conn
	_mtx              sync.RWMutex
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// WebSockNew creates an instance of webSock
func WebSockNew(webSocketService interfaces.IWebSocketService) interfaces.IWebSock {
	webSock := &webSock{}
	webSock._webSocketService = webSocketService

	return webSock
}

// Create initializes an instance of webSock
func (sock *webSock) Create(w http.ResponseWriter, r *http.Request, uid string) interfaces.IWebSock {
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
			//sock._mtx.Lock()
			if sock._conn == nil {
				//sock._mtx.Unlock()
				return
			}

			sock._conn.SetReadDeadline(time.Now().Add(timeoutDuration))

			// Read message from browser
			_, msg, err := sock._conn.ReadMessage()
			if _, ok := err.(*websocket.CloseError); ok {
				sock._webSocketService.ClientClosed(sock._uid)
				sock.onClose()
				//sock._mtx.Unlock()
				return
			}

			if err != nil {
				fmt.Printf("websocket readmessage error met\n")
				//sock._mtx.Unlock()
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
			//sock._mtx.Unlock()
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
					sock._mtx.Unlock()
					return
				}

				err := sock._conn.WriteMessage(websocket.PingMessage, []byte{})
				if err != nil {
					log.Println("write:", err)
					sock._mtx.Unlock()
					return
				}

				log.Println("ping sent")
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

// WriteMessage write a message to the web socket
func (sock *webSock) WriteMessage(data []byte) error {
	sock._mtx.Lock()
	defer sock._mtx.Unlock()
	return sock._conn.WriteMessage(1, data)
}

// IsValid checks if the web socket is valid
func (sock *webSock) IsValid() bool {
	return sock._conn != nil
}

func (sock *webSock) onClose() {
	sock._conn.Close()
	sock._conn = nil
}
