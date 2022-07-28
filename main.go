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

	"fix-workshop-ue/wrongs"
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
			&models.OrganizationSectionModel{},            // 区间
			&models.OrganizationCenterModel{},             // 中心
			&models.OrganizationRailroadGradeCrossModel{}, // 道口
			&models.OrganizationStationModel{},            // 站场

			// 仓储
			&models.LocationDepotStorehouseModel{}, // 仓储仓库
			&models.LocationDepotSectionModel{},    // 仓储仓库区域
			&models.LocationDepotRowTypeModel{},    // 仓储仓库排类型
			&models.LocationDepotRowModel{},        // 仓储仓库排
			&models.LocationDepotCabinetModel{},    // 仓储柜架
			&models.LocationDepotTierModel{},       // 仓储柜架层
			&models.LocationDepotCellModel{},       // 仓储柜架格位

			// 室内上道位置
			&models.LocationIndoorRoomTypeModel{}, // 机房类型
			&models.LocationIndoorRoomModel{},     // 机房
			&models.LocationIndoorRowModel{},      // 排
			&models.LocationIndoorCabinetModel{},  // 架
			&models.LocationIndoorTierModel{},     // 层
			&models.LocationIndoorCellModel{},     // 位

			// 种类型
			&models.KindCategoryModel{},   // 种类
			&models.KindEntireTypeModel{}, // 类型
			&models.KindSubTypeModel{},    // 型号

		); errAutoMigrate != nil {
		fmt.Println("自动迁移错误：", errAutoMigrate)
		os.Exit(1)
	}

	engine := gin.Default()
	engine.Use(wrongs.RecoverHandler) // 异常处理

	(&v1.Router{}).Load(engine) // 加载v1路由

	initServer(engine, setting.App.Section("app").Key("addr").MustString(":8080")) // 启动服务
}
