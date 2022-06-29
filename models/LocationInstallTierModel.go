package models

type LocationInstallTierModel struct {
	BaseModel
	UniqueCode                     string                         `gorm:"type:CHAR(13);UNIQUE;NOT NULL;COMMENT:层代码;" json:"unique_code"`
	Name                           string                         `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:层名称;" json:"name"`
	LocationInstallShelfUniqueCode string                         `gorm:"type:CHAR(11);NOT NULL;COMMENT:所属柜架代码;" json:"location_install_shelf_unique_code"`
	LocationInstallShelf           LocationInstallShelfModel      `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallShelfUniqueCode;references:UniqueCode;COMMENT:所属柜架;" json:"location_install_shelf"`
	LocationInstallPositions       []LocationInstallPositionModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallTierUniqueCode;references:UniqueCode;COMMENT:相关位;" json:"location_install_tier"`
}

// TableName 表名称
func (cls *LocationInstallTierModel) TableName() string {
	return "locationInstallTiers"
}