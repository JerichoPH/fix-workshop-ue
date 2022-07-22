package KindModels

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/models/EntireInstanceModels"
)

type KindEntireTypeModel struct {
	models.BaseModel
	UniqueCode             string                                     `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:类型代码;" json:"unique_code"`
	Name                   string                                     `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:类型名称;" json:"name"`
	Nickname               string                                     `gorm:"type:VARCHAR(128);COMMENT:打印别名;" json:"nickname"`
	BeEnable               string                                     `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	KindCategoryUniqueCode string                                     `gorm:"type:CHAR(3);NOT NULL;COMMENT:所属种类代码;" json:"kind_category_unique_code"`
	KindCategory           KindCategoryModel                          `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindCategoryUniqueCode;references:UniqueCode;" json:"kind_category"`
	KindSubTypes           []KindSubTypeModel                         `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindEntireTypeUniqueCode;references:UniqueCode;COMMENT:相关型号;" json:"kind_sub_types"`
	EntireInstances        []EntireInstanceModels.EntireInstanceModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindEntireTypeUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *KindEntireTypeModel) TableName() string {
	return "kind_entire_types"
}