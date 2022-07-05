package database

import (
	"fix-workshop-ue/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySql struct {
	Schema   string
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
}

var mySqlConn *gorm.DB

func (cls *MySql) getMySqlConn() (tx *gorm.DB) {
	ctf := config.Config{}
	config := ctf.Init()

	cls.Username = config.DB.Section("mysql").Key("username").MustString("root")
	cls.Password = config.DB.Section("mysql").Key("password").MustString("root")
	cls.Host = config.DB.Section("mysql").Key("host").MustString("127.0.0.1")
	cls.Port = config.DB.Section("mysql").Key("port").MustString("3306")
	cls.Database = config.DB.Section("mysql").Key("database").MustString("FixWorkshop")
	cls.Charset = config.DB.Section("mysql").Key("charset").MustString("utf8mb4")

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

	tx = mySqlConn.Session(&gorm.Session{
		SkipDefaultTransaction: true,
		QueryFields:            true,
		PrepareStmt:            true,
	})

	return
}

// GetMySqlConn 获取数据库链接
func (cls *MySql) GetMySqlConn() *gorm.DB {
	if mySqlConn == nil {
		mySqlConn = cls.getMySqlConn()
	}
	return mySqlConn
}

// GetNewMySqlConn 获取新数据库链接
func (cls *MySql) GetNewMySqlConn() *gorm.DB {
	return cls.getMySqlConn()
}