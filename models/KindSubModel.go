package models

import (
	"gorm.io/gorm"
)

type KindSubModel struct {
	BaseModel
	Preloads                  []string
	Omits                     []string
	Selects                   []string
	UniqueCode                string           `gorm:"type:CHAR(7);UNIQUE;NOT NULL;COMMENT:型号代码;" json:"unique_code"`
	Name                      string           `gorm:"type:VARCHAR(128);NOT NULL;COMMENT:型号名称;" json:"name"`
	Nickname                  string           `gorm:"type:VARCHAR(128);COMMENT:打印别名;" json:"nickname"`
	IsShow                    string           `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否显示;" json:"is_show"`
	KindEntireModelUniqueCode string           `gorm:"type:CHAR(3);NOT NULL;COMMENT:所属类型代码;" json:"kind_entire_model_unique_code"`
	KindEntireModel           KindEntireModel  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindEntireModelUniqueCode;references:UniqueCode;" json:"kind_entire_model"`
	EntireInstances           []EntireInstance `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindSubModelUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *KindSubModel) FindOneByUniqueCode(db *gorm.DB, uniqueCode string) (kindSubModel KindSubModel) {
	cls.Boot(db, cls.Preloads, cls.Selects, cls.Omits).Where(map[string]interface{}{"unique_code": uniqueCode}).First(&kindSubModel)

	return
}
