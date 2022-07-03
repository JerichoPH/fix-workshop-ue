package model

// OrganizationRailwayModel 路局
type OrganizationRailwayModel struct {
	BaseModel
	UniqueCode             string                       `gorm:"type:CHAR(3);UNIQUE;NOT NULL;COMMENT:路局代码;" json:"unique_code"` // A12
	Name                   string                       `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:路局名称;" json:"name"`
	ShotName               string                       `gorm:"type:VARCHAR(64);COMMENT:路局简称;" json:"shot_name"`
	BeEnable               bool                         `gorm:"type:BOOLEAN;DEFAULT:1;COMMENT:是否启用;" json:"be_enable"`
	OrganizationParagraphs []OrganizationParagraphModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationRailwayUniqueCode;references:UniqueCode;COMMENT:相关站段;" json:"organization_paragraphs"`
	EntireInstances        []EntireInstanceModel        `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationRailwayUniqueCode;references:UniqueCode;COMMENT:相关器材;" json:"entire_instances"`
}

// TableName 表名称
func (cls *OrganizationRailwayModel) TableName() string {
	return "organization_railways"
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *OrganizationRailwayModel) FindOneByUniqueCode(uniqueCode string) (organizationRailway OrganizationRailwayModel) {
	cls.Boot().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&organizationRailway)

	return
}
