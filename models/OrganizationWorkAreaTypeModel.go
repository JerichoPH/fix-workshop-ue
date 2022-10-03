package models

import (
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
)

type OrganizationWorkAreaTypeModel struct {
	BaseModel
	UniqueCode            string                      `gorm:"type:VARCHAR(64);COMMENT:工区类型代码;" json:"unique_code"`
	Name                  string                      `gorm:"type:VARCHAR(64);COMMENT:工区类型名称;" json:"name"`
	OrganizationWorkAreas []OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaTypeUuid;references:Uuid;COMMENT:相关工区;" json:"organization_work_areas"`
}

// TableName 表名称
func (OrganizationWorkAreaTypeModel) TableName() string {
	return "organization_work_area_types"
}

// FindOneByUUID 根据UUID获取单条数据
//  @receiver ins
//  @param uuid
//  @return OrganizationWorkAreaTypeModel
func (ins OrganizationWorkAreaTypeModel) FindOneByUUID(uuid string) OrganizationWorkAreaTypeModel {
	if ret := BootByModel(ins).SetWheres(tools.Map{"uuid": uuid}).Prepare("").First(&ins); ret.Error != nil {
		panic(wrongs.PanicWhenIsEmpty(ret, "工区类型"))
	}

	return ins
}
