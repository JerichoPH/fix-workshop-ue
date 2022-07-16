package models

type LocationIndoorCabinetModel struct {
	BaseModel
	UniqueCode            string                    `gorm:"type:CHAR(11);UNIQUE;NOT NULL;COMMENT:柜架代码;" json:"unique_code"`
	Name                  string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:柜架名称;" json:"name"`
	LocationIndoorRowUUID string                    `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属排;" json:"location_indoor_row_uuid"`
	LocationIndoorRow     LocationIndoorRowModel    `gorm:"foreignKey:LocationIndoorRowUUID;references:UUID;COMMENT:所属排;" json:"location_indoor_row"`
	LocationIndoorTiers   []LocationIndoorTierModel `gorm:"foreignKey:LocationIndoorCabinetUUID;references:UUID;COMMENT:相关层;" json:"location_indoor_tiers"`
}

// TableName 表名称
func (cls *LocationIndoorCabinetModel) TableName() string {
	return "location_indoor_cabinets"
}