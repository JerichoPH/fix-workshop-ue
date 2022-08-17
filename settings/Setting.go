package settings

import "gopkg.in/ini.v1"

// Setting 设置
type Setting struct {
	App *ini.File
	DB  *ini.File
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

	return cls
}
