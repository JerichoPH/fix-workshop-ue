package models

type PositionIndoorCabinetModel struct {
	BaseModel
	UniqueCode            string                    `gorm:"type:CHAR(11);COMMENT:柜架代码;" json:"unique_code"`
	Name                  string                    `gorm:"type:VARCHAR(64);COMMENT:柜架名称;" json:"name"`
	PositionIndoorRowUUID string                    `gorm:"type:VARCHAR(36);COMMENT:所属排;" json:"position_indoor_row_uuid"`
	PositionIndoorRow     PositionIndoorRowModel    `gorm:"foreignKey:PositionIndoorRowUUID;references:UUID;COMMENT:所属排;" json:"position_indoor_row"`
	PositionIndoorTiers   []PositionIndoorTierModel `gorm:"foreignKey:PositionIndoorCabinetUUID;references:UUID;COMMENT:相关层;" json:"position_indoor_tiers"`
}

// TableName 表名称
func (PositionIndoorCabinetModel) TableName() string {
	return "position_indoor_cabinets"
}