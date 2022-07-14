package models

import "gorm.io/gorm"

// OrganizationWorkshopModel 车间
type OrganizationWorkshopModel struct {
	BaseModel
	UniqueCode                         string                        `gorm:"type:CHAR(7);UNIQUE;NOT NULL;COMMENT:车间代码;" json:"unique_code"`
	Name                               string                        `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:车间名称;" json:"name"`
	BeEnable                           bool                          `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopTypeUniqueCode string                        `gorm:"type:VARCHAR(64);COMMENT:车间类型;" json:"organization_workshop_type_unique_code"`
	OrganizationWorkshopType           OrganizationWorkshopTypeModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopTypeUniqueCode;references:UniqueCode;COMMENT:所属类型;" json:"organization_workshop_type"`
	OrganizationParagraphUniqueCode    string                        `gorm:"type:CHAR(4);COMMENT:所属站段;" json:"organization_paragraph_unique_code"`
	OrganizationParagraph              OrganizationParagraphModel    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationParagraphUniqueCode;references:UniqueCode;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationSections               []OrganizationSectionModel    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:相关区间;" json:"organization_sections"`
	OrganizationWorkAreas              []OrganizationWorkAreaModel   `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:相关工区;" json:"organization_work_areas"`
	OrganizationStations               []OrganizationStationModel    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:相关站场;" json:"organization_stations"`
	EntireInstances                    []EntireInstanceModel         `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:所属器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *OrganizationWorkshopModel) TableName() string {
	return "organization_workshops"
}

// ScopeBeEnable 获取启用的数据
func (cls *OrganizationWorkshopModel) ScopeBeEnable(db *gorm.DB) *gorm.DB {
	return db.Where("be_enable is ?", true)
}
