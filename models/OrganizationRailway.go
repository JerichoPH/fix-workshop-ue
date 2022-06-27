package models

import (
	"gorm.io/gorm"
)

// OrganizationRailway 路局
type OrganizationRailway struct {
	BaseModel
	Preloads               []string
	Selects                []string
	Omits                  []string
	UniqueCode             string                  `gorm:"type:CHAR(3);UNIQUE;NOT NULL;COMMENT:路局代码;" json:"unique_code"` // A12
	Name                   string                  `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:路局名称;" json:"name"`
	ShotName               string                  `gorm:"type:VARCHAR(64);COMMENT:路局简称;" json:"shot_name"`
	BeEnable               bool                    `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationParagraphs []OrganizationParagraph `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationRailwayUniqueCode;references:UniqueCode;COMMENT:相关站段;" json:"organization_paragraphs"`
	EntireInstances        []EntireInstance        `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationRailwayUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *OrganizationRailway) FindOneByUniqueCode(db *gorm.DB, uniqueCode string) (organizationRailway OrganizationRailway) {
	cls.Boot(db, cls.Preloads, cls.Selects, cls.Omits).
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&organizationRailway)

	return
}
