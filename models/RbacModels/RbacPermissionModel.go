package RbacModels

import "fix-workshop-ue/models"

type RbacPermissionModel struct {
	models.BaseModel
	Name                    string                   `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:权限名称;" json:"name"`
	URI                     string                   `gorm:"type:VARCHAR(128);INDEX;NOT NULL;COMMENT:指向路由;" json:"uri"`
	Method                  string                   `gorm:"type:VARCHAR(64);INDEX;NOT NULL;COMMENT:请求方法;" json:"method"`
	RbacRoles               []*RbacRoleModel         `gorm:"many2many:pivot_rbac_role_and_rbac_permissions;foreignKey:id;joinForeignKey:rbac_permission_id;references:id;joinReferences:rbac_role_id;COMMENT:角色与权限多对多;" json:"rbac_roles"`
	RbacPermissionGroupUUID string                   `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属权限分组编号;" json:"rbac_permission_group_uuid"`
	RbacPermissionGroup     RbacPermissionGroupModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:RbacPermissionGroupUUID;references:UUID;COMMENT:所属权限分组;" json:"rbac_permission_group"`
}

// TableName 表名称
func (cls *RbacPermissionModel) TableName() string {
	return "rbac_permissions"
}
