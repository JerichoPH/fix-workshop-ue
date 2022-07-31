package models

// OrganizationWorkAreaModel 工区
type OrganizationWorkAreaModel struct {
	BaseModel
	UniqueCode                          string                                `gorm:"type:CHAR(8);UNIQUE;NOT NULL;COMMENT:工区代码;" json:"unique_code"` //B049D001
	Name                                string                                `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:工区名称;" json:"name"`
	BeEnable                            bool                                  `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkAreaTypeUUID        string                                `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属工区类型uuid;" json:"organization_work_area_type_uuid"`
	OrganizationWorkAreaType            OrganizationWorkAreaTypeModel         `gorm:"foreignKey:OrganizationWorkAreaTypeUUID;references:UUID;COMMENT:所属工区类型;" json:"organization_work_area_type"`
	OrganizationWorkshopUUID            string                                `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属车间uuid;" json:"organization_workshop_uuid"`
	OrganizationWorkshop                OrganizationWorkshopModel             `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationLines                   []*OrganizationLineModel              `gorm:"many2many:pivot_organization_line_and_organization_work_areas;foreignKey:id;joinForeignKey:organization_work_area_id;references:id;joinReferences:organization_line_id;COMMENT:线别与工区多对多;" json:"organization_lines"`
	OrganizationSections                []OrganizationSectionModel            `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:相关区间;" json:"organization_sections"`
	OrganizationCenters                 []OrganizationCenterModel             `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:相关中心;" json:"organization_centers"`
	OrganizationRailroadGradeCrossModel []OrganizationRailroadGradeCrossModel `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:相关道口;" json:"organization_railroad_grade_crosses"`
	OrganizationStations                []OrganizationStationModel            `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:相关站场;" json:"organization_stations"`
	//LocationDepotStorehouse             []LocationDepotStorehouseModel             `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;NOT NULL;COMMENT:相关库房;" json:"location_depot_storehouses"`
	EntireInstances []EntireInstanceModel `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (OrganizationWorkAreaModel) TableName() string {
	return "organization_work_areas"
}
