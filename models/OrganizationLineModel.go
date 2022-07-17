package models

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/tools"
	"gorm.io/gorm"
)

type OrganizationLineModel struct {
	BaseModel
	UniqueCode             string                        `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:线别代码;" json:"unique_code"` // E0001
	Name                   string                        `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:线别名称;" json:"name"`
	BeEnable               bool                          `gorm:"type:BOOLEAN;NOT NULL;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationRailways   []*OrganizationRailwayModel   `gorm:"many2many:pivot_organization_line_and_organization_railways;foreignKey:id;joinForeignKey:organization_line_id;references:id;joinReferences:organization_railway_id;COMMENT:线别与路局多对多;" json:"organization_railways"`
	OrganizationParagraphs []*OrganizationParagraphModel `gorm:"many2many:pivot_organization_line_and_organization_paragraphs;foreignKey:id;joinForeignKey:organization_line_id;references:id;joinReferences:organization_paragraph_id;COMMENT:线别与站段多对多;" json:"organization_paragraphs"`
	OrganizationStations   []*OrganizationStationModel   `gorm:"many2many:pivot_organization_line_and_organization_stations;foreignKey:id;joinForeignKey:organization_line_id;references:id;joinReferences:organization_station_id;COMMENT:线别与车站多对多;" json:"organization_stations"`
}

// TableName 表名称
func (cls *OrganizationLineModel) TableName() string {
	return "organization_lines"
}

// ScopeBeEnable 获取启用的数据
func (cls *OrganizationLineModel) ScopeBeEnable(db *gorm.DB) *gorm.DB {
	return db.Where("be_enable = ?", 1)
}

// FindOneByUUID 根据UUID获取单条数据
//  @receiver cls
//  @param uuid
//  @return OrganizationLineModel
func (cls OrganizationLineModel) FindOneByUUID(uuid string) OrganizationLineModel {
	if ret := Init(cls).SetWheres(tools.Map{"uuid": uuid}).Prepare().First(&cls); ret.Error != nil {
		panic(abnormals.PanicWhenIsEmpty(ret, "线别"))
	}

	return cls
}
