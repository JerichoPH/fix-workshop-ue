package models

type OrganizationWorkAreaTypeModel struct {
	BaseModel
	UniqueCode            string                      `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:工区类型代码;" json:""`
	Name                  string                      `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:工区类型名称;" json:""`
	OrganizationWorkAreas []OrganizationWorkAreaModel `gorm:"constraint:OnUpdate;CASCADE;foreignKey:OrganizationWorkAreaTypeUniqueCode;references:UniqueCode;COMMENT:相关工区;" json:"organization_work_areas"`
}

// TableName 表名称
func (cls *OrganizationWorkAreaTypeModel) TableName() string {
	return "organization_work_area_types"
}
