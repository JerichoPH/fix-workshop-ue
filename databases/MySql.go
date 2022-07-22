package databases

import (
	"fix-workshop-ue/settings"
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

func (cls *MySql) getConn() (db *gorm.DB) {
	ctf := settings.Setting{}
	config := ctf.Init()

	cls.Username = config.DB.Section("mysql").Key("username").MustString("root")
	cls.Password = config.DB.Section("mysql").Key("password").MustString("root")
	cls.Host = config.DB.Section("mysql").Key("host").MustString("127.0.0.1")
	cls.Port = config.DB.Section("mysql").Key("port").MustString("3306")
	cls.Database = config.DB.Section("mysql").Key("databases").MustString("FixWorkshop")
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
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	db = mySqlConn.Session(&gorm.Session{
		SkipDefaultTransaction: true,
		QueryFields:            true,
		PrepareStmt:            true,
	})

	return
}

// GetConn 获取数据库链接
func (cls *MySql) GetConn() *gorm.DB {
	if mySqlConn == nil {
		mySqlConn = cls.getConn()
	}
	return mySqlConn
}

// NewConn 获取新数据库链接
func (cls *MySql) NewConn() *gorm.DB {
	return cls.getConn()
}
