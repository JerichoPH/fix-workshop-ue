package models

type LocationIndoorRoomTypeModel struct {
	BaseModel
	UniqueCode           string                    `gorm:"type:CHAR(2);UNIQUE;NOT NULL;COMMENT:机房类型代码;" json:"unique_code"`
	Name                 string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:机房类型名称;" json:"name"`
	LocationInstallRooms []LocationIndoorRoomModel `gorm:"foreignKey:LocationIndoorRoomTypeUUID;references:UUID;COMMENT:相关机房;" json:"location_install_rooms"`
}

// TableName 表名称
func (cls *LocationIndoorRoomTypeModel) TableName() string {
	return "location_indoor_room_types"
}
