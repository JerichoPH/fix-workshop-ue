package main

import (
	"context"
	"fix-workshop-ue/configs"
	"fix-workshop-ue/databases"
	"fix-workshop-ue/models"
	v1 "fix-workshop-ue/routes/v1"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"fix-workshop-ue/errors"
	"github.com/gin-gonic/gin"
)

// initServer 启动服务
func initServer(router *gin.Engine, addr string) {
	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	serverErr := server.ListenAndServe()
	if serverErr != nil {
		log.Println("服务器启动错误：", serverErr)
	}

	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("关闭服务中……")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("服务无法关闭", err)
	}
	log.Println("服务关闭")

}

func main() {
	// 获取参数
	config := (&configs.Config{}).Init()

	//mssqlConn := (&MsSql{
	//	Schema:   "sqlserver",
	//	Username: "sa",
	//	Password: "JW087073yjz..",
	//	Host:     "127.0.0.1:14332",
	//	Database: "Dwqcgl",
	//}).
	//	InitDB() // 创建mssql链接

	mySqlConn := (&databases.MySql{}).GetMySqlConn()
	errAutoMigrate := mySqlConn.
		Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").
		Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			// 用户与权鉴
			&models.AccountStatusModel{},  // 用户状态
			&models.AccountModel{},        // 用户
			&models.RbacRoleModel{},       // 角色
			&models.RbacPermissionModel{}, // 权限

		)
	if errAutoMigrate != nil {
		fmt.Println("自动迁移错误：", errAutoMigrate)
		os.Exit(1)
	}

	router := gin.Default()
	router.Use(errors.RecoverHandler) // 异常处理

	(&v1.AuthorizationRouter{}).Load(router) // 权鉴
	(&v1.AccountRouter{}).Load(router)       // 用户                                                                                                                                                          // 用户
	(&v1.AccountStatusRouter{}).Load(router) // 用户状态
	(&v1.RbacRoleRouter{}).Load(router)      // 角色

	initServer(router, config.App.Section("app").Key("addr").MustString(":8080")) // 启动服务
}
