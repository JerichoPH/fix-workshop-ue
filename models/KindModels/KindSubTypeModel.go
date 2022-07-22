package KindModels

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/models/EntireInstanceModels"
)

type KindSubTypeModel struct {
	models.BaseModel
	UniqueCode               string                                     `gorm:"type:CHAR(7);UNIQUE;NOT NULL;COMMENT:型号代码;" json:"unique_code"`
	Name                     string                                     `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:型号名称;" json:"name"`
	Nickname                 string                                     `gorm:"type:VARCHAR(128);COMMENT:打印别名;" json:"nickname"`
	BeEnable                 string                                     `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	KindEntireTypeUniqueCode string                                     `gorm:"type:CHAR(3);NOT NULL;COMMENT:所属类型代码;" json:"kind_entire_type_unique_code"`
	KindEntireType           KindEntireTypeModel                        `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindEntireTypeUniqueCode;references:UniqueCode;" json:"kind_entire_type"`
	EntireInstances          []EntireInstanceModels.EntireInstanceModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindSubTypeUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *KindSubTypeModel) TableName() string {
	return "kind_sub_types"
}
