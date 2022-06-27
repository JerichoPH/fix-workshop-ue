package models

type LocationSignalPostMainOrIndicator struct {
	BaseModel
	Preloads   []string
	Selects    []string
	Omits      []string
	UniqueCode string `gorm:"type:CHAR(6);UNIQUE;NOT NULL;COMMENT:信号灯主机或表示器代码;" json:"unique_code"`
	Name       string `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:信号机主机或表示器名称;" json:"name"`
}
