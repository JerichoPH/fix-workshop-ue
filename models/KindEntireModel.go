package models

type KindEntireModel struct {
	BaseModel
	UniqueCode             string           `gorm:"<-;type:CHAR(5);UNIQUE;NOT NULL;COMMENT:类型代码;" json:"unique_code"`
	Name                   string           `gorm:"<-;type:VARCHAR(128);NOT NULL;COMMENT:类型名称;" json:"name"`
	Nickname               string           `gorm:"<-;type:VARCHAR(128);COMMENT:打印别名;" json:"nickname"`
	IsShow                 string           `gorm:"<-;type:BOOLEAN;DEFAULT:0;COMMENT:是否显示;" json:"is_show"`
	KindCategoryUniqueCode string           `gorm:"<-;type:CHAR(3);NOT NULL;COMMENT:所属种类代码;" json:"kind_category_unique_code"`
	KindCategory           KindCategory     `gorm:"<-;constraint:OnUpdate:CASCADE;foreignKey:KindCategoryUniqueCode;references:UniqueCode;" json:"kind_category"`
	KindSubModels          []KindSubModel   `gorm:"<-;constraint:OnUpdate:CASCADE;foreignKey:KindEntireModelUniqueCode;references:UniqueCode;COMMENT:相关型号;" json:"kind_sub_models"`
	EntireInstances        []EntireInstance `gorm:"<-;constraint:OnUpdate:CASCADE;foreignKey:KindEntireModelUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *KindEntireModel) FindOneByUniqueCode(uniqueCode string) (kindEntireModel KindEntireModel) {
	cls.Boot().Where(map[string]interface{}{"unique_code": uniqueCode}).First(&kindEntireModel)

	return
}
