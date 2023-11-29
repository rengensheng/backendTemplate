package database

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func Init(driverName, dbURL string) (*xorm.Engine, error) {
	var err error
	engine, err = xorm.NewEngine(driverName, dbURL)
	if err != nil {
		log.Panic("数据库加载失败!", err.Error())
	}
	err = engine.Ping()
	if err != nil {
		log.Panic("数据库无响应，请检查网络!", err.Error())
	}
	return engine, err
}

func GetDB() *xorm.Engine {
	return engine
}
