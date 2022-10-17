package models

type LocationLineModel struct {
	BaseModel
	UniqueCode        string                   `gorm:"type:CHAR(5);COMMENT:线别代码;" json:"unique_code"` // E0001
	Name              string                   `gorm:"type:VARCHAR(64);COMMENT:线别名称;" json:"name"`
	BeEnable          bool                     `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	LocationSections  []*LocationSectionModel  `gorm:"foreignKey:LocationLineUuid;references:Uuid;COMMENT:相关中心;" json:"location_sections"`
	LocationStations  []*LocationStationModel  `gorm:"foreignKey:LocationLineUuid;references:Uuid;COMMENT:相关中心;" json:"location_stations"`
	LocationRailroads []*LocationRailroadModel `gorm:"foreignKey:LocationLineUuid;references:Uuid;COMMENT:相关中心;" json:"location_railroads"`
	LocationCenters   []*LocationCenterModel   `gorm:"foreignKey:LocationLineUuid;references:Uuid;COMMENT:相关中心;" json:"location_centers"`
}

// TableName 表名称
func (LocationLineModel) TableName() string {
	return "location_lines"
}
