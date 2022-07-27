package models

// LocationDepotSectionModel 仓储仓库区域模型
type LocationDepotSectionModel struct {
	BaseModel
	UniqueCode                  string                       `gorm:"type:CHAR(6);UNIQUE;NOT NULL;COMMENT:仓储仓库区域代码;" json:"unique_code"`
	Name                        string                       `gorm:"type:VARCHAR(64);NOT NULL;仓储仓库区域名称;" json:"name"`
	LocationDepotStorehouseUUID string                       `gorm:"type:CHAR(36);NOT NULL;COMMENT:仓储仓库UUID;" json:"location_depot_storehouse_uuid"`
	LocationDepotStorehouse     LocationDepotStorehouseModel `gorm:"foreignKey:LocationDepotStorehouseUUID;references:UUID;COMMENT:所属仓储仓库;" json:"location_depot_storehouse"`
	LocationDepotRows           []LocationDepotRowModel      `gorm:"foreignKey:LocationDepotSectionUUID;references:UUID;COMMENT:相关仓储仓库排;" json:"location_depot_rows"`
}

// TableName 表名称
//  @receiver cls
//  @return string
func (LocationDepotSectionModel) TableName() string {
	return "location_depot_sections"
}
