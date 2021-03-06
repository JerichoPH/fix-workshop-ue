package models

type LocationWarehousePosition struct {
	BaseModel
	Preloads                        []string
	Selects                         []string
	Omits                           []string
	UniqueCode                      string                `gorm:"type:CHAR(18);UNIQUE;NOT NULL;COMMENT:位代码;" json:"unique_code"`
	Name                            string                `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:位名称;" json:"name"`
	LocationWarehouseTierUniqueCode string                `gorm:"type:CHAR(10);COMMENT:所属层代码;" json:"location_warehouse_tier_unique_code"`
	LocationWarehouseTier           LocationWarehouseTier `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehouseTierUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属层;" json:"location_warehouse_tier"`
	EntireInstances                 []EntireInstance      `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehousePositionUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}
