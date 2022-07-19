package models

// 排模型
type LocationWarehousePlatoonTypeModel struct {
	BaseModel
	UniqueCode                string                  `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:排类型代码;" json:"unique_code"`
	Name                      string                  `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:排类型名称;" json:"name"`
	LocationWarehousePlatoons []LocationDepotRowModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehousePlatoonTypeUniqueCode;references:UniqueCode;NOT NULL;COMMENT:相关排;" json:"location_warehouse_platoons"`
}

// TableName 表名称
func (cls *LocationWarehousePlatoonTypeModel) TableName() string {
	return "location_warehouse_platoon_types"
}
