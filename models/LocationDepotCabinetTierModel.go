package models

// LocationDepotCabinetTierModel 库房柜架层
type LocationDepotCabinetTierModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(16);UNIQUE;NOT NULL;COMMENT:层代码;" json:"unique_code"`
	Name                     string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:层名称;" json:"name"`
	LocationDepotCabinetUUID string                    `gorm:"type:CHAR(36);COMMENT:所属柜架代码;" json:"location_depot_cabinet_uuid"`
	LocationDepotCabinet     LocationDepotCabinetModel `gorm:"foreignKey:LocationDepotCabinetUUID;references:UUID;NOT NULL;COMMENT:所属柜架;" json:"location_depot_cabinet"`
	LocationDepotCells       []LocationDepotCellsModel `gorm:"foreignKey:LocationDepotCabinetTierUUID;references:UUID;NOT NULL;COMMENT:相关位;" json:"location_depot_cells"`
}

// TableName 表名称
func (cls *LocationDepotCabinetTierModel) TableName() string {
	return "location_depot_tiers"
}
