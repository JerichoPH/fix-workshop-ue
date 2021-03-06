package models

// EntireInstanceLog 器材日志模型
type EntireInstanceLog struct {
	BaseModel
	Preloads                                 []string
	Selects                                  []string
	Omits                                    []string
	EntireInstanceLogTypeUniqueCode          string                         `gorm:"type:VARCHAR(64);DEFAULT:'fa-envelope-o';COMMENT:器材日志类型代码;" json:"entire_instance_type_unique_code"`
	EntireInstanceLogType                    EntireInstanceLogType          `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceLogTypeUniqueCode;references:UniqueCode;COMMENT:所属器材日志类型;" json:"entire_instance_log_type"`
	Name                                     string                         `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:器材日志名称;" json:"name"`
	Url                                      string                         `gorm:"type:LONGTEXT;COMMENT:器材日志相关跳转统一资源定位器;" json:"url"`
	OperatorId                               uint64                         `gorm:"type:BIGINT(20) UNSIGNED;COMMENT:操作人编号;" json:"operator_id"`
	Operator                                 Account                        `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OperatorId;references:ID;COMMENT:所属操作人;" json:"operator"`
	OrganizationParagraphUniqueCode          string                         `gorm:"type:CHAR(4);NOT NULL;COMMENT:所属站段代码;" json:"organization_paragraph_unique_code"`
	OrganizationParagraph                    OrganizationParagraph          `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationParagraphUniqueCode;references:UniqueCode;NOT NULL;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationWorkshopUniqueCode           string                         `gorm:"type:CHAR(7);COMMENT:所属车间代码;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop                     OrganizationWorkshop           `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUniqueCode           string                         `gorm:"type:CHAR(8);COMMENT:所属工区代码;" json:"organization_work_area_unique_code"`
	OrganizationWorkArea                     OrganizationWorkArea           `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkAreaUniqueCode;references:UniqueCode;COMMENT:所属工区;" json:"organization_work_area"`
	OrganizationRailroadGradeCrossUniqueCode string                         `gorm:"type:CHAR(5);COMMENT:所属道口代码;" json:"organization_railroad_grade_cross_unique_code"`
	OrganizationRailroadGradeCross           OrganizationRailroadGradeCross `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationRailroadGradeCrossUniqueCode;references:UniqueCode;COMMENT:所属道口;" json:"organization_railroad_grade_cross"`
	OrganizationCenterUniqueCode             string                         `gorm:"type:CHAR(6);COMMENT:所属中心代码;" json:"organization_center_unique_code"`
	OrganizationCenter                       OrganizationCenter             `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationCenterUniqueCode;references:UniqueCode;COMMENT:所属中心;" json:"organization_center"`
	OrganizationStationUniqueCode            string                         `gorm:"type:CHAR(6);COMMENT:所属站场代码;" json:"organization_station_unique_code"`
	OrganizationStation                      OrganizationStation            `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationStationUniqueCode;references:UniqueCode;COMMENT:所属站场;" json:"organization_station"`
	EntireInstanceIdentityCode               string                         `gorm:"type:VARCHAR(20);COMMENT:所属器材;" json:"entire_instance_identity_code"`
	EntireInstance                           EntireInstance                 `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:所属器材;" json:"entire_instance"`
}
