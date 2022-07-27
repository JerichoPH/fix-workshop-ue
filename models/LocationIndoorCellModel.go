package models

type LocationIndoorCellModel struct {
	BaseModel
	UniqueCode             string                  `gorm:"type:CHAR(15);UNIQUE;NOT NULL;COMMENT:位代码;" json:"unique_code"`
	Name                   string                  `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:位名称;" json:"name"`
	LocationIndoorTierUUID string                  `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属层代码;" json:"location_indoor_tier_uuid"`
	LocationIndoorTier     LocationIndoorTierModel `gorm:"foreignKey:LocationIndoorTierUUID;references:UUID;COMMENT:所属层;" json:"location_indoor_tier"`
}

// TableName 表名称
func (LocationIndoorCellModel) TableName() string {
	return "location_indoor_cells"
}
