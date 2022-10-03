package databases

import (
	"fix-workshop-ue/settings"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSql struct {
	Host     string
	Port     string
	Username string
	Database string
	Password string
	SSLMode  string
}

var postgresqlConn *gorm.DB

// getConn 获取链接
//  @receiver ins
//  @return db
func (ins *PostgreSql) getConn() (db *gorm.DB) {
	ctf := settings.Setting{}
	config := ctf.Init()

	ins.Host = config.DB.Section("postgresql").Key("host").MustString("127.0.0.1")
	ins.Port = config.DB.Section("postgresql").Key("port").MustString("5432")
	ins.Username = config.DB.Section("postgresql").Key("username").MustString("postgres")
	ins.Password = config.DB.Section("postgresql").Key("password").MustString("zces@1234")
	ins.Database = config.DB.Section("postgresql").Key("database").MustString("postgres")
	ins.SSLMode = config.DB.Section("postgresql").Key("ssl_mode").MustString("disable")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		ins.Host,
		ins.Port,
		ins.Username,
		ins.Database,
		ins.Password,
		ins.SSLMode,
	)

	mySqlConn, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
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
func (ins *PostgreSql) GetConn() *gorm.DB {
	if postgresqlConn == nil {
		postgresqlConn = ins.getConn()
	}
	return postgresqlConn
}

// NewConn 获取新数据库链接
func (ins *PostgreSql) NewConn() *gorm.DB {
	return ins.getConn()
}
