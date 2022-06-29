package main

import (
	"fix-workshop-go/models"
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//type MsSql struct {
//	Schema   string
//	Username string
//	Password string
//	Host     string
//	Database string
//	DB       *gorm.DB
//}
//
//var MsSqlConn *gorm.DB
//
//func (cls *MsSql) InitConfig() *MsSql {
//	dbConfigFile, dbConfigErr := ini.Load("./configs/db.ini")
//	if dbConfigErr != nil {
//		panic(dbConfigErr)
//	}
//
//	cls.Username = dbConfigFile.Section("mssql").Key("username").MustString("")
//	cls.Password = dbConfigFile.Section("mssql").Key("password").MustString("")
//	cls.Host = dbConfigFile.Section("mssql").Key("host").MustString("127.0.0.1")
//	cls.Database = dbConfigFile.Section("mssql").Key("database").MustString("")
//
//	return cls
//}
//
//func (cls *MsSql) InitDB() *gorm.DB {
//	//dsn := "sqlserver://sa:JW087073yjz..@127.0.0.1?Database=Dwqcgl"
//
//	cls.InitConfig()
//
//	dsn := fmt.Sprintf(
//		"%s://%s:%s@%s?database=%s",
//		cls.Schema,
//		cls.Username,
//		cls.Password,
//		cls.Host,
//		cls.Database,
//	)
//	fmt.Println(dsn)
//	msSqlConn, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic(err)
//	}
//
//	return msSqlConn
//}
//
//func GetMsSqlConn() *gorm.DB {
//	if MsSqlConn == nil {
//		MsSqlConn = (&MsSql{}).InitConfig().InitDB()
//	}
//
//	return MsSqlConn
//}
//
//func GetNewMsSqlConn() *gorm.DB {
//	return (&MsSql{}).InitConfig().InitDB()
//}

type MySql struct {
	Schema   string
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
}

var MySqlConn *gorm.DB

func (cls *MySql) InitConfig() *MySql {
	dbConfigFile, dbConfigErr := ini.Load("./configs/db.ini")
	if dbConfigErr != nil {
		panic(dbConfigErr)
	}

	cls.Username = dbConfigFile.Section("mysql").Key("username").MustString("")
	cls.Password = dbConfigFile.Section("mysql").Key("password").MustString("")
	cls.Host = dbConfigFile.Section("mysql").Key("host").MustString("127.0.0.1")
	cls.Port = dbConfigFile.Section("mysql").Key("port").MustString("3306")
	cls.Database = dbConfigFile.Section("mysql").Key("database").MustString("")
	cls.Charset = dbConfigFile.Section("mysql").Key("charset").MustString("")

	return cls
}

func (cls *MySql) InitDB() *gorm.DB {
	//dsn := "root:root@tcp(127.0.0.1:3307)/detector_already_upload?charset=utf8mb4&parseTime=True&loc=Local"

	cls.InitConfig()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		cls.Username,
		cls.Password,
		cls.Host,
		cls.Port,
		cls.Database,
		cls.Charset,
	)

	mySqlConn, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000,
		//SkipDefaultTransaction: true,
		//PrepareStmt:            true,
	})

	tx := mySqlConn.Session(&gorm.Session{
		SkipDefaultTransaction: true,
		QueryFields:            true,
		PrepareStmt:            true,
	})

	errAutoMigrate := tx.
		//Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").
		Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			// 用户
			&models.AccountModel{},       // 用户主表
			&models.AccountStatusModel{}, // 用户状态

			// 种类型
			&models.KindCategoryModel{},   // 种类
			&models.KindEntireTypeModel{}, // 类型
			&models.KindSubTypeModel{},    // 型号

			// 组织机构
			&models.OrganizationRailwayModel{},            // 路局
			&models.OrganizationParagraphModel{},          // 站段
			&models.OrganizationWorkshopModel{},           // 车间
			&models.OrganizationWorkshopTypeModel{},       // 车间类型
			&models.OrganizationSectionModel{},            // 区间
			&models.OrganizationRailroadGradeCrossModel{}, // 道口
			&models.OrganizationWorkAreaModel{},           // 工区
			&models.OrganizationStationModel{},            // 站场
			&models.OrganizationLineModel{},               // 线别
			&models.OrganizationCenterModel{},             // 中心

			// 器材
			&models.EntireInstanceModel{},        // 器材主表
			&models.EntireInstanceStatusModel{},  // 器材状态
			&models.EntireInstanceUseModel{},     // 器材使用数据
			&models.EntireInstanceLogModel{},     // 器材日志
			&models.EntireInstanceLogTypeModel{}, // 器材日志类型
			&models.EntireInstanceRepairModel{},  // 器材检修记录

			// 检修单
			&models.FixWorkflowReportModel{}, // 检测单主表
			//&models.FixWorkflowProcessModel{}, // 检测过程
			//&models.FixWorkflowRecodeModel{},  // 实测值

			// 仓库位置
			&models.LocationWarehouseStorehouseModel{}, // 仓
			&models.LocationWarehouseAreaModel{},       // 区
			&models.LocationWarehousePlatoonModel{},    // 排
			&models.LocationWarehouseShelfModel{},      // 柜架
			&models.LocationWarehouseTierModel{},       // 层
			&models.LocationWarehousePositionModel{},   // 位

			// 上道位置
			&models.LocationInstallRoomModel{},                      // 机房
			&models.LocationInstallRoomTypeModel{},                  // 机房类型
			&models.LocationInstallPlatoonModel{},                   // 排
			&models.LocationInstallShelfModel{},                     // 柜架
			&models.LocationInstallTierModel{},                      // 层
			&models.LocationInstallPositionModel{},                  // 位
			&models.LocationSignalPostMainOrIndicatorModel{},        // 信号机主体或表示器
			&models.LocationSignalPostMainLightPositionModel{},      // 信号机主体灯位
			&models.LocationSignalPostIndicatorLightPositionModel{}, // 信号机表示器灯位

			// 供应商
			&models.FactoryModel{},

			// 来源
			&models.SourceTypeModel{}, // 来源类型
			&models.SourceNameModel{}, // 来源名称

		)

	if errAutoMigrate != nil {
		fmt.Println("自动迁移错误：", errAutoMigrate)
		return nil
	}

	return tx
}

func GetMySqlConn() *gorm.DB {
	if MySqlConn == nil {
		MySqlConn = (&MySql{}).InitConfig().InitDB()
	}

	return MySqlConn
}

func GetNewMySqlConn() *gorm.DB {
	return (&MySql{}).InitConfig().InitDB()
}
