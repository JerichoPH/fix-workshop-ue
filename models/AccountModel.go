package models

import (
	"fix-workshop-ue/abnormals"
	"fix-workshop-ue/tools"
)

// AccountModel 用户模型
type AccountModel struct {
	BaseModel
	Username              string                                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:登录账号;" json:"username"`
	Password              string                                      `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:登录密码;" json:"password"`
	Nickname              string                                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:昵称;" json:"nickname"`
	DeleteEntireInstances []*EntireInstanceModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:DeleteOperatorUUID;references:UUID;COMMENT:相关删除的器材;" json:"delete_entire_instances"`
	RbacRoles             []*RbacRoleModel                 `gorm:"many2many:pivot_rbac_role_and_accounts;foreignKey:id;joinForeignKey:account_id;References:id;joinReferences:rbac_role_id;COMMENT:角色与用户多对多;" json:"rbac_roles"`
}

// TableName 表名称
func (cls *AccountModel) TableName() string {
	return "accounts"
}

// FindOneByUUID 根据UUID获取单个对象
//  @receiver cls
//  @param uuid
//  @return AccountModel
func (cls AccountModel) FindOneByUUID(uuid string) AccountModel {
	if ret := Init(cls).SetWheres(tools.Map{"uuid": uuid}).Prepare().First(&cls); ret.Error != nil {
		panic(abnormals.PanicWhenIsEmpty(ret, "用户"))
	}

	return cls
}
