package BL

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"../Models"
)

type WebSocketService struct {
}

// keys - sessionUids
var webSocketConnections = map[string]*WebSock{}

// keys - main equipment topics
// values - slice of session uids
var topicConnections = map[string][]string{}

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

			//find equipment name of a new message
			topicParts := strings.Split(d.Topic, "/")
			activatedEquipInfo := strings.Join([]string{topicParts[0], topicParts[1]}, "/")

			//find all sessions activated this equipment
			if sessionUids, ok := topicConnections[activatedEquipInfo]; ok {
				for _, uid := range sessionUids {

					//find websocket
					v := webSocketConnections[uid]
					b, err := json.Marshal(d)
					if err = v.Conn.WriteMessage(1, b); err != nil {
						// return
					}
				}
			}

		}
	}()

	http.ListenAndServe(":8080", nil)
}

func (*WebSocketService) Activate(sessionUid string, activatedEquipInfo string, deactivatedEquipInfo string) {
	if deactivatedEquipInfo != "" {
	}

	topicConnections[activatedEquipInfo] = append(topicConnections[activatedEquipInfo], sessionUid)

	return
}
