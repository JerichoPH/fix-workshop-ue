package main

import (
	"context"
	"fix-workshop-ue/config"
	"fix-workshop-ue/database"
	"fix-workshop-ue/model"
	v1 "fix-workshop-ue/router/v1"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"fix-workshop-ue/error"
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
	config := (&config.Config{}).Init()

	//mssqlConn := (&MsSql{
	//	Schema:   "sqlserver",
	//	Username: "sa",
	//	Password: "JW087073yjz..",
	//	Host:     "127.0.0.1:14332",
	//	Database: "Dwqcgl",
	//}).
	//	InitDB() // 创建mssql链接

	mySqlConn := (&database.MySql{}).GetMySqlConn()
	errAutoMigrate := mySqlConn.
		//Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").
		Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			// 用户
			&model.AccountModel{},       // 用户主表
			&model.AccountStatusModel{}, // 用户状态

			// RBAC
			&model.RbacRoleModel{},       // 角色
			&model.RbacPermissionModel{}, // 权限

			// 种类型
			&model.KindCategoryModel{},   // 种类
			&model.KindEntireTypeModel{}, // 类型
			&model.KindSubTypeModel{},    // 型号

			// 组织机构
			&model.OrganizationRailwayModel{},            // 路局
			&model.OrganizationParagraphModel{},          // 站段
			&model.OrganizationWorkshopModel{},           // 车间
			&model.OrganizationWorkshopTypeModel{},       // 车间类型
			&model.OrganizationSectionModel{},            // 区间
			&model.OrganizationRailroadGradeCrossModel{}, // 道口
			&model.OrganizationWorkAreaModel{},           // 工区
			&model.OrganizationStationModel{},            // 站场
			&model.OrganizationLineModel{},               // 线别
			&model.OrganizationCenterModel{},             // 中心

			// 器材
			&model.EntireInstanceModel{},        // 器材主表
			&model.EntireInstanceStatusModel{},  // 器材状态
			&model.EntireInstanceUseModel{},     // 器材使用数据
			&model.EntireInstanceLogModel{},     // 器材日志
			&model.EntireInstanceLogTypeModel{}, // 器材日志类型
			&model.EntireInstanceRepairModel{},  // 器材检修记录

			// 检修单
			&model.FixWorkflowReportModel{}, // 检测单主表
			//&model.FixWorkflowProcessModel{}, // 检测过程
			//&model.FixWorkflowRecodeModel{},  // 实测值

			// 仓库位置
			&model.LocationWarehouseStorehouseModel{}, // 仓
			&model.LocationWarehouseAreaModel{},       // 区
			&model.LocationWarehousePlatoonModel{},    // 排
			&model.LocationWarehouseShelfModel{},      // 柜架
			&model.LocationWarehouseTierModel{},       // 层
			&model.LocationWarehousePositionModel{},   // 位

			// 上道位置
			&model.LocationInstallRoomModel{},                      // 机房
			&model.LocationInstallRoomTypeModel{},                  // 机房类型
			&model.LocationInstallPlatoonModel{},                   // 排
			&model.LocationInstallShelfModel{},                     // 柜架
			&model.LocationInstallTierModel{},                      // 层
			&model.LocationInstallPositionModel{},                  // 位
			&model.LocationSignalPostMainOrIndicatorModel{},        // 信号机主体或表示器
			&model.LocationSignalPostMainLightPositionModel{},      // 信号机主体灯位
			&model.LocationSignalPostIndicatorLightPositionModel{}, // 信号机表示器灯位

			// 供应商
			&model.FactoryModel{}, // 供应商

			// 来源
			&model.SourceTypeModel{}, // 来源类型
			&model.SourceNameModel{}, // 来源名称

		)

	if errAutoMigrate != nil {
		fmt.Println("自动迁移错误：", errAutoMigrate)
		os.Exit(1)
	}

	router := gin.Default()

	router.Use(error.RecoverHandler)      // 异常处理
	(&v1.V1Router{Router: router}).Load() // 加载v1路由

	initServer(router, config.App.Section("app").Key("addr").MustString(":8080")) // 启动服务
}
