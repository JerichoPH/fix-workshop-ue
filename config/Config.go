package config

import "gopkg.in/ini.v1"

type Config struct {
	App *ini.File
	DB  *ini.File
}

func (cls *Config) Init() *Config {

	appConfigFile, appConfigErr := ini.Load("./config/app.ini")
	if appConfigErr != nil {
		panic(appConfigErr)
	}

	dbConfigFile, dbConfigErr := ini.Load("./config/db.ini")
	if dbConfigErr != nil {
		panic(dbConfigErr)
	}

	cls.App = appConfigFile
	cls.DB = dbConfigFile

	return cls
}
