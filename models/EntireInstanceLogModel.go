package models

// EntireInstanceLogModel 器材日志模型
type EntireInstanceLogModel struct {
	BaseModel
	EntireInstanceLogTypeUniqueCode string                          `gorm:"type:VARCHAR(64);DEFAULT:;COMMENT:器材日志类型代码;" json:"entire_instance_type_unique_code"`
	EntireInstanceLogType           EntireInstanceLogTypeModel      `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceLogTypeUniqueCode;references:UniqueCode;COMMENT:所属器材日志类型;" json:"entire_instance_log_type"`
	Name                            string                          `gorm:"type:VARCHAR(64);COMMENT:器材日志名称;" json:"name"`
	Url                             string                          `gorm:"type:TEXT;COMMENT:器材日志相关跳转统一资源定位器;" json:"url"`
	OperatorUuid                    uint64                          `gorm:"type:VARCHAR(36);COMMENT:操作人编号;" json:"operator_uuid"`
	Operator                        AccountModel                    `gorm:"foreignKey:OperatorUuid;references:Id;COMMENT:所属操作人;" json:"operator"`
	EntireInstanceIdentityCode      string                          `gorm:"type:VARCHAR(20);COMMENT:所属器材UUID;" json:"entire_instance_identity_code"`
	EntireInstance                  EntireInstanceModel             `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:所属器材;" json:"entire_instance"`
	OrganizationRailwayUuid         string                          `gorm:"type:VARCHAR(36);COMMENT:所属路局UUID;organization_railway_uuid"`
	OrganizationRailway             OrganizationRailwayModel        `gorm:"foreignKey:OrganizationRailwayUuid;references:Uuid;COMMENT:所属路局;" json:"organization_railway"`
	OrganizationParagraphUuid       string                          `gorm:"type:VARCHAR(36);COMMENT:所属站段UUID;organization_paragraph_uuid"`
	OrganizationParagraph           OrganizationParagraphModel      `gorm:"foreignKey:OrganizationParagraphUuid;references:Uuid;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationWorkshopUuid        string                          `gorm:"type:VARCHAR(36);COMMENT:所属车间UUID;organization_workshop_uuid"`
	OrganizationWorkshop            OrganizationWorkshopModel       `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:所属车间;" json:"organization_workshop"`
	OrganizationWorkAreaUuid        string                          `gorm:"type:VARCHAR(36);COMMENT:所属工区;organization_work_area_uuid"`
	OrganizationWorkArea            OrganizationWorkAreaModel       `gorm:"foreignKey:OrganizationWorkAreaUuid;references:Uuid;COMMENT:所属工区;" json:"organization_work_area"`
	LocationLineUuid                string                          `gorm:"VARCHAR(36);COMMENT:所属线别UUID;" json:"location_line_uuid"`
	LocationLine                    LocationLineModel               `gorm:"foreignKey:LocationLineUuid;references:Uuid;COMMENT:所属线别;" json:"location_line"`
	LocationStationUuid             string                          `gorm:"VARCHAR(36);COMMENT:所属站场UUID;" json:"location_station_uuid"`
	LocationStation                 LocationStationModel            `gorm:"foreignKey:LocationStationUuid;references:Uuid;COMMENT:所属站场;" json:"location_station"`
	LocationSectionUuid             string                          `gorm:"VARCHAR(36);COMMENT:所属区间UUID;" json:"location_section_uuid"`
	LocationSection                 LocationSectionModel            `gorm:"foreignKey:LocationSectionUuid;references:Uuid;COMMENT:所属区间;" json:"location_section"`
	LocationCenterUuid              string                          `gorm:"VARCHAR(36);COMMENT:所属中心UUID;" json:"location_center_uuid"`
	LocationCenter                  LocationCenterModel             `gorm:"foreignKey:LocationCenterUuid;references:Uuid;COMMENT:所属中心;" json:"location_center"`
	LocationRailroadGradeCrossUuid  string                          `gorm:"VARCHAR(36);COMMENT:所属道口UUID;" json:"location_railroad_grade_cross_uuid"`
	LocationRailroadGradeCross      LocationRailroadGradeCrossModel `gorm:"foreignKey:LocationRailroadGradeCrossUuid;references:Uuid;COMMENT:所属道口;" json:"location_railroad_grade_cross"`
}

// TableName 表名称
func (EntireInstanceLogModel) TableName() string {
	return "entire_instance_logs"
}
