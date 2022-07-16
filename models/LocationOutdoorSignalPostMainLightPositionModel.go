package models

type LocationOutdoorSignalPostMainLightPositionModel struct {
	BaseModel
	UniqueCode string `gorm:"type:CHAR(2);UNIQUE;NOT NULL;COMMENT:信号机主体灯位代码;" json:"unique_code"`
	Name       string `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:信号机主体灯位名称;" json:"name"`
}

// TableName 表名称
func (cls *LocationOutdoorSignalPostMainLightPositionModel) TableName() string {
	return "location_signal_post_main_light_positions"
}
