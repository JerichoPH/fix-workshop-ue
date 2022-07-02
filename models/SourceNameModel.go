package models

type SourceNameModel struct {
	BaseModel
	UniqueCode           string                `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:来源名称代码;" json:"unique_code"`
	Name                 string                `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:来源名称名称;" json:"name"`
	SourceTypeUniqueCode string                `gorm:"type:CHAR(2);COMMENT:所属来源类型代码;" json:"source_type_unique_code"`
	SourceType           SourceTypeModel       `gorm:"constraint:OnUpdate:CASCADE;foreignKey:SourceTypeUniqueCode;references:UniqueCode;COMMENT:所属来源类型;" json:"source_type"`
	EntireInstance       []EntireInstanceModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:SourceNameUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *SourceNameModel) TableName() string {
	return "source_names"
}
