package models

type FactoryModel struct {
	BaseModel
	UniqueCode      string                `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:生产厂家代码;" json:"unique_code"`
	Name            string                `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:生产厂家名称;" json:"name"`
	ShotName        string                `gorm:"type:VARCHAR(64);COMMENT:生产厂家名称;" json:"shot_name"`
	EntireInstances []EntireInstanceModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:FactoryUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *FactoryModel) TableName() string {
	return "factories"
}
