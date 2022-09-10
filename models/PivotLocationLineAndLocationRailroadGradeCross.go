package models

import "gorm.io/gorm"

// PivotLocationLineAndLocationRailroadGradeCross 线别对道口多对多
type PivotLocationLineAndLocationRailroadGradeCross struct {
	LocationLineId               uint64                          `json:"location_line_id"`
	LocationLine                 LocationLineModel               `json:"location_line"`
	LocationRailroadGradeCrossId uint64                          `json:"location_railroad_grade_cross_id"`
	LocationRailroadGradeCross   LocationRailroadGradeCrossModel `json:"location_railroad_grade_cross"`
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationRailroadGradeCross
//  @param db
//  @return string
func (PivotLocationLineAndLocationRailroadGradeCross) TableName(db *gorm.DB) string {
	return "pivot_location_line_and_location_railroad_grade_crosses"
}
