package models

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/tools"
	"gorm.io/gorm"
)

type OrganizationCenterModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(6);UNIQUE;NOT NULL;COMMENT:中心代码;"` // A12F01
	Name                     string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:中心名称;"`
	BeEnable                 bool                      `gorm:"type:BOOLEAN;NOT NULL;DEFAULT:0;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUUID string                    `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属车间UUID;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUUID string                    `gorm:"type:CHAR(36);COMMENT:所属工区UUID;" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:所属工区;" json:"organization_work_area"`
}

// TableName 表名称
func (cls *OrganizationCenterModel) TableName() string {
	return "organization_centers"
}

// ScopeBeEnable 获取启用的数据
func (cls *OrganizationCenterModel) ScopeBeEnable(db *gorm.DB) *gorm.DB {
	return db.Where("be_enable = ?", 1)
}

// FindOneByUUID 根据UUID获取单条数据
//  @receiver cls
//  @param uuid
//  @return OrganizationCenterModel
func (cls OrganizationCenterModel) FindOneByUUID(uuid string) OrganizationCenterModel {
	if ret := Init(cls).SetWheres(tools.Map{"uuid": uuid}).Prepare().First(&cls); ret.Error != nil {
		panic(abnormals.PanicWhenIsEmpty(ret, "中心"))
	}

	return cls
}
