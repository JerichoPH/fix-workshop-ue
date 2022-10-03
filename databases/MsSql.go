package databases

import (
	"fix-workshop-ue/settings"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type MsSql struct {
	Schema   string
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

var msSqlConn *gorm.DB

func (ins *MsSql) getConn() (db *gorm.DB) {
	ctf := settings.Setting{}
	config := ctf.Init()

	ins.Username = config.DB.Section("mssql").Key("username").MustString("sa")
	ins.Password = config.DB.Section("mssql").Key("password").MustString("JW087073yjz..")
	ins.Host = config.DB.Section("mssql").Key("host").MustString("127.0.0.1")
	ins.Port = config.DB.Section("mssql").Key("port").MustString("1433")
	ins.Database = config.DB.Section("mssql").Key("databases").MustString("Dwqcgl")

	dsn := fmt.Sprintf(
		"%s://%s:%s@%s:%s?database=%s",
		ins.Schema,
		ins.Username,
		ins.Password,
		ins.Host,
		ins.Port,
		ins.Database,
	)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return
}

// GetConn 获取数据库链接
func (ins *MsSql) GetConn() *gorm.DB {
	if msSqlConn == nil {
		msSqlConn = ins.getConn()
	}
	return msSqlConn
}

// NewConn 获取新数据库链接
func (ins *MsSql) NewConn() *gorm.DB {
	return ins.getConn()
}
