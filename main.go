package main

import (
	"github.com/rengensheng/backend/app"
	"github.com/rengensheng/backend/app/config"
	"github.com/rengensheng/backend/app/constant"
	"github.com/rengensheng/backend/logger"
)

func main() {
	var appInstance *app.App
	appConfig, err := config.LoadConfig("./config.yaml")
	if err != nil {
		panic(constant.LOAD_CONFIG_FILE_ERROR + ":" + err.Error())
	}
	logger.InitLogger(appConfig.Log.LogFilePath)
	appInstance = appInstance.Initialize(nil, nil, appConfig)
	appInstance.SetupRoutes()
	appInstance.Engine.Run(":8080")
}
