package models

type Factory struct {
	BaseModel
	UniqueCode      string           `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:生产厂家代码;" json:"unique_code"`
	Name            string           `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:生产厂家名称;" json:"name"`
	ShotName        string           `gorm:"type:VARCHAR(64);COMMENT:生产厂家名称;" json:"shot_name"`
	EntireInstances []EntireInstance `gorm:"constraint:OnUpdate:CASCADE;foreignKey:FactoryUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}
