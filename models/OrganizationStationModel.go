package models

type OrganizationStationModel struct {
	BaseModel
	UniqueCode                     string                     `gorm:"type:CHAR(6);UNIQUE;NOT NULL;COMMENT:站场代码;" json:"unique_code"` // G00001
	Name                           string                     `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:站场名称;" json:"name"`
	BeEnable                       bool                       `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUniqueCode string                     `gorm:"type:CHAR(7);COMMENT:所属车间代码;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop           OrganizationWorkshopModel  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUniqueCode string                     `gorm:"type:CHAR(8);COMMENT:所属工区代码;" json:"organization_work_area_unique_code"`
	OrganizationWorkArea           OrganizationWorkAreaModel  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkAreaUniqueCode;references:UniqueCode;COMMENT:所属工区;" json:"organization_work_area"`
	OrganizationLines              []*OrganizationLineModel   `gorm:"many2many:pivot_line_stations;COMMENT:线别与车站多对多;"`
	LocationInstallRooms           []LocationInstallRoomModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationStationUniqueCode;references:UniqueCode;COMMENT:相关机房;" json:"location_install_rooms"`
}

// TableName 表名称
func (cls *OrganizationStationModel) TableName() string {
	return "organization_stations"
}