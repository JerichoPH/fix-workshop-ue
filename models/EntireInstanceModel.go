package models

import "time"

type EntireInstanceModel struct {
	BaseModel
	IdentityCode                   string                     `gorm:"type:VARCHAR(20);UNIQUE;NOT NULL;COMMENT:唯一编号;" json:"identity_code"`
	EntireInstanceStatusUniqueCode string                     `gorm:"type:VARCHAR(64);COMMENT:所属类型;" json:"entire_instance_status_unique_code"`
	EntireInstanceStatus           EntireInstanceStatusModel  `gorm:"foreignKey:EntireInstanceStatusUniqueCode;references:UniqueCode;COMMENT:所属状态;" json:"entire_instance_status"`
	KindSubTypeUUID                string                     `gorm:"type:CHAR(36);COMMENT:所属型号UUID;" json:"kind_sub_model_uuid"`
	KindSubModel                   KindSubTypeModel           `gorm:"foreignKey:KindSubTypeUUID;references:UUID;COMMENT:所属型号;" json:"kind_sub_model"`
	ParentIdentityCode             string                     `gorm:"type:VARCHAR(20);COMMENT:所属整机唯一编号;" json:"parent_identity_code"`
	Parent                         *EntireInstanceModel       `gorm:"foreignKey:ParentIdentityCode;references:IdentityCode;COMMENT:所属整机;" json:"parent"`
	Parts                          []*EntireInstanceModel     `gorm:"foreignKey:ParentIdentityCode;references:IdentityCode;COMMENT:相关部件;" json:"parts"`
	BePart                         bool                       `gorm:"type:BOOLEAN;DEFAULT:0;COMMENT:是否是部件;" json:"be_part"`
	HasExtrusionShroud             bool                       `gorm:"type:BOOLEAN;DEFAULT:0;COMMENT:是否具备防挤压防护装置;" json:"has_extrusion_shroud"`
	SaidRod                        string                     `gorm:"type:VARCHAR(64);COMMENT:表示杆特征;" json:"said_rod"`
	FixCycleYear                   int16                      `gorm:"type:TINYINT;COMMENT:周期修年;" json:"fix_cycle_year"`
	OrganizationRailwayUUID        string                     `gorm:"type:CHAR(36);COMMENT:所属路局;" json:"organization_railway_uuid"`
	OrganizationRailway            OrganizationRailwayModel   `gorm:"foreignKey:OrganizationRailwayUUID;references:UUID;COMMENT:所属路局;" json:"organization_railway"`
	OrganizationParagraphUUID      string                     `gorm:"type:CHAR(36);COMMENT:所属站段;" json:"organization_paragraph_unique_code"`
	OrganizationParagraph          OrganizationParagraphModel `gorm:"foreignKey:OrganizationParagraphUUID;references:UUID;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationWorkshopUUID       string                     `gorm:"type:CHAR(36);COMMENT:所属车间;" json:"organization_workshop_uuid"`
	OrganizationWorkshop           OrganizationWorkshopModel  `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属专业车间;" json:"organization_workshop"`
	OrganizationWorkAreaUUID       string                     `gorm:"type:CHAR(36);COMMENT:所属工区;" json:"organization_work_area_uuid"`
	OrganizationWorkArea           OrganizationWorkAreaModel  `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:所属专业工区;" json:"organization_work_area"`
	AssetCode                      string                     `gorm:"type:VARCHAR(128);COMMENT:物资编码;" json:"asset_code"`
	FixedAssetCode                 string                     `gorm:"type:VARCHAR(128);COMMENT:固资编码;" json:"fixed_asset_code"`
	SourceTypeUUID                 string                     `gorm:"type:CHAR(36);COMMENT:来源类型代码;" json:"source_type_UUID"`
	SourceType                     SourceTypeModel            `gorm:"foreignKey:SourceTypeUUID;references:UUID;COMMENT:所属来源类型;" json:"source_type"`
	SourceNameUUID                 string                     `gorm:"type:CHAR(36);COMMENT:来源名称代码;" json:"source_name_UUID"`
	SourceName                     SourceNameModel            `gorm:"foreignKey:SourceNameUUID;references:UUID;COMMENT:所属来源名称;" json:"source_name"`
	Note                           string                     `gorm:"type:LONGTEXT;COMMENT:备注;" json:"Note"`
	InWarehouseAt                  time.Time                  `gorm:"type:DATETIME;COMMENT:入库时间;'" json:"in_warehouse_at"`
	MadeAt                         time.Time                  `gorm:"type:DATETIME;COMMENT:出厂日期;" json:"made_at"`
	FactoryUUID                    string                     `gorm:"type:CHAR(36);COMMENT:所属供应商代码;" json:"factory_uuid"`
	Factory                        FactoryModel               `gorm:"foreignKey:FactoryUUID;references:UUID;COMMENT:所属供应商;" json:"factory"`
	LockName                       string                     `gorm:"type:VARCHAR(64);COMMENT:锁名称;" json:"lock_name"`
	LockDescription                string                     `gorm:"type:TEXT;COMMENT:锁说明;" json:"lock_description"`
	//EntireInstanceUses             []EntireInstanceUseModel   `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:使用信息;" json:"entire_instance_uses"`
	//EntireInstanceLogs []EntireInstanceLogModel `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:相关日志;" json:"entire_instance_logs"`
	//EntireInstanceRepairs               []EntireInstanceRepairModel    `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:相关检修记录;" json:"entire_instance_repairs"`
	DeleteOperatorUUID string       `gorm:"type:CHAR(36);COMMENT:删除操作人UUID;" json:"delete_operator_uuid"`
	DeleteOperator     AccountModel `gorm:"foreignKey:DeleteOperatorUUID;references:UUID;COMMENT:删除操作人;" json:"delete_operator"`
}

// TableName 表名称
func (EntireInstanceModel) TableName() string {
	return "entire_instances"
}
