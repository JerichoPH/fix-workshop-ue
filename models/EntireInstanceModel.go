package models

import (
	"time"
)

type EntireInstanceModel struct {
	BaseModel
	IdentityCode                    string                     `gorm:"type:VARCHAR(20);UNIQUE;NOT NULL;COMMENT:唯一编号;" json:"identity_code"`
	EntireInstanceStatusUniqueCode  string                     `gorm:"type:VARCHAR(64);COMMENT:所属类型;" json:"entire_instance_status_unique_code"`
	EntireInstanceStatus            EntireInstanceStatusModel  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceStatusUniqueCode;references:UniqueCode;COMMENT:所属状态;" json:"entire_instance_status"`
	KindCategoryUniqueCode          string                     `gorm:"type:CHAR(3);COMMENT:所属类型;" json:"kind_category_unique_code"`
	KindCategory                    KindCategoryModel          `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindCategoryUniqueCode;references:UniqueCode;COMMENT:所属种类;" json:"kind_category"`
	KindEntireTypeUniqueCode        string                     `gorm:"type:CHAR(5);COMMENT:所属类型;" json:"kind_entire_model_unique_code"`
	KindEntireType                  KindEntireTypeModel        `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindEntireTypeUniqueCode;references:UniqueCode;COMMENT:所属类型;" json:"kind_entire_model"`
	KindSubTypeUniqueCode           string                     `gorm:"type:CHAR(7);COMMENT:所属型号;" json:"kind_sub_model_unique_code"`
	KindSubModel                    KindSubTypeModel           `gorm:"constraint:OnUpdate:CASCADE;foreignKey:KindSubTypeUniqueCode;references:UniqueCode;COMMENT:所属型号;" json:"kind_sub_model"`
	ParentIdentityCode              string                     `gorm:"type:VARCHAR(20);COMMENT:所属整机唯一编号;" json:"parent_identity_code"`
	Parent                          *EntireInstanceModel       `gorm:"constraint:OnUpdate:CASCADE;foreignKey:ParentIdentityCode;references:IdentityCode;COMMENT:所属整机;" json:"parent"`
	Parts                           []*EntireInstanceModel     `gorm:"constraint:OnUpdate:CASCADE;foreignKey:ParentIdentityCode;references:IdentityCode;COMMENT:相关部件;" json:"parts"`
	BePart                          bool                       `gorm:"type:BOOLEAN;DEFAULT:0;COMMENT:是否是部件;" json:"be_part"`
	HasExtrusionShroud              bool                       `gorm:"type:BOOLEAN;DEFAULT:0;COMMENT:是否具备防挤压防护装置;" json:"has_extrusion_shroud"`
	SaidRod                         string                     `gorm:"type:VARCHAR(64);COMMENT:表示杆特征;" json:"said_rod"`
	FixCycleYear                    int16                      `gorm:"type:TINYINT;COMMENT:周期修年;" json:"fix_cycle_year"`
	OrganizationRailwayUniqueCode   string                     `gorm:"type:CHAR(3);COMMENT:所属路局;" json:"organization_railway_unique_code"`
	OrganizationRailway             OrganizationRailwayModel   `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationRailwayUniqueCode;references:UniqueCode;COMMENT:所属路局;" json:"organization_railway"`
	OrganizationParagraphUniqueCode string                     `gorm:"type:CHAR(4);COMMENT:所属站段;" json:"organization_paragraph_unique_code"`
	OrganizationParagraph           OrganizationParagraphModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationParagraphUniqueCode;references:UniqueCode;COMMENT:所属站段;" json:"organization_paragraph"`
	OrganizationWorkshopUniqueCode  string                     `gorm:"type:CHAR(7);COMMENT:所属车间;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop            OrganizationWorkshopModel  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUUID;references:UniqueCode;COMMENT:所属专业车间;" json:"organization_workshop"`
	OrganizationWorkAreaUniqueCode  string                     `gorm:"type:CHAR(8);COMMENT:所属工区;" json:"organization_work_area_unique_code"`
	OrganizationWorkArea            OrganizationWorkAreaModel  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkAreaUniqueCode;references:UniqueCode;COMMENT:所属专业工区;" json:"organization_work_area"`
	DeleteProcessorId               string                     `gorm:"type:INT;COMMENT:删除器材操作人;" json:"delete_processor_id"`
	AssetCode                       string                     `gorm:"type:VARCHAR(128);COMMENT:物资编码;" json:"asset_code"`
	FixedAssetCode                  string                     `gorm:"type:VARCHAR(128);COMMENT:固资编码;" json:"fixed_asset_code"`
	SourceTypeUniqueCode            string                     `gorm:"type:CHAR(2);COMMENT:来源类型代码;" json:"source_type_unique_code"`
	SourceType                      SourceTypeModel            `gorm:"constraint:OnUpdate:CASCADE;foreignKey:SourceTypeUniqueCode;references:UniqueCode;COMMENT:所属来源类型;" json:"source_type"`
	SourceNameUniqueCode            string                     `gorm:"type:VARCHAR(64);COMMENT:来源名称代码;" json:"source_name_unique_code"`
	SourceName                      SourceNameModel            `gorm:"constraint:OnUpdate:CASCADE;foreignKey:SourceNameUniqueCode;references:UniqueCode;COMMENT:所属来源名称;" json:"source_name"`
	Note                            string                     `gorm:"type:LONGTEXT;COMMENT:备注;" json:"Note"`
	InWarehouseAt                   time.Time                  `gorm:"type:DATETIME;COMMENT:入库时间;'" json:"in_warehouse_at"`
	MadeAt                          time.Time                  `gorm:"type:DATETIME;COMMENT:出厂日期;" json:"made_at"`
	FactoryUniqueCode               string                     `gorm:"type:CHAR(5);COMMENT:所属供应商代码;" json:"factory_unique_code"`
	Factory                         FactoryModel               `gorm:"constraint:OnUpdate:CASCADE;foreignKey:FactoryUniqueCode;references:UniqueCode;COMMENT:所属供应商;" json:"factory"`
	LockName                        string                     `gorm:"type:VARCHAR(64);COMMENT:锁名称;" json:"lock_name"`
	LockDescription                 string                     `gorm:"type:TEXT;COMMENT:锁说明;" json:"lock_description"`
	EntireInstanceUses              []EntireInstanceUseModel   `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:使用信息;" json:"entire_instance_uses"`
	EntireInstanceLogs              []EntireInstanceLogModel   `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:相关日志;" json:"entire_instance_logs"`
	//EntireInstanceRepairs               []EntireInstanceRepairModel    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:相关检修记录;" json:"entire_instance_repairs"`
	LocationWarehousePositionUniqueCode string                         `gorm:"type:CHAR(18);COMMENT:所属仓库位置代码;" json:"location_warehouse_position_unique_code"`
	LocationWarehousePosition           LocationWarehousePositionModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationWarehousePositionUniqueCode;references:UniqueCode;COMMENT:所属仓库位置;" json:"location_warehouse_position"`
	DeleteOperatorUUID                  string                         `gorm:"type:CHAR(36);COMMENT:删除操作人UUID;" json:"delete_operator_uuid"`
	DeleteOperator                      AccountModel                   `gorm:"constraint:OnUpdate:CASCADE;foreignKey:DeleteOperatorUUID;references:UUID;COMMENT:删除操作人;" json:"delete_operator"`
}

// TableName 表名称
func (cls *EntireInstanceModel) TableName() string {
	return "entire_instances"
}