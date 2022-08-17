package models

type PositionOutdoorSignalPostIndicatorLightPositionModel struct {
	BaseModel
	UniqueCode string `gorm:"type:CHAR(2);NOT NULL;COMMENT:信号机表示器灯位代码;" json:"unique_code"`
	Name       string `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:信号机表示器灯位名称;" json:"name"`
}

// TableName 表名称
func (PositionOutdoorSignalPostIndicatorLightPositionModel) TableName() string {
	return "position_signal_post_indicator_light_positions"
}
