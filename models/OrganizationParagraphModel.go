package models

import "gorm.io/gorm"

type OrganizationParagraphModel struct {
	BaseModel
	UniqueCode                    string                      `gorm:"type:CHAR(4);UNIQUE;NOT NULL;COMMENT:站段代码;" json:"unique_code"` // B049
	Name                          string                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:站段名称;" json:"name"`
	ShotName                      string                      `gorm:"type:VARCHAR(64);COMMENT:站段简称;" json:"shot_name"`
	BeEnable                      bool                        `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationRailwayUniqueCode string                      `gorm:"type:CHAR(3);COMMENT:所属路局;" json:"organization_railway_unique_code"`
	OrganizationRailway           OrganizationRailwayModel    `gorm:"foreignKey:OrganizationRailwayUniqueCode;references:UniqueCode;COMMENT:所属路局;" json:"organization_railway"`
	OrganizationWorkshops         []OrganizationWorkshopModel `gorm:"foreignKey:OrganizationParagraphUniqueCode;references:UniqueCode;COMMENT:相关车间;" json:"organization_workshops"`
	OrganizationLines []*OrganizationLineModel `gorm:"many2many:pivot_organization_line_and_organization_paragraphs;foreignKey:id;joinForeignKey:organization_paragraph_id;references:id;joinReferences:organization_line_id;COMMENT:线别与站段多对多;" json:"organization_lines"`
	EntireInstances               []EntireInstanceModel       `gorm:"foreignKey:OrganizationParagraphUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *OrganizationParagraphModel) TableName() string {
	return "organization_paragraphs"
}

// ScopeBeEnable 获取启用的数据
func (cls *OrganizationParagraphModel) ScopeBeEnable(db *gorm.DB) *gorm.DB {
	return db.Where("be_enable = ?", 1)
}