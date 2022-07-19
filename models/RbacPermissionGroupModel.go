package models

type RbacPermissionGroupModel struct {
	BaseModel
	Name            string                 `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:权限分组名称;" json:"name"`
	RbacPermissions []*RbacPermissionModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:RbacPermissionGroupUUID;references:UUID;COMMENT:相关权限;" json:"rbac_permissions"`
}

// TableName 表名称
func (cls *RbacPermissionGroupModel) TableName() string {
	return "rbac_permission_groups"
}
