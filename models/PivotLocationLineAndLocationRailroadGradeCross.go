package models

import "gorm.io/gorm"

// PivotLocationLineAndLocationRailroadGradeCross 线别对道口多对多
type PivotLocationLineAndLocationRailroadGradeCross struct {
	LocationLineID    uint64
	LocationLine      LocationLineModel
	LocationRailroadGradeCrossID uint64
	LocationRailroadGradeCross   LocationRailroadGradeCrossModel
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationRailroadGradeCross
//  @param db
//  @return string
func (PivotLocationLineAndLocationRailroadGradeCross) TableName(db *gorm.DB) string {
	return "pivot_location_line_and_location_railroad_grade_crosses"
}
