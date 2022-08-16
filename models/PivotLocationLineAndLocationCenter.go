package models

import "gorm.io/gorm"

// PivotLocationLineAndLocationCenter 线别对中心多对多
type PivotLocationLineAndLocationCenter struct {
	LocationLineID   uint64
	LocationLine     LocationLineModel
	LocationCenterID uint64
	LocationCenter   LocationCenterModel
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationCenter
//  @param db
//  @return string
func (PivotLocationLineAndLocationCenter) TableName(db *gorm.DB) string {
	return "pivot_location_line_and_location_centers"
}
