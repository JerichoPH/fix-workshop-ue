package models

// OrganizationWorkArea 工区
type OrganizationWorkArea struct {
	BaseModel
	UniqueCode                     string                `gorm:"type:CHAR(8);UNIQUE;NOT NULL;COMMENT:工区代码;" json:"unique_code"` //B049D001
	Name                           string                `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:工区名称;" json:"name"`
	BeEnable                       bool                  `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUniqueCode string                `gorm:"type:CHAR(7);NOT NULL;COMMENT:所属车间;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop           OrganizationWorkshop  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationStations           []OrganizationStation `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkAreaUniqueCode;references:UniqueCode;COMMENT:相关站场;" json:"organization_stations"`
	EntireInstances                []EntireInstance      `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkAreaUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *OrganizationWorkArea) FindOneByUniqueCode(uniqueCode string) (organizationWorkArea OrganizationWorkArea) {
	cls.Boot().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&organizationWorkArea)

	return
}
