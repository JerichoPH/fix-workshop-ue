package models

import (
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
)

// AccountModel 用户模型
type AccountModel struct {
	BaseModel
	Username                  string                     `gorm:"type:VARCHAR(64);COMMENT:登录账号;" json:"username"`
	Password                  string                     `gorm:"type:VARCHAR(128);COMMENT:登录密码;" json:"password"`
	Nickname                  string                     `gorm:"type:VARCHAR(64);COMMENT:昵称;" json:"nickname"`
	DeleteEntireInstances     []*EntireInstanceModel     `gorm:"constraint:OnUpdate:CASCADE;foreignKey:DeleteOperatorUUID;references:UUID;COMMENT:相关删除的器材;" json:"delete_entire_instances"`
	RbacRoles                 []*RbacRoleModel           `gorm:"many2many:pivot_rbac_role_and_accounts;foreignKey:id;joinForeignKey:account_id;References:id;joinReferences:rbac_role_id;COMMENT:角色与用户多对多;" json:"rbac_roles"`
	OrganizationRailwayUUID   string                     `gorm:"type:CHAR(36);COMMENT:所属路局UUID;" json:"organization_railway_uuid"`
	OrganizationRailway       OrganizationRailwayModel   `gorm:"foreignKey:OrganizationRailwayUUID;references:UUID;COMMENT:所属路局;" json:"organization_railway"`
	OrganizationParagraphUUID string                     `gorm:"type:CHAR(36);COMMENT:所属站段UUID;" json:"organization_paragraph_uuid"`
	OrganizationParagraph     OrganizationParagraphModel `gorm:"foreignKey:OrganizationParagraphUUID;references:UUID;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationWorkshopUUID  string                     `gorm:"type:CHAR(36);COMMENT:所属车间UUID;" json:"organization_workshop_uuid"`
	OrganizationWorkshop      OrganizationWorkshopModel  `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUUID  string                     `gorm:"type:CHAR(36);COMMENT:所属工区UUID;" json:"organization_work_area_uuid"`
	OrganizationWorkArea      OrganizationWorkAreaModel  `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:所属工区;" json:"organization_work_area"`
}

// TableName 表名称
func (AccountModel) TableName() string {
	return "accounts"
}

// FindOneByUUID 根据UUID获取单个对象
//  @receiver cls
//  @param uuid
//  @return AccountModel
func (cls AccountModel) FindOneByUUID(uuid string) AccountModel {
	if ret := Init(cls).SetWheres(tools.Map{"uuid": uuid}).Prepare("").First(&cls); ret.Error != nil {
		panic(wrongs.PanicWhenIsEmpty(ret, "用户"))
	}

	return cls
}
