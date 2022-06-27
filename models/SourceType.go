package models

type SourceType struct {
	BaseModel
	Preloads        []string
	Selects         []string
	Omits           []string
	UniqueCode      string           `gorm:"type:CHAR(2);UNIQUE;NOT NULL;COMMENT:来源类型代码;" json:"unique_code"`
	Name            string           `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:来源类型名称;" json:"name"`
	SourceNames     []SourceName     `gorm:"constraint:OnUpdate:CASCADE;foreignKey:SourceTypeUniqueCode;references:UniqueCode;COMMENT:相关来源名称;" json:"source_names"`
	EntireInstances []EntireInstance `gorm:"constraint:OnUpdate:CASCADE;foreignKey:SourceTypeUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}
