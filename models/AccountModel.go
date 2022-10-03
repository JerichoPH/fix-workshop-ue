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
	DeleteEntireInstances     []*EntireInstanceModel     `gorm:"constraint:OnUpdate:CASCADE;foreignKey:DeleteOperatorUuid;references:Uuid;COMMENT:相关删除的器材;" json:"delete_entire_instances"`
	BeSuperAdmin              bool                       `gorm:"BOOLEAN;NOT NULL;DEFAULT:0;COMMENT:超级管理员;" json:"be_super_admin"`
	RbacRoles                 []*RbacRoleModel           `gorm:"many2many:pivot_rbac_role_and_accounts;foreignKey:id;joinForeignKey:account_id;References:id;joinReferences:rbac_role_id;COMMENT:角色与用户多对多;" json:"rbac_roles"`
	OrganizationRailwayUuid   string                     `gorm:"type:VARCHAR(36);COMMENT:所属路局UUID;" json:"organization_railway_uuid"`
	OrganizationRailway       OrganizationRailwayModel   `gorm:"foreignKey:OrganizationRailwayUuid;references:Uuid;COMMENT:所属路局;" json:"organization_railway"`
	OrganizationParagraphUuid string                     `gorm:"type:VARCHAR(36);COMMENT:所属站段UUID;" json:"organization_paragraph_uuid"`
	OrganizationParagraph     OrganizationParagraphModel `gorm:"foreignKey:OrganizationParagraphUuid;references:Uuid;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationWorkshopUuid  string                     `gorm:"type:VARCHAR(36);COMMENT:所属车间UUID;" json:"organization_workshop_uuid"`
	OrganizationWorkshop      OrganizationWorkshopModel  `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUuid  string                     `gorm:"type:VARCHAR(36);COMMENT:所属工区UUID;" json:"organization_work_area_uuid"`
	OrganizationWorkArea      OrganizationWorkAreaModel  `gorm:"foreignKey:OrganizationWorkAreaUuid;references:Uuid;COMMENT:所属工区;" json:"organization_work_area"`
}

// TableName 表名称
func (AccountModel) TableName() string {
	return "accounts"
}

// FindOneByUUID 根据UUID获取单个对象
//  @receiver ins
//  @param uuid
//  @return AccountModel
func (ins AccountModel) FindOneByUUID(uuid string) AccountModel {
	if ret := BootByModel(ins).SetWheres(tools.Map{"uuid": uuid}).Prepare("").First(&ins); ret.Error != nil {
		panic(wrongs.PanicWhenIsEmpty(ret, "用户"))
	}

	return ins
}
