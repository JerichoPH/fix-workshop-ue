package model

type FactoryModel struct {
	BaseModel
	UniqueCode      string                `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:供应商代码;" json:"unique_code"`
	Name            string                `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:供应商名称;" json:"name"`
	ShotName        string                `gorm:"type:VARCHAR(64);COMMENT:供应商简称;" json:"shot_name"`
	EntireInstances []EntireInstanceModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:FactoryUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *FactoryModel) TableName() string {
	return "factories"
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *FactoryModel) FindOneByUniqueCode(uniqueCode string) (factory FactoryModel) {
	cls.Boot().Where(map[string]interface{}{"unique_code": uniqueCode}).First(&factory)
	return
}
