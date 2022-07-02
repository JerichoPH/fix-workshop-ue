package models

type RbacPermissionModel struct {
	BaseModel
	Name      string           `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:权限名称;" json:"name"`
	Uri       string           `gorm:"type:VARCHAR(128);UNIQUE;NOT NULL;COMMENT:指向路由;" json:"uri"`
	RbacRoles []*RbacRoleModel `gorm:"many2many:pivot_rbac_role_and_rbac_permissions;COMMENT:角色与权限多对多;" json:"rbac_roles"`
}

// TableName 表名称
func (cls RbacPermissionModel) TableName() string {
	return "rbac_roles"
}
