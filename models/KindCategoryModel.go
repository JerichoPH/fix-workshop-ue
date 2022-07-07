package models

type KindCategoryModel struct {
	BaseModel
	UniqueCode      string                `gorm:"type:CHAR(3);UNIQUE;NOT NULL;COMMENT:种类代码;" json:"unique_code"`
	Name            string                `gorm:"type:VARCHAR(128);UNIQUE;NOT NULL;COMMENT:种类名称;" json:"name"`
	Nickname        string                `gorm:"type:VARCHAR(128);UNIQUE:NOT NULL;COMMENT:打印别名;" json:"nickname"`
	BeEnable        string                `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	KindEntireTypes []KindEntireTypeModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindCategoryUniqueCode;references:UniqueCode;" json:"kind_entire_types"`
	EntireInstances []EntireInstanceModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindCategoryUniqueCode;references:UniqueCode;NOT NULL;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *KindCategoryModel) TableName() string {
	return "kind_categories"
}
