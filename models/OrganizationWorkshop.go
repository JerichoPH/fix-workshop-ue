package models

// OrganizationWorkshop 车间
type OrganizationWorkshop struct {
	BaseModel
	UniqueCode                         string                   `gorm:"type:CHAR(7);UNIQUE;NOT NULL;COMMENT:车间代码;" json:"unique_code"`
	Name                               string                   `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:车间名称;" json:"name"`
	BeEnable                           bool                     `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopTypeUniqueCode string                   `gorm:"type:VARCHAR(64);COMMENT:车间类型;" json:"organization_workshop_type_unique_code"`
	OrganizationWorkshopType           OrganizationWorkshopType `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopTypeUniqueCode;references:UniqueCode;COMMENT:所属类型;" json:"organization_workshop_type"`
	OrganizationParagraphUniqueCode    string                   `gorm:"type:CHAR(4);COMMENT:所属站段;" json:"organization_paragraph_unique_code"`
	OrganizationParagraph              OrganizationParagraph    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationParagraphUniqueCode;references:UniqueCode;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationSections               []OrganizationSection    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:相关区间;" json:"organization_sections"`
	OrganizationStations               []OrganizationStation    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:相关站场;" json:"organization_stations"`
	EntireInstances                    []EntireInstance         `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:所属器材;" json:"entire_instances"`
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *OrganizationWorkshop) FindOneByUniqueCode(uniqueCode string) (organizationWorkshop OrganizationWorkshop) {
	cls.Boot().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&organizationWorkshop)

	return
}
