package models

import (
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
)

type OrganizationWorkAreaTypeModel struct {
	BaseModel
	UniqueCode            string                      `gorm:"type:VARCHAR(64);COMMENT:工区类型代码;" json:"unique_code"`
	Name                  string                      `gorm:"type:VARCHAR(64);COMMENT:工区类型名称;" json:"name"`
	OrganizationWorkAreas []OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaTypeUUID;references:UUID;COMMENT:相关工区;" json:"organization_work_areas"`
}

// TableName 表名称
func (OrganizationWorkAreaTypeModel) TableName() string {
	return "organization_work_area_types"
}

// FindOneByUUID 根据UUID获取单条数据
//  @receiver cls
//  @param uuid
//  @return OrganizationWorkAreaTypeModel
func (cls OrganizationWorkAreaTypeModel) FindOneByUUID(uuid string) OrganizationWorkAreaTypeModel {
	if ret := Init(cls).SetWheres(tools.Map{"uuid": uuid}).Prepare("").First(&cls); ret.Error != nil {
		panic(wrongs.PanicWhenIsEmpty(ret, "工区类型"))
	}

	return cls
}
