package models

type LocationCenterModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(6);UNIQUE;NOT NULL;COMMENT:中心代码;"` // A12F01
	Name                     string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:中心名称;"`
	BeEnable                 bool                      `gorm:"type:BOOLEAN;NOT NULL;DEFAULT:0;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUUID string                    `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属车间UUID;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUUID string                    `gorm:"type:CHAR(36);COMMENT:所属工区UUID;" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:所属工区;" json:"organization_work_area"`
	LocationLines            []*LocationLineModel      `gorm:"many2many:pivot_location_line_and_location_centers;foreignKey:id;joinForeignKey:location_line_id;references:id;joinReferences:location_center_id;COMMENT:线别与中心多对多;" json:"location_centers"`
	LocationIndoorRooms      []PositionIndoorRoomModel `gorm:"foreignKey:LocationCenterUUID;references:UUID;COMMENT:相关机房;" json:"location_indoor_rooms"`
}

// TableName 表名称
func (LocationCenterModel) TableName() string {
	return "location_centers"
}
