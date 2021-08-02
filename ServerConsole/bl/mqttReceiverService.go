package bl

import (
	"fmt"
	"sync"
	"time"

	"../interfaces"
	"../models"
	Models "../models"
)

// mqtt receiver service implementation type
type mqttReceiverService struct {
	//logger
	_log interfaces.ILogger

	// synchronization mutex
	_mtx              sync.RWMutex

	// IoC provider
	_ioCProvider      interfaces.IIoCProvider

	// web socket service
	_webSocketService interfaces.IWebSocketService

	// DAL service
	_dalService    interfaces.IDalService

	// equipment service
	_equipsService interfaces.IEquipsService

	// events service
	_eventsService interfaces.IEventsService

	// chanel for DAL communications
	_dalCh         chan *models.RawMqttMessage

	// chanel for communications with websocket services
	_webSockCh     chan *models.RawMqttMessage

	// chanel for communications with events services
	_eventsCh     chan *models.RawMqttMessage

	// mqtt connections map
	// key - topic
	// value - mqtt client
	_mqttConnections map[string]interfaces.IMqttClient
	_topicStorage    interfaces.ITopicStorage

	// topics : server may communicate with a client
	_supportedTopics []string
}

// MqttReceiverServiceNew creates an instance of mqttReceiverService
func MqttReceiverServiceNew(
	log interfaces.ILogger,
	ioCProvider interfaces.IIoCProvider,
	webSocketService interfaces.IWebSocketService,
	dalService interfaces.IDalService,
	equipsService interfaces.IEquipsService,
	eventsService interfaces.IEventsService,
	topicStorage interfaces.ITopicStorage,
	dalCh chan *models.RawMqttMessage,
	webSockCh chan *models.RawMqttMessage,
	eventsCh  chan *models.RawMqttMessage) interfaces.IMqttReceiverService {
	service := &mqttReceiverService{}

	service._log = log
	service._ioCProvider = ioCProvider
	service._webSocketService = webSocketService
	service._dalService = dalService
	service._equipsService = equipsService
	service._eventsService = eventsService
	service._topicStorage = topicStorage
	service._dalCh = dalCh
	service._webSockCh = webSockCh
	service._eventsCh = eventsCh
	service._mqttConnections = map[string]interfaces.IMqttClient{}

	service._supportedTopics = topicStorage.GetTopics()

	go service.startActiveConnectionsCheck()

	return service
}

//UpdateMqttConnections updates mqtt connections map for an equipment connection state
func (service *mqttReceiverService) UpdateMqttConnections(state *models.EquipConnectionState) {
	rootTopic := state.Name
	isOff := !state.Connected
	topics := service._supportedTopics
	mqttConnections := service._mqttConnections
	ioCProvider := service._ioCProvider
	// dalService := service._dalService
	equipsService := service._equipsService
	eventsService := service._eventsService

	fmt.Printf("UpdateMqttConnections from topic: %s\n", rootTopic)

	service._mtx.Lock()
	defer service._mtx.Unlock()

	fmt.Printf("UpdateMqttConnections unlocked")

	if client, ok := mqttConnections[rootTopic]; ok {
		fmt.Println(rootTopic + " already exists")
		if isOff {
			go client.Disconnect()
			delete(mqttConnections, rootTopic)
			fmt.Println(rootTopic + " deleted")
		}

		// if the topic is observed by any client -> send activate command
		if service._webSocketService.HasActiveClients(rootTopic) {
			go service.SendCommand(rootTopic, "activate")
		}

		return
	}

	if !isOff {
		mqttConnections[rootTopic] = ioCProvider.GetMqttClient().Create(rootTopic, topics)
		go equipsService.CheckEquipment(rootTopic) 
		go service.SendCommand(rootTopic, "serverReady")
		go eventsService.InsertConnectEvent(rootTopic)
	}

	fmt.Println(rootTopic + " created")
}

