package models

import "time"

type EntireInstanceModel struct {
	BaseModel
	// 基础信息
	IdentityCode                   string                    `gorm:"type:VARCHAR(20);UNIQUE;NOT NULL;COMMENT:唯一编号;" json:"identity_code"`
	SerialNumber                   string                    `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:'';COMMENT:所编号;" json:"serial_number"`
	EntireInstanceStatusUniqueCode string                    `gorm:"type:VARCHAR(64);COMMENT:所属类型;" json:"entire_instance_status_unique_code"`
	EntireInstanceStatus           EntireInstanceStatusModel `gorm:"foreignKey:EntireInstanceStatusUniqueCode;references:UniqueCode;COMMENT:所属状态;" json:"entire_instance_status"`
	KindCategoryUUID               string                    `gorm:"type:CHAR(36);COMMENT:所属种类UUID;" json:"kind_category_uuid"`
	KindCategory                   KindCategoryModel         `gorm:"foreignKey:KindCategoryUUID;references:UUID;COMMENT:所属种类;" json:"kind_category"`
	KindEntireTypeUUID             string                    `gorm:"type:CHAR(36);COMMENT:所属类型UUID;" json:"kind_entire_type_uuid"`
	KindEntireType                 KindEntireTypeModel       `gorm:"foreignKey:KindEntireTypeUUID;references:UUID;COMMENT:所属类型;" json:"kind_entire_model"`
	KindSubTypeUUID                string                    `gorm:"type:CHAR(36);COMMENT:所属型号UUID;" json:"kind_sub_model_uuid"`
	KindSubModel                   KindSubTypeModel          `gorm:"foreignKey:KindSubTypeUUID;references:UUID;COMMENT:所属型号;" json:"kind_sub_model"`
	FactoryUUID                    string                    `gorm:"type:CHAR(36);COMMENT:所属供应商代码;" json:"factory_uuid"`
	Factory                        FactoryModel              `gorm:"foreignKey:FactoryUUID;references:UUID;COMMENT:所属供应商;" json:"factory"`
	FactoryMadeSerialNumber        string                    `gorm:"type:VARCHAR(64);COMMENT:出厂编号;" json:"factory_made_serial_number"`
	FactoryMadeAt                  time.Time                 `gorm:"COMMENT:出厂日期;" json:"factory_made_at"`
	AssetCode                      string                    `gorm:"type:VARCHAR(128);COMMENT:物资编码;" json:"asset_code"`
	FixedAssetCode                 string                    `gorm:"type:VARCHAR(128);COMMENT:固资编码;" json:"fixed_asset_code"`
	ParentIdentityCode             string                    `gorm:"type:VARCHAR(20);COMMENT:所属整机唯一编号;" json:"parent_identity_code"`
	Parent                         *EntireInstanceModel      `gorm:"foreignKey:ParentIdentityCode;references:IdentityCode;COMMENT:所属整机;" json:"parent"`
	Parts                          []*EntireInstanceModel    `gorm:"foreignKey:ParentIdentityCode;references:IdentityCode;COMMENT:相关部件;" json:"parts"`
	BePart                         bool                      `gorm:"type:BOOLEAN;DEFAULT:0;COMMENT:是否是部件;" json:"be_part"`
	Note                           string                    `gorm:"type:TEXT;COMMENT:备注;" json:"note"`
	SourceNameUUID                 string                    `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:来源名称代码;" json:"source_name_UUID"`
	SourceName                     SourceNameModel           `gorm:"foreignKey:SourceNameUUID;references:UUID;COMMENT:所属来源名称;" json:"source_name"`
	DeleteOperatorUUID             string                    `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:删除操作人UUID;" json:"delete_operator_uuid"`
	DeleteOperator                 AccountModel              `gorm:"foreignKey:DeleteOperatorUUID;references:UUID;COMMENT:删除操作人;" json:"delete_operator"`
	WiringSystem                   string                    `gorm:"type:VARCHAR(64);NOT NULL;COMMENT:线制;" json:"wiring_system"`
	HasExtrusionShroud             bool                      `gorm:"type:BOOLEAN;NOT NULL;DEFAULT:0;COMMENT:是否具备防挤压防护装置;" json:"has_extrusion_shroud"`
	SaidRod                        string                    `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:'';COMMENT:表示杆特征;" json:"said_rod"`

	// 使用信息
	UseExpirationAt                   time.Time              `gorm:"COMMENT:到期日期;" json:"use_expiration_at"`
	UseDestroyAt                      time.Time              `gorm:"COMMENT:报废日期；" json:"use_destroy_at"`
	UseNextCycleRepairAt              time.Time              `gorm:"COMMENT:下次周期修日期;" json:"use_next_cycle_repair_at"`
	UseWarehouseInAt                  time.Time              `gorm:"COMMENT:入库时间;'" json:"use_warehouse_in_at"`
	UseWarehousePositionDepotCellUUID string                 `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:'';COMMENT:存放位置-仓储位置;" json:"use_warehouse_position_depot_cell_uuid"`
	UseWarehousePositionDepotCell     PositionDepotCellModel `gorm:"foreignKey:UseWarehousePositionDepotCellUUID;references:UUID;COMMENT:存放位置-仓储位置;" json:"use_warehouse_position_depot_cell"`
	UseRepairCurrentFixedAt           time.Time              `gorm:"COMMENT:当前检修时间;" json:"use_repair_current_fixed_at"`
	UseRepairCurrentFixerName         string                 `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:'';COMMENT:当前检修人;" json:"use_repair_current_fixer_name"`
	UseRepairCurrentCheckedAt         time.Time              `gorm:"COMMENT:当前验收时间;" json:"use_repair_current_checked_at"`
	UseRepairCurrentCheckerName       string                 `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:'';COMMENT:当前验收人;" json:"use_repair_current_checker_name"`
	UseRepairCurrentSpotCheckedAt     time.Time              `gorm:"COMMENT:当前抽验时间" json:"use_repair_current_spot_checked_at"`
	UseRepairCurrentSpotCheckerName   string                 `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:'';COMMENT:当前抽验人;" json:"use_repair_current_spot_checker_name"`
	UseRepairLastFixedAt              time.Time              `gorm:"COMMENT:上次检修时间;" json:"use_repair_last_fixed_at"`
	UseRepairLastFixerName            string                 `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:'';COMMENT:上次检修人;" json:"use_repair_last_fixer_name"`
	UseRepairLastCheckedAt            time.Time              `gorm:"COMMENT:上次验收时间;" json:"use_repair_last_checked_at"`
	UseRepairLastCheckerName          string                 `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:'';COMMENT:上次验收人;" json:"use_repair_last_checker_name"`
	UseRepairLastSpotCheckedAt        time.Time              `gorm:"COMMENT:上次抽验时间" json:"use_repair_last_spot_checked_at"`
	UseRepairLastSpotCheckerName      string                 `gorm:"type:VARCHAR(64);NOT NULL;DEFAULT:'';COMMENT:上次抽验人;" json:"use_repair_last_spot_checker_name"`

	// 资产归属
	BelongToOrganizationRailwayUUID   string                     `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:资产归属-所属路局;" json:"belong_to_organization_railway_uuid"`
	BelongToOrganizationRailway       OrganizationRailwayModel   `gorm:"foreignKey:BelongToOrganizationRailwayUUID;references:UUID;COMMENT:资产归属-所属路局;" json:"belong_to_organization_railway"`
	BelongToOrganizationParagraphUUID string                     `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:资产归属-所属站段;" json:"belong_to_organization_paragraph_uuid"`
	BelongToOrganizationParagraph     OrganizationParagraphModel `gorm:"foreignKey:BelongToOrganizationParagraphUUID;references:UUID;COMMENT:资产归属-所属站段;" json:"belong_to_organization_paragraph"`
	BelongToOrganizationWorkshopUUID  string                     `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:资产归属-所属车间;" json:"belong_to_organization_workshop_uuid"`
	BelongToOrganizationWorkshop      OrganizationWorkshopModel  `gorm:"foreignKey:BelongToOrganizationWorkshopUUID;references:UUID;COMMENT:资产归属-所属专业车间;" json:"belong_to_organization_workshop"`
	BelongToOrganizationWorkAreaUUID  string                     `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:资产归属-所属工区;" json:"belong_to_organization_work_area_uuid"`
	BelongToOrganizationWorkArea      OrganizationWorkAreaModel  `gorm:"foreignKey:BelongToOrganizationWorkAreaUUID;references:UUID;COMMENT:资产归属-所属专业工区;" json:"belong_to_organization_work_area"`

	// 使用处所
	UsePlaceCurrentOrganizationWorkshopUUID       string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:当前使用处所-车间UUID;" json:"use_place_current_organization_workshop_uuid"`
	UsePlaceCurrentOrganizationWorkshop           OrganizationWorkshopModel       `gorm:"foreignKey:UsePlaceCurrentOrganizationWorkshopUUID;references:UUID;COMMENT:当前使用处所-车间;" json:"use_place_current_organization_workshop"`
	UsePlaceCurrentOrganizationWorkAreaUUID       string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:当前使用处所-工区UUID;" json:"use_place_current_organization_work_area_uuid"`
	UsePlaceCurrentOrganizationWorkArea           OrganizationWorkAreaModel       `gorm:"foreignKey:UsePlaceCurrentOrganizationWorkAreaUUID;references:UUID;COMMENT:当前使用处所-工区;" json:"use_place_current_organization_work_area"`
	UsePlaceCurrentLocationLineUUID               string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:当前使用处所-线别UUID;" json:"use_place_current_location_line_uuid"`
	UsePlaceCurrentLocationLine                   LocationLineModel               `gorm:"foreignKey:UsePlaceCurrentLocationLineUUID;references:UUID;COMMENT:当前使用处所-线别" json:"use_place_current_location_line"`
	UsePlaceCurrentLocationStationUUID            string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:当前使用处所-站场UUID;" json:"use_place_current_location_station_uuid"`
	UsePlaceCurrentLocationStation                LocationStationModel            `gorm:"foreignKey:UsePlaceCurrentLocationStationUUID;references:UUID;COMMENT:当前使用处所-站场" json:"use_place_current_location_station"`
	UsePlaceCurrentLocationSectionUUID            string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:当前使用处所-区间UUID;" json:"use_place_current_location_section_uuid"`
	UsePlaceCurrentLocationSection                LocationSectionModel            `gorm:"foreignKey:UsePlaceCurrentLocationSectionUUID;references:UUID;COMMENT:当前使用处所-区间" json:"use_place_current_location_section"`
	UsePlaceCurrentLocationCenterUUID             string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:当前使用处所-中心UUID;" json:"use_place_current_location_center_uuid"`
	UsePlaceCurrentLocationCenter                 LocationCenterModel             `gorm:"foreignKey:UsePlaceCurrentLocationCenterUUID;references:UUID;COMMENT:当前使用处所-中心" json:"use_place_current_location_center"`
	UsePlaceCurrentLocationRailroadGradeCrossUUID string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:当前使用处所-道口UUID;" json:"use_place_current_location_railroad_grade_cross_uuid"`
	UsePlaceCurrentLocationRailroadGradeCross     LocationRailroadGradeCrossModel `gorm:"foreignKey:UsePlaceCurrentLocationRailroadGradeCrossUUID;references:UUID;COMMENT:当前使用处所-道口" json:"use_place_current_location_railroad_grade_cross"`
	UsePlaceLastOrganizationWorkshopUUID          string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:上次使用处所-车间UUID;" json:"use_place_last_organization_workshop_uuid"`
	UsePlaceLastOrganizationWorkshop              OrganizationWorkshopModel       `gorm:"foreignKey:UsePlaceLastOrganizationWorkshopUUID;references:UUID;COMMENT:上次使用处所-车间;" json:"use_place_last_organization_workshop"`
	UsePlaceLastOrganizationWorkAreaUUID          string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:上次使用处所-工区UUID;" json:"use_place_last_organization_work_area_uuid"`
	UsePlaceLastOrganizationWorkArea              OrganizationWorkAreaModel       `gorm:"foreignKey:UsePlaceLastOrganizationWorkAreaUUID;references:UUID;COMMENT:上次使用处所-工区;" json:"use_place_last_organization_work_area"`
	UsePlaceLastLocationLineUUID                  string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:上次使用处所-线别UUID;" json:"use_place_last_location_line_uuid"`
	UsePlaceLastLocationLine                      LocationLineModel               `gorm:"foreignKey:UsePlaceLastLocationLineUUID;references:UUID;COMMENT:上次使用处所-线别" json:"use_place_last_location_line"`
	UsePlaceLastLocationStationUUID               string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:上次使用处所-站场UUID;" json:"use_place_last_location_station_uuid"`
	UsePlaceLastLocationStation                   LocationStationModel            `gorm:"foreignKey:UsePlaceLastLocationStationUUID;references:UUID;COMMENT:上次使用处所-站场" json:"use_place_last_location_station"`
	UsePlaceLastLocationSectionUUID               string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:上次使用处所-区间UUID;" json:"use_place_last_location_section_uuid"`
	UsePlaceLastLocationSection                   LocationSectionModel            `gorm:"foreignKey:UsePlaceLastLocationSectionUUID;references:UUID;COMMENT:上次使用处所-区间" json:"use_place_last_location_section"`
	UsePlaceLastLocationCenterUUID                string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:上次使用处所-中心UUID;" json:"use_place_last_location_center_uuid"`
	UsePlaceLastLocationCenter                    LocationCenterModel             `gorm:"foreignKey:UsePlaceLastLocationCenterUUID;references:UUID;COMMENT:上次使用处所-中心" json:"use_place_last_location_center"`
	UsePlaceLastLocationRailroadGradeCrossUUID    string                          `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:上次使用处所-道口UUID;" json:"use_place_last_location_railroad_grade_cross_uuid"`
	UsePlaceLastLocationRailroadGradeCross        LocationRailroadGradeCrossModel `gorm:"foreignKey:UsePlaceLastLocationRailroadGradeCrossUUID;references:UUID;COMMENT:上次使用处所-道口" json:"use_place_last_location_railroad_grade_cross"`

	// 使用位置
	UsePlaceCurrentPositionIndoorCellUUID string                  `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:当前使用位置-室内上道位置UUID;" json:"use_place_current_position_indoor_cell_uuid"`
	UsePlaceCurrentPositionIndoorCell     PositionIndoorCellModel `gorm:"foreignKey:UsePlaceCurrentPositionIndoorCellUUID;references:UUID;COMMENT:当前使用位置-室内上道位置" json:"use_place_current_position_indoor_cell"`
	UsePlaceLastPositionIndoorCellUUID    string                  `gorm:"type:CHAR(36);NOT NULL;DEFAULT:'';COMMENT:上次使用位置-室内上道位置UUID;" json:"use_place_last_position_indoor_cell_uuid"`
	UsePlaceLastPositionIndoorCell        PositionIndoorCellModel `gorm:"foreignKey:UsePlaceLastPositionIndoorCellUUID;references:UUID;COMMENT:上次使用位置-室内上道位置" json:"use_place_last_position_indoor_cell"`

	// 其他
	ExCycleRepairYear int16 `gorm:"type:INT2;NOT NULL;DEFAULT:0;COMMENT:周期修年;" json:"ex_cycle_repair_year"`
	ExLifeYear        int16 `gorm:"type:INT2;NOT NULL;DEFAULT:15;COMMENT:;" json:"ex_life_year"`

	//LockName        string `gorm:"type:VARCHAR(64);COMMENT:锁名称;" json:"lock_name"`
	//LockDescription string `gorm:"type:TEXT;COMMENT:锁说明;" json:"lock_description"`
	//EntireInstanceUses             []EntireInstanceUseModel   `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:使用信息;" json:"entire_instance_uses"`
	//EntireInstanceLogs []EntireInstanceLogModel `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:相关日志;" json:"entire_instance_logs"`
	//EntireInstanceRepairs               []EntireInstanceRepairModel    `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:相关检修记录;" json:"entire_instance_repairs"`
}

// TableName 表名称
func (EntireInstanceModel) TableName() string {
	return "entire_instances"
}
