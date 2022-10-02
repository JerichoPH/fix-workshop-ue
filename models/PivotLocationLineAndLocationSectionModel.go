package models

// PivotLocationLineAndLocationSectionModel 线别对区间多对多
type PivotLocationLineAndLocationSectionModel struct {
	LocationLineId    uint64               `json:"location_line_id"`
	LocationLine      LocationLineModel    `json:"location_line"`
	LocationSectionId uint64               `json:"location_section_id"`
	LocationSection   LocationSectionModel `json:"location_section"`
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationSectionModel
//  @param db
//  @return string
func (PivotLocationLineAndLocationSectionModel) TableName() string {
	return "pivot_location_line_and_location_sections"
}
