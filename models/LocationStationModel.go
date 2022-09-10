package models

type LocationStationModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(6);COMMENT:站场代码;" json:"unique_code"` // G00001
	Name                     string                    `gorm:"type:VARCHAR(64);COMMENT:站场名称;" json:"name"`
	BeEnable                 bool                      `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUuid string                    `gorm:"type:VARCHAR(36);COMMENT:所属车间uuid;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUuid string                    `gorm:"type:VARCHAR(36);COMMENT:所属工区uuid;" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaUuid;references:Uuid;COMMENT:所属工区;" json:"organization_work_area"`
	LocationLines            []*LocationLineModel      `gorm:"many2many:pivot_location_line_and_location_stations;foreignKey:id;joinForeignKey:location_station_id;references:id;joinReferences:location_line_id;COMMENT:线别与车站多对多;" json:"location_lines"`
	LocationIndoorRooms      []PositionIndoorRoomModel `gorm:"foreignKey:LocationStationUuid;references:Uuid;COMMENT:所属站场;" json:"organization_station"`
}

// TableName 表名称
func (LocationStationModel) TableName() string {
	return "location_stations"
}
