package models

type LocationWarehouseStorehouseModel struct {
	BaseModel
	UniqueCode                               string                              `gorm:"type:CHAR(8);UNIQUE;NOT NULL;COMMENT:仓库代码;" json:"unique_code"`
	Name                                     string                              `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:仓库名称;" json:"name"`
	OrganizationParagraphUniqueCode          string                              `gorm:"type:CHAR(4);NOT NULL;COMMENT:所属站段代码;" json:"organization_paragraph_unique_code"`
	OrganizationParagraph                    OrganizationParagraphModel          `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationParagraphUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationWorkshopUniqueCode           string                              `gorm:"type:CHAR(7);COMMENT:所属车间代码;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop                     OrganizationWorkshopModel           `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUniqueCode           string                              `json:"organization_work_area_unique_code"`
	OrganizationWorkArea                     OrganizationWorkAreaModel           `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkAreaUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属工区;" json:"organization_work_area"`
	OrganizationRailroadGradeCrossUniqueCode string                              `gorm:"type:CHAR(5);COMMENT:所属道口;" json:"organization_railroad_grade_cross_unique_code"`
	OrganizationRailroadGradeCross           OrganizationRailroadGradeCrossModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationRailroadGradeCrossUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属道口;" json:"organization_railroad_grade_cross"`
	OrganizationCenterUniqueCode             string                              `gorm:"type:CHAR(6);COMMENT:所属中心代码;" json:"organization_center_unique_code"`
	OrganizationCenter                       OrganizationCenterModel             `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationCenterUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属中心;" json:"organization_center"`
	OrganizationStationUniqueCode            string                              `gorm:"type:CHAR(6);COMMENT:所属车站代码;" json:"organization_station_unique_code"`
	OrganizationStation                      OrganizationStationModel            `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationStationUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属站场;" json:"organization_station"`
	LocationWarehouseAreas                   []LocationWarehouseAreaModel        `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehouseStorehouseUniqueCode;references:UniqueCode;NOT NULL;COMMENT:相关区;" json:"location_warehouse_areas"`
}

// TableName 表名称
func (cls *LocationWarehouseStorehouseModel) TableName() string {
	return "location_warehouse_storehouses"
}
