package main

import (
	"context"
	"fix-workshop-ue/configs"
	"fix-workshop-ue/databases"
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
		//Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").
		Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			//// 用户
			//&models.AccountModel{},       // 用户主表
			//&models.AccountStatusModel{}, // 用户状态
			//
			//// RBAC
			//&models.RbacRoleModel{},       // 角色
			//&models.RbacPermissionModel{}, // 权限
			//
			//// 种类型
			//&models.KindCategoryModel{},   // 种类
			//&models.KindEntireTypeModel{}, // 类型
			//&models.KindSubTypeModel{},    // 型号
			//
			//// 组织机构
			//&models.OrganizationRailwayModel{},            // 路局
			//&models.OrganizationParagraphModel{},          // 站段
			//&models.OrganizationWorkshopModel{},           // 车间
			//&models.OrganizationWorkshopTypeModel{},       // 车间类型
			//&models.OrganizationSectionModel{},            // 区间
			//&models.OrganizationRailroadGradeCrossModel{}, // 道口
			//&models.OrganizationWorkAreaModel{},           // 工区
			//&models.OrganizationStationModel{},            // 站场
			//&models.OrganizationLineModel{},               // 线别
			//&models.OrganizationCenterModel{},             // 中心
			//
			//// 器材
			//&models.EntireInstanceModel{},        // 器材主表
			//&models.EntireInstanceStatusModel{},  // 器材状态
			//&models.EntireInstanceUseModel{},     // 器材使用数据
			//&models.EntireInstanceLogModel{},     // 器材日志
			//&models.EntireInstanceLogTypeModel{}, // 器材日志类型
			//&models.EntireInstanceRepairModel{},  // 器材检修记录
			//
			//// 检修单
			//&models.FixWorkflowReportModel{}, // 检测单主表
			////&models.FixWorkflowProcessModel{}, // 检测过程
			////&models.FixWorkflowRecodeModel{},  // 实测值
			//
			//// 仓库位置
			//&models.LocationWarehouseStorehouseModel{}, // 仓
			//&models.LocationWarehouseAreaModel{},       // 区
			//&models.LocationWarehousePlatoonModel{},    // 排
			//&models.LocationWarehouseShelfModel{},      // 柜架
			//&models.LocationWarehouseTierModel{},       // 层
			//&models.LocationWarehousePositionModel{},   // 位
			//
			//// 上道位置
			//&models.LocationInstallRoomModel{},                      // 机房
			//&models.LocationInstallRoomTypeModel{},                  // 机房类型
			//&models.LocationInstallPlatoonModel{},                   // 排
			//&models.LocationInstallShelfModel{},                     // 柜架
			//&models.LocationInstallTierModel{},                      // 层
			//&models.LocationInstallPositionModel{},                  // 位
			//&models.LocationSignalPostMainOrIndicatorModel{},        // 信号机主体或表示器
			//&models.LocationSignalPostMainLightPositionModel{},      // 信号机主体灯位
			//&models.LocationSignalPostIndicatorLightPositionModel{}, // 信号机表示器灯位
			//
			//// 供应商
			//&models.FactoryModel{}, // 供应商
			//
			//// 来源
			//&models.SourceTypeModel{}, // 来源类型
			//&models.SourceNameModel{}, // 来源名称

		)

	if errAutoMigrate != nil {
		fmt.Println("自动迁移错误：", errAutoMigrate)
		os.Exit(1)
	}

	router := gin.Default()

	router.Use(errors.RecoverHandler)     // 异常处理
	(&v1.V1Router{Router: router}).Load() // 加载v1路由

	initServer(router, config.App.Section("app").Key("addr").MustString(":8080")) // 启动服务
}
