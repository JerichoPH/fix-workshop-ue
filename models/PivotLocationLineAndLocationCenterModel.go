package models

// PivotLocationLineAndLocationCenterModel 线别对中心多对多
type PivotLocationLineAndLocationCenterModel struct {
	Id                 uint64              `json:"id"`
	LocationLineUuid   string              `json:"location_line_uuid"`
	LocationLine       LocationLineModel   `json:"location_line"`
	LocationCenterUuid string              `json:"location_center_uuid"`
	LocationCenter     LocationCenterModel `json:"location_center"`
}

// TableName 表名称
//  @receiver PivotLocationLineAndLocationCenterModel
//  @param db
//  @return string
func (PivotLocationLineAndLocationCenterModel) TableName() string {
	return "pivot_location_line_and_location_centers"
}
