package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rengensheng/backend/app/handlers"
	"github.com/rengensheng/backend/app/repositories"
	"github.com/rengensheng/backend/app/services"
	"xorm.io/xorm"
)

func SetupDeptRoute(rootRouter *gin.RouterGroup, db *xorm.Engine) {
	deptRepository := repositories.NewDeptRepository(db)
	deptService := services.NewDeptService(deptRepository)
	deptHandler := handlers.NewDeptHandler(deptService)
	router := rootRouter.Group("/dept")
	{
		router.POST("/", deptHandler.CreateDept)
		router.POST("/list", deptHandler.GetDeptList)
		router.GET("/:id", deptHandler.GetDeptById)
		router.PUT("/:id", deptHandler.UpdateDeptById)
		router.DELETE("/:id", deptHandler.DeleteDeptById)
	}
}
