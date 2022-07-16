package models

type LocationIndoorRowModel struct {
	BaseModel
	UniqueCode             string                       `gorm:"type:CHAR(9);UNIQUE;NOT NULL;COMMENT:排代码;" json:"unique_code"`
	Name                   string                       `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:排名称;" json:"name"`
	LocationIndoorRoomUUID string                       `gorm:"type:CHAR(36);NOT NULL;" json:"location_install_room_unique_code"`
	LocationIndoorRoom     LocationIndoorRoomModel      `gorm:"foreignKey:LocationIndoorRoomUUID;references:UUID;COMMENT:所属机房;" json:"location_indoor_room"`
	LocationIndoorCabinets []LocationIndoorCabinetModel `gorm:"foreignKey:LocationIndoorRowUUID;references:UUID;COMMENT:相关柜架;" json:"location_indoor_cabinets"`
}

// TableName 表名称
func (cls *LocationIndoorRowModel) TableName() string {
	return "location_indoor_rows"
}
