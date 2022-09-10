package models

type PositionIndoorCabinetModel struct {
	BaseModel
	UniqueCode            string                    `gorm:"type:CHAR(11);COMMENT:柜架代码;" json:"unique_code"`
	Name                  string                    `gorm:"type:VARCHAR(64);COMMENT:柜架名称;" json:"name"`
	PositionIndoorRowUuid string                    `gorm:"type:VARCHAR(36);COMMENT:所属排;" json:"position_indoor_row_uuid"`
	PositionIndoorRow     PositionIndoorRowModel    `gorm:"foreignKey:PositionIndoorRowUuid;references:Uuid;COMMENT:所属排;" json:"position_indoor_row"`
	PositionIndoorTiers   []PositionIndoorTierModel `gorm:"foreignKey:PositionIndoorCabinetUuid;references:Uuid;COMMENT:相关层;" json:"position_indoor_tiers"`
}

// TableName 表名称
func (PositionIndoorCabinetModel) TableName() string {
	return "position_indoor_cabinets"
}
