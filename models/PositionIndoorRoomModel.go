package models

type PositionIndoorRoomModel struct {
	BaseModel
	UniqueCode                 string                      `gorm:"type:CHAR(7);COMMENT:机房代码;" json:"unique_code"`
	Name                       string                      `gorm:"type:VARCHAR(64);COMMENT:机房名称;" json:"name"`
	PositionIndoorRoomTypeUUID string                      `gorm:"type:CHAR(36);COMMENT:所属机房类型;" json:"position_indoor_room_type_uuid"`
	PositionIndoorRoomType     PositionIndoorRoomTypeModel `gorm:"foreignKey:PositionIndoorRoomTypeUUID;references:UUID;COMMENT:所属机房类型;" json:"position_indoor_room_type"`
	LocationStationUUID        string                      `gorm:"type:CHAR(36);COMMENT:所属站场UUID;" json:"organization_station_uuid"`
	LocationStation            LocationStationModel        `gorm:"foreignKey:LocationStationUUID;references:UUID;COMMENT:所属站场;" json:"location_station"`
	LocationSectionUUID        string                      `gorm:"type:CHAR(36);COMMENT:所属区间UUID;" json:"organization_section_uuid"`
	LocationSection            LocationSectionModel        `gorm:"foreignKey:LocationSectionUUID;references:UUID;COMMENT:所属区间;" json:"location_section"`
	LocationCenterUUID         string                      `gorm:"type:CHAR(36);COMMENT:所属中心UUID;" json:"organization_center_uuid"`
	LocationCenter             LocationCenterModel         `gorm:"foreignKey:LocationCenterUUID;references:UUID;COMMENT:所属中心;" json:"location_center"`
	OwnerType                  string                      `gorm:"type:VARCHAR(64);COMMENT:所属上级类型;" json:"owner_type"`
	PositionIndoorRows         []PositionIndoorRowModel    `gorm:"foreignKey:PositionIndoorRoomUUID;references:UUID;COMMENT:相关排;" json:"position_indoor_rows"`
}

// TableName 表名称
func (PositionIndoorRoomModel) TableName() string {
	return "position_indoor_rooms"
}
