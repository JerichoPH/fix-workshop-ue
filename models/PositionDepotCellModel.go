package models

// PositionDepotCellModel 仓储柜架格位模型
type PositionDepotCellModel struct {
	BaseModel
	UniqueCode            string                 `gorm:"type:CHAR(14);UNIQUE;NOT NULL;COMMENT:仓储柜架格位代码;" json:"unique_code"`
	Name                  string                 `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:仓储柜架格位名称;" json:"name"`
	PositionDepotTierUUID string                 `gorm:"type:CHAR(36);NOT NULL;COMMENT:仓储柜架层UUID;" json:"position_depot_tier_uuid"`
	PositionDepotTier     PositionDepotTierModel `gorm:"foreignKey:PositionDepotTierUUID;references:UUID;COMMENT:所属仓储柜架层;" json:"position_depot_tier"`
}

// TableName 表名称
//  @receiver PositionDepotCellModel
//  @return string
func (PositionDepotCellModel) TableName() string {
	return "position_depot_cells"
}