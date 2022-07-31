package models

type OrganizationLineModel struct {
	BaseModel
	UniqueCode                       string                                 `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:线别代码;" json:"unique_code"` // E0001
	Name                             string                                 `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:线别名称;" json:"name"`
	BeEnable                         bool                                   `gorm:"type:BOOLEAN;NOT NULL;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationRailways             []*OrganizationRailwayModel            `gorm:"many2many:pivot_organization_line_and_organization_railways;foreignKey:id;joinForeignKey:organization_line_id;references:id;joinReferences:organization_railway_id;COMMENT:线别与路局多对多;" json:"organization_railways"`
	OrganizationParagraphs           []*OrganizationParagraphModel          `gorm:"many2many:pivot_organization_line_and_organization_paragraphs;foreignKey:id;joinForeignKey:organization_line_id;references:id;joinReferences:organization_paragraph_id;COMMENT:线别与站段多对多;" json:"organization_paragraphs"`
	OrganizationWorkshops            []*OrganizationWorkshopModel           `gorm:"many2many:pivot_organization_line_and_organization_workshops;foreignKey:id;joinForeignKey:organization_line_id;references:id;organization_workshop_id;COMMENT:线别与车间多对多;" json:"organization_workshops"`
	OrganizationWorkAreas            []*OrganizationWorkAreaModel           `gorm:"many2many:pivot_organization_line_and_organization_work_areas;foreignKey:id;joinForeignKey:organization_line_id;references:id;joinReferences:organization_work_area_id;COMMENT:线别与工区多对多;" json:"organization_work_areas"`
	OrganizationSections             []*OrganizationSectionModel            `gorm:"many2many:pivot_organization_line_and_organization_sections;foreignKey:id;joinForeignKey:organization_line_id;references:id;joinReferences:organization_section_id;COMMENT:线别与区间多对多;" json:"organization_sections"`
	OrganizationStations             []*OrganizationStationModel            `gorm:"many2many:pivot_organization_line_and_organization_stations;foreignKey:id;joinForeignKey:organization_line_id;references:id;joinReferences:organization_station_id;COMMENT:线别与车站多对多;" json:"organization_stations"`
	OrganizationRailroadGradeCrosses []*OrganizationRailroadGradeCrossModel `gorm:"many2many:pivot_organization_line_and_organization_railroad_grade_crosses;foreignKey:id;joinForeignKey:organization_line_id;references:id;joinReferences:organization_railroad_grade_cross_id;COMMENT:线别与道口多对多;" join:"organization_railroad_grade_crosses"`
	OrganizationCenters              []*OrganizationCenterModel             `gorm:"many2many:pivot_organization_line_and_organization_centers;foreignKey:id;joinForeignKey:organization_line_id;references:id;joinReferences:organization_center_id;COMMENT:线别与中心多对多;" json:"organization_centers"`
}

// TableName 表名称
func (OrganizationLineModel) TableName() string {
	return "organization_lines"
}
