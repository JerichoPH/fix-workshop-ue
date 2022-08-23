package models

type EntireInstanceStatusModel struct {
	BaseModel
	UniqueCode      string                 `gorm:"type:VARCHAR(64);COMMENT:状态代码;" json:"unique_code"`
	Name            string                 `gorm:"type:VARCHAR(64);COMMENT:状态名称;" json:"name"`
	Number          string                 `gorm:"type:CHAR(2);COMMENT:状态数字代码;" json:"number"`
}

// TableName 表名称
func (EntireInstanceStatusModel) TableName() string {
	return "entire_instance_statuses"
}