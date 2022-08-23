package models

// PositionDepotSectionModel 仓储仓库区域模型
type PositionDepotSectionModel struct {
	BaseModel
	UniqueCode                  string                       `gorm:"type:CHAR(6);COMMENT:仓储仓库区域代码;" json:"unique_code"`
	Name                        string                       `gorm:"type:VARCHAR(64);仓储仓库区域名称;" json:"name"`
	PositionDepotStorehouseUUID string                       `gorm:"type:CHAR(36);COMMENT:仓储仓库UUID;" json:"position_depot_storehouse_uuid"`
	PositionDepotStorehouse     PositionDepotStorehouseModel `gorm:"foreignKey:PositionDepotStorehouseUUID;references:UUID;COMMENT:所属仓储仓库;" json:"position_depot_storehouse"`
	PositionDepotRows           []PositionDepotRowModel      `gorm:"foreignKey:PositionDepotSectionUUID;references:UUID;COMMENT:相关仓储仓库排;" json:"position_depot_rows"`
}

// TableName 表名称
//  @receiver cls
//  @return string
func (PositionDepotSectionModel) TableName() string {
	return "position_depot_sections"
}
