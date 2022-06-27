package models

import (
	"gorm.io/gorm"
)

type AccountStatus struct {
	BaseModel
	Preloads   []string
	Selects    []string
	Omits      []string
	UniqueCode string    `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:用户状态代码;"`
	Name       string    `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:用户状态名称;"`
	Accounts   []Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:AccountStatusUniqueCode;references:UniqueCode;"`
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *AccountStatus) FindOneByUniqueCode(db *gorm.DB, uniqueCode string) (accountStatus AccountStatus) {
	cls.Boot(db, cls.Preloads, cls.Selects, cls.Omits).
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&accountStatus)
	return
}
