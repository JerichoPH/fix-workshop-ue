package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// AccountModel 用户模型
type AccountModel struct {
	BaseModel
	UUID                    string                 `gorm:"type:CHAR(36);UNIQUE;NOT NULL;COMMENT:UUID;" json:"uuid"`
	Username                string                 `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:登录账号;" json:"username"`
	Password                string                 `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:登录密码;" json:"password"`
	Nickname                string                 `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:昵称;" json:"nickname"`
	AccountStatusUniqueCode string                 `gorm:"type:VARCHAR(64);COMMENT:所属状态代码;" json:"account_status_unique_code"`
	AccountStatus           AccountStatusModel     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:AccountStatusUniqueCode;references:UniqueCode;COMMENT:所属状态;" json:"account_status"`
	DeleteEntireInstances   []*EntireInstanceModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:DeleteOperatorUUID;references:UUID;COMMENT:相关删除的器材;" json:"delete_entire_instances"`
	RbacRoles               []*RbacRoleModel       `gorm:"many2many:pivot_rbac_role_and_accounts;foreignKey:id;joinForeignKey:account_id;References:id;joinReferences:rbac_role_id;COMMENT:角色与用户多对多;" json:"rbac_roles"`
}

// AccountUpdateForm 用户编辑表单
type AccountUpdateForm struct {
	Username                string `form:"username" json:"string" uri:"username"`
	Nickname                string `form:"nickname" json:"nickname" uri:"nickname"`
	AccountStatusUniqueCode string `form:"account_status_unique_code" json:"account_status_unique_code" uri:"account_status_unique_code"`
}

// TableName 表名称
func (cls *AccountModel) TableName() string {
	return "accounts"
}

// BeforeCreate 自动生成UniqueCode
func (cls *AccountModel) BeforeCreate(tx *gorm.DB) (err error) {
	cls.UUID = uuid.NewV4().String()
	return
}
