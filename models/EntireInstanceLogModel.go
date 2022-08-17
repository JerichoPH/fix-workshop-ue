package models

// EntireInstanceLogModel 器材日志模型
type EntireInstanceLogModel struct {
	BaseModel
	EntireInstanceLogTypeUniqueCode string                          `gorm:"type:VARCHAR(64);DEFAULT:;COMMENT:器材日志类型代码;" json:"entire_instance_type_unique_code"`
	EntireInstanceLogType           EntireInstanceLogTypeModel      `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceLogTypeUniqueCode;references:UniqueCode;COMMENT:所属器材日志类型;" json:"entire_instance_log_type"`
	Name                            string                          `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:器材日志名称;" json:"name"`
	Url                             string                          `gorm:"type:TEXT;COMMENT:器材日志相关跳转统一资源定位器;" json:"url"`
	OperatorUUID                    uint64                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:;COMMENT:操作人编号;" json:"operator_uuid"`
	Operator                        AccountModel                    `gorm:"foreignKey:OperatorUUID;references:ID;COMMENT:所属操作人;" json:"operator"`
	EntireInstanceIdentityCode      string                          `gorm:"type:VARCHAR(20);COMMENT:所属器材;" json:"entire_instance_identity_code"`
	EntireInstance                  EntireInstanceModel             `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:所属器材;" json:"entire_instance"`
	LocationLineUUID                string                          `gorm:"CHAR(36);NOT NULL;DEFAULT:;COMMENT:所属线别UUID;" json:""`
	LocationLine                    LocationLineModel               `gorm:"foreignKey:LocationLineUUID;references:UUID;COMMENT:所属线别;" json:""`
	LocationStationUUID             string                          `gorm:"CHAR(36);NOT NULL;DEFAULT:;COMMENT:所属站场UUID;" json:""`
	LocationStation                 LocationStationModel            `gorm:"foreignKey:LocationStationUUID;references:UUID;COMMENT:所属站场;" json:""`
	LocationSectionUUID             string                          `gorm:"CHAR(36);NOT NULL;DEFAULT:;COMMENT:所属区间UUID;" json:""`
	LocationSection                 LocationSectionModel            `gorm:"foreignKey:LocationSectionUUID;references:UUID;COMMENT:所属区间;" json:""`
	LocationCenterUUID              string                          `gorm:"CHAR(36);NOT NULL;DEFAULT:;COMMENT:所属中心UUID;" json:""`
	LocationCenter                  LocationCenterModel             `gorm:"foreignKey:LocationCenterUUID;references:UUID;COMMENT:所属中心;" json:""`
	LocationRailroadGradeCrossUUID  string                          `gorm:"CHAR(36);NOT NULL;DEFAULT:;COMMENT:所属道口UUID;" json:""`
	LocationRailroadGradeCross      LocationRailroadGradeCrossModel `gorm:"foreignKey:LocationRailroadGradeCrossUUID;references:UUID;COMMENT:所属道口;" json:""`
}

// TableName 表名称
func (EntireInstanceLogModel) TableName() string {
	return "entire_instance_logs"
}
