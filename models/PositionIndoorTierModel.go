package models

type PositionIndoorTierModel struct {
	BaseModel
	UniqueCode                string                     `gorm:"type:CHAR(13);COMMENT:层代码;" json:"unique_code"`
	Name                      string                     `gorm:"type:VARCHAR(64);COMMENT:层名称;" json:"name"`
	PositionIndoorCabinetUUID string                     `gorm:"type:CHAR(36);COMMENT:所属柜架代码;" json:"position_indoor_cabinet_uuid"`
	PositionIndoorCabinet     PositionIndoorCabinetModel `gorm:"foreignKey:PositionIndoorCabinetUUID;references:UUID;COMMENT:所属柜架;" json:"position_indoor_cabinet"`
	PositionIndoorCells       []PositionIndoorCellModel  `gorm:"foreignKey:PositionIndoorTierUUID;references:UUID;COMMENT:相关位;" json:"position_indoor_tier"`
}

// TableName 表名称
func (PositionIndoorTierModel) TableName() string {
	return "position_indoor_tiers"
}
