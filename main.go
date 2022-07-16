package main

import (
	"context"
	"fix-workshop-ue/databases"
	"fix-workshop-ue/models"
	v1 "fix-workshop-ue/routes/v1"
	"fix-workshop-ue/settings"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"fix-workshop-ue/exceptions"
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
	if serverErr := server.ListenAndServe(); serverErr != nil {
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
	setting := (&settings.Setting{}).Init()

	//mssqlConn := (&MsSql{
	//	Schema:   "sqlserver",
	//	Username: "sa",
	//	Password: "JW087073yjz..",
	//	Host:     "127.0.0.1:14332",
	//	Database: "Dwqcgl",
	//}).
	//	InitDB() // 创建mssql链接

	if errAutoMigrate := (&databases.MySql{}).GetConn().
		Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").
		Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			// 用户与权鉴
			&models.AccountModel{},             // 用户
			&models.RbacRoleModel{},            // 角色
			&models.RbacPermissionModel{},      // 权限
			&models.RbacPermissionGroupModel{}, //权限分组

			// 组织机构
			&models.OrganizationRailwayModel{},            //路局
			&models.OrganizationParagraphModel{},          // 站段
			&models.OrganizationLineModel{},               // 线别
			&models.OrganizationWorkshopTypeModel{},       // 车间类型
			&models.OrganizationWorkshopModel{},           // 车间
			&models.OrganizationWorkAreaTypeModel{},       // 工区类型
			&models.OrganizationWorkAreaModel{},           // 工区
			&models.OrganizationStationModel{},            // 站场
			&models.OrganizationRailroadGradeCrossModel{}, // 道口
			&models.OrganizationSectionModel{},            // 区间
			&models.OrganizationCenterModel{},             // 中心

		); errAutoMigrate != nil {
		fmt.Println("自动迁移错误：", errAutoMigrate)
		os.Exit(1)
	}

	router := gin.Default()
	router.Use(exceptions.RecoverHandler) // 异常处理

	// 用户与权鉴
	(&v1.AuthorizationRouter{}).Load(router)       // 权鉴
	(&v1.AccountRouter{}).Load(router)             // 用户                                                                                                                                                          // 用户
	(&v1.RbacRoleRouter{}).Load(router)            // 角色
	(&v1.RbacPermissionGroupRouter{}).Load(router) //权限分组
	(&v1.RbacPermissionRouter{}).Load(router)      // 权限
	(&v1.MenuRouter{}).Load(router)                // 菜单

	// 组织机构
	(&v1.OrganizationLineRouter{}).Load(router)         // 线别
	(&v1.OrganizationRailwayRouter{}).Load(router)      // 路局
	(&v1.OrganizationParagraphRouter{}).Load(router)    // 站段
	(&v1.OrganizationWorkshopTypeRouter{}).Load(router) // 车间类型
	(&v1.OrganizationWorkshopRouter{}).Load(router)     // 车间

	initServer(router, setting.App.Section("app").Key("addr").MustString(":8080")) // 启动服务
}
