package models

// OrganizationSection 区间
type OrganizationSection struct {
	BaseModel
	UniqueCode                     string               `gorm:"type:CHAR(6);UNIQUE;NOT NULL;COMMENT:区间代码;" json:"unique_code"` // H07675
	Name                           string               `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:区间名称;" json:"name"`
	BeEnable                       bool                 `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUniqueCode string               `gorm:"type:CHAR(7);NOT NULL;COMMENT:所属车间;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop           OrganizationWorkshop `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:所属车间;" json:"organization_workshop"`
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *OrganizationSection) FindOneByUniqueCode(uniqueCode string) (organizationSection OrganizationSection) {
	cls.Boot().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&organizationSection)

	return
}
