package main

import (
	v1 "fix-workshop-go/routes/v1"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"os"
	"time"

	"fix-workshop-go/errors"
	"github.com/gin-gonic/gin"
)

// initServer 启动服务
func initServer(router *gin.Engine) {
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	serverErr := server.ListenAndServe()
	if serverErr != nil {
		log.Println("服务器启动错误：", serverErr)
	}
}

// initConfig 初始化配置
func initConfig() (appConfigFile *ini.File, dbConfigFile *ini.File) {
	appConfigFile, appConfigErr := ini.Load("./configs/app.ini")
	if appConfigErr != nil {
		fmt.Println("加载主配置文件失败")
		os.Exit(1)
	}

	dbConfigFile, dbConfigErr := ini.Load("./configs/db.ini")
	if dbConfigErr != nil {
		fmt.Println("加载数据库配置文件失败")
		os.Exit(1)
	}

	return
}

func main() {
	appConfigFile, dbConfigFile := initConfig()

	mySqlConn := (&MySql{
		Username: dbConfigFile.Section("mysql").Key("username").MustString(""),
		Password: dbConfigFile.Section("mysql").Key("password").MustString(""),
		Host:     dbConfigFile.Section("mysql").Key("host").MustString(""),
		Port:     dbConfigFile.Section("mysql").Key("port").MustString(""),
		Database: dbConfigFile.Section("mysql").Key("database").MustString(""),
		Charset:  dbConfigFile.Section("mysql").Key("charset").MustString(""),
	}).InitDB() // 创建mysql链接

	//mssqlConn := (&MsSql{
	//	Schema:   "sqlserver",
	//	Username: "sa",
	//	Password: "JW087073yjz..",
	//	Host:     "127.0.0.1:14332",
	//	Database: "Dwqcgl",
	//}).
	//	InitDB() // 创建mssql链接

	router := gin.Default()

	router.Use(errors.RecoverHandler)                                                                             // 异常处理
	(&v1.V1Router{Router: router, MySqlConn: mySqlConn, AppConfig: appConfigFile, DBConfig: dbConfigFile}).Load() // 加载v1路由

	initServer(router) // 启动服务
}
