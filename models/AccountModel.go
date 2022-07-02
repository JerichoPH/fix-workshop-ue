package models

import (
	uuid "github.com/satori/go.uuid"
)

type AccountModel struct {
	BaseModel
	UUID                    string                 `gorm:"type:CHAR(36);UNIQUE;NOT NULL;COMMENT:UUID;" json:"uuid"`
	Username                string                 `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:登录账号;"`
	Password                string                 `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:登录密码;"`
	Nickname                string                 `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:昵称;"`
	AccountStatusUniqueCode string                 `gorm:"type:VARCHAR(64);COMMENT:所属状态;"`
	AccountStatus           AccountStatusModel     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:AccountStatusUniqueCode;references:UniqueCode;"`
	DeleteEntireInstances   []*EntireInstanceModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:DeleteOperatorUUID;references:UUID;COMMENT:相关删除的器材;" json:"delete_entire_instances"`
}

// TableName 表名称
func (cls *AccountModel) TableName() string {
	return "accounts"
}

// BeforeCreate 自动生成UniqueCode
func (cls *AccountModel) BeforeCreate() (err error) {
	cls.UUID = uuid.NewV4().String()
	return
}

// FindOneByUUID 根据uuid获取单条数据
func (cls *AccountModel) FindOneByUUID(uuid string) (account AccountModel) {
	cls.Boot().
		Where(map[string]interface{}{"uuid": uuid}).
		First(&account)
	return
}

// FindOneById 根据id获取单条数据
func (cls *AccountModel) FindOneById(id int) (account AccountModel) {
	cls.Boot().
		Where(map[string]interface{}{"id": id}).
		First(&account)
	return
}

// FindOneByUsername 根据username获取单条数据
func (cls *AccountModel) FindOneByUsername(username string) (account AccountModel) {
	cls.Boot().
		Where(map[string]interface{}{"username": username}).
		First(&account)
	return
}

// FindOneByNickname 根据nickname获取单条数据
func (cls *AccountModel) FindOneByNickname(nickname string) (account AccountModel) {
	cls.Boot().
		Where(map[string]interface{}{"nickname": nickname}).
		First(&account)
	return
}
