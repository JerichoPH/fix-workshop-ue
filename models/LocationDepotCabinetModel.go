package models

type LocationDepotCabinetModel struct {
	BaseModel
	UniqueCode                string                          `gorm:"type:CHAR(14);UNIQUE;NOT NULL;COMMENT:柜架代码;" json:"unique_code"`
	Name                      string                          `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:柜架名称;" json:"name"`
	LocationDepotRowUUID      string                          `gorm:"type:CHAR(36);COMMENT:所属排代码;" json:"location_depot_row_uuid"`
	LocationDepotRow          LocationDepotRowModel           `gorm:"foreignKey:LocationDepotRowUUID;references:UUID;NOT NULL;COMMENT:所属排;" json:"location_depot_row"`
	LocationDepotCabinetTiers []LocationDepotCabinetTierModel `gorm:"foreignKey:LocationDepotCabinetUUID;references:UUID;NOT NULL;COMMENT:相关层;" json:"location_depot_tiers"`
}

// TableName 表名称
func (cls *LocationDepotCabinetModel) TableName() string {
	return "location_depot_cabinets"
}
