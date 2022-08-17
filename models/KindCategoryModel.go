package models

type KindCategoryModel struct {
	BaseModel
	UniqueCode      string                `gorm:"type:CHAR(3);NOT NULL;COMMENT:种类代码;" json:"unique_code"`
	Name            string                `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:种类名称;" json:"name"`
	Nickname        string                `gorm:"type:VARCHAR(128);UNIQUE:NOT NULL;COMMENT:打印别名;" json:"nickname"`
	BeEnable        bool                  `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	KindEntireTypes []KindEntireTypeModel `gorm:"foreignKey:KindCategoryUUID;references:UUID;" json:"kind_entire_types"`
	Race            string                `gorm:"type:CHAR(1);NOT NULL;COMMENT:设备、器材分类，S设备、Q器材;" json:"race"`
}

// TableName 表名称
func (KindCategoryModel) TableName() string {
	return "kind_categories"
}
