package models

// 排模型
type LocationDepotRowModel struct {
	BaseModel
	UniqueCode                             string                            `gorm:"type:CHAR(12);UNIQUE;NOT NULL;COMMENT:排代码;" json:"unique_code"`
	Name                                   string                            `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:排名称;" json:"name"`
	LocationWarehouseAreaUniqueCode        string                            `gorm:"type:CHAR(8);COMMENT:所属区代码;" json:"location_warehouse_area_unique_code"`
	LocationWarehouseArea                  LocationDepotSectionModel         `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehouseAreaUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属区;" json:"location_warehouse_area"`
	LocationWarehousePlatoonTypeUniqueCode string                            `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:所属排类型;"`
	LocationWarehousePlatoonType           LocationWarehousePlatoonTypeModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehousePlatoonTypeUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属排类型;" json:"location_warehouse_platoon_type"`
	LocationWarehouseShelves               []LocationWarehouseShelfModel     `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehousePlatoonUniqueCode;references:UniqueCode;NOT NULL;COMMENT:相关柜架;" json:"location_warehouse_shelves"`
}

// TableName 表名称
func (cls *LocationDepotRowModel) TableName() string {
	return "location_warehouse_platoons"
}
