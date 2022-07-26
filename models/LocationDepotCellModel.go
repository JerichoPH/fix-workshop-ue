package models

// LocationDepotCellModel 仓储柜架格位模型
type LocationDepotCellModel struct {
	BaseModel
	UniqueCode            string                 `gorm:"type:CHAR(14);UNIQUE;NOT NULL;COMMENT:仓储柜架格位代码;" json:"unique_code"`
	Name                  string                 `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:仓储柜架格位名称;" json:"name"`
	LocationDepotTierUUID string                 `gorm:"type:CHAR(36);NOT NULL;COMMENT:仓储柜架层UUID;" json:"location_depot_tier_uuid"`
	LocationDepotTier     LocationDepotTierModel `gorm:"foreignKey:LocationDepotTierUUID;references:UUID;COMMENT:所属仓储柜架层;" json:"location_depot_tier"`
}

// TableName 表名称
//  @receiver LocationDepotCellModel
//  @return string
func (LocationDepotCellModel) TableName() string {
	return "location_depot_cells"
}
