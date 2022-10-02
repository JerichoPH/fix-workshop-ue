package models

// PivotLocationLineAndLocationRailroadGradeCrossModel 线别对道口多对多
type PivotLocationLineAndLocationRailroadGradeCrossModel struct {
	LocationLineId               uint64                          `json:"location_line_id"`
	LocationLine                 LocationLineModel               `json:"location_line"`
	LocationRailroadGradeCrossId uint64                          `json:"location_railroad_grade_cross_id"`
	LocationRailroadGradeCross   LocationRailroadGradeCrossModel `json:"location_railroad_grade_cross"`
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationRailroadGradeCrossModel
//  @param db
//  @return string
func (PivotLocationLineAndLocationRailroadGradeCrossModel) TableName() string {
	return "pivot_location_line_and_location_railroad_grade_crosses"
}
