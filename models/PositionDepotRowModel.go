package models

// PositionDepotRowModel 仓储仓库排
type PositionDepotRowModel struct {
	BaseModel
	UniqueCode               string                      `gorm:"type:CHAR(8);COMMENT:仓储仓库排代码;" json:"unique_code"`
	Name                     string                      `gorm:"type:VARCHAR(64);COMMENT:仓储仓库排名称;" json:"name"`
	PositionDepotRowTypeUuid string                      `gorm:"type:VARCHAR(36);COMMENT:所属仓储仓库排类型UUID;" json:"position_depot_row_type_uuid"`
	PositionDepotRowType     PositionDepotRowTypeModel   `gorm:"foreignKey:PositionDepotRowTypeUuid;references:Uuid;COMMENT:所属仓储仓库排类型;" json:"position_depot_row_type"`
	PositionDepotSectionUuid string                      `gorm:"type:VARCHAR(36);COMMENT:所属仓储仓库区域UUID;" json:"position_depot_section_uuid"`
	PositionDepotSection     PositionDepotSectionModel   `gorm:"foreignKey:PositionDepotSectionUuid;references:Uuid;COMMENT:所属仓储仓库区域;" json:"position_depot_section"`
	PositionDepotCabinets    []PositionDepotCabinetModel `gorm:"foreignKey:PositionDepotRowUuid;references:Uuid;COMMENT:相关仓储柜架;" json:"position_depot_cabinet"`
}

// TableName 表名称
//  @receiver ins
//  @return string
func (PositionDepotRowModel) TableName() string {
	return "position_depot_rows"
}
