package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rengensheng/backend/app/config"
	"github.com/rengensheng/backend/app/constant"
	middleware "github.com/rengensheng/backend/app/middlewares"
	"github.com/rengensheng/backend/app/routes"
	"github.com/rengensheng/backend/infrastructure/database"
	"github.com/rengensheng/backend/infrastructure/migrations"
	"xorm.io/xorm"
)

type App struct {
	DB     *xorm.Engine
	Engine *gin.Engine
	Config *config.Config
}

var appInstance *App

func (app *App) Initialize(db *xorm.Engine, engine *gin.Engine, sourceConfig *config.Config) *App {
	var err error
	if engine == nil {
		engine = gin.Default()
	}
	if sourceConfig == nil {
		sourceConfig, err = config.LoadConfig("")
		if err != nil {
			panic(constant.LOAD_CONFIG_FILE_ERROR)
		}
	}
	if db == nil {
		driverName := sourceConfig.Database.DriverName
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			sourceConfig.Database.Username,
			sourceConfig.Database.Pwd,
			sourceConfig.Database.Host,
			sourceConfig.Database.Port,
			sourceConfig.Database.DatabaseName,
		)
		db, err = database.Init(driverName, dataSourceName)
		if err != nil {
			panic(constant.LOAD_DATABASE_ERROR)
		}
	}
	if sourceConfig.Database.Sync {
		migrations.Sync(db)
	}
	appInstance = &App{
		DB:     db,
		Engine: engine,
		Config: sourceConfig,
	}
	return appInstance
}

func (app *App) SetupRoutes() {
	rootRoute := app.Engine.Group("/api")
	{
		// 启动中间件
		rootRoute.Use(middleware.Auth())
		// 注册静态资源服务

		rootRoute.Static("/upload", app.Config.Service.UploadDir)
		// 注册业务路由
		routes.SetupRoutes(rootRoute, app.DB)
	}
}

func (app *App) RunApplication() {
	err := app.Engine.Run(fmt.Sprintf(":%d", app.Config.Service.Port))
	if err != nil {
		panic(constant.RUN_APPLICATION_ERROR + ":" + err.Error())
	}
}

func GetAppInstance() *App {
	return appInstance
}
