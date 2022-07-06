package models

type LocationWarehouseAreaModel struct {
	BaseModel
	UniqueCode                            string                           `gorm:"type:CHAR(10);UNIQUE;NOT NULL;COMMENT:区代码;" json:"unique_code"`
	Name                                  string                           `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:区名称;" json:"name"`
	LocationWarehouseStorehouseUniqueCode string                           `gorm:"type:CHAR(8);COMMENT:所属仓库代码;" json:"location_warehouse_storehouse_unique_code"`
	LocationWarehouseStorehouse           LocationWarehouseStorehouseModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehouseStorehouseUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属仓库;" json:"location_warehouse_storehouse"`
	LocationWarehousePlatoons             []LocationWarehousePlatoonModel  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehouseAreaUniqueCode;references:UniqueCode;NOT NULL;COMMENT:相关排;" json:"location_warehouse_platoons"`
}

// TableName 表名称
func (cls *LocationWarehouseAreaModel) TableName() string {
	return "location_warehouse_areas"
}
