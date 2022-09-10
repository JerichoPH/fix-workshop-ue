package models

type RbacPermissionGroupModel struct {
	BaseModel
	Name            string                 `gorm:"type:VARCHAR(64);COMMENT:权限分组名称;" json:"name"`
	RbacPermissions []*RbacPermissionModel `gorm:"foreignKey:RbacPermissionGroupUuid;references:Uuid;COMMENT:相关权限;" json:"rbac_permissions"`
}

// TableName 表名称
func (RbacPermissionGroupModel) TableName() string {
	return "rbac_permission_groups"
}
