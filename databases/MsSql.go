package databases

import (
	"fix-workshop-ue/settings"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type MsSql struct{
	Schema string
	Username string
	Password string
	Host string
	Database string
}

var msSqlConn *gorm.DB

func(cls *MsSql) getConn() (db *gorm.DB){
	ctf := settings.Setting{}
	config := ctf.Init()

	cls.Username = config.DB.Section("mssql").Key("username").MustString("sa")
	cls.Password = config.DB.Section("mssql").Key("password").MustString("JW087073yjz..")
	cls.Host = config.DB.Section("mssql").Key("host").MustString("127.0.0.1")
	cls.Database = config.DB.Section("mssql").Key("databases").MustString("Dwqcgl")

	dsn := fmt.Sprintf(
		"%s://%s:%s@%s?database=%s",
		cls.Schema,
		cls.Username,
		cls.Password,
		cls.Host,
		cls.Database,
	)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return
}

// GetConn 获取数据库链接
func (cls *MsSql) GetConn() *gorm.DB {
	if msSqlConn == nil {
		msSqlConn = cls.getConn()
	}
	return msSqlConn
}

// NewConn 获取新数据库链接
func (cls *MsSql) NewConn() *gorm.DB {
	return cls.getConn()
}