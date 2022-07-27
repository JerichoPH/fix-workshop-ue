package models

// LocationDepotTierModel 仓储柜架层模型
type LocationDepotTierModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(12);UNIQUE;NOT NULL;COMMENT:仓储柜架层代码;" json:"unique_code"`
	Name                     string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:仓储柜架层名称;" json:"name"`
	LocationDepotCabinetUUID string                    `gorm:"CHAR(36);NOT NULL;COMMENT:仓储柜架UUID;" json:"location_depot_row_uuid"`
	LocationDepotCabinet     LocationDepotCabinetModel `gorm:"foreignKey:LocationDepotCabinetUUID;references:UUID;COMMENT:所属仓储柜架;" json:"location_depot_cabinet"`
	LocationDepotCells       []LocationDepotCellModel  `gorm:"foreignKey:LocationDepotTierUUID;references:UUID;COMMENT:相关仓储柜架格位;" json:"location_depot_cells"`
}

// TableName 表名称
//  @receiver cls
//  @return string
func (LocationDepotTierModel) TableName() string {
	return "location_depot_tiers"
}
