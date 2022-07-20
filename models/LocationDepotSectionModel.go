package models

type LocationDepotSectionModel struct {
	BaseModel
	UniqueCode                  string                       `gorm:"type:CHAR(10);UNIQUE;NOT NULL;COMMENT:区代码;" json:"unique_code"`
	Name                        string                       `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:区名称;" json:"name"`
	LocationDepotStorehouseUUID string                       `gorm:"type:CHAR(36);COMMENT:所属仓库代码;" json:"location_warehouse_storehouse_uuid"`
	LocationDepotStorehouse     LocationDepotStorehouseModel `gorm:"foreignKey:LocationDepotStorehouseUUID;references:UUID;NOT NULL;COMMENT:所属仓库;" json:"location_warehouse_storehouse"`
	LocationWarehouseRows       []LocationDepotRowModel      `gorm:"foreignKey:LocationDepotSectionUUID;references:UUID;NOT NULL;COMMENT:相关排;" json:"location_warehouse_platoons"`
}

// TableName 表名称
func (cls *LocationDepotSectionModel) TableName() string {
	return "location_depot_sections"
}
