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
			// ??????
			&models.Account{},       // ????????????
			&models.AccountStatus{}, // ????????????

			// ?????????
			&models.KindCategory{},    // ??????
			&models.KindEntireModel{}, // ??????
			&models.KindSubModel{},    // ??????

			// ????????????
			&models.OrganizationRailway{},            // ??????
			&models.OrganizationParagraph{},          // ??????
			&models.OrganizationWorkshop{},           // ??????
			&models.OrganizationWorkshopType{},       // ????????????
			&models.OrganizationSection{},            // ??????
			&models.OrganizationRailroadGradeCross{}, // ??????
			&models.OrganizationWorkArea{},           // ??????
			&models.OrganizationStation{},            // ??????
			&models.OrganizationLine{},               // ??????
			&models.OrganizationCenter{},             // ??????

			// ??????
			&models.EntireInstance{},        // ????????????
			&models.EntireInstanceStatus{},  // ????????????
			&models.EntireInstanceUse{},     // ??????????????????
			&models.EntireInstanceLog{},     // ????????????
			&models.EntireInstanceLogType{}, // ??????????????????
			&models.EntireInstanceRepair{},  // ??????????????????

			// ?????????
			&models.FixWorkflowReport{},  // ???????????????
			&models.FixWorkflowProcess{}, // ????????????
			&models.FixWorkflowRecode{},  // ?????????

			// ????????????
			&models.LocationWarehouseStorehouse{}, // ???
			&models.LocationWarehouseArea{},       // ???
			&models.LocationWarehousePlatoon{},    // ???
			&models.LocationWarehouseShelf{},      // ??????
			&models.LocationWarehouseTier{},       // ???
			&models.LocationWarehousePosition{},   // ???

			// ????????????
			&models.LocationInstallRoom{},                      // ??????
			&models.LocationInstallRoomType{},                  // ????????????
			&models.LocationInstallPlatoon{},                   // ???
			&models.LocationInstallShelf{},                     // ??????
			&models.LocationInstallTier{},                      // ???
			&models.LocationInstallPosition{},                  // ???
			&models.LocationSignalPostMainOrIndicator{},        // ???????????????????????????
			&models.LocationSignalPostMainLightPosition{},      // ?????????????????????
			&models.LocationSignalPostIndicatorLightPosition{}, // ????????????????????????

			// ?????????
			&models.Factory{},

			// ??????
			&models.SourceType{}, // ????????????
			&models.SourceName{}, // ????????????

		)

	if errAutoMigrate != nil {
		fmt.Println("?????????????????????", errAutoMigrate)
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
