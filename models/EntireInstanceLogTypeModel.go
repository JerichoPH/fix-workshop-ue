package models

// EntireInstanceLogTypeModel 器材日志类型模型
type EntireInstanceLogTypeModel struct {
	BaseModel
	UniqueCode         string                   `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:器材日志类型代码;" json:"unique_code"`
	Name               string                   `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:器材日志类型名称;" json:"name"`
	Number             string                   `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:器材日志类型数字代码;" json:"number"`
	EntireInstanceLogs []EntireInstanceLogModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceLogTypeUniqueCode;references:UniqueCode;COMMENT:相关器材日志;" json:"entire_instance_logs"`
}

// TableName 表名称
func (EntireInstanceLogTypeModel) TableName() string {
	return "entire_instance_log_types"
}