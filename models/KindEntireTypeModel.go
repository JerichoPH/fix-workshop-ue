package models

type KindEntireTypeModel struct {
	BaseModel
	UniqueCode       string                `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:类型代码;" json:"unique_code"`
	Name             string                `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:类型名称;" json:"name"`
	Nickname         string                `gorm:"type:VARCHAR(128);COMMENT:打印别名;" json:"nickname"`
	BeEnable         bool                `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	KindCategoryUUID string                `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属种类代码;" json:"kind_category_uuid"`
	KindCategory     KindCategoryModel     `gorm:"foreignKey:KindCategoryUUID;references:UUID;" json:"kind_category"`
	KindSubTypes     []KindSubTypeModel    `gorm:"foreignKey:KindEntireTypeUUID;references:UUID;COMMENT:相关型号;" json:"kind_sub_types"`
}

// TableName 表名称
func (KindEntireTypeModel) TableName() string {
	return "kind_entire_types"
}