package models

type LocationWarehouseShelf struct {
	BaseModel
	Preloads                           []string
	Selects                            []string
	Omits                              []string
	UniqueCode                         string                   `gorm:"type:CHAR(14);UNIQUE;NOT NULL;COMMENT:柜架代码;" json:"unique_code"`
	Name                               string                   `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:柜架名称;" json:"name"`
	LocationWarehousePlatoonUniqueCode string                   `gorm:"type:CHAR(10);COMMENT:所属排代码;" json:"location_warehouse_platoon_unique_code"`
	LocationWarehousePlatoon           LocationWarehousePlatoon `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehousePlatoonUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属排;" json:"location_warehouse_platoon"`
	LocationWarehouseTiers             []LocationWarehouseTier  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehouseShelfUniqueCode;references:UniqueCode;NOT NULL;COMMENT:相关层;" json:"location_warehouse_tiers"`
}
