package websockets

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hyahm/golog"
	"github.com/rengensheng/backend/app/utils"
	"net/http"
)

type WebSocketHandle struct {
	upgrader websocket.Upgrader
	service  *WebSocketService
}

func NewWebSocketHandler(service *WebSocketService) *WebSocketHandle {
	return &WebSocketHandle{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		service: service,
	}
}

func (handle *WebSocketHandle) HandleWebSocket(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		golog.Info("User token is empty, create websocket connection failed")
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	// parse token
	claims, err := utils.ParseToken(token)
	if err != nil || claims.Username == "" {
		golog.Info("Parse token error %s", err.Error())
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	userId := claims.Username
	conn, err := handle.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := NewWebSocketClient(conn)
	golog.Info("User %s creates a new websocket connection", userId)
	handle.service.AddClient(client, userId)
}
