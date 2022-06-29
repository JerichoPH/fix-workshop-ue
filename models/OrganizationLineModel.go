package models

type OrganizationLineModel struct {
	BaseModel
	UniqueCode           string                      `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:线别代码;" json:"unique_code"` // E0001
	Name                 string                      `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:线别名称;" json:"name"`
	OrganizationStations []*OrganizationStationModel `gorm:"many2many:PivotLineStations;COMMENT:线别与车站多对多;" json:"organization_stations"`
}

// TableName 表名称
func (cls *OrganizationLineModel) TableName() string {
	return "OrganizationLines"
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *OrganizationLineModel) FindOneByUniqueCode(uniqueCode string) (organizationLine OrganizationLineModel) {
	cls.Boot().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&organizationLine)

	return
}
