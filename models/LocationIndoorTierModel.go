package models

type LocationIndoorTierModel struct {
	BaseModel
	UniqueCode                string                     `gorm:"type:CHAR(13);UNIQUE;NOT NULL;COMMENT:层代码;" json:"unique_code"`
	Name                      string                     `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:层名称;" json:"name"`
	LocationIndoorCabinetUUID string                     `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属柜架代码;" json:"location_indoor_cabinet_uuid"`
	LocationIndoorCabinet     LocationIndoorCabinetModel `gorm:"foreignKey:LocationIndoorCabinetUUID;references:UUID;COMMENT:所属柜架;" json:"location_indoor_cabinet"`
	LocationIndoorCells       []LocationIndoorCellModel  `gorm:"foreignKey:LocationIndoorTierUUID;references:UUID;COMMENT:相关位;" json:"location_indoor_tier"`
}

// TableName 表名称
func (cls *LocationIndoorTierModel) TableName() string {
	return "location_indoor_tiers"
}
