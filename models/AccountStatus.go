package models

type AccountStatus struct {
	BaseModel
	UniqueCode string    `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:用户状态代码;"`
	Name       string    `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:用户状态名称;"`
	Accounts   []Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:AccountStatusUniqueCode;references:UniqueCode;"`
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *AccountStatus) FindOneByUniqueCode(uniqueCode string) (accountStatus AccountStatus) {
	cls.Boot().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&accountStatus)
	return
}
