package models

type OrganizationCenterModel struct {
	BaseModel
	UniqueCode                     string                    `gorm:"type:CHAR(6);UNIQUE;NOT NULL;COMMENT:中心代码;"` // A12F01
	Name                           string                    `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:中心名称;"`
	OrganizationWorkshopUniqueCode string                    `gorm:"type:CHAR(7);NOT NULL;COMMENT:所属车间;"`
	OrganizationWorkshop           OrganizationWorkshopModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属车间;"`
}

// TableName 表名称
func (cls *OrganizationCenterModel) TableName() string {
	return "organization_centers"
}
