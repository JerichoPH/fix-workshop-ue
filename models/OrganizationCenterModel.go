package models

import "gorm.io/gorm"

type OrganizationCenterModel struct {
	BaseModel
	UniqueCode                     string                    `gorm:"type:CHAR(6);UNIQUE;NOT NULL;COMMENT:中心代码;"` // A12F01
	Name                           string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:中心名称;"`
	BeEnable                       bool                      `gorm:"type:BOOLEAN;NOT NULL;DEFAULT:0;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUniqueCode string                    `gorm:"type:CHAR(7);NOT NULL;COMMENT:所属车间;"`
	OrganizationWorkshop           OrganizationWorkshopModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属车间;"`
}

// TableName 表名称
func (cls *OrganizationCenterModel) TableName() string {
	return "organization_centers"
}

// ScopeBeEnable 获取启用的数据
func (cls *OrganizationCenterModel) ScopeBeEnable(db *gorm.DB) *gorm.DB {
	return db.Where("be_enable is ?", true)
}