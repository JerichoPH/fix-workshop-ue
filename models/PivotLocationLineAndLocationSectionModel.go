package models

// PivotLocationLineAndLocationSectionModel 线别对区间多对多
type PivotLocationLineAndLocationSectionModel struct {
	Id                  uint64               `json:"id"`
	LocationLineUuid    string               `json:"location_line_uuid"`
	LocationLine        LocationLineModel    `json:"location_line"`
	LocationSectionUuid string               `json:"location_section_uuid"`
	LocationSection     LocationSectionModel `json:"location_section"`
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationSectionModel
//  @param db
//  @return string
func (PivotLocationLineAndLocationSectionModel) TableName() string {
	return "pivot_location_line_and_location_sections"
}
