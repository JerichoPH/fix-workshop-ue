package models

import (
	"gorm.io/gorm"
)

type Account struct {
	BaseModel
	Preloads                []string
	Selects                 []string
	Omits                   []string
	Username                string        `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:登录账号;"`
	Password                string        `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:登录密码;"`
	Nickname                string        `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:昵称;"`
	AccountStatusUniqueCode string        `gorm:"type:VARCHAR(64);COMMENT:所属状态;"`
	AccountStatus           AccountStatus `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:AccountStatusUniqueCode;references:UniqueCode;"`
}

// FindOneById 根据id获取单条数据
func (cls *Account) FindOneById(db *gorm.DB, id int) (account Account) {
	cls.Boot(db, cls.Preloads, cls.Selects, cls.Omits).
		Where(map[string]interface{}{"id": id}).
		First(&account, id)
	return account
}

// FindOneByUsername 根据username获取单条数据
func (cls *Account) FindOneByUsername(db *gorm.DB, username string) (account Account) {
	cls.Boot(db, cls.Preloads, cls.Selects, cls.Omits).
		Where(map[string]interface{}{"username": username}).
		First(&account)
	return
}

// FindOneByNickname 根据nickname获取单条数据
func (cls *Account) FindOneByNickname(db *gorm.DB, nickname string) (account Account) {
	cls.Boot(db, cls.Preloads, cls.Selects, cls.Omits).
		Where(map[string]interface{}{"nickname": nickname}).
		First(&account)
	return
}
