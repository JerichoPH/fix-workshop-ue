package models

type OrganizationParagraphModel struct {
	BaseModel
	UniqueCode                    string                      `gorm:"type:CHAR(4);UNIQUE;NOT NULL;COMMENT:站段代码;" json:"unique_code"` // B049
	Name                          string                      `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:站段名称;" json:"name"`
	ShotName                      string                      `gorm:"type:VARCHAR(64);COMMENT:站段简称;" json:"shot_name"`
	BeEnable                      bool                        `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationRailwayUniqueCode string                      `gorm:"type:CHAR(3);COMMENT:所属路局;" json:"organization_railway_unique_code"`
	OrganizationRailway           OrganizationRailwayModel    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationRailwayUniqueCode;references:UniqueCode;COMMENT:所属路局;" json:"organization_railway"`
	OrganizationWorkshops         []OrganizationWorkshopModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationParagraphUniqueCode;references:UniqueCode;COMMENT:相关车间;" json:"organization_workshops"`
	EntireInstances               []EntireInstanceModel       `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationParagraphUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *OrganizationParagraphModel) TableName() string {
	return "organization_paragraphs"
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *OrganizationParagraphModel) FindOneByUniqueCode(uniqueCode string) (organizationParagraph OrganizationParagraphModel) {
	cls.Preare().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		Find(&organizationParagraph)

	return
}
