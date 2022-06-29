package models

type LocationInstallPositionModel struct {
	BaseModel
	UniqueCode                    string                   `gorm:"type:CHAR(15);UNIQUE;NOT NULL;COMMENT:位代码;" json:"unique_code"`
	Name                          string                   `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:位名称;" json:"name"`
	LocationInstallTierUniqueCode string                   `gorm:"type:CHAR(13);NOT NULL;COMMENT:所属层代码;" json:"location_install_tier_unique_code"`
	LocationInstallTier           LocationInstallTierModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallTierUniqueCode;references:UniqueCode;COMMENT:所属层;" json:"location_install_tier"`
}


// TableName 表名称
func (cls *LocationInstallPositionModel) TableName() string {
	return "LocationInstallPositions"
}
