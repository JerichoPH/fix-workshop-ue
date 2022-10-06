package models

// PivotLocationLineAndLocationRailroadGradeCrossModel 线别对道口多对多
type PivotLocationLineAndLocationRailroadGradeCrossModel struct {
	Id                 uint64                `json:"id"`
	LocationLineId     uint64                `json:"location_line_uuid"`
	LocationLine       LocationLineModel     `json:"location_line"`
	LocationRailroadId uint64                `json:"location_railroad_uuid"`
	LocationRailroad   LocationRailroadModel `json:"location_railroad"`
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationRailroadGradeCrossModel
//  @param db
//  @return string
func (PivotLocationLineAndLocationRailroadGradeCrossModel) TableName() string {
	return "pivot_location_line_and_location_railroads"
}
