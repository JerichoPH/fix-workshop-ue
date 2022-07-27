package models

// LocationDepotStorehouseModel 仓储仓库模型
type LocationDepotStorehouseModel struct {
	BaseModel
	UniqueCode               string                      `gorm:"type:CHAR(4);UNIQUE;NOT NULL;COMMENT:仓储库房代码;" json:"unique_code"`
	Name                     string                      `gorm:"type:VARCHAR(36);NOT NULL;COMMENT:仓储库房名称;" json:"name"`
	OrganizationWorkshopUUID string                      `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属车间UUID;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     OrganizationWorkshopModel   `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属车间;" json:"organization_workshop"`
	LocationDepotSections    []LocationDepotSectionModel `gorm:"foreignKey:LocationDepotStorehouseUUID;references:UUID;COMMENT:相关仓储仓库区域;" json:"location_depot_sections"`
}

// TableName 表名称
//  @receiver cls
//  @return strung
func (LocationDepotStorehouseModel) TableName() string {
	return "location_depot_storehouses"
}
