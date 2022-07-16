package models

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/tools"
)

type RbacPermissionGroupModel struct {
	BaseModel
	Name            string                 `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:权限分组名称;" json:"name"`
	RbacPermissions []*RbacPermissionModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:RbacPermissionGroupUUID;references:UUID;COMMENT:相关权限;" json:"rbac_permissions"`
}

// TableName 表名称
func (cls *RbacPermissionGroupModel) TableName() string {
	return "rbac_permission_groups"
}

// FindOneByUUID 根据UUID获取单条数据
//  @receiver cls
//  @param uuid
//  @return OrganizationLineModel
func (cls RbacPermissionGroupModel) FindOneByUUID(uuid string) RbacPermissionGroupModel {
	if ret := Init(cls).SetWheres(tools.Map{"uuid": uuid}).Prepare().First(&cls); ret.Error != nil {
		panic(exceptions.ThrowWhenIsEmptyByDB(ret, "权限分组"))
	}

	return cls
}
