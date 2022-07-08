package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type RbacPermissionModel struct {
	BaseModel
	UUID                    string                   `gorm:"type:CHAR(36);UNIQUE;NOT NULL;COMMENT:uuid;" json:"uuid"`
	Name                    string                   `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:权限名称;" json:"name"`
	URI                     string                   `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:指向路由;" json:"uri"`
	Method                  string                   `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:请求方法;" json:"method"`
	RbacRoles               []*RbacRoleModel         `gorm:"many2many:pivot_rbac_role_and_rbac_permissions;COMMENT:角色与权限多对多;" json:"rbac_roles"`
	RbacPermissionGroupUUID string                   `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属权限分组编号;" json:"rbac_permission_group_uuid"`
	RbacPermissionGroup     RbacPermissionGroupModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:RbacPermissionGroupUUID;references:UUID;COMMENT:所属权限分组;" json:"rbac_permission_group"`
}

// TableName 表名称
func (cls *RbacPermissionModel) TableName() string {
	return "rbac_permissions"
}

// BeforeCreate 新建前
func (cls *RbacPermissionModel) BeforeCreate(db *gorm.DB) (err error) {
	cls.UUID = uuid.NewV4().String()
	return
}
