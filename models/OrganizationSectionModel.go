package models

// OrganizationSectionModel 区间
type OrganizationSectionModel struct {
	BaseModel
	UniqueCode                     string                    `gorm:"type:CHAR(6);NOT NULL;COMMENT:区间代码;" json:"unique_code"` // H07675
	Name                           string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:区间名称;" json:"name"`
	BeEnable                       bool                      `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUniqueCode string                    `gorm:"type:CHAR(7);NOT NULL;COMMENT:所属车间;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop           OrganizationWorkshopModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:所属车间;" json:"organization_workshop"`
}

// TableName 表名称
func (cls *OrganizationSectionModel) TableName() string {
	return "organization_sections"
}