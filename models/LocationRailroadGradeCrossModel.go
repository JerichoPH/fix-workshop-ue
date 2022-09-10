package models

// LocationRailroadGradeCrossModel 道口
type LocationRailroadGradeCrossModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(5);COMMENT:道口代码;" json:"unique_code"` // I0100
	Name                     string                    `gorm:"type:VARCHAR(64);COMMENT:道口名称;" json:"name"`
	BeEnable                 bool                      `gorm:"type:BOOLEAN;DEFAULT:0;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUuid string                    `gorm:"type:VARCHAR(36);COMMENT:所属车间UUID;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUuid string                    `gorm:"type:VARCHAR(36);COMMENT:所属工区UUID;" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaUuid;references:Uuid;COMMENT:所属工区;" json:"organization_work_area"`
	LocationLines            []*LocationLineModel      `gorm:"many2many:pivot_location_line_and_location_railroad_grade_crosses;foreignKey:id;joinForeignKey:location_railroad_grade_crossroad_id;references:id;joinReferences:location_line_id;COMMENT:线别与道口多对多;" join:"location_lines"`
}

// TableName 表名称
func (LocationRailroadGradeCrossModel) TableName() string {
	return "location_railroad_grade_crosses"
}