// CreateCommonConnections reates common mqtt connections
func (service *mqttReceiverService) CreateCommonConnections() {
	mqttConnections := service._mqttConnections
	ioCProvider := service._ioCProvider

	service._mtx.Lock()
	defer service._mtx.Unlock()
	mqttConnections[Models.CommonTopicPath] = ioCProvider.GetMqttClient().Create(models.CommonTopicPath, []string{})
	mqttConnections[Models.BroadcastCommandsTopic] = ioCProvider.GetMqttClient().Create(models.BroadcastCommandsTopic, []string{})
	mqttConnections[Models.CommonChatsPath] = ioCProvider.GetMqttClient().Create(models.CommonChatsPath, []string{})

	return
}

// SendCommand sends a command to equipment via mqtt
func (service *mqttReceiverService) SendCommand(equipment string, command string) {
	fmt.Printf("SendCommand from topic: %s %s\n", equipment, command)

	service._mtx.Lock()
	defer service._mtx.Unlock()

	if client, ok := service._mqttConnections[equipment]; ok {
		go client.SendCommand(command)
	}

	return
}

// PublishChatNote sends a chat note to equipment via mqtt
func (service *mqttReceiverService) PublishChatNote(equipment string, message string, user string, isInternal bool) {	

	service._mtx.Lock()
	defer service._mtx.Unlock()

	mqttConnections := service._mqttConnections
	if client, ok := mqttConnections[models.CommonChatsPath]; ok {
		go client.SendChatMessage(equipment, user, message, isInternal)
	}

	/*// we may have no connection to this client
	topics := service._supportedTopics
	ioCProvider := service._ioCProvider

	if _, ok := mqttConnections[equipment]; !ok {
		mqttConnections[equipment] = ioCProvider.GetMqttClient().Create(equipment, topics)
	}	

	if client, ok := mqttConnections[equipment]; ok {
		go client.SendChatMessage(user, message)
	}*/

	return
}

// SendCommand sends a broadcast command to equipments via mqtt
func (service *mqttReceiverService) SendBroadcastCommand(command string) {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	if client, ok := service._mqttConnections[models.BroadcastCommandsTopic]; ok {
		go client.SendCommand(command)
	}

	return
}

// GetConnectionNames returns connected equipment names
func (service *mqttReceiverService) GetConnectionNames() []string {
	service._mtx.Lock()
	defer service._mtx.Unlock()

	mqttConnections := service._mqttConnections

	keys := make([]string, len(mqttConnections))

	i := 0
	for k, d := range mqttConnections {
		if d.IsEquipTopic() {
			keys[i] = k
			i++
		}
	}

	return keys
}

// Activate activates a specified connection to equipment and deactivates the other
func (service *mqttReceiverService) Activate(activatedEquipInfo string, deactivatedEquipInfo string) {

	if deactivatedEquipInfo != "" && deactivatedEquipInfo != activatedEquipInfo &&
		!service._webSocketService.HasActiveClients(deactivatedEquipInfo){
		service.SendCommand(deactivatedEquipInfo, "deactivate")
	}
	
	service.SendCommand(activatedEquipInfo, "activate")

	return
}

func (service *mqttReceiverService) startActiveConnectionsCheck(){
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
				case _ = <-ticker.C:
					service._mtx.Lock()				

					mqttConnections := service._mqttConnections
					checkTime := time.Now().Add(time.Duration(-models.KeepAliveCheckPeriod) * time.Second)

					for t, c := range mqttConnections {
						if(!c.IsEquipTopic()){
							continue
						}

						lastTime := c.GetLastAliveMessage()
						if lastTime.Before(checkTime){
							state := &models.EquipConnectionState{t, false}
							go service.UpdateMqttConnections(state)
							go service._webSocketService.UpdateWebClients(state)
						}
					}

					service._mtx.Unlock()

					fmt.Printf("Active connectins: %v\n", len(mqttConnections))
			}
		}
}