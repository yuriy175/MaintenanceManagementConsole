package BL

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../Models"
	"../Utils"
)

type IWebSocketService interface {
	Start()
	Activate(sessionUid string, activatedEquipInfo string, deactivatedEquipInfo string)
	UpdateWebClients(state *Models.EquipConnectionState)
}

type webSocketService struct {
	_ioCProvider IIoCProvider
	_webSockCh   chan *Models.RawMqttMessage
	// keys - sessionUids
	_webSocketConnections map[string]IWebSock

	// keys - main equipment topics
	// values - slice of session uids
	_topicConnections map[string][]string
}

func WebSocketServiceNew(
	ioCProvider IIoCProvider,
	webSockCh chan *Models.RawMqttMessage) IWebSocketService {
	service := &webSocketService{}

	service._ioCProvider = ioCProvider
	service._webSockCh = webSockCh
	service._webSocketConnections = map[string]IWebSock{}
	service._topicConnections = map[string][]string{}

	return service
}

func (service *webSocketService) Start() {
	http.HandleFunc(Models.WebSocketQueryString, func(w http.ResponseWriter, r *http.Request) {
		uids, ok := r.URL.Query()["uid"]

		if !ok || len(uids[0]) < 1 {
			log.Println("Url Param 'uid' is missing")
			return
		}
		uid := uids[0]
		fmt.Printf("created uid: %s \n", uid)

		service._webSocketConnections[uid] = service._ioCProvider.GetWebSocket().Create(w, r, uid)

		/*msgType, msg, err := webSocketConnections[uid].Conn.ReadMessage()
		if err != nil {
			fmt.Printf("sent: %s %d\n", string(msg), msgType)
		}*/
	})

	go func() {
		for d := range service._webSockCh {

			//find equipment name of a new message
			//topicParts := strings.Split(d.Topic, "/")
			activatedEquipInfo := Utils.GetEquipFromTopic(d.Topic) //strings.Join([]string{topicParts[0], topicParts[1]}, "/")

			//find all sessions activated this equipment
			if sessionUids, ok := service._topicConnections[activatedEquipInfo]; ok {
				for _, uid := range sessionUids {

					//find websocket
					log.Println(" message topic %s data %s to web sock %s", d.Topic, d.Data, uid)

					v := service._webSocketConnections[uid]
					b, err := json.Marshal(d)
					if v == nil || v.Conn == nil {
						log.Println(" no connection for  %s", uid)
					} else if err = v.WriteMessage(b); err != nil {
						// return
					}
				}
			}

		}
	}()

	http.ListenAndServe(":8080", nil)
}

func (service *webSocketService) Activate(sessionUid string, activatedEquipInfo string, deactivatedEquipInfo string) {
	topicConnections := service._topicConnections

	if deactivatedEquipInfo != "" {
	}

	topicConnections[activatedEquipInfo] = append(topicConnections[activatedEquipInfo], sessionUid)

	return
}

func (service *webSocketService) UpdateWebClients(state *Models.EquipConnectionState) {
	for _, ws := range service._webSocketConnections {
		b, _ := json.Marshal(state)
		ws.WriteMessage(b)
	}
}
