package OrganizationModels

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/models/EntireInstanceModels"
)

// OrganizationRailwayModel 路局
type OrganizationRailwayModel struct {
	models.BaseModel
	UniqueCode             string                                     `gorm:"type:CHAR(3);UNIQUE;NOT NULL;COMMENT:路局代码;" json:"unique_code"` // A12
	Name                   string                                     `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:路局名称;" json:"name"`
	ShortName              string                                     `gorm:"type:VARCHAR(64);COMMENT:路局简称;" json:"short_name"`
	BeEnable               bool                                       `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationLines      []*OrganizationLineModel                   `gorm:"many2many:pivot_organization_line_and_organization_railways;foreignKey:id;joinForeignKey:organization_railway_id;references:id;joinReferences:organization_line_id;COMMENT:线别与站段多对多;" json:"organization_lines"`
	OrganizationParagraphs []OrganizationParagraphModel               `gorm:"foreignKey:OrganizationRailwayUUID;references:UUID;COMMENT:相关站段;" json:"organization_paragraphs"`
	EntireInstances        []EntireInstanceModels.EntireInstanceModel `gorm:"foreignKey:OrganizationRailwayUUID;references:UUID;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *OrganizationRailwayModel) TableName() string {
	return "organization_railways"
}