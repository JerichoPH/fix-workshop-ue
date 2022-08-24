package settings

import (
	"gopkg.in/ini.v1"
	"time"
)

// Setting 设置
type Setting struct {
	App      *ini.File
	DB       *ini.File
	Time     time.Time
	Timezone *time.Location
}

// Init 获取配置文件
//  @receiver cls
//  @return *Setting
func (cls *Setting) Init() *Setting {

	appConfigFile, appConfigErr := ini.Load("./settings/app.ini")
	if appConfigErr != nil {
		panic(appConfigErr)
	}

	dbConfigFile, dbConfigErr := ini.Load("./settings/db.ini")
	if dbConfigErr != nil {
		panic(dbConfigErr)
	}

	cls.App = appConfigFile
	cls.DB = dbConfigFile
	cls.Timezone, _ = time.LoadLocation(cls.App.Section("app").Key("timezone").MustString("Asia/Shanghai"))
	cls.Time = (&time.Time{}).In(cls.Timezone)

	return cls
}

// Boot 获取配置
//  @return *Setting
func Boot() *Setting {
	return (&Setting{}).Init()
}
