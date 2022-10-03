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

// getConn 获取数据库链接
//  @receiver ins
//  @return db
func (ins *MySql) getConn() (db *gorm.DB) {
	ctf := settings.Setting{}
	config := ctf.Init()

	ins.Username = config.DB.Section("mysql").Key("username").MustString("root")
	ins.Password = config.DB.Section("mysql").Key("password").MustString("root")
	ins.Host = config.DB.Section("mysql").Key("host").MustString("127.0.0.1")
	ins.Port = config.DB.Section("mysql").Key("port").MustString("3306")
	ins.Database = config.DB.Section("mysql").Key("database").MustString("FixWorkshop")
	ins.Charset = config.DB.Section("mysql").Key("charset").MustString("utf8mb4")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		ins.Username,
		ins.Password,
		ins.Host,
		ins.Port,
		ins.Database,
		ins.Charset,
	)

	mySqlConn, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize:                          1000,
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
func (ins *MySql) GetConn() *gorm.DB {
	if mySqlConn == nil {
		mySqlConn = ins.getConn()
	}
	return mySqlConn
}

// NewConn 获取新数据库链接
func (ins *MySql) NewConn() *gorm.DB {
	return ins.getConn()
}
