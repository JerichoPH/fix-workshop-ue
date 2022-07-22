package models

type SourceTypeModel struct {
	BaseModel
	UniqueCode      string                                     `gorm:"type:CHAR(2);UNIQUE;NOT NULL;COMMENT:来源类型代码;" json:"unique_code"`
	Name            string                                     `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:来源类型名称;" json:"name"`
	SourceNames     []SourceNameModel                          `gorm:"foreignKey:SourceTypeUUID;references:UUID;COMMENT:相关来源名称;" json:"source_names"`
	EntireInstances []EntireInstanceModel `gorm:"foreignKey:SourceTypeUUID;references:UUID;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *SourceTypeModel) TableName() string {
	return "source_types"
}
