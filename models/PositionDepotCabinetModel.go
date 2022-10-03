package models

// PositionDepotCabinetModel 仓储柜架模型
type PositionDepotCabinetModel struct {
	BaseModel
	UniqueCode           string                   `gorm:"type:CHAR(10);COMMENT:仓储柜架代码;" json:"unique_code"`
	Name                 string                   `gorm:"type:VARCHAR(64);COMMENT:仓储柜架名称;" json:"name"`
	PositionDepotRowUuid string                   `gorm:"type:VARCHAR(36);COMMENT:仓储柜架排UUID;" json:"position_depot_row_uuid"`
	PositionDepotRow     PositionDepotRowModel    `gorm:"foreignKey:PositionDepotRowUuid;references:Uuid;COMMENT:所属仓储排;" json:"position_depot_row"`
	PositionDepotTiers   []PositionDepotTierModel `gorm:"foreignKey:PositionDepotCabinetUuid;references:Uuid;COMMENT:相关仓储柜架层;" json:"position_depot_tiers"`
}

// TableName 表名称
//  @receiver ins
//  @return string
func (PositionDepotCabinetModel) TableName() string {
	return "position_depot_cabinets"
}
