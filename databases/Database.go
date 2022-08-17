package databases

import (
	"fix-workshop-ue/settings"
	"gorm.io/gorm"
)

type Database struct {
	DBDriver string
}

// GetDatabase 获取当前数据库链接
//  @receiver Database
//  @return *gorm.DB
func (cls Database) GetDatabase() (dbSession *gorm.DB) {
	var dbDriver string

	if cls.DBDriver != "" {
		dbDriver = cls.DBDriver
	} else {
		setting := (&settings.Setting{}).Init()
		setting.App.Section("app").Key("db_driver").MustString("")
	}

	switch dbDriver {
	case "mysql":
		dbSession = (&MySql{}).GetConn()
	case "mssql":
		dbSession = (&MsSql{}).GetConn()
	case "postgresql":
	default:
		dbSession = (&Postgresql{}).GetConn()
	}

	return
}
