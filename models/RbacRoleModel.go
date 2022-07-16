package models

import (
	"fix-workshop-ue/exceptions"
	"fix-workshop-ue/tools"
)

// RbacRoleModel 角色模型
type RbacRoleModel struct {
	BaseModel
	Name            string                 `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:角色名称;" json:"name"`
	RbacPermissions []*RbacPermissionModel `gorm:"many2many:pivot_rbac_role_and_rbac_permissions;foreignKey:id;joinForeignKey:rbac_role_id;references:id;joinReferences:rbac_permission_id;COMMENT:角色与权限多对多;" json:"rbac_permissions"`
	Accounts        []*AccountModel        `gorm:"many2many:pivot_rbac_role_and_accounts;foreignKey:id;joinForeignKey:rbac_role_id;references:id;joinReferences:account_id;COMMENT:角色与用户多对多;" json:"accounts"`
	Menus           []*MenuModel           `gorm:"many2many:pivot_rbac_role_and_menus;foreignKey:id;joinForeignKey:rbac_role_id;references:id;joinReferences:menu_id;COMMENT:角色与菜单多对多;" json:"menus"`
}

// TableName 表名称
func (cls *RbacRoleModel) TableName() string {
	return "rbac_roles"
}

// FindOneByUUID 根据UUID获取单挑数据
//  @receiver cls
//  @param uuid
//  @return RbacRoleModel
func (cls RbacRoleModel) FindOneByUUID(uuid string) RbacRoleModel {
	if ret := Init(cls).SetWheres(tools.Map{"uuid": uuid}).Prepare().First(&cls); ret.Error != nil {
		panic(exceptions.ThrowWhenIsEmptyByDB(ret, "角色"))
	}

	return cls
}
