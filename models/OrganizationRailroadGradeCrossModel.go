package models

// OrganizationRailroadGradeCrossModel 道口
type OrganizationRailroadGradeCrossModel struct {
	BaseModel
	UniqueCode                     string                    `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:道口代码;" json:"unique_code"` // I0100
	Name                           string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:道口名称;" json:"name"`
	BeEnable                       bool                      `gorm:"type:BOOLEAN;NOT NULL;DEFAULT:0;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUniqueCode string                    `gorm:"type:CHAR(7);NOT NULL;COMMENT:车间代码;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop           OrganizationWorkshopModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属车间;" json:"organization_workshop"`
}

// TableName 表名称
func (cls *OrganizationRailroadGradeCrossModel) TableName() string {
	return "organization_railroad_grade_crosses"
}