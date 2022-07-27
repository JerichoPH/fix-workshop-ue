package models

import (
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
)

// OrganizationWorkshopTypeModel 车间类型
type OrganizationWorkshopTypeModel struct {
	BaseModel
	UniqueCode            string                      `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:车间类型代码;" json:"unique_code"`
	Name                  string                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:车间类型名称;" json:"name"`
	Number                string                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:车间类型数字代码;" json:"number"`
	OrganizationWorkshops []OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopTypeUUID;references:UUID;" json:"organization_workshops"`
}

// TableName 表名称
func (OrganizationWorkshopTypeModel) TableName() string {
	return "organization_workshop_types"
}

// FindOneByUUID 根据uuid获取单条数据
//  @receiver cls
//  @param uuid
//  @return OrganizationWorkshopTypeModel
func (cls OrganizationWorkshopTypeModel) FindOneByUUID(uuid string) OrganizationWorkshopTypeModel {
	if ret := Init(cls).SetWheres(tools.Map{"uuid": uuid}).Prepare().First(&cls); ret.Error != nil {
		panic(wrongs.PanicWhenIsEmpty(ret, "车间类型"))
	}

	return cls
}
