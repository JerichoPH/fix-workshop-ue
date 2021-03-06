package models

import "gorm.io/gorm"

// OrganizationRailroadGradeCross 道口
type OrganizationRailroadGradeCross struct {
	BaseModel
	Preloads                       []string
	Selects                        []string
	Omits                          []string
	UniqueCode                     string               `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:道口代码;" json:"unique_code"` // I0100
	Name                           string               `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:道口名称;" json:"name"`
	OrganizationWorkshopUniqueCode string               `gorm:"type:CHAR(7);NOT NULL;COMMENT:车间代码;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop           OrganizationWorkshop `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属车间;" json:"organization_workshop"`
}

// FindOneByUniqueCode 通过unique_code获取单条数据
func (cls *OrganizationRailroadGradeCross) FindOneByUniqueCode(db *gorm.DB, uniqueCode string) (organizationRailroadGradeCross OrganizationRailroadGradeCross) {
	cls.Boot(db, cls.Preloads, cls.Selects, cls.Omits).
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&organizationRailroadGradeCross)

	return
}
