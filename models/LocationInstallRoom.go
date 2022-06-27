package models

import (
	"gorm.io/gorm"
)

type LocationInstallRoom struct {
	BaseModel
	Preloads                          []string
	Selects                           []string
	Omits                             []string
	UniqueCode                        string                  `gorm:"type:CHAR(7);UNIQUE;NOT NULL;COMMENT:机房代码;" json:"unique_code"`
	Name                              string                  `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:机房名称;" json:"name"`
	LocationInstallRoomTypeUniqueCode string                  `gorm:"type:CHAR(2);COMMENT:所属机房类型;" json:"location_install_room_type_unique_code"`
	LocationInstallRoomType           LocationInstallRoomType `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallRoomTypeUniqueCode;references:UniqueCode;COMMENT:所属机房类型;" json:"location_install_room_type"`
	OrganizationStationUniqueCode     string                  `gorm:"type:CHAR(6);COMMENT:所属车站代码;" json:"organization_station_unique_code"`
	OrganizationStation               OrganizationStation     `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationStationUniqueCode;references:UniqueCode;COMMENT:所属车站;" json:"organization_station"`
	//LocationInstallPlatoons           []LocationInstallPlatoon `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallRoomUniqueCode;references:UniqueCode;COMMENT:相关;" json:"location_install_platoons"`
}

// FindOneByUniqueCode 根据unique_code
func (cls *LocationInstallRoom) FindOneByUniqueCode(db *gorm.DB, uniqueCode string) (locationInstallRoom LocationInstallRoom) {
	cls.Boot(db, cls.Preloads, cls.Selects, cls.Omits).
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&locationInstallRoom)

	return
}
