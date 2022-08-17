package main

import (
	"context"
	"fix-workshop-ue/databases"
	"fix-workshop-ue/models"
	v1 "fix-workshop-ue/routes/v1"
	web "fix-workshop-ue/routes/web"
	"fix-workshop-ue/settings"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
)

// runServer 启动服务
func runServer(router *gin.Engine, addr string) {
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

// runAutoMigrate 初始化数据库迁移
func runAutoMigrate() {
	if errAutoMigrate := (&databases.DatabaseLaunch{}).GetDatabase().AutoMigrate(
		// 用户与权鉴
		&models.AccountModel{},             // 用户
		&models.RbacRoleModel{},            // 角色
		&models.RbacPermissionModel{},      // 权限
		&models.RbacPermissionGroupModel{}, //权限分组

		// 组织机构
		&models.OrganizationRailwayModel{},            //路局
		&models.OrganizationParagraphModel{},          // 站段
		&models.LocationLineModel{},                   // 线别
		&models.OrganizationWorkshopTypeModel{},       // 车间类型
		&models.OrganizationWorkshopModel{},           // 车间
		&models.OrganizationWorkAreaTypeModel{},       // 工区类型
		&models.OrganizationWorkAreaProfessionModel{}, // 工区专业
		&models.OrganizationWorkAreaModel{},           // 工区
		&models.LocationSectionModel{},                // 区间
		&models.LocationCenterModel{},                 // 中心
		&models.LocationRailroadGradeCrossModel{},     // 道口
		&models.LocationStationModel{},                // 站场

		// 仓储
		&models.PositionDepotStorehouseModel{}, // 仓储仓库
		&models.PositionDepotSectionModel{},    // 仓储仓库区域
		&models.PositionDepotRowTypeModel{},    // 仓储仓库排类型
		&models.PositionDepotRowModel{},        // 仓储仓库排
		&models.PositionDepotCabinetModel{},    // 仓储柜架
		&models.PositionDepotTierModel{},       // 仓储柜架层
		&models.PositionDepotCellModel{},       // 仓储柜架格位

		// 室内上道位置
		&models.PositionIndoorRoomTypeModel{}, // 机房类型
		&models.PositionIndoorRoomModel{},     // 机房
		&models.PositionIndoorRowModel{},      // 排
		&models.PositionIndoorCabinetModel{},  // 架
		&models.PositionIndoorTierModel{},     // 层
		&models.PositionIndoorCellModel{},     // 位

		// 种类型
		&models.KindCategoryModel{},   // 种类
		&models.KindEntireTypeModel{}, // 类型
		&models.KindSubTypeModel{},    // 型号

		// 器材
		&models.EntireInstanceModel{},        // 器材
		&models.EntireInstanceStatusModel{},  // 器材状态
		&models.EntireInstanceLockModel{},    // 器材锁
		&models.EntireInstanceLogTypeModel{}, // 器材日志类型
		&models.EntireInstanceLogModel{},     // 器材日志
	); errAutoMigrate != nil {
		fmt.Println("数据库迁移错误")
		os.Exit(1)
	}
}

// main 程序入口
func main() {
	settingApp := (&settings.Setting{}).Init().App // 加载参数（程序）
	runAutoMigrate()                               // 数据库迁移
	engine := gin.Default()                        // 启动服务引擎
	engine.Use(wrongs.RecoverHandler)              // 异常处理
	(&web.Router{}).Load(engine)                   // 加载web路由
	(&v1.Router{}).Load(engine)                    // 加载v1路由

	runServer(engine, settingApp.Section("app").Key("addr").MustString(":8080")) // 启动服务
}
