package models

type LocationLineModel struct {
	BaseModel
	UniqueCode                   string                             `gorm:"type:CHAR(5);NOT NULL;COMMENT:线别代码;" json:"unique_code"` // E0001
	Name                         string                             `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:线别名称;" json:"name"`
	BeEnable                     bool                               `gorm:"type:BOOLEAN;NOT NULL;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	LocationSections             []*LocationSectionModel            `gorm:"many2many:pivot_location_line_and_location_sections;foreignKey:id;joinForeignKey:location_line_id;references:id;joinReferences:location_section_id;COMMENT:线别与区间多对多;" json:"location_sections"`
	LocationStations             []*LocationStationModel            `gorm:"many2many:pivot_location_line_and_location_stations;foreignKey:id;joinForeignKey:location_line_id;references:id;joinReferences:location_station_id;COMMENT:线别与车站多对多;" json:"location_stations"`
	LocationRailroadGradeCrosses []*LocationRailroadGradeCrossModel `gorm:"many2many:pivot_location_line_and_location_railroad_grade_crosses;foreignKey:id;joinForeignKey:location_line_id;references:id;joinReferences:location_railroad_grade_cross_id;COMMENT:线别与道口多对多;" join:"location_railroad_grade_crosses"`
	LocationCenters              []*LocationCenterModel             `gorm:"many2many:pivot_location_line_and_location_centers;foreignKey:id;joinForeignKey:location_line_id;references:id;joinReferences:location_center_id;COMMENT:线别与中心多对多;" json:"location_centers"`
}

// TableName 表名称
func (LocationLineModel) TableName() string {
	return "location_lines"
}
