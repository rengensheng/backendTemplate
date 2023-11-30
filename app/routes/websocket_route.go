package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rengensheng/backend/app/websockets"
)

func SetupWebSocketRouter(rootRouter *gin.RouterGroup) {
	websocketService := websockets.NewWebSocketService()
	websocketHandle := websockets.NewWebSocketHandler(websocketService)
	router := rootRouter.Group("/websocket")
	{
		router.GET("/:token", websocketHandle.HandleWebSocket)
	}
}
