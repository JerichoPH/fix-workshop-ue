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
//  @receiver ins
//  @return *Setting
func (ins *Setting) Init() *Setting {

	appConfigFile, appConfigErr := ini.Load("./settings/app.ini")
	if appConfigErr != nil {
		panic(appConfigErr)
	}

	dbConfigFile, dbConfigErr := ini.Load("./settings/db.ini")
	if dbConfigErr != nil {
		panic(dbConfigErr)
	}

	ins.App = appConfigFile
	ins.DB = dbConfigFile
	ins.Timezone, _ = time.LoadLocation(ins.App.Section("app").Key("timezone").MustString("Asia/Shanghai"))
	ins.Time = (&time.Time{}).In(ins.Timezone)

	return ins
}

// Boot 获取配置
//  @return *Setting
func Boot() *Setting {
	return (&Setting{}).Init()
}
