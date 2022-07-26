package models

// LocationDepotCabinetModel 仓储柜架模型
type LocationDepotCabinetModel struct {
	BaseModel
	UniqueCode           string                   `gorm:"type:CHAR(10);UNIQUE;NOT NULL;COMMENT:仓储柜架代码;" json:"unique_code"`
	Name                 string                   `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:仓储柜架名称;" json:"name"`
	LocationDepotRowUUID string                   `gorm:"type:CHAR(36);NOT NULL;COMMENT:仓储柜架排UUID;" json:"location_depot_row_uuid"`
	LocationDepotRow     LocationDepotRowModel    `gorm:"foreignKey:LocationDepotRowUUID;references:UUID;COMMENT:所属仓储排;" json:"location_depot_row"`
	LocationDepotTiers   []LocationDepotTierModel `gorm:"foreignKey:LocationDepotCabinetUUID;references:UUID;COMMENT:相关仓储柜架层;" json:"location_depot_tiers"`
}

// TableName 表名称
//  @receiver cls
//  @return string
func (cls LocationDepotCabinetModel) TableName() string {
	return "location_depot_cabinets"
}
