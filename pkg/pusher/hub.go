package pusher

import (
	"encoding/json"
	"log"
)

type Service interface {
	RegisterClient(c Clienter)
	UnregisterClient(c Clienter)
	HandleMsg(msg *Message, c Clienter)
}

type Message struct {
	Type string
	Data map[string]interface{}
}

type Hub struct {
	clients          map[Clienter]bool
	register         chan Clienter
	unregister       chan Clienter
	services         map[Service]bool
	msgTypeToService map[string]Service
}

func NewHub() *Hub {
	return &Hub{
		clients:          make(map[Clienter]bool),
		register:         make(chan Clienter),
		unregister:       make(chan Clienter),
		services:         make(map[Service]bool),
		msgTypeToService: make(map[string]Service),
	}
}

func (h *Hub) RegisterService(s Service, msgTypes []string) {
	h.services[s] = true

	for _, msgType := range msgTypes {
		if _, ok := h.msgTypeToService[msgType]; ok {
			panic("msgType has existed")
		}

		h.msgTypeToService[msgType] = s
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			for service := range h.services {
				service.RegisterClient(client)
			}
			log.Printf("client register")
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				//close(client )
			}

			for service := range h.services {
				service.UnregisterClient(client)
			}
			log.Printf("client unregister")
		}
	}
}

func (h *Hub) handleMsg(data []byte, c Clienter) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Printf("can't unmarshal data %v", data)
		return
	}

	log.Printf("handle msg %v", msg)
	ttype := msg.Type
	if _, ok := h.msgTypeToService[ttype]; !ok {
		log.Printf("unknown msgType %s\n", ttype)
		return
	}

	service := h.msgTypeToService[ttype]
	service.HandleMsg(&msg, c)
}
