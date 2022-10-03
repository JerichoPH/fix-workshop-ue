package databases

import (
	"fix-workshop-ue/settings"
	"fix-workshop-ue/wrongs"
	"gorm.io/gorm"
)

// Launcher 数据库启动器
type Launcher struct {
	DbDriver string
}

// GetDatabaseConn 获取当前数据库链接
//  @receiver Launcher
//  @return *gorm.DB
func (ins Launcher) GetDatabaseConn() (dbSession *gorm.DB) {
	var dbDriver string

	if ins.DbDriver != "" {
		dbDriver = ins.DbDriver
	} else {
		setting := (&settings.Setting{}).Init()
		dbDriver = setting.DB.Section("db").Key("db_driver").MustString("")
	}

	if "postgresql" == dbDriver {
		dbSession = new(PostgreSql).GetConn()
	} else if "mysql" == dbDriver {
		dbSession = new(MySql).GetConn()
	} else if "mssql" == dbDriver {
		dbSession = new(MsSql).GetConn()
	} else {
		wrongs.PanicForbidden("没有配置数据库")
	}

	return
}
