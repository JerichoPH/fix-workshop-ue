package models

type OrganizationParagraphModel struct {
	BaseModel
	UniqueCode              string                      `gorm:"type:CHAR(4);COMMENT:站段代码;" json:"unique_code"` // B049
	Name                    string                      `gorm:"type:VARCHAR(64);COMMENT:站段名称;" json:"name"`
	ShortName               string                      `gorm:"type:VARCHAR(64);COMMENT:站段简称;" json:"short_name"`
	BeEnable                bool                        `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationRailwayUUID string                      `gorm:"type:VARCHAR(36);COMMENT:所属路局uuid;" json:"organization_railway_uuid"`
	OrganizationRailway     OrganizationRailwayModel    `gorm:"foreignKey:OrganizationRailwayUUID;references:UUID;COMMENT:所属路局;" json:"organization_railway"`
	OrganizationWorkshops   []OrganizationWorkshopModel `gorm:"foreignKey:OrganizationParagraphUUID;references:UUID;COMMENT:相关车间;" json:"organization_workshops"`
}

// TableName 表名称
func (OrganizationParagraphModel) TableName() string {
	return "organization_paragraphs"
}
