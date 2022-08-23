package models

// PositionDepotStorehouseModel 仓储仓库模型
type PositionDepotStorehouseModel struct {
	BaseModel
	UniqueCode               string                      `gorm:"type:CHAR(4);COMMENT:仓储库房代码;" json:"unique_code"`
	Name                     string                      `gorm:"type:VARCHAR(36);COMMENT:仓储库房名称;" json:"name"`
	OrganizationWorkshopUUID string                      `gorm:"type:CHAR(36);COMMENT:所属车间UUID;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     OrganizationWorkshopModel   `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属车间;" json:"organization_workshop"`
	PositionDepotSections    []PositionDepotSectionModel `gorm:"foreignKey:PositionDepotStorehouseUUID;references:UUID;COMMENT:相关仓储仓库区域;" json:"position_depot_sections"`
}

// TableName 表名称
//  @receiver cls
//  @return strung
func (PositionDepotStorehouseModel) TableName() string {
	return "position_depot_storehouses"
}
