package models

type KindCategory struct {
	BaseModel
	UniqueCode       string            `gorm:"type:CHAR(3);UNIQUE;NOT NULL;COMMENT:种类代码;" json:"unique_code"`
	Name             string            `gorm:"type:VARCHAR(128);UNIQUE;NOT NULL;COMMENT:种类名称;" json:"name"`
	Nickname         string            `gorm:"type:VARCHAR(128);UNIQUE:NOT NULL;COMMENT:打印别名;" json:"nickname"`
	IsShow           string            `gorm:"type:BOOLEAN;DEFAULT:0;COMMENT:是否显示;" json:"is_show"`
	KindEntireModels []KindEntireModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindCategoryUniqueCode;references:UniqueCode;" json:"kind_entire_models"`
	EntireInstances  []EntireInstance  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindCategoryUniqueCode;references:UniqueCode;NOT NULL;COMMENT:相关器材;" json:"entire_instances"`
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *KindCategory) FindOneByUniqueCode(uniqueCode string) (kindCategory KindCategory) {
	cls.Boot().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&kindCategory)
	return
}
