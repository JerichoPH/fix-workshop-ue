package models

// OrganizationSectionModel 区间
type OrganizationSectionModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(6);NOT NULL;COMMENT:区间代码;" json:"unique_code"` // H07675
	Name                     string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:区间名称;" json:"name"`
	BeEnable                 bool                      `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUUID string                    `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属车间;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop     OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUUID string                    `gorm:"type:CHAR(36);COMMENT:所护工区;" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:所属工区;" json:"organization_work_area"`
	OrganizationLines        []*OrganizationLineModel  `gorm:"many2many:pivot_organization_line_and_organization_sections;foreignKey:id;joinForeignKey:organization_section_id;references:id;joinReferences:organization_line_id;COMMENT:线别与区间多对多;" json:"organization_lines"`
	LocationIndoorRooms      []LocationIndoorRoomModel `gorm:"foreignKey:OrganizationSectionUUID;references:UUID;COMMENT:相关机房;" json:"location_indoor_rooms"`
}

// TableName 表名称
func (OrganizationSectionModel) TableName() string {
	return "organization_sections"
}
