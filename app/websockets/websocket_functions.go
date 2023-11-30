package websockets

import (
	"github.com/gorilla/websocket"
	"github.com/hyahm/golog"
)

type WebSocketFunctions struct{}

func NewWebSocketFunctions() *WebSocketFunctions {
	return &WebSocketFunctions{}
}

func (functions *WebSocketFunctions) HandleWebSocketMessage(messageType int, message []byte) {
	switch messageType {
	case websocket.TextMessage:
		golog.Info("Client send TextMessage: %s", string(message))
	default:
	}
}
