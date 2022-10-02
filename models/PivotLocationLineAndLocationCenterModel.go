package models

// PivotLocationLineAndLocationCenterModel 线别对中心多对多
type PivotLocationLineAndLocationCenterModel struct {
	LocationLineId   uint64              `json:"location_line_id"`
	LocationLine     LocationLineModel   `json:"location_line"`
	LocationCenterId uint64              `json:"location_center_id"`
	LocationCenter   LocationCenterModel `json:"location_center"`
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationCenterModel
//  @param db
//  @return string
func (PivotLocationLineAndLocationCenterModel) TableName() string {
	return "pivot_location_line_and_location_centers"
}
