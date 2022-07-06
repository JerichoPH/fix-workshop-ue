package models

type LocationInstallRoomModel struct {
	BaseModel
	UniqueCode                        string                       `gorm:"type:CHAR(7);UNIQUE;NOT NULL;COMMENT:机房代码;" json:"unique_code"`
	Name                              string                       `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:机房名称;" json:"name"`
	LocationInstallRoomTypeUniqueCode string                       `gorm:"type:CHAR(2);COMMENT:所属机房类型;" json:"location_install_room_type_unique_code"`
	LocationInstallRoomType           LocationInstallRoomTypeModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallRoomTypeUniqueCode;references:UniqueCode;COMMENT:所属机房类型;" json:"location_install_room_type"`
	OrganizationStationUniqueCode     string                       `gorm:"type:CHAR(6);COMMENT:所属车站代码;" json:"organization_station_unique_code"`
	OrganizationStation               OrganizationStationModel     `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationStationUniqueCode;references:UniqueCode;COMMENT:所属车站;" json:"organization_station"`
	LocationInstallPlatoons           []LocationInstallPlatoonModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallRoomUniqueCode;references:UniqueCode;COMMENT:相关排;" json:"location_install_platoons"`
}

// TableName 表名称
func (cls *LocationInstallRoomModel) TableName() string {
	return "location_install_rooms"
}

// FindOneByUniqueCode 根据unique_code
func (cls *LocationInstallRoomModel) FindOneByUniqueCode(uniqueCode string) (locationInstallRoom LocationInstallRoomModel) {
	cls.Boot().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&locationInstallRoom)

	return
}
