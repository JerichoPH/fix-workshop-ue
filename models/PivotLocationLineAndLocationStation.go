package models

import "gorm.io/gorm"

// PivotLocationLineAndLocationStation 线别对站场多对多
type PivotLocationLineAndLocationStation struct {
	LocationLineId    uint64               `json:"location_line_id"`
	LocationLine      LocationLineModel    `json:"location_line"`
	LocationStationId uint64               `json:"location_station_id"`
	LocationStation   LocationStationModel `json:"location_station"`
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationStation
//  @param db
//  @return string
func (PivotLocationLineAndLocationStation) TableName(db *gorm.DB) string {
	return "pivot_location_line_and_location_stations"
}
