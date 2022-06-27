package models

type EntireInstanceStatus struct {
	BaseModel
	Preloads  []string
	Selects   []string
	Omits     []string
	UniqueCode      string            `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:状态代码;" json:"unique_code"`
	Name            string            `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:状态名称;" json:"name"`
	Number          string            `gorm:"type:CHAR(2);UNIQUE;NOT NULL;COMMENT:状态数字代码;" json:"number"`
	EntireInstances []*EntireInstance `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceStatusUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}
