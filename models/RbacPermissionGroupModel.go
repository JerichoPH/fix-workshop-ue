package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type RbacPermissionGroupModel struct {
	BaseModel
	UUID            string                 `gorm:"type:CHAR(36);UNIQUE;NOT NULL;COMMENT:uuid;" json:"uuid"`
	Name            string                 `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:权限分组名称;" json:"name"`
	RbacPermissions []*RbacPermissionModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:RbacPermissionGroupUUID;references:UUID;COMMENT:相关权限;" json:"rbac_permissions"`
}

// TableName 表名称
func (cls *RbacPermissionGroupModel) TableName() string {
	return "rbac_permission_groups"
}

// BeforeCreate 插入数据前
func (cls *RbacPermissionGroupModel) BeforeCreate(db *gorm.DB) (err error) {
	cls.UUID = uuid.NewV4().String()
	return
}
