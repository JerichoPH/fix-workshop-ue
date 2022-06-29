package models

type LocationWarehouseTier struct {
	BaseModel
	UniqueCode                       string                      `gorm:"type:CHAR(16);UNIQUE;NOT NULL;COMMENT:层代码;" json:"unique_code"`
	Name                             string                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:层名称;" json:"name"`
	LocationWarehouseShelfUniqueCode string                      `gorm:"type:CHAR(10);COMMENT:所属柜架代码;" json:"location_warehouse_shelf_unique_code"`
	LocationWarehouseShelf           LocationWarehouseShelf      `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehouseShelfUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属柜架;" json:"location_warehouse_shelf"`
	LocationWarehousePositions       []LocationWarehousePosition `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehouseTierUniqueCode;references:UniqueCode;NOT NULL;COMMENT:相关位;" json:"location_warehouse_positions"`
}
