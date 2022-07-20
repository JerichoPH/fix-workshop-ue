package models

type LocationDepotCellsModel struct {
	BaseModel
	UniqueCode                   string                        `gorm:"type:CHAR(18);UNIQUE;NOT NULL;COMMENT:位代码;" json:"unique_code"`
	Name                         string                        `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:位名称;" json:"name"`
	LocationDepotCabinetTierUUID string                        `gorm:"type:CHAR(36);COMMENT:所属层UUID;" json:"location_depot_cabinet_tier_uuid"`
	LocationDepotCabinetTier     LocationDepotCabinetTierModel `gorm:"foreignKey:LocationDepotCabinetTierUUID;references:UUID;NOT NULL;COMMENT:所属层;" json:"location_depot_cabinet_tier"`
	EntireInstances              []EntireInstanceModel         `gorm:"foreignKey:LocationDepotCellUUID;references:UUID;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *LocationDepotCellsModel) TableName() string {
	return "location_depot_cells"
}
