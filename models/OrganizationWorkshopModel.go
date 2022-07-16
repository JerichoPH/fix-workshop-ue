package models

import "gorm.io/gorm"

// OrganizationWorkshopModel 车间
type OrganizationWorkshopModel struct {
	BaseModel
	UniqueCode                   string                        `gorm:"type:CHAR(7);UNIQUE;NOT NULL;COMMENT:车间代码;" json:"unique_code"`
	Name                         string                        `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:车间名称;" json:"name"`
	BeEnable                     bool                          `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopTypeUUID string                        `gorm:"type:CHAR(36);COMMENT:车间类型UUID;" json:"organization_workshop_type_uuid"`
	OrganizationWorkshopType     OrganizationWorkshopTypeModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopTypeUUID;references:UUID;COMMENT:所属类型;" json:"organization_workshop_type"`
	OrganizationParagraphUUID    string                        `gorm:"type:CHAR(36);COMMENT:所属站段UUID;" json:"organization_paragraph_uuid"`
	OrganizationParagraph        OrganizationParagraphModel    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationParagraphUUID;references:UUID;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationSections         []OrganizationSectionModel    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:相关区间;" json:"organization_sections"`
	OrganizationWorkAreas        []OrganizationWorkAreaModel   `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:相关工区;" json:"organization_work_areas"`
	OrganizationStations         []OrganizationStationModel    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:相关站场;" json:"organization_stations"`
	//EntireInstances              []EntireInstanceModel         `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:所属器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *OrganizationWorkshopModel) TableName() string {
	return "organization_workshops"
}

// ScopeBeEnable 获取启用的数据
func (cls *OrganizationWorkshopModel) ScopeBeEnable(db *gorm.DB) *gorm.DB {
	return db.Where("be_enable = ?", 1)
}
