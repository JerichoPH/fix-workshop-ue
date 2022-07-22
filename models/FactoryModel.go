package models

import "fix-workshop-ue/models/EntireInstanceModels"

type FactoryModel struct {
	BaseModel
	UniqueCode      string                                     `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:供应商代码;" json:"unique_code"`
	Name            string                                     `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:供应商名称;" json:"name"`
	ShotName        string                                     `gorm:"type:VARCHAR(64);COMMENT:供应商简称;" json:"shot_name"`
	EntireInstances []EntireInstanceModels.EntireInstanceModel `gorm:"foreignKey:FactoryUUID;references:UUID;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *FactoryModel) TableName() string {
	return "factories"
}
