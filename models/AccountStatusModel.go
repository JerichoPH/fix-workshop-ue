package models

type AccountStatusModel struct {
	BaseModel
	UniqueCode string         `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:用户状态代码;" json:"unique_code"`
	Name       string         `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:用户状态名称;" json:"name"`
	Accounts   []AccountModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:AccountStatusUniqueCode;references:UniqueCode;COMMENT:相关账号;" json:"accounts"`
}

type AccountStatusStoreForm struct {
	UniqueCode string `form:"unique_code" json:"unique_code"`
	Name       string `form:"name" json:"name"`
}

type AccountStatusUpdateForm struct {
	UniqueCode string `form:"unique_code" json:"unique_code"`
	Name       string `form:"name" json:"name"`
}

// TableName 表名称
func (cls *AccountStatusModel) TableName() string {
	return "account_statuses"
}
