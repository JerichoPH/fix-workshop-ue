package models

type AccountStatusModel struct {
	BaseModel
	UniqueCode string         `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:用户状态代码;" json:"unique_code"`
	Name       string         `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:用户状态名称;" json:"name"`
	Accounts   []AccountModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:AccountStatusUniqueCode;references:UniqueCode;" json:"accounts"`
}

// TableName 表名称
func (cls *AccountStatusModel) TableName() string {
	return "account_statuses"
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *AccountStatusModel) FindOneByUniqueCode(uniqueCode string) (accountStatus AccountStatusModel) {
	cls.Boot().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&accountStatus)
	return
}
