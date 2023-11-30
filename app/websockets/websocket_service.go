package websockets

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hyahm/golog"
	"sync"
)

type WebSocketService struct {
	clients   map[*WebSocketClient]string
	functions *WebSocketFunctions
	mutex     sync.Mutex
}

func NewWebSocketService() *WebSocketService {
	return &WebSocketService{
		clients:   make(map[*WebSocketClient]string),
		functions: NewWebSocketFunctions(),
	}
}

func (service *WebSocketService) AddClient(client *WebSocketClient, userId string) {
	service.RemoveClientByUserId(userId)
	service.mutex.Lock()
	service.clients[client] = userId
	service.mutex.Unlock()
	go client.Start(
		service.functions.HandleWebSocketMessage,
		func(err error) {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				golog.Info("remove client using userId [%s]", userId)
				client.conn.Close()
				service.RemoveClient(client)
			}
		})
}

func (service *WebSocketService) RemoveClient(client *WebSocketClient) {
	service.mutex.Lock()
	delete(service.clients, client)
	service.mutex.Unlock()
}

func (service *WebSocketService) RemoveClientByUserId(userId string) {
	service.mutex.Lock()
	for client := range service.clients {
		if service.clients[client] == userId {
			golog.Info("delete client connection %s", userId)
			client.Close()
			delete(service.clients, client)
		}
	}
	service.mutex.Unlock()
}

func (service *WebSocketService) BroadCase(message []byte) {
	service.mutex.Lock()
	for client := range service.clients {
		client.SendMessage(message)
	}
	service.mutex.Unlock()
}

func (service *WebSocketService) SendMessageToUser(userId string, message []byte) error {
	service.mutex.Lock()
	for client := range service.clients {
		if service.clients[client] == userId {
			service.mutex.Unlock()
			return client.SendMessage(message)
		}
	}
	service.mutex.Unlock()
	return fmt.Errorf("user not found")
}
