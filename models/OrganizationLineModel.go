package models

type OrganizationLineModel struct {
	BaseModel
	UniqueCode           string                      `gorm:"type:CHAR(5);UNIQUE;NOT NULL;COMMENT:线别代码;" json:"unique_code"` // E0001
	Name                 string                      `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:线别名称;" json:"name"`
	OrganizationStations []*OrganizationStationModel `gorm:"many2many:pivot_line_stations;COMMENT:线别与车站多对多;" json:"organization_stations"`
}

// TableName 表名称
func (cls *OrganizationLineModel) TableName() string {
	return "organization_lines"
}