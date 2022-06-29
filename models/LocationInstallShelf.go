package models

type LocationInstallShelf struct {
	BaseModel
	UniqueCode                       string                 `gorm:"type:CHAR(11);UNIQUE;NOT NULL;COMMENT:柜架代码;" json:"unique_code"`
	Name                             string                 `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:柜架名称;" json:"name"`
	LocationInstallPlatoonUniqueCode string                 `gorm:"type:CHAR(9);NOT NULL;COMMENT:所属排;" json:"location_install_platoon_unique_code"`
	LocationInstallPlatoon           LocationInstallPlatoon `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallPlatoonUniqueCode;references:UniqueCode;COMMENT:所属排;" json:"location_install_platoon"`
	LocationInstallTiers             []LocationInstallTier  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallShelfUniqueCode;references:UniqueCode;COMMENT:相关层;" json:"location_install_tiers"`
}
