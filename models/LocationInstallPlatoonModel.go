package models

type LocationInstallPlatoonModel struct {
	BaseModel
	UniqueCode                    string                      `gorm:"type:CHAR(9);UNIQUE;NOT NULL;COMMENT:排代码;" json:"unique_code"`
	Name                          string                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:排名称;" json:"name"`
	LocationInstallRoomUniqueCode string                      `gorm:"type:CHAR(7);NOT NULL;" json:"location_install_room_unique_code"`
	LocationInstallRoom           LocationInstallRoomModel    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallRoomUniqueCode;references:UniqueCode;COMMENT:所属机房;" json:"location_install_room"`
	LocationInstallShelves        []LocationInstallShelfModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallPlatoonUniqueCode;references:UniqueCode;COMMENT:相关柜架;" json:"location_install_shelves"`
}

// TableName 表名称
func (cls *LocationInstallPlatoonModel) TableName() string {
	return "location_install_platoons"
}
