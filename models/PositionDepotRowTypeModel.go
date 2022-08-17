package models

// PositionDepotRowTypeModel 仓储仓库排类型模型
type PositionDepotRowTypeModel struct {
	BaseModel
	UniqueCode        string                  `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:所属仓储仓库排类型代码;" json:"unique_code"`
	Name              string                  `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:所属仓储仓库排类型名称;" json:"name"`
	PositionDepotRows []PositionDepotRowModel `gorm:"foreignKey:PositionDepotRowTypeUUID;references:UUID;COMMENT:相关仓储仓库排;" json:"position_depot_rows"`
}

// TableName 表名称
//  @receiver PositionDepotRowTypeModel
//  @return string
func (PositionDepotRowTypeModel) TableName() string {
	return "position_depot_row_types"
}
