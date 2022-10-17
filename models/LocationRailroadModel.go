package models

// LocationRailroadModel 道口
type LocationRailroadModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(5);COMMENT:道口代码;" json:"unique_code"` // I0100
	Name                     string                    `gorm:"type:VARCHAR(64);COMMENT:道口名称;" json:"name"`
	BeEnable                 bool                      `gorm:"type:BOOLEAN;DEFAULT:0;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUuid string                    `gorm:"type:VARCHAR(36);COMMENT:所属车间UUID;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUuid string                    `gorm:"type:VARCHAR(36);COMMENT:所属工区UUID;" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaUuid;references:Uuid;COMMENT:所属工区;" json:"organization_work_area"`
	LocationIndoorRooms      []PositionIndoorRoomModel `gorm:"foreignKey:LocationRailroadUuid;COMMENT:所属道口;" json:"location_indoor_rooms"`
	LocationLineUuid         string                    `gorm:"type:CHAR(36);COMMENT:线别Uuid;" json:"location_line_uuid"`
	LocationLine             LocationLineModel         `gorm:"foreignKey:LocationLineUuid;references:Uuid;COMMENT:所属线别;" json:"location_line"`
}

// TableName 表名称
func (LocationRailroadModel) TableName() string {
	return "location_railroad_grade_crosses"
}