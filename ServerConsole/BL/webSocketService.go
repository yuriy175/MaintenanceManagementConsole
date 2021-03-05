package BL

import (
	"encoding/json"
	"log"
	"net/http"

	"../Models"
)

type WebSocketService struct {
}

var webSocketConnections = map[string]*WebSock{}

func WebServer(equipWebSockCh chan *Models.RawMqttMessage) {
	http.HandleFunc(Models.WebSocketQueryString, func(w http.ResponseWriter, r *http.Request) {
		uids, ok := r.URL.Query()["uid"]

		if !ok || len(uids[0]) < 1 {
			log.Println("Url Param 'uid' is missing")
			return
		}
		uid := uids[0]
		webSocketConnections[uid] = CreateWebSock(w, r, uid)

		/*msgType, msg, err := webSocketConnections[uid].Conn.ReadMessage()
		if err != nil {
			fmt.Printf("sent: %s %d\n", string(msg), msgType)
		}*/
	})

	go func() {
		for d := range equipWebSockCh {
			for _, v := range webSocketConnections {
				// if strings.Contains(d.Topic, "/stand/state") {
				// 	b, err := json.Marshal(d)
				// 	if err = v.Conn.WriteMessage(1, b); err != nil {
				// 		// return
				// 	}
				// }
				b, err := json.Marshal(d)
				if err = v.Conn.WriteMessage(1, b); err != nil {
					// return
				}
				/*if strings.Contains(d.Topic, "ARM/Hardware") {
					b, err := json.Marshal(d)
					if err = v.Conn.WriteMessage(1, b); err != nil {
						// return
					}
				} else if strings.Contains(d.Topic, "/organauto") {
					b, err := json.Marshal(d)
					if err = v.Conn.WriteMessage(1, b); err != nil {
						// return
					}
				} else if strings.Contains(d.Topic, "/generator/state") {
					b, err := json.Marshal(d)
					if err = v.Conn.WriteMessage(1, b); err != nil {
						// return
					}
				} else if strings.Contains(d.Topic, "/detector/state") {
					b, err := json.Marshal(d)
					if err = v.Conn.WriteMessage(1, b); err != nil {
						// return
					}
				}*/
			}
		}
	}()

	http.ListenAndServe(":8080", nil)
}
