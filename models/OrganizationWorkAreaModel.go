package models

// OrganizationWorkAreaModel 工区
type OrganizationWorkAreaModel struct {
	BaseModel
	UniqueCode                         string                        `gorm:"type:CHAR(8);UNIQUE;NOT NULL;COMMENT:工区代码;" json:"unique_code"` //B049D001
	Name                               string                        `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:工区名称;" json:"name"`
	BeEnable                           bool                          `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkAreaTypeUniqueCode string                        `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:所属工区类型;" json:"organization_work_area_type_unique_code"`
	OrganizationWorkAreaType           OrganizationWorkAreaTypeModel `gorm:"constraint:OnUpdate;CASCADE;foreignKey:OrganizationWorkAreaTypeUniqueCode;references:UniqueCode;COMMENT:所属工区类型;" json:"organization_work_area_type"`
	OrganizationWorkshopUniqueCode     string                        `gorm:"type:CHAR(7);NOT NULL;COMMENT:所属车间;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop               OrganizationWorkshopModel     `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationStations               []OrganizationStationModel    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkAreaUniqueCode;references:UniqueCode;COMMENT:相关站场;" json:"organization_stations"`
	EntireInstances                    []EntireInstanceModel         `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkAreaUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *OrganizationWorkAreaModel) TableName() string {
	return "organization_work_areas"
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *OrganizationWorkAreaModel) FindOneByUniqueCode(uniqueCode string) (organizationWorkArea OrganizationWorkAreaModel) {
	cls.Boot().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&organizationWorkArea)

	return
}
