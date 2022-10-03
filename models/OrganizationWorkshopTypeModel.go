package models

import (
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
)

// OrganizationWorkshopTypeModel 车间类型
type OrganizationWorkshopTypeModel struct {
	BaseModel
	UniqueCode            string                      `gorm:"type:VARCHAR(64);COMMENT:车间类型代码;" json:"unique_code"`
	Name                  string                      `gorm:"type:VARCHAR(64);COMMENT:车间类型名称;" json:"name"`
	NumberCode            string                      `gorm:"type:VARCHAR(64);COMMENT:车间类型数字代码;" json:"number_code"`
	OrganizationWorkshops []OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopTypeUuid;references:Uuid;" json:"organization_workshops"`
}

// TableName 表名称
func (OrganizationWorkshopTypeModel) TableName() string {
	return "organization_workshop_types"
}

// FindOneByUUID 根据uuid获取单条数据
//  @receiver ins
//  @param uuid
//  @return OrganizationWorkshopTypeModel
func (ins OrganizationWorkshopTypeModel) FindOneByUUID(uuid string) OrganizationWorkshopTypeModel {
	if ret := BootByModel(ins).SetWheres(tools.Map{"uuid": uuid}).Prepare("").First(&ins); ret.Error != nil {
		panic(wrongs.PanicWhenIsEmpty(ret, "车间类型"))
	}

	return ins
}
