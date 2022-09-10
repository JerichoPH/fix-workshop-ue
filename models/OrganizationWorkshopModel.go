package models

// OrganizationWorkshopModel 车间
type OrganizationWorkshopModel struct {
	BaseModel
	UniqueCode                      string                            `gorm:"type:CHAR(7);COMMENT:车间代码;" json:"unique_code"`
	Name                            string                            `gorm:"type:VARCHAR(64);COMMENT:车间名称;" json:"name"`
	BeEnable                        bool                              `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopTypeUuid    string                            `gorm:"type:VARCHAR(36);COMMENT:车间类型UUID;" json:"organization_workshop_type_uuid"`
	OrganizationWorkshopType        OrganizationWorkshopTypeModel     `gorm:"foreignKey:OrganizationWorkshopTypeUuid;references:Uuid;COMMENT:所属类型;" json:"organization_workshop_type"`
	OrganizationParagraphUuid       string                            `gorm:"type:VARCHAR(36);COMMENT:所属站段UUID;" json:"organization_paragraph_uuid"`
	OrganizationParagraph           OrganizationParagraphModel        `gorm:"foreignKey:OrganizationParagraphUuid;references:Uuid;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationWorkAreas           []OrganizationWorkAreaModel       `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:相关工区;" json:"organization_work_areas"`
	LocationSections                []LocationSectionModel            `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:相关区间;" json:"location_sections"`
	LocationStations                []LocationStationModel            `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:相关站场;" json:"location_stations"`
	LocationCenters                 []LocationCenterModel             `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:相关中心;" json:"location_centers"`
	LocationRailroadGradeCrossModel []LocationRailroadGradeCrossModel `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:相关道口;" json:"location_railroad_grade_crosses"`
	LocationDepotStorehouses        []PositionDepotStorehouseModel    `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:相关仓储库房;" json:"location_depot_storehouses"`
}

// TableName 表名称
//  @receiver cls
//  @return string
func (cls *OrganizationWorkshopModel) TableName() string {
	return "organization_workshops"
}
