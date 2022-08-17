package models

// EntireInstanceLogTypeModel 器材日志类型模型
type EntireInstanceLogTypeModel struct {
	BaseModel
	UniqueCode             string `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;DEFAULT:;COMMENT:器材日志类型代码;" json:"unique_code"`
	Name                   string `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:;COMMENT:器材日志类型名称;" json:"name"`
	UniqueCodeForParagraph string `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:;COMMENT:器材日志对应段中心代码;" json:"unique_code_for_paragraph"`
	Number                 string `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:;COMMENT:器材日志类型数字代码;" json:"number"`
	Icon                   string `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:;COMMENT:器材日志类型图标;" json:"icon"`
}

// TableName 表名称
func (EntireInstanceLogTypeModel) TableName() string {
	return "entire_instance_log_types"
}
