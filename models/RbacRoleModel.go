package models

// RbacRoleModel 角色模型
type RbacRoleModel struct {
	BaseModel
	Name            string                 `gorm:"type:VARCHAR(64);COMMENT:角色名称;" json:"name"`
	RbacPermissions []*RbacPermissionModel `gorm:"many2many:pivot_rbac_role_and_rbac_permissions;foreignKey:Uuid;joinForeignKey:rbac_role_uuid;references:id;joinReferences:rbac_permission_uuid;COMMENT:角色与权限多对多;" json:"rbac_permissions"`
	Accounts        []*AccountModel        `gorm:"many2many:pivot_rbac_role_and_accounts;foreignKey:Uuid;joinForeignKey:rbac_role_uuid;references:uuid;joinReferences:account_uuid;COMMENT:角色与用户多对多;" json:"accounts"`
	Menus           []*MenuModel           `gorm:"many2many:pivot_rbac_role_and_menus;foreignKey:Uuid;joinForeignKey:rbac_role_uuid;references:uuid;joinReferences:menu_uuid;COMMENT:角色与菜单多对多;" json:"menus"`
}

// TableName 表名称
func (RbacRoleModel) TableName() string {
	return "rbac_roles"
}
