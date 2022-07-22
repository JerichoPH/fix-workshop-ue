package models

// EntireInstanceLogModel 器材日志模型
type EntireInstanceLogModel struct {
	BaseModel
	EntireInstanceLogTypeUniqueCode          string                                                 `gorm:"type:VARCHAR(64);DEFAULT:'fa-envelope-o';COMMENT:器材日志类型代码;" json:"entire_instance_type_unique_code"`
	EntireInstanceLogType                    EntireInstanceLogTypeModel                             `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceLogTypeUniqueCode;references:UniqueCode;COMMENT:所属器材日志类型;" json:"entire_instance_log_type"`
	Name                                     string                                                 `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:器材日志名称;" json:"name"`
	Url                                      string                                                 `gorm:"type:LONGTEXT;COMMENT:器材日志相关跳转统一资源定位器;" json:"url"`
	OperatorId                               uint64                                                 `gorm:"type:BIGINT(20) UNSIGNED;COMMENT:操作人编号;" json:"operator_id"`
	Operator                                 AccountModel                                    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OperatorId;references:ID;COMMENT:所属操作人;" json:"operator"`
	OrganizationParagraphUniqueCode          string                                                 `gorm:"type:CHAR(4);NOT NULL;COMMENT:所属站段代码;" json:"organization_paragraph_unique_code"`
	OrganizationParagraph                    OrganizationParagraphModel          `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationParagraphUUID;references:UniqueCode;NOT NULL;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationWorkshopUniqueCode           string                                                 `gorm:"type:CHAR(7);COMMENT:所属车间代码;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop                     OrganizationWorkshopModel           `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUUID;references:UniqueCode;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUniqueCode           string                                                 `gorm:"type:CHAR(8);COMMENT:所属工区代码;" json:"organization_work_area_unique_code"`
	OrganizationWorkArea                     OrganizationWorkAreaModel           `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkAreaUUID;references:UniqueCode;COMMENT:所属工区;" json:"organization_work_area"`
	OrganizationRailroadGradeCrossUniqueCode string                                                 `gorm:"type:CHAR(5);COMMENT:所属道口代码;" json:"organization_railroad_grade_cross_unique_code"`
	OrganizationRailroadGradeCross           OrganizationRailroadGradeCrossModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationRailroadGradeCrossUUID;references:UniqueCode;COMMENT:所属道口;" json:"organization_railroad_grade_cross"`
	OrganizationCenterUniqueCode             string                                                 `gorm:"type:CHAR(6);COMMENT:所属中心代码;" json:"organization_center_unique_code"`
	OrganizationCenter                       OrganizationCenterModel             `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationCenterUUID;references:UniqueCode;COMMENT:所属中心;" json:"organization_center"`
	OrganizationStationUniqueCode            string                                                 `gorm:"type:CHAR(6);COMMENT:所属站场代码;" json:"organization_station_unique_code"`
	OrganizationStation                      OrganizationStationModel            `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationStationUUID;references:UniqueCode;COMMENT:所属站场;" json:"organization_station"`
	EntireInstanceIdentityCode               string                                                 `gorm:"type:VARCHAR(20);COMMENT:所属器材;" json:"entire_instance_identity_code"`
	EntireInstance                           EntireInstanceModel                                    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:所属器材;" json:"entire_instance"`
}

// TableName 表名称
func (cls *EntireInstanceLogModel) TableName() string {
	return "entire_instance_logs"
}
