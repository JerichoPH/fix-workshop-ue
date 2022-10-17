package models

type LocationCenterModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(6);COMMENT:中心代码;" json:"unique_code"` // A12F01
	Name                     string                    `gorm:"type:VARCHAR(64);COMMENT:中心名称;" json:"name"`
	BeEnable                 bool                      `gorm:"type:BOOLEAN;DEFAULT:0;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUuid string                    `gorm:"type:VARCHAR(36);COMMENT:所属车间UUID;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUuid string                    `gorm:"type:VARCHAR(36);COMMENT:所属工区UUID;" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaUuid;references:Uuid;COMMENT:所属工区;" json:"organization_work_area"`
	LocationLineUuid         string                    `gorm:"type:CHAR(36);COMMENT:线别Uuid;" json:"location_line_uuid"`
	LocationLine             LocationLineModel         `gorm:"foreignKey:LocationLineUuid;references:Uuid;COMMENT:所属线别;" json:"location_line"`
	LocationIndoorRooms      []PositionIndoorRoomModel `gorm:"foreignKey:LocationCenterUuid;references:Uuid;COMMENT:相关机房;" json:"location_indoor_rooms"`
}

// TableName 表名称
func (LocationCenterModel) TableName() string {
	return "location_centers"
}
