package models

// OrganizationWorkshopType 车间类型
type OrganizationWorkshopType struct {
	BaseModel
	UniqueCode            string                 `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:车间类型代码;" json:"unique_code"`
	Name                  string                 `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:车间类型名称;" json:"name"`
	Number                string                 `gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:车间类型数字代码;" json:"number"`
	OrganizationWorkshops []OrganizationWorkshop `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopTypeUniqueCode;references:UniqueCode;" json:"organization_workshops"`
}

// FindOneByUniqueCode 根据unique_code获取单条数据
func (cls *OrganizationWorkshopType) FindOneByUniqueCode(uniqueCode string) (organizationWorkshopType OrganizationWorkshopType) {
	cls.Boot().Where(map[string]interface{}{"unique_code": uniqueCode}).First(&organizationWorkshopType)

	return
}
