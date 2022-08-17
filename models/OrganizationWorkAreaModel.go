package models

// OrganizationWorkAreaModel 工区
type OrganizationWorkAreaModel struct {
	BaseModel
	UniqueCode                         string                              `gorm:"type:CHAR(8);NOT NULL;COMMENT:工区代码;" json:"unique_code"` //B049D001
	Name                               string                              `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:工区名称;" json:"name"`
	BeEnable                           bool                                `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkAreaTypeUUID       string                              `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属工区类型UUID;" json:"organization_work_area_type_uuid"`
	OrganizationWorkAreaType           OrganizationWorkAreaTypeModel       `gorm:"foreignKey:OrganizationWorkAreaTypeUUID;references:UUID;COMMENT:所属工区类型;" json:"organization_work_area_type"`
	OrganizationWorkAreaProfessionUUID string                              `gorm:"type:CHAR(36);NOT NULL;DEFAULT:;COMMENT:所属工区专业UUID;" json:"organization_work_area_profession_uuid"`
	OrganizationWorkAreaProfession     OrganizationWorkAreaProfessionModel `gorm:"foreignKey:OrganizationWorkAreaProfessionUUID;references:UUID;COMMENT:所属工区类型;" json:"organization_work_area_profession"`
	OrganizationWorkshopUUID           string                              `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属车间uuid;" json:"organization_workshop_uuid"`
	OrganizationWorkshop               OrganizationWorkshopModel           `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属车间;" json:"organization_workshop"`
	LocationSections                   []LocationSectionModel              `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:相关区间;" json:"organization_sections"`
	LocationCenters                    []LocationCenterModel               `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:相关中心;" json:"organization_centers"`
	LocationRailroadGradeCrossModel    []LocationRailroadGradeCrossModel   `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:相关道口;" json:"organization_railroad_grade_crosses"`
	LocationStations                   []LocationStationModel              `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:相关站场;" json:"organization_stations"`
}

// TableName 表名称
func (OrganizationWorkAreaModel) TableName() string {
	return "organization_work_areas"
}
