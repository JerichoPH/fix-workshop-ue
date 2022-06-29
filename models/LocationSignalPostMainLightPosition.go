package models

type LocationSignalPostMainLightPosition struct {
	BaseModel
	UniqueCode string `gorm:"type:CHAR(2);UNIQUE;NOT NULL;COMMENT:信号机主体灯位代码;" json:"unique_code"`
	Name       string `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:信号机主体灯位名称;" json:"name"`
}
