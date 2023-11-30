package routes

import (
	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func SetupRoutes(rootRouter *gin.RouterGroup, db *xorm.Engine) {
	SetupDeptRoute(rootRouter, db)
	SetupUserRoute(rootRouter, db)
	SetupWebSocketRouter(rootRouter)
}
