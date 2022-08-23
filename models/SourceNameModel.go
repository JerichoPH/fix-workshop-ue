package models

type SourceNameModel struct {
	BaseModel
	UniqueCode     string          `gorm:"type:VARCHAR(64);COMMENT:来源名称代码;" json:"unique_code"`
	Name           string          `gorm:"type:VARCHAR(64);COMMENT:来源名称名称;" json:"name"`
	SourceTypeUUID string          `gorm:"type:CHAR(36);COMMENT:所属来源类型代码;" json:"source_type_uuid"`
	SourceType     SourceTypeModel `gorm:"foreignKey:SourceTypeUUID;references:UUID;COMMENT:所属来源类型;" json:"source_type"`
}

// TableName 表名称
func (SourceNameModel) TableName() string {
	return "source_names"
}
