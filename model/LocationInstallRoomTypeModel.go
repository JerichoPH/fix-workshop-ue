package model

type LocationInstallRoomTypeModel struct {
	BaseModel
	UniqueCode           string                     `gorm:"type:CHAR(2);UNIQUE;NOT NULL;COMMENT:机房类型代码;" json:"unique_code"`
	Name                 string                     `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:机房类型名称;" json:"name"`
	LocationInstallRooms []LocationInstallRoomModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallRoomTypeUniqueCode;references:UniqueCode;COMMENT:相关机房;" json:"location_install_rooms"`
}

// TableName 表名称
func (cls *LocationInstallRoomTypeModel) TableName() string {
	return "location_install_room_types"
}

// FindOneByUniqueCode 根据unique_code
func (cls *LocationInstallRoomTypeModel) FindOneByUniqueCode(uniqueCode string) (locationInstallRoomType LocationInstallRoomTypeModel) {
	cls.Boot().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&locationInstallRoomType)

	return
}
