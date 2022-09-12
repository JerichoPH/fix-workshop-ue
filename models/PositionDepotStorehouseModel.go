package models

// PositionDepotStorehouseModel 仓储仓库模型
type PositionDepotStorehouseModel struct {
	BaseModel
	UniqueCode               string                      `gorm:"type:CHAR(4);COMMENT:仓储库房代码;" json:"unique_code"`
	Name                     string                      `gorm:"type:VARCHAR(36);COMMENT:仓储库房名称;" json:"name"`
	OrganizationWorkshopUuid string                      `gorm:"type:VARCHAR(36);COMMENT:所属车间UUID;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     OrganizationWorkshopModel   `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUuid string                      `gorm:"type:VARCHAR(36);COMMENT:所属工区UUID;" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationWorkAreaModel   `gorm:"foreignKey:OrganizationWorkAreaUuid;references:Uuid;COMMENT:所属工区;" json:"organization_work_area"`
	PositionDepotSections    []PositionDepotSectionModel `gorm:"foreignKey:PositionDepotStorehouseUuid;references:Uuid;COMMENT:相关仓储仓库区域;" json:"position_depot_sections"`
}

// TableName 表名称
//  @receiver cls
//  @return strung
func (PositionDepotStorehouseModel) TableName() string {
	return "position_depot_storehouses"
}
