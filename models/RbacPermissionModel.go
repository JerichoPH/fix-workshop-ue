package models

type RbacPermissionModel struct {
	BaseModel
	Name                    string                   `gorm:"type:VARCHAR(64);COMMENT:权限名称;" json:"name"`
	URI                     string                   `gorm:"type:VARCHAR(128);INDEX;COMMENT:指向路由;" json:"uri"`
	Method                  string                   `gorm:"type:VARCHAR(64);INDEX;COMMENT:请求方法;" json:"method"`
	RbacRoles               []*RbacRoleModel         `gorm:"many2many:pivot_rbac_role_and_rbac_permissions;foreignKey:uuid;joinForeignKey:rbac_permission_uuid;references:uuid;joinReferences:rbac_role_uuid;COMMENT:角色与权限多对多;" json:"rbac_roles"`
	RbacPermissionGroupUuid string                   `gorm:"type:VARCHAR(36);COMMENT:所属权限分组编号;" json:"rbac_permission_group_uuid"`
	RbacPermissionGroup     RbacPermissionGroupModel `gorm:"foreignKey:RbacPermissionGroupUuid;references:Uuid;COMMENT:所属权限分组;" json:"rbac_permission_group"`
}

// TableName 表名称
func (RbacPermissionModel) TableName() string {
	return "rbac_permissions"
}
