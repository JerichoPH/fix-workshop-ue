package models

import "gorm.io/gorm"

// PivotLocationLineAndLocationSection 线别对区间多对多
type PivotLocationLineAndLocationSection struct {
	LocationLineID    uint64
	LocationLine      LocationLineModel
	LocationSectionID uint64
	LocationSection   LocationSectionModel
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationSection
//  @param db
//  @return string
func (PivotLocationLineAndLocationSection) TableName(db *gorm.DB) string {
	return "pivot_location_line_and_location_sections"
}
