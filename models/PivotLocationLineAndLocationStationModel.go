package models

// PivotLocationLineAndLocationStationModel 线别对站场多对多
type PivotLocationLineAndLocationStationModel struct {
	Id                  uint64               `json:"id"`
	LocationLineUuid    string               `json:"location_line_uuid"`
	LocationLine        LocationLineModel    `json:"location_line"`
	LocationStationUuid string               `json:"location_station_uuid"`
	LocationStation     LocationStationModel `json:"location_station"`
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationStationModel
//  @param db
//  @return string
func (PivotLocationLineAndLocationStationModel) TableName() string {
	return "pivot_location_line_and_location_stations"
}
