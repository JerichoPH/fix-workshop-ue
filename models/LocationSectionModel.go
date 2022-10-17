package models

// LocationSectionModel 区间
type LocationSectionModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(6);COMMENT:区间代码;" json:"unique_code"` // H07675
	Name                     string                    `gorm:"type:VARCHAR(64);COMMENT:区间名称;" json:"name"`
	BeEnable                 bool                      `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUuid string                    `gorm:"type:VARCHAR(36);COMMENT:所属车间;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop     OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUuid string                    `gorm:"type:VARCHAR(36);COMMENT:所护工区;" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaUuid;references:Uuid;COMMENT:所属工区;" json:"organization_work_area"`
	LocationLineUuid         string                    `gorm:"type:CHAR(36);COMMENT:线别Uuid;" json:"location_line_uuid"`
	LocationLine             LocationLineModel         `gorm:"foreignKey:LocationLineUuid;references:Uuid;COMMENT:所属线别;" json:"location_line"`
	LocationIndoorRooms      []PositionIndoorRoomModel `gorm:"foreignKey:LocationSectionUuid;references:Uuid;COMMENT:相关机房;" json:"location_indoor_rooms"`
}

// TableName 表名称
func (LocationSectionModel) TableName() string {
	return "location_sections"
}
