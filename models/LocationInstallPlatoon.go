package models

type LocationInstallPlatoon struct {
	BaseModel
	Preloads  []string
	Omits     []string
	Selects   []string
	UniqueCode                    string                 `gorm:"type:CHAR(9);UNIQUE;NOT NULL;COMMENT:排代码;" json:"unique_code"`
	Name                          string                 `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:排名称;" json:"name"`
	LocationInstallRoomUniqueCode string                 `gorm:"type:CHAR(7);UNIQUE;NOT NULL;" json:"location_install_room_unique_code"`
	LocationInstallRoom           LocationInstallRoom    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallRoomUniqueCode;references:UniqueCode;COMMENT:所属机房;" json:"location_install_room"`
	LocationInstallShelves        []LocationInstallShelf `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallPlatoonUniqueCode;references:UniqueCode;COMMENT:相关柜架;" json:"location_install_shelves"`
}
