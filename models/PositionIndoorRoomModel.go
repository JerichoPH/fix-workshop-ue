package models

type PositionIndoorRoomModel struct {
	BaseModel
	UniqueCode                 string                      `gorm:"type:CHAR(7);COMMENT:机房代码;" json:"unique_code"`
	Name                       string                      `gorm:"type:VARCHAR(64);COMMENT:机房名称;" json:"name"`
	PositionIndoorRoomTypeUuid string                      `gorm:"type:VARCHAR(36);COMMENT:所属机房类型;" json:"position_indoor_room_type_uuid"`
	PositionIndoorRoomType     PositionIndoorRoomTypeModel `gorm:"foreignKey:PositionIndoorRoomTypeUuid;references:Uuid;COMMENT:所属机房类型;" json:"position_indoor_room_type"`
	LocationStationUuid        string                      `gorm:"type:VARCHAR(36);COMMENT:所属站场UUID;" json:"organization_station_uuid"`
	LocationStation            LocationStationModel        `gorm:"foreignKey:LocationStationUuid;references:Uuid;COMMENT:所属站场;" json:"location_station"`
	LocationSectionUuid        string                      `gorm:"type:VARCHAR(36);COMMENT:所属区间UUID;" json:"organization_section_uuid"`
	LocationSection            LocationSectionModel        `gorm:"foreignKey:LocationSectionUuid;references:Uuid;COMMENT:所属区间;" json:"location_section"`
	LocationCenterUuid         string                      `gorm:"type:VARCHAR(36);COMMENT:所属中心UUID;" json:"organization_center_uuid"`
	LocationCenter             LocationCenterModel         `gorm:"foreignKey:LocationCenterUuid;references:Uuid;COMMENT:所属中心;" json:"location_center"`
	LocationRailroadUuid       string                      `gorm:"type:VARCHAR(36);COMMENT:所属道口UUID;" json:"location_railroad_uuid"`
	LocationRailroad           LocationRailroadModel       `gorm:"foreignKey:LocationRailroadUuid;COMMENT:所属道口;" json:"location_railroad"`
	PositionIndoorRows         []PositionIndoorRowModel    `gorm:"foreignKey:PositionIndoorRoomUuid;references:Uuid;COMMENT:相关排;" json:"position_indoor_rows"`
	OwnerType                  string                      `gorm:"type:VARCHAR(64);COMMENT:归属类型;" json:"owner_type"`
}

// TableName 表名称
func (PositionIndoorRoomModel) TableName() string {
	return "position_indoor_rooms"
}
