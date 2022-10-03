package models

// PositionDepotTierModel 仓储柜架层模型
type PositionDepotTierModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(12);COMMENT:仓储柜架层代码;" json:"unique_code"`
	Name                     string                    `gorm:"type:VARCHAR(64);COMMENT:仓储柜架层名称;" json:"name"`
	PositionDepotCabinetUuid string                    `gorm:"VARCHAR(36);COMMENT:仓储柜架UUID;" json:"position_depot_row_uuid"`
	PositionDepotCabinet     PositionDepotCabinetModel `gorm:"foreignKey:PositionDepotCabinetUuid;references:Uuid;COMMENT:所属仓储柜架;" json:"position_depot_cabinet"`
	PositionDepotCells       []PositionDepotCellModel  `gorm:"foreignKey:PositionDepotTierUuid;references:Uuid;COMMENT:相关仓储柜架格位;" json:"position_depot_cells"`
}

// TableName 表名称
//  @receiver ins
//  @return string
func (PositionDepotTierModel) TableName() string {
	return "position_depot_tiers"
}
