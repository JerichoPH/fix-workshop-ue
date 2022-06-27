package models

// 排模型
type LocationWarehousePlatoon struct {
	BaseModel
	Preloads                        []string
	Selects                         []string
	Omits                           []string
	UniqueCode                      string                   `gorm:"type:CHAR(12);UNIQUE;NOT NULL;COMMENT:排代码;" json:"unique_code"`
	Name                            string                   `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:排名称;" json:"name"`
	LocationWarehouseAreaUniqueCode string                   `gorm:"type:CHAR(8);COMMENT:所属区代码;" json:"location_warehouse_area_unique_code"`
	LocationWarehouseArea           LocationWarehouseArea    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehouseAreaUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属区;" json:"location_warehouse_area"`
	LocationWarehouseShelves        []LocationWarehouseShelf `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehousePlatoonUniqueCode;references:UniqueCode;NOT NULL;COMMENT:相关柜架;" json:"location_warehouse_shelves"`
}
