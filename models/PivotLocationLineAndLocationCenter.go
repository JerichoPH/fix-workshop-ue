package models

import "gorm.io/gorm"

// PivotLocationLineAndLocationCenter 线别对中心多对多
type PivotLocationLineAndLocationCenter struct {
	LocationLineId   uint64              `json:"location_line_id"`
	LocationLine     LocationLineModel   `json:"location_line"`
	LocationCenterId uint64              `json:"location_center_id"`
	LocationCenter   LocationCenterModel `json:"location_center"`
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationCenter
//  @param db
//  @return string
func (PivotLocationLineAndLocationCenter) TableName(db *gorm.DB) string {
	return "pivot_location_line_and_location_centers"
}
