package OrganizationModels

import "fix-workshop-ue/models"

// OrganizationSectionModel 区间
type OrganizationSectionModel struct {
	models.BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(6);NOT NULL;COMMENT:区间代码;" json:"unique_code"` // H07675
	Name                     string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:区间名称;" json:"name"`
	BeEnable                 bool                      `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUUID string                    `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属车间;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop     OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUUID string                    `gorm:"type:CHAR(36);COMMENT:所护工区;" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:所属工区;" json:"organization_work_area"`
}

// TableName 表名称
func (cls *OrganizationSectionModel) TableName() string {
	return "organization_sections"
}