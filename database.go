package main

import (
	"fix-workshop-ue/model"
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
//	dbConfigFile, dbConfigErr := ini.Load("./config/db.ini")
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
	dbConfigFile, dbConfigErr := ini.Load("./config/db.ini")
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
			&model.AccountModel{},       // 用户主表
			&model.AccountStatusModel{}, // 用户状态

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
			&model.FactoryModel{},

			// 来源
			&model.SourceTypeModel{}, // 来源类型
			&model.SourceNameModel{}, // 来源名称

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
