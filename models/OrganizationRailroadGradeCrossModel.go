package models

// OrganizationRailroadGradeCrossModel 道口
type OrganizationRailroadGradeCrossModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:道口代码;" json:"unique_code"` // I0100
	Name                     string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:道口名称;" json:"name"`
	BeEnable                 bool                      `gorm:"type:BOOLEAN;NOT NULL;DEFAULT:0;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUUID string                    `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属车间UUID;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;NOT NULL;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUUID string                    `gorm:"type:CHAR(36);COMMENT:所属工区UUID;" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:所属工区;" json:"organization_work_area"`
}

// TableName 表名称
func (cls *OrganizationRailroadGradeCrossModel) TableName() string {
	return "organization_railroad_grade_crosses"
}
