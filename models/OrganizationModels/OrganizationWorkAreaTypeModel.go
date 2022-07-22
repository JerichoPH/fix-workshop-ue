package OrganizationModels

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
)

type OrganizationWorkAreaTypeModel struct {
	models.BaseModel
	UniqueCode            string                      `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:工区类型代码;" json:""`
	Name                  string                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:工区类型名称;" json:""`
	OrganizationWorkAreas []OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaTypeUUID;references:UUID;COMMENT:相关工区;" json:"organization_work_areas"`
}

// TableName 表名称
func (cls *OrganizationWorkAreaTypeModel) TableName() string {
	return "organization_work_area_types"
}

// FindOneByUUID 根据UUID获取单条数据
//  @receiver cls
//  @param uuid
//  @return OrganizationWorkAreaTypeModel
func (cls OrganizationWorkAreaTypeModel) FindOneByUUID(uuid string) OrganizationWorkAreaTypeModel {
	if ret := models.Init(cls).SetWheres(tools.Map{"uuid": uuid}).Prepare().First(&cls); ret.Error != nil {
		panic(abnormals.PanicWhenIsEmpty(ret, "工区类型"))
	}

	return cls
}
