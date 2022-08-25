package databases

import (
	"fix-workshop-ue/settings"
	"gorm.io/gorm"
)

// DatabaseLaunch 数据库启动器
type DatabaseLaunch struct {
	DBDriver string
}

// GetDatabase 获取当前数据库链接
//  @receiver DatabaseLaunch
//  @return *gorm.DB
func (cls DatabaseLaunch) GetDatabase() (dbSession *gorm.DB) {
	var dbDriver string

	if cls.DBDriver != "" {
		dbDriver = cls.DBDriver
	} else {
		setting := (&settings.Setting{}).Init()
		dbDriver = setting.DB.Section("db").Key("db_driver").MustString("")
	}

	switch dbDriver {
	case "postgresql":
	default:
		dbSession = (&Postgresql{}).GetConn()
	case "mysql":
		dbSession = (&MySql{}).GetConn()
	case "mssql":
		dbSession = (&MsSql{}).GetConn()
	}

	return
}
