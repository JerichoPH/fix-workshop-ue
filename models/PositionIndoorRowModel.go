package models

type PositionIndoorRowModel struct {
	BaseModel
	UniqueCode             string                       `gorm:"type:CHAR(9);COMMENT:排代码;" json:"unique_code"`
	Name                   string                       `gorm:"type:VARCHAR(64);COMMENT:排名称;" json:"name"`
	PositionIndoorRoomUUID string                       `gorm:"type:VARCHAR(36);" json:"position_install_room_unique_code"`
	PositionIndoorRoom     PositionIndoorRoomModel      `gorm:"foreignKey:PositionIndoorRoomUUID;references:UUID;COMMENT:所属机房;" json:"position_indoor_room"`
	PositionIndoorCabinets []PositionIndoorCabinetModel `gorm:"foreignKey:PositionIndoorRowUUID;references:UUID;COMMENT:相关柜架;" json:"position_indoor_cabinets"`
}

// TableName 表名称
func (PositionIndoorRowModel) TableName() string {
	return "position_indoor_rows"
}
