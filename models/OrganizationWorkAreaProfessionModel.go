package models

// OrganizationWorkAreaProfessionModel 工区专业
type OrganizationWorkAreaProfessionModel struct {
	BaseModel
	UniqueCode string `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:工区专业代码;" json:"unique_code"`
	Name       string `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:工区专业名称;" json:"name"`
	//OrganizationWorkAreas     []OrganizationWorkAreaModel `gorm:"foreignKey:OrganizationWorkAreaProfessionUUID:references:UUID;COMMENT:所属工区专业;" json:"organization_work_area"`
}

// TableName 表名称
//  @receiver OrganizationWorkAreaProfessionModel
//  @return string
func (OrganizationWorkAreaProfessionModel) TableName() string {
	return "organization_work_area_professions"
}
