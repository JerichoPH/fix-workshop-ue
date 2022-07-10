package models

// AccountStatusModel 用户状态模型
type AccountStatusModel struct {
	BaseModel
	UniqueCode string         `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:用户状态代码;" json:"unique_code"`
	Name       string         `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:用户状态名称;" json:"name"`
	Accounts   []AccountModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:AccountStatusUniqueCode;references:UniqueCode;COMMENT:相关账号;" json:"accounts"`
}

// TableName 表名称
func (cls *AccountStatusModel) TableName() string {
	return "account_statuses"
}
