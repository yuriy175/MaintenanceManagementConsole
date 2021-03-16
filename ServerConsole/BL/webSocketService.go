package BL

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../Models"
	"../Utils"
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
		fmt.Printf("created uid: %s \n", uid)

		webSocketConnections[uid] = CreateWebSock(w, r, uid)

		/*msgType, msg, err := webSocketConnections[uid].Conn.ReadMessage()
		if err != nil {
			fmt.Printf("sent: %s %d\n", string(msg), msgType)
		}*/
	})

	go func() {
		for d := range equipWebSockCh {

			//find equipment name of a new message
			//topicParts := strings.Split(d.Topic, "/")
			activatedEquipInfo := Utils.GetEquipFromTopic(d.Topic) //strings.Join([]string{topicParts[0], topicParts[1]}, "/")

			//find all sessions activated this equipment
			if sessionUids, ok := topicConnections[activatedEquipInfo]; ok {
				for _, uid := range sessionUids {

					//find websocket
					log.Println(" message topic %s data %s to web sock %s", d.Topic, d.Data, uid)

					v := webSocketConnections[uid]
					b, err := json.Marshal(d)
					if v == nil || v.Conn == nil {
						log.Println(" no connection for  %s", uid)
					} else if err = v.Conn.WriteMessage(1, b); err != nil {
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
