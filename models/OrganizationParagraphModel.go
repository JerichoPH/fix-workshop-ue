package models

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/tools"
	"gorm.io/gorm"
)

type OrganizationParagraphModel struct {
	BaseModel
	UniqueCode              string                      `gorm:"type:CHAR(4);UNIQUE;NOT NULL;COMMENT:站段代码;" json:"unique_code"` // B049
	Name                    string                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:站段名称;" json:"name"`
	ShortName               string                      `gorm:"type:VARCHAR(64);COMMENT:站段简称;" json:"short_name"`
	BeEnable                bool                        `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationRailwayUUID string                      `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属路局uuid;" json:"organization_railway_uuid"`
	OrganizationRailway     OrganizationRailwayModel    `gorm:"foreignKey:OrganizationRailwayUUID;references:UUID;COMMENT:所属路局;" json:"organization_railway"`
	OrganizationWorkshops   []OrganizationWorkshopModel `gorm:"foreignKey:OrganizationParagraphUniqueCode;references:UniqueCode;COMMENT:相关车间;" json:"organization_workshops"`
	OrganizationLines       []*OrganizationLineModel    `gorm:"many2many:pivot_organization_line_and_organization_paragraphs;foreignKey:id;joinForeignKey:organization_paragraph_id;references:id;joinReferences:organization_line_id;COMMENT:线别与站段多对多;" json:"organization_lines"`
	//EntireInstances         []EntireInstanceModel       `gorm:"foreignKey:OrganizationParagraphUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *OrganizationParagraphModel) TableName() string {
	return "organization_paragraphs"
}

// ScopeBeEnable 获取启用的数据
func (cls *OrganizationParagraphModel) ScopeBeEnable(db *gorm.DB) *gorm.DB {
	return db.Where("be_enable = ?", 1)
}

// FindOneByUUID 根据UUID获取单条数据
//  @receiver cls
//  @param uuid
//  @return OrganizationParagraphModel
func (cls OrganizationParagraphModel) FindOneByUUID(uuid string) OrganizationParagraphModel {
	if ret := Init(cls).SetWheres(tools.Map{"uuid": uuid}).Prepare().First(&cls); ret.Error != nil {
		panic(exceptions.ThrowWhenIsEmptyByDB(ret, "站段"))
	}

	return cls
}
