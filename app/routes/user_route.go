package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rengensheng/backend/app/handlers"
	"github.com/rengensheng/backend/app/repositories"
	"github.com/rengensheng/backend/app/services"
	"xorm.io/xorm"
)

func SetupUserRoute(rootRouter *gin.RouterGroup, db *xorm.Engine) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	roleRepository := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepository)

	menuRepository := repositories.NewMenuRepository(db)
	menuService := services.NewMenuService(menuRepository)

	userHandler := handlers.NewUserHandler(userService, roleService, menuService)

	router := rootRouter.Group("/user")
	{
		router.POST("/", userHandler.CreateUser)
		router.POST("/list", userHandler.GetUserList)
		router.GET("/:id", userHandler.GetUserById)
		router.PUT("/:id", userHandler.UpdateUserById)
		router.DELETE("/:id", userHandler.DeleteUserById)
		router.POST("/login", userHandler.UserLogin)
	}
}
