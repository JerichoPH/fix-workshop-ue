package models

type LocationIndoorRoomModel struct {
	BaseModel
	UniqueCode                 string                      `gorm:"type:CHAR(7);UNIQUE;NOT NULL;COMMENT:机房代码;" json:"unique_code"`
	Name                       string                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:机房名称;" json:"name"`
	LocationIndoorRoomTypeUUID string                      `gorm:"type:CHAR(36);COMMENT:所属机房类型;" json:"location_indoor_room_type_uuid"`
	LocationIndoorRoomType     LocationIndoorRoomTypeModel `gorm:"foreignKey:LocationIndoorRoomTypeUUID;references:UUID;COMMENT:所属机房类型;" json:"location_indoor_room_type"`
	OrganizationStationUUID    string                      `gorm:"type:CHAR(36);COMMENT:所属站场UUID;" json:"organization_station_uuid"`
	OrganizationStation        OrganizationStationModel    `gorm:"foreignKey:OrganizationStationUUID;references:UUID;COMMENT:所属站场;" json:"organization_station"`
	OrganizationSectionUUID    string                      `gorm:"type:CHAR(36);COMMENT:所属区间UUID;" json:"organization_section_uuid"`
	OrganizationSection        OrganizationSectionModel    `gorm:"foreignKey:OrganizationSectionUUID;references:UUID;COMMENT:所属区间;" json:"organization_section"`
	OrganizationCenterUUID     string                      `gorm:"type:CHAR(36);COMMENT:所属中心UUID;" json:"organization_center_uuid"`
	OrganizationCenter         OrganizationCenterModel     `gorm:"foreignKey:OrganizationCenterUUID;references:UUID;COMMENT:所属中心;" json:"organization_center"`
	OwnerType                  string                      `gorm:"type:VARCHAR(64);COMMENT:所属上级类型;" json:"owner_type"`
	LocationIndoorRows         []LocationIndoorRowModel    `gorm:"foreignKey:LocationIndoorRoomUUID;references:UUID;COMMENT:相关排;" json:"location_indoor_rows"`
}

// TableName 表名称
func (LocationIndoorRoomModel) TableName() string {
	return "location_indoor_rooms"
}
