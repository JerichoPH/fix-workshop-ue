package settings

import "gopkg.in/ini.v1"

type Setting struct {
	App *ini.File
	DB  *ini.File
}

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
