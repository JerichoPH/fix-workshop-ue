package models

// LocationDepotRowTypeModel 库房排类型
type LocationDepotRowTypeModel struct {
	BaseModel
	UniqueCode            string                  `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:排类型代码;" json:"unique_code"`
	Name                  string                  `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:排类型名称;" json:"name"`
	LocationWarehouseRows []LocationDepotRowModel `gorm:"foreignKey:LocationDepotRowTypeUUID;references:UUID;NOT NULL;COMMENT:相关排;" json:"location_warehouse_rows"`
}

// TableName 表名称
func (cls *LocationDepotRowTypeModel) TableName() string {
	return "location_depot_row_types"
}
