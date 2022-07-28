package models

// LocationDepotRowTypeModel 仓储仓库排类型模型
type LocationDepotRowTypeModel struct {
	BaseModel
	UniqueCode        string                  `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:所属仓储仓库排类型代码;" json:"unique_code"`
	Name              string                  `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:所属仓储仓库排类型名称;" json:"name"`
	LocationDepotRows []LocationDepotRowModel `gorm:"foreignKey:LocationDepotRowTypeUUID;references:UUID;COMMENT:相关仓储仓库排;" json:"location_depot_rows"`
}

// TableName 表名称
//  @receiver LocationDepotRowTypeModel
//  @return string
func (LocationDepotRowTypeModel) TableName() string {
	return "location_depot_row_types"
}
