package models

type OrganizationLine struct {
	BaseModel
	UniqueCode           string                 `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:线别代码;" json:"unique_code"` // E0001
	Name                 string                 `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:线别名称;" json:"name"`
	OrganizationStations []*OrganizationStation `gorm:"many2many:pivot_line_stations;COMMENT:线别与车站多对多;" json:"organization_stations"`
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *OrganizationLine) FindOneByUniqueCode(uniqueCode string) (organizationLine OrganizationLine) {
	cls.Boot().
		Where(map[string]interface{}{"unique_code": uniqueCode}).
		First(&organizationLine)

	return
}
