package models

import "time"

type EntireInstanceModel struct {
	BaseModel
	// 基础信息
	IdentityCode                   string                    `gorm:"type:VARCHAR(20);UNIQUE;COMMENT:唯一编号;" json:"identity_code"`
	SerialNumber                   string                    `gorm:"type:VARCHAR(64);COMMENT:所编号;" json:"serial_number"`
	EntireInstanceStatusUniqueCode string                    `gorm:"type:VARCHAR(64);COMMENT:所属类型;" json:"entire_instance_status_unique_code"`
	EntireInstanceStatus           EntireInstanceStatusModel `gorm:"foreignKey:EntireInstanceStatusUniqueCode;references:UniqueCode;COMMENT:所属状态;" json:"entire_instance_status"`
	KindCategoryUuid               string                    `gorm:"type:VARCHAR(36);COMMENT:所属种类UUID;" json:"kind_category_uuid"`
	KindCategory                   KindCategoryModel         `gorm:"foreignKey:KindCategoryUuid;references:Uuid;COMMENT:所属种类;" json:"kind_category"`
	KindEntireTypeUuid             string                    `gorm:"type:VARCHAR(36);COMMENT:所属类型UUID;" json:"kind_entire_type_uuid"`
	KindEntireType                 KindEntireTypeModel       `gorm:"foreignKey:KindEntireTypeUuid;references:Uuid;COMMENT:所属类型;" json:"kind_entire_model"`
	KindSubTypeUuid                string                    `gorm:"type:VARCHAR(36);COMMENT:所属型号UUID;" json:"kind_sub_model_uuid"`
	KindSubModel                   KindSubTypeModel          `gorm:"foreignKey:KindSubTypeUuid;references:Uuid;COMMENT:所属型号;" json:"kind_sub_model"`
	FactoryUuid                    string                    `gorm:"type:VARCHAR(36);COMMENT:所属供应商代码;" json:"factory_uuid"`
	Factory                        FactoryModel              `gorm:"foreignKey:FactoryUuid;references:Uuid;COMMENT:所属供应商;" json:"factory"`
	FactoryMadeSerialNumber        string                    `gorm:"type:VARCHAR(64);COMMENT:出厂编号;" json:"factory_made_serial_number"`
	FactoryMadeAt                  time.Time                 `gorm:"COMMENT:出厂日期;" json:"factory_made_at"`
	AssetCode                      string                    `gorm:"type:VARCHAR(128);COMMENT:物资编码;" json:"asset_code"`
	FixedAssetCode                 string                    `gorm:"type:VARCHAR(128);COMMENT:固资编码;" json:"fixed_asset_code"`
	ParentIdentityCode             string                    `gorm:"type:VARCHAR(20);COMMENT:所属整机唯一编号;" json:"parent_identity_code"`
	Parent                         *EntireInstanceModel      `gorm:"foreignKey:ParentIdentityCode;references:IdentityCode;COMMENT:所属整机;" json:"parent"`
	Parts                          []*EntireInstanceModel    `gorm:"foreignKey:ParentIdentityCode;references:IdentityCode;COMMENT:相关部件;" json:"parts"`
	BePart                         bool                      `gorm:"type:BOOLEAN;DEFAULT:0;COMMENT:是否是部件;" json:"be_part"`
	Note                           string                    `gorm:"type:TEXT;COMMENT:备注;" json:"note"`
	SourceNameUuid                 string                    `gorm:"type:VARCHAR(36);COMMENT:来源名称代码;" json:"source_name_UUID"`
	SourceName                     SourceNameModel           `gorm:"foreignKey:SourceNameUuid;references:Uuid;COMMENT:所属来源名称;" json:"source_name"`
	DeleteOperatorUuid             string                    `gorm:"type:VARCHAR(36);COMMENT:删除操作人UUID;" json:"delete_operator_uuid"`
	DeleteOperator                 AccountModel              `gorm:"foreignKey:DeleteOperatorUuid;references:Uuid;COMMENT:删除操作人;" json:"delete_operator"`
	WiringSystem                   string                    `gorm:"type:VARCHAR(64);COMMENT:线制;" json:"wiring_system"`
	HasExtrusionShroud             bool                      `gorm:"type:BOOLEAN;DEFAULT:0;COMMENT:是否具备防挤压防护装置;" json:"has_extrusion_shroud"`
	SaidRod                        string                    `gorm:"type:VARCHAR(64);COMMENT:表示杆特征;" json:"said_rod"`

	// 使用信息
	UseExpireAt                       time.Time              `gorm:"COMMENT:到期日期;" json:"use_expire_at"`
	UseDestroyAt                      time.Time              `gorm:"COMMENT:报废日期；" json:"use_destroy_at"`
	UseNextCycleRepairAt              time.Time              `gorm:"COMMENT:下次周期修日期;" json:"use_next_cycle_repair_at"`
	UseWarehouseInAt                  time.Time              `gorm:"COMMENT:入库时间;'" json:"use_warehouse_in_at"`
	UseWarehousePositionDepotCellUuid string                 `gorm:"type:VARCHAR(64);COMMENT:存放位置-仓储位置;" json:"use_warehouse_position_depot_cell_uuid"`
	UseWarehousePositionDepotCell     PositionDepotCellModel `gorm:"foreignKey:UseWarehousePositionDepotCellUuid;references:Uuid;COMMENT:存放位置-仓储位置;" json:"use_warehouse_position_depot_cell"`
	UseRepairCurrentFixedAt           time.Time              `gorm:"COMMENT:当前检修时间;" json:"use_repair_current_fixed_at"`
	UseRepairCurrentFixerName         string                 `gorm:"type:VARCHAR(64);COMMENT:当前检修人;" json:"use_repair_current_fixer_name"`
	UseRepairCurrentCheckedAt         time.Time              `gorm:"COMMENT:当前验收时间;" json:"use_repair_current_checked_at"`
	UseRepairCurrentCheckerName       string                 `gorm:"type:VARCHAR(64);COMMENT:当前验收人;" json:"use_repair_current_checker_name"`
	UseRepairCurrentSpotCheckedAt     time.Time              `gorm:"COMMENT:当前抽验时间" json:"use_repair_current_spot_checked_at"`
	UseRepairCurrentSpotCheckerName   string                 `gorm:"type:VARCHAR(64);COMMENT:当前抽验人;" json:"use_repair_current_spot_checker_name"`
	UseRepairLastFixedAt              time.Time              `gorm:"COMMENT:上次检修时间;" json:"use_repair_last_fixed_at"`
	UseRepairLastFixerName            string                 `gorm:"type:VARCHAR(64);COMMENT:上次检修人;" json:"use_repair_last_fixer_name"`
	UseRepairLastCheckedAt            time.Time              `gorm:"COMMENT:上次验收时间;" json:"use_repair_last_checked_at"`
	UseRepairLastCheckerName          string                 `gorm:"type:VARCHAR(64);COMMENT:上次验收人;" json:"use_repair_last_checker_name"`
	UseRepairLastSpotCheckedAt        time.Time              `gorm:"COMMENT:上次抽验时间" json:"use_repair_last_spot_checked_at"`
	UseRepairLastSpotCheckerName      string                 `gorm:"type:VARCHAR(64);COMMENT:上次抽验人;" json:"use_repair_last_spot_checker_name"`

	// 资产归属
	BelongToOrganizationRailwayUuid   string                     `gorm:"type:VARCHAR(36);COMMENT:资产归属-所属路局;" json:"belong_to_organization_railway_uuid"`
	BelongToOrganizationRailway       OrganizationRailwayModel   `gorm:"foreignKey:BelongToOrganizationRailwayUuid;references:Uuid;COMMENT:资产归属-所属路局;" json:"belong_to_organization_railway"`
	BelongToOrganizationParagraphUuid string                     `gorm:"type:VARCHAR(36);COMMENT:资产归属-所属站段;" json:"belong_to_organization_paragraph_uuid"`
	BelongToOrganizationParagraph     OrganizationParagraphModel `gorm:"foreignKey:BelongToOrganizationParagraphUuid;references:Uuid;COMMENT:资产归属-所属站段;" json:"belong_to_organization_paragraph"`
	BelongToOrganizationWorkshopUuid  string                     `gorm:"type:VARCHAR(36);COMMENT:资产归属-所属车间;" json:"belong_to_organization_workshop_uuid"`
	BelongToOrganizationWorkshop      OrganizationWorkshopModel  `gorm:"foreignKey:BelongToOrganizationWorkshopUuid;references:Uuid;COMMENT:资产归属-所属专业车间;" json:"belong_to_organization_workshop"`
	BelongToOrganizationWorkAreaUuid  string                     `gorm:"type:VARCHAR(36);COMMENT:资产归属-所属工区;" json:"belong_to_organization_work_area_uuid"`
	BelongToOrganizationWorkArea      OrganizationWorkAreaModel  `gorm:"foreignKey:BelongToOrganizationWorkAreaUuid;references:Uuid;COMMENT:资产归属-所属专业工区;" json:"belong_to_organization_work_area"`

	// 使用地点
	UsePlaceCurrentOrganizationWorkshopUuid string                    `gorm:"type:VARCHAR(36);COMMENT:当前使用地点-车间UUID;" json:"use_place_current_organization_workshop_uuid"`
	UsePlaceCurrentOrganizationWorkshop     OrganizationWorkshopModel `gorm:"foreignKey:UsePlaceCurrentOrganizationWorkshopUuid;references:Uuid;COMMENT:当前使用地点-车间;" json:"use_place_current_organization_workshop"`
	UsePlaceCurrentOrganizationWorkAreaUuid string                    `gorm:"type:VARCHAR(36);COMMENT:当前使用地点-工区UUID;" json:"use_place_current_organization_work_area_uuid"`
	UsePlaceCurrentOrganizationWorkArea     OrganizationWorkAreaModel `gorm:"foreignKey:UsePlaceCurrentOrganizationWorkAreaUuid;references:Uuid;COMMENT:当前使用地点-工区;" json:"use_place_current_organization_work_area"`
	UsePlaceCurrentLocationLineUuid         string                    `gorm:"type:VARCHAR(36);COMMENT:当前使用地点-线别UUID;" json:"use_place_current_location_line_uuid"`
	UsePlaceCurrentLocationLine             LocationLineModel         `gorm:"foreignKey:UsePlaceCurrentLocationLineUuid;references:Uuid;COMMENT:当前使用地点-线别" json:"use_place_current_location_line"`
	UsePlaceCurrentLocationStationUuid      string                    `gorm:"type:VARCHAR(36);COMMENT:当前使用地点-站场UUID;" json:"use_place_current_location_station_uuid"`
	UsePlaceCurrentLocationStation          LocationStationModel      `gorm:"foreignKey:UsePlaceCurrentLocationStationUuid;references:Uuid;COMMENT:当前使用地点-站场" json:"use_place_current_location_station"`
	UsePlaceCurrentLocationSectionUuid      string                    `gorm:"type:VARCHAR(36);COMMENT:当前使用地点-区间UUID;" json:"use_place_current_location_section_uuid"`
	UsePlaceCurrentLocationSection          LocationSectionModel      `gorm:"foreignKey:UsePlaceCurrentLocationSectionUuid;references:Uuid;COMMENT:当前使用地点-区间" json:"use_place_current_location_section"`
	UsePlaceCurrentLocationCenterUuid       string                    `gorm:"type:VARCHAR(36);COMMENT:当前使用地点-中心UUID;" json:"use_place_current_location_center_uuid"`
	UsePlaceCurrentLocationCenter           LocationCenterModel       `gorm:"foreignKey:UsePlaceCurrentLocationCenterUuid;references:Uuid;COMMENT:当前使用地点-中心" json:"use_place_current_location_center"`
	UsePlaceCurrentLocationRailroadUuid     string                    `gorm:"type:VARCHAR(36);COMMENT:当前使用地点-道口UUID;" json:"use_place_current_location_railroad_grade_cross_uuid"`
	UsePlaceCurrentLocationRailroad         LocationRailroadModel     `gorm:"foreignKey:UsePlaceCurrentLocationRailroadUuid;references:Uuid;COMMENT:当前使用地点-道口" json:"use_place_current_location_railroad_grade_cross"`
	UsePlaceLastOrganizationWorkshopUuid    string                    `gorm:"type:VARCHAR(36);COMMENT:上次使用地点-车间UUID;" json:"use_place_last_organization_workshop_uuid"`
	UsePlaceLastOrganizationWorkshop        OrganizationWorkshopModel `gorm:"foreignKey:UsePlaceLastOrganizationWorkshopUuid;references:Uuid;COMMENT:上次使用地点-车间;" json:"use_place_last_organization_workshop"`
	UsePlaceLastOrganizationWorkAreaUuid    string                    `gorm:"type:VARCHAR(36);COMMENT:上次使用地点-工区UUID;" json:"use_place_last_organization_work_area_uuid"`
	UsePlaceLastOrganizationWorkArea        OrganizationWorkAreaModel `gorm:"foreignKey:UsePlaceLastOrganizationWorkAreaUuid;references:Uuid;COMMENT:上次使用地点-工区;" json:"use_place_last_organization_work_area"`
	UsePlaceLastLocationLineUuid            string                    `gorm:"type:VARCHAR(36);COMMENT:上次使用地点-线别UUID;" json:"use_place_last_location_line_uuid"`
	UsePlaceLastLocationLine                LocationLineModel         `gorm:"foreignKey:UsePlaceLastLocationLineUuid;references:Uuid;COMMENT:上次使用地点-线别" json:"use_place_last_location_line"`
	UsePlaceLastLocationStationUuid         string                    `gorm:"type:VARCHAR(36);COMMENT:上次使用地点-站场UUID;" json:"use_place_last_location_station_uuid"`
	UsePlaceLastLocationStation             LocationStationModel      `gorm:"foreignKey:UsePlaceLastLocationStationUuid;references:Uuid;COMMENT:上次使用地点-站场" json:"use_place_last_location_station"`
	UsePlaceLastLocationSectionUuid         string                    `gorm:"type:VARCHAR(36);COMMENT:上次使用地点-区间UUID;" json:"use_place_last_location_section_uuid"`
	UsePlaceLastLocationSection             LocationSectionModel      `gorm:"foreignKey:UsePlaceLastLocationSectionUuid;references:Uuid;COMMENT:上次使用地点-区间" json:"use_place_last_location_section"`
	UsePlaceLastLocationCenterUuid          string                    `gorm:"type:VARCHAR(36);COMMENT:上次使用地点-中心UUID;" json:"use_place_last_location_center_uuid"`
	UsePlaceLastLocationCenter              LocationCenterModel       `gorm:"foreignKey:UsePlaceLastLocationCenterUuid;references:Uuid;COMMENT:上次使用地点-中心" json:"use_place_last_location_center"`
	UsePlaceLastLocationRailroadUuid        string                    `gorm:"type:VARCHAR(36);COMMENT:上次使用地点-道口UUID;" json:"use_place_last_location_railroad_uuid"`
	UsePlaceLastLocationRailroad            LocationRailroadModel     `gorm:"foreignKey:UsePlaceLastLocationRailroadUuid;references:Uuid;COMMENT:上次使用地点-道口" json:"use_place_last_location_railroad"`

	// 使用位置
	UsePlaceCurrentPositionIndoorCellUuid string                  `gorm:"type:VARCHAR(36);COMMENT:当前使用位置-室内上道位置UUID;" json:"use_place_current_position_indoor_cell_uuid"`
	UsePlaceCurrentPositionIndoorCell     PositionIndoorCellModel `gorm:"foreignKey:UsePlaceCurrentPositionIndoorCellUuid;references:Uuid;COMMENT:当前使用位置-室内上道位置" json:"use_place_current_position_indoor_cell"`
	UsePlaceLastPositionIndoorCellUuid    string                  `gorm:"type:VARCHAR(36);COMMENT:上次使用位置-室内上道位置UUID;" json:"use_place_last_position_indoor_cell_uuid"`
	UsePlaceLastPositionIndoorCell        PositionIndoorCellModel `gorm:"foreignKey:UsePlaceLastPositionIndoorCellUuid;references:Uuid;COMMENT:上次使用位置-室内上道位置" json:"use_place_last_position_indoor_cell"`

	// 其他
	ExCycleRepairYear int16 `gorm:"type:INT2;DEFAULT:0;COMMENT:周期修年;" json:"ex_cycle_repair_year"`
	ExLifeYear        int16 `gorm:"type:INT2;DEFAULT:15;COMMENT:;" json:"ex_life_year"`

	// 关联
	EntireInstanceLock *EntireInstanceLockModel `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:相关锁;" json:"entire_instance_lock"`

	//LockName        string `gorm:"type:VARCHAR(64);COMMENT:锁名称;" json:"lock_name"`
	//LockDescription string `gorm:"type:TEXT;COMMENT:锁说明;" json:"lock_description"`
	//EntireInstanceUses             []EntireInstanceUseModel   `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:使用信息;" json:"entire_instance_uses"`
	//EntireInstanceLogs []EntireInstanceLogModel `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:相关日志;" json:"entire_instance_logs"`
	//EntireInstanceRepairs               []EntireInstanceRepairModel    `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:相关检修记录;" json:"entire_instance_repairs"`
}

// TableName 表名称
//  @receiver EntireInstanceModel
//  @return string
func (EntireInstanceModel) TableName() string {
	return "entire_instances"
}
