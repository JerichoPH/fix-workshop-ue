package models

type LocationOutdoorSignalPostMainOrIndicatorModel struct {
	BaseModel
	UniqueCode string `gorm:"type:CHAR(6);UNIQUE;NOT NULL;COMMENT:信号灯主机或表示器代码;" json:"unique_code"`
	Name       string `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:信号机主机或表示器名称;" json:"name"`
}

// TableName 表名称
func (LocationOutdoorSignalPostMainOrIndicatorModel) TableName() string {
	return "location_signal_post_main_or_indicators"
}
