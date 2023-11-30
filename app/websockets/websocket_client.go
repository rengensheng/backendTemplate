package websockets

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hyahm/golog"
)

type WebSocketClient struct {
	conn *websocket.Conn
	send chan []byte
}

func NewWebSocketClient(conn *websocket.Conn) *WebSocketClient {
	return &WebSocketClient{
		conn: conn,
		send: make(chan []byte, 256),
	}
}

func (client *WebSocketClient) Start(handle func(messageType int, message []byte), errHandle func(err error)) {
	go client.Listen(handle, errHandle)
	go client.Write()
}

func (client *WebSocketClient) Close() {
	client.conn.Close()
	close(client.send)
}

func (client *WebSocketClient) Write() {
	defer func() {
		client.conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.send:
			if !ok {
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := client.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Println("Error writing to WebSocket:", err.Error())
				return
			}
		}
	}
}

func (client *WebSocketClient) Listen(handle func(messageType int, message []byte), errHandle func(err error)) {
	defer func() {
		client.conn.Close()
	}()
	for {
		messageType, message, err := client.conn.ReadMessage()
		if err != nil {
			errHandle(err)
			fmt.Println("Error reading from WebSocket:", err.Error(), websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway))
			// 如果是断开连接，则删除客户端
			break
		}
		handle(messageType, message)
	}
}

func (client *WebSocketClient) SendMessage(message []byte) error {
	golog.Info("send message [%s] to client", string(message))
	if client.conn == nil {
		return fmt.Errorf("websocket client conn is nil")
	}
	if client.send == nil {
		return fmt.Errorf("websocket client send is nil")
	}
	select {
	case client.send <- message:
	default:
		golog.Info("Websocket client send buffer full")
	}
	return nil
}
