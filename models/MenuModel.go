package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type MenuModel struct {
	BaseModel
	UUID       string          `gorm:"type:CHAR(36);UNIQUE;NOT NULL;COMMENT:uuid;" json:"uuid"`
	Name       string          `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:菜单名称;" json:"name"`
	URL        string          `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:url;" json:"url"`
	URIName    string          `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:路由标识;" json:"uri_name"`
	ParentUUID string          `gorm:"type:CHAR(36);COMMENT:父级编码;" json:"parent_uuid"`
	Icon       string          `gorm:"type:VARCHAR(64);COMMENT:图标;" json:"icon"`
	Parent     *MenuModel      `gorm:"foreignKey:ParentUUID;references:UUID;COMMENT:所属父级;" json:"parent"`
	Subs       []*MenuModel    `gorm:"foreignKey:ParentUUID;references:UUID;COMMENT:相关子集;" json:"subs"`
	RbacRoles  []RbacRoleModel `gorm:"many2many:pivot_rbac_role_and_menus;COMMENT:角色与菜单多对多;" json:"rbac_roles"`
}

// TableName 表名称
func (cls *MenuModel) TableName() string {
	return "menus"
}

// BeforeCreate 新建前
func (cls *MenuModel) BeforeCreate(db *gorm.DB) (err error) {
	cls.UUID = uuid.NewV4().String()
	return
}
