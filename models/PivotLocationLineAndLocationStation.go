package models

import "gorm.io/gorm"

// PivotLocationLineAndLocationStation 线别对站场多对多
type PivotLocationLineAndLocationStation struct {
	LocationLineID    uint64
	LocationLine      LocationLineModel
	LocationStationID uint64
	LocationStation   LocationStationModel
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationStation
//  @param db
//  @return string
func (PivotLocationLineAndLocationStation) TableName(db *gorm.DB) string {
	return "pivot_location_line_and_location_stations"
}
