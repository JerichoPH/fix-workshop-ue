package models

type PositionIndoorCellModel struct {
	BaseModel
	UniqueCode             string                  `gorm:"type:CHAR(15);COMMENT:位代码;" json:"unique_code"`
	Name                   string                  `gorm:"type:VARCHAR(64);COMMENT:位名称;" json:"name"`
	PositionIndoorTierUuid string                  `gorm:"type:VARCHAR(36);COMMENT:所属层代码;" json:"position_indoor_tier_uuid"`
	PositionIndoorTier     PositionIndoorTierModel `gorm:"foreignKey:PositionIndoorTierUuid;references:Uuid;COMMENT:所属层;" json:"position_indoor_tier"`
}

// TableName 表名称
func (PositionIndoorCellModel) TableName() string {
	return "position_indoor_cells"
}
