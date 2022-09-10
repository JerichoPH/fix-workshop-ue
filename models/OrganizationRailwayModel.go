package models

// OrganizationRailwayModel 路局
type OrganizationRailwayModel struct {
	BaseModel
	UniqueCode             string                       `gorm:"type:CHAR(3);COMMENT:路局代码;" json:"unique_code"` // A12
	Name                   string                       `gorm:"type:VARCHAR(64);COMMENT:路局名称;" json:"name"`
	ShortName              string                       `gorm:"type:VARCHAR(64);COMMENT:路局简称;" json:"short_name"`
	BeEnable               bool                         `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationParagraphs []OrganizationParagraphModel `gorm:"foreignKey:OrganizationRailwayUuid;references:Uuid;COMMENT:相关站段;" json:"organization_paragraphs"`
}

// TableName 表名称
func (OrganizationRailwayModel) TableName() string {
	return "organization_railways"
}
