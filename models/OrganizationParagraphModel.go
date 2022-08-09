package models

type OrganizationParagraphModel struct {
	BaseModel
	UniqueCode              string                      `gorm:"type:CHAR(4);UNIQUE;NOT NULL;COMMENT:站段代码;" json:"unique_code"` // B049
	Name                    string                      `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:站段名称;" json:"name"`
	ShortName               string                      `gorm:"type:VARCHAR(64);COMMENT:站段简称;" json:"short_name"`
	BeEnable                bool                        `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationRailwayUUID string                      `gorm:"type:CHAR(36);NOT NULL;COMMENT:所属路局uuid;" json:"organization_railway_uuid"`
	OrganizationRailway     OrganizationRailwayModel    `gorm:"foreignKey:OrganizationRailwayUUID;references:UUID;COMMENT:所属路局;" json:"organization_railway"`
	OrganizationWorkshops   []OrganizationWorkshopModel `gorm:"foreignKey:OrganizationParagraphUUID;references:UUID;COMMENT:相关车间;" json:"organization_workshops"`
	LocationLines           []*LocationLineModel        `gorm:"many2many:pivot_location_line_and_organization_paragraphs;foreignKey:id;joinForeignKey:organization_paragraph_id;references:id;joinReferences:location_line_id;COMMENT:线别与站段多对多;" json:"location_lines"`
	EntireInstances         []EntireInstanceModel       `gorm:"foreignKey:OrganizationParagraphUUID;references:UUID;COMMENT:相关器材;" json:"entire_instances"`
	//PositionDepotStorehouse []PositionDepotStorehouseModel             `gorm:"foreignKey:OrganizationParagraphUUID;references:UUID;NOT NULL;COMMENT:相关仓库;" json:"location_depots"`
}

// TableName 表名称
func (OrganizationParagraphModel) TableName() string {
	return "organization_paragraphs"
}
