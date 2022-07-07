package models

// RbacRoleModel 角色模型
type RbacRoleModel struct {
	BaseModel
	Name            string                 `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:角色名称;" json:"name"`
	RbacPermissions []*RbacPermissionModel `gorm:"many2many:pivot_rbac_role_and_rbac_permissions" json:"rbac_permissions"`
	Accounts        []*AccountModel        `gorm:"many2many:pivot_rbac_role_and_accounts;foreignKey:id;joinForeignKey:rbac_role_id;References:id;joinReferences:account_id;COMMENT:角色与用户多对多;" json:"accounts"`
}

// TableName 表名称
func (cls RbacRoleModel) TableName() string {
	return "rbac_roles"
}
