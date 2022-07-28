package models

type OrganizationStationModel struct {
	BaseModel
	UniqueCode               string                    `gorm:"type:CHAR(6);UNIQUE;NOT NULL;COMMENT:站场代码;" json:"unique_code"` // G00001
	Name                     string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:站场名称;" json:"name"`
	BeEnable                 bool                      `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationWorkshopUUID string                    `gorm:"type:CHAR(36);COMMENT:所属车间uuid;" json:"organization_workshop_uuid"`
	OrganizationWorkshop     OrganizationWorkshopModel `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUUID string                    `gorm:"type:CHAR(36);COMMENT:所属工区uuid;" json:"organization_work_area_uuid"`
	OrganizationWorkArea     OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:所属工区;" json:"organization_work_area"`
	OrganizationLines        []*OrganizationLineModel  `gorm:"many2many:pivot_organization_line_and_organization_stations;foreignKey:id;joinForeignKey:organization_station_id;references:id;joinReferences:organization_line_id;COMMENT:线别与车站多对多;" json:"organization_lines"`
	LocationIndoorRooms      []LocationIndoorRoomModel `gorm:"foreignKey:OrganizationStationUUID;references:UUID;COMMENT:所属站场;" json:"organization_station"`
}

// TableName 表名称
func (OrganizationStationModel) TableName() string {
	return "organization_stations"
}
