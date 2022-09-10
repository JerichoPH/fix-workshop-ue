package models

type SourceTypeModel struct {
	BaseModel
	UniqueCode  string            `gorm:"type:CHAR(2);COMMENT:来源类型代码;" json:"unique_code"`
	Name        string            `gorm:"type:VARCHAR(64);COMMENT:来源类型名称;" json:"name"`
	SourceNames []SourceNameModel `gorm:"foreignKey:SourceTypeUuid;references:Uuid;COMMENT:相关来源名称;" json:"source_names"`
}

// TableName 表名称
func (SourceTypeModel) TableName() string {
	return "source_types"
}
