package models

type LocationDepotStorehouseModel struct {
	BaseModel
	UniqueCode                string                      `gorm:"type:CHAR(8);UNIQUE;NOT NULL;COMMENT:仓库代码;" json:"unique_code"`
	Name                      string                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:仓库名称;" json:"name"`
	OrganizationParagraphUUID string                      `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属站段代码;" json:"organization_paragraph_uuid"`
	OrganizationParagraph     OrganizationParagraphModel  `gorm:"foreignKey:OrganizationParagraphUUID;references:UUID;NOT NULL;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationWorkshopUUID  string                      `gorm:"type:CHAR(36);COMMENT:所属车间代码;" json:"organization_workshop_uuid"`
	OrganizationWorkshop      OrganizationWorkshopModel   `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;NOT NULL;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUUID  string                      `gorm:"type:CHAR(36);COMMENT:所属工区代码;" json:"organization_work_area_uuid"`
	OrganizationWorkArea      OrganizationWorkAreaModel   `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;NOT NULL;COMMENT:所属工区;" json:"organization_work_area"`
	OrganizationCenterUUID    string                      `gorm:"type:CHAR(36);COMMENT:所属中心代码;" json:"organization_center_uuid"`
	OrganizationCenter        OrganizationCenterModel     `gorm:"foreignKey:OrganizationCenterUUID;references:UUID;NOT NULL;COMMENT:所属中心;" json:"organization_center"`
	OrganizationStationUUID   string                      `gorm:"type:CHAR(36);COMMENT:所属车站代码;" json:"organization_station_uuid"`
	OrganizationStation       OrganizationStationModel    `gorm:"foreignKey:OrganizationStationUUID;references:UUID;NOT NULL;COMMENT:所属站场;" json:"organization_station"`
	LocationDepotSections     []LocationDepotSectionModel `gorm:"foreignKey:LocationDepotStorehouseUUID;references:UniqueCode;NOT NULL;COMMENT:相关区;" json:"location_warehouse_areas"`
}

// TableName 表名称
func (cls *LocationDepotStorehouseModel) TableName() string {
	return "location_depot_storehouses"
}
