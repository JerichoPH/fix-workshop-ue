package models

// LocationDepotStorehouseModel 仓储仓库模型
type LocationDepotStorehouseModel struct {
	BaseModel
	UniqueCode            string                      `gorm:"type:CHAR(4);UNIQUE;NOT NULL;COMMENT:仓储仓库代码;" json:"unique_code"`
	Name                  string                      `gorm:"type:VARCHAR(36);NOT NULL;COMMENT:仓储仓库名称;" json:"name"`
	LocationDepotSections []LocationDepotSectionModel `gorm:"foreignKey:LocationDepotStorehouseUUID;references:UUID;COMMENT:相关仓储仓库区域;" json:"location_depot_sections"`
}

// TableName 表名称
//  @receiver cls
//  @return strung
func (cls LocationDepotStorehouseModel) TableName() string {
	return "location_depot_storehouses"
}
