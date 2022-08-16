package databases

import (
	"fix-workshop-ue/settings"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgresql struct {
	Host     string
	Port     string
	Username string
	Database string
	Password string
	SSLMode  string
}

var postgresqlConn *gorm.DB

// getConn 获取链接
//  @receiver cls
//  @return db
func (cls *Postgresql) getConn() (db *gorm.DB) {
	ctf := settings.Setting{}
	config := ctf.Init()

	cls.Host = config.DB.Section("postgresql").Key("host").MustString("127.0.0.1")
	cls.Port = config.DB.Section("postgresql").Key("port").MustString("5432")
	cls.Username = config.DB.Section("postgresql").Key("username").MustString("postgres")
	cls.Password = config.DB.Section("postgresql").Key("password").MustString("zces@1234")
	cls.Database = config.DB.Section("postgresql").Key("database").MustString("postgres")
	cls.SSLMode = config.DB.Section("postgresql").Key("ssl_mode").MustString("disable")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cls.Host,
		cls.Port,
		cls.Username,
		cls.Database,
		cls.Password,
		cls.SSLMode,
	)

	db, _ = gorm.Open(postgres.Open(dsn))

	return
}

// GetConn 获取数据库链接
func (cls *Postgresql) GetConn() *gorm.DB {
	if postgresqlConn == nil {
		postgresqlConn = cls.getConn()
	}
	return postgresqlConn
}

// NewConn 获取新数据库链接
func (cls *Postgresql) NewConn() *gorm.DB {
	return cls.getConn()
}
