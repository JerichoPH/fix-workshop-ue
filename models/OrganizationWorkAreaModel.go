package models

import "gorm.io/gorm"

// OrganizationWorkAreaModel 工区
type OrganizationWorkAreaModel struct {
	BaseModel
	UniqueCode                   string                        `gorm:"type:CHAR(8);UNIQUE;NOT NULL;COMMENT:工区代码;" json:"unique_code"` //B049D001
	Name                         string                        `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:工区名称;" json:"name"`
	BeEnable                     bool                          `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkAreaTypeUUID string                        `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属工区类型uuid;" json:"organization_work_area_type_uuid"`
	OrganizationWorkAreaType     OrganizationWorkAreaTypeModel `gorm:"CASCADE;foreignKey:OrganizationWorkAreaTypeUUID;references:UUID;COMMENT:所属工区类型;" json:"organization_work_area_type"`
	OrganizationWorkshopUUID     string                        `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属车间uuid;" json:"organization_workshop_uuid"`
	OrganizationWorkshop         OrganizationWorkshopModel     `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationStations         []OrganizationStationModel    `gorm:"foreignKey:OrganizationWorkAreaUniqueCode;references:UniqueCode;COMMENT:相关站场;" json:"organization_stations"`
	EntireInstances              []EntireInstanceModel         `gorm:"foreignKey:OrganizationWorkAreaUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *OrganizationWorkAreaModel) TableName() string {
	return "organization_work_areas"
}

// ScopeBeEnable 获取启用的数据
func (cls *OrganizationWorkAreaModel) ScopeBeEnable(db *gorm.DB) *gorm.DB {
	return db.Where("be_enable = ?", 1)
}
