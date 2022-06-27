package models

import (
	"gorm.io/gorm"
)

type OrganizationStation struct {
	BaseModel
	Preloads                       []string
	Selects                        []string
	Omits                          []string
	UniqueCode                     string                `gorm:"type:CHAR(6);UNIQUE;NOT NULL;COMMENT:站场代码;" json:"unique_code"` // G00001
	Name                           string                `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:站场名称;" json:"name"`
	BeEnable                       bool                  `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUniqueCode string                `gorm:"type:CHAR(7);COMMENT:所属车间代码;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop           OrganizationWorkshop  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUniqueCode string                `gorm:"type:CHAR(8);COMMENT:所属工区代码;" json:"organization_work_area_unique_code"`
	OrganizationWorkArea           OrganizationWorkArea  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkAreaUniqueCode;references:UniqueCode;COMMENT:所属工区;" json:"organization_work_area"`
	OrganizationLines              []*OrganizationLine   `gorm:"many2many:pivot_line_stations;COMMENT:线别与车站多对多;"`
	LocationInstallRooms           []LocationInstallRoom `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationStationUniqueCode;references:UniqueCode;COMMENT:相关机房;" json:"location_install_rooms"`
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *OrganizationStation) FindOneByUniqueCode(db *gorm.DB, uniqueCode string) (organizationStation OrganizationStation) {
	cls.Boot(db, cls.Preloads, cls.Selects, cls.Omits).
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&organizationStation)

	return
}
