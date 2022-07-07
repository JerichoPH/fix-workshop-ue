package models

// OrganizationWorkshopTypeModel 车间类型
type OrganizationWorkshopTypeModel struct {
	BaseModel
	UniqueCode            string                      `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:车间类型代码;" json:"unique_code"`
	Name                  string                      `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:车间类型名称;" json:"name"`
	Number                string                      `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:车间类型数字代码;" json:"number"`
	OrganizationWorkshops []OrganizationWorkshopModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopTypeUniqueCode;references:UniqueCode;" json:"organization_workshops"`
}

// TableName 表名称
func (cls *OrganizationWorkshopTypeModel) TableName() string {
	return "organization_workshop_types"
}
