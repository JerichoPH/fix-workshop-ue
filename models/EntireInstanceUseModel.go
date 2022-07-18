package models

import (
	"time"
)

type EntireInstanceUseModel struct {
	BaseModel
	EntireInstanceIdentityCode         string                `gorm:"type:VARCHAR(20);COMMENT:所属器材;" json:"entire_instance_identity_code"`
	EntireInstance                     EntireInstanceModel          `gorm:"constraint:OnUpdate:CASCADE;foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:所属器材;" json:"entire_instance"`
	InAt                               time.Time                    `gorm:"type:DATETIME;COMMENT:入所时间;" json:"in_at"`
	PrevInAt                           time.Time                    `gorm:"type:DATETIME;COMMENT:上一次入所时间;" json:"prev_in_at"`
	OutAt                              time.Time                    `gorm:"type:DATETIME;COMMENT:出所时间;" json:"out_at"`
	PrevOutAt                          time.Time                    `gorm:"type:DATETIME;COMMENT:上一次出所时间;" json:"prev_out_at"`
	OrganizationLineUniqueCode         string                       `gorm:"type:CHAR(5);COMMENT:所属线别代码;" json:"organization_line_unique_code"`
	OrganizationLine                   OrganizationLineModel        `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationLineUniqueCode;references:UniqueCode;COMMENT:所属线别;" json:"organization_line"`
	PrevOrganizationLineUniqueCode     string                       `gorm:"type:CHAR(5);COMMENT:上一次线别代码;" json:"prev_organization_line_unique_code"`
	PrevOrganizationLine               OrganizationLineModel        `gorm:"constraint:OnUpdate:CASCADE;foreignKey:PrevOrganizationLineUniqueCode;references:UniqueCode;COMMENT:上一次所属线别;" json:"prev_organization_line"`
	OrganizationWorkshopUniqueCode     string                       `gorm:"type:CHAR(7);COMMENT:所属车间代码;" json:"organization_workshop_unique_code"`
	OrganizationWorkshop               OrganizationWorkshopModel    `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkshopUUID;references:UniqueCode;COMMENT:所属车间;" json:"organization_workshop"`
	PrevOrganizationWorkshopUniqueCode string                       `gorm:"type:CHAR(7);COMMENT:上一次所属车间代码;" json:"prev_organization_workshop_unique_code"`
	PrevOrganizationWorkshop              OrganizationWorkshopModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:PrevOrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:上一次所属车间;" json:"prev_organization_workshop"`
	OrganizationStationUniqueCode         string                    `gorm:"type:CHAR(6);COMMENT:所属车站代码;" json:"organization_station_unique_code"`
	OrganizationStation                   OrganizationStationModel  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationStationUniqueCode;references:UniqueCode;COMMENT:所属车间;" json:"organization_station"`
	PrevOrganizationStationUniqueCode     string                    `gorm:"type:CHAR(6);COMMENT:上一次所属车站代码;" json:"prev_organization_station_unique_code"`
	PrevOrganizationStation               OrganizationStationModel  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:PrevOrganizationWorkshopUniqueCode;references:UniqueCode;COMMENT:上一次所属车间;" json:"prev_organization_station"`
	OrganizationWorkAreaUniqueCode        string                    `gorm:"type:CHAR(8);COMMENT:所属现场工区代码;" json:"organization_work_area_unique_code"`
	OrganizationWorkArea                  OrganizationWorkAreaModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationWorkAreaUniqueCode;references:UniqueCode;COMMENT:所属工区;" json:"organization_work_area"`
	PrevOrganizationWorkAreaUniqueCode    string                    `gorm:"type:CHAR(8);COMMENT:上一次所属现场工区代码;" json:"prev_organization_work_area_unique_code"`
	PrevOrganizationWorkArea              OrganizationWorkAreaModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:PrevOrganizationWorkAreaUniqueCode;references:UniqueCode;COMMENT:上一次所属工区;" json:"prev_organization_work_area"`
	LocationInstallPositionUniqueCode     string                    `gorm:"type:VARCHAR(64);COMMENT:所属室内上道位置代码;" json:"location_install_position_unique_code"`
	LocationInstallPosition               LocationIndoorCellModel   `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationInstallPositionUniqueCode;references:UniqueCode;COMMENT:所属室内上道位置;" json:"location_install_position"`
	PrevLocationInstallPositionUniqueCode string                    `gorm:"type:VARCHAR(64);COMMENT:上一次所属室内上道位置代码;" json:"prev_location_install_position_unique_code"`
	PrevLocationInstallPosition           LocationIndoorCellModel   `gorm:"constraint:OnUpdate:CASCADE;foreignKey:PrevLocationInstallPositionUniqueCode;references:UniqueCode;COMMENT:上一次所属室内上道位置;" json:"prev_location_install_position"`
	CrossroadName                         string                    `gorm:"type:VARCHAR(64);COMMENT:所属道岔号名称;" json:"crossroad_name"`
	PrevCrossroadName                     string                    `gorm:"type:VARCHAR(64);COMMENT:上一次所属道岔号名称;" json:"prev_crossroad_name"`
	OpenDirection                         string                    `gorm:"type:VARCHAR(64);COMMENT:开向;" json:"open_direction"`
	PrevOpenDirection                     string                    `gorm:"type:VARCHAR(64);COMMENT:上一次开向;" json:"prev_open_direction"`
	OrganizationSectionName               string                    `gorm:"type:VARCHAR(64);COMMENT:所属区间名称;" json:"organization_section_name"`
	PrevOrganizationSectionName           string                    `gorm:"type:VARCHAR(64);COMMENT:上一次所属区间名称;" json:"prev_organization_section_name"`
	OrganizationSection                   OrganizationSectionModel  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationSectionUniqueCode;references:UniqueCode;COMMENT:所属区间;" json:"organization_section"`
	OrganizationSectionUniqueCode         string                    `gorm:"type:CHAR(6);COMMENT:所属区间代码;" json:"organization_section_unique_code"`
	PrevOrganizationSection               OrganizationSectionModel  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:PrevOrganizationSectionUniqueCode;references:UniqueCode;COMMENT:上一次所属区间;" json:"prev_organization_section"`
	PrevOrganizationSectionUniqueCode     string                    `gorm:"type:CHAR(6);COMMENT:上一次所属区间代码;" json:"prev_organization_section_unique_code"`
	SendOrReceive                                          string                                               `gorm:"type:VARCHAR(64);COMMENT:送/受端;" json:"send_or_receive"`
	PrevSendOrReceive                                      string                                               `gorm:"type:VARCHAR(64);COMMENT:上一次送/受端;" json:"prev_send_or_receive"`
	LocationSignalPostMainOrIndicatorUniqueCode            string                                               `gorm:"type:CHAR(6);COMMENT:信号机主机或表示器;" json:"location_signal_post_main_or_indicator_unique_code"`
	LocationSignalPostMainOrIndicator                      LocationOutdoorSignalPostMainOrIndicatorModel        `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationSignalPostMainOrIndicatorUniqueCode;references:UniqueCode;COMMENT:所属信号灯主机或表示器;" json:"location_signal_post_main_or_indicator"`
	PrevLocationSignalPostMainOrIndicatorUniqueCode        string                                               `gorm:"type:CHAR(6);COMMENT:上一次信号机主机或表示器;" json:"prev_location_signal_post_main_or_indicator_unique_code"`
	PrevLocationSignalPostMainOrIndicator                  LocationOutdoorSignalPostMainOrIndicatorModel        `gorm:"constraint:OnUpdate:CASCADE;foreignKey:PrevLocationSignalPostMainOrIndicatorUniqueCode;references:UniqueCode;COMMENT:上一次所属信号灯主机或表示器;" json:"prev_location_signal_post_main_or_indicator"`
	LocationSignalPostMainLightPositionUniqueCode          string                                               `gorm:"type:CHAR(2);COMMENT:信号机主机灯位代码;" json:"location_signal_post_main_light_position_unique_code"`
	LocationSignalPostMainLightPosition                    LocationOutdoorSignalPostMainLightPositionModel      `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationSignalPostMainLightPositionUniqueCode;references:UniqueCode;COMMENT:所属信号灯主机灯位代码;" json:"location_signal_post_main_light_position"`
	PrevLocationSignalPostMainLightPositionUniqueCode      string                                               `gorm:"type:CHAR(2);COMMENT:上一次信号机主机灯位代码;" json:"prev_location_signal_post_main_light_position_unique_code"`
	PrevLocationSignalPostMainLightPosition                LocationOutdoorSignalPostMainLightPositionModel      `gorm:"constraint:OnUpdate:CASCADE;foreignKey:PrevLocationSignalPostMainLightPositionUniqueCode;references:UniqueCode;COMMENT:上一次所属信号灯主机灯位代码;" json:"prev_location_signal_post_main_light_position"`
	LocationSignalPostIndicatorLightPositionUniqueCode     string                                               `gorm:"type:CHAR(2);COMMENT:信号机表示器灯位代码;" json:"location_signal_post_indicator_light_position_unique_code"`
	LocationSignalPostIndicatorLightPosition               LocationOutdoorSignalPostIndicatorLightPositionModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:LocationSignalPostIndicatorLightPositionUniqueCode;references:UniqueCode;COMMENT:所属信号灯表示器灯位代码;" json:"location_signal_post_indicator_light_position"`
	PrevLocationSignalPostIndicatorLightPositionUniqueCode string                                               `gorm:"type:CHAR(2);COMMENT:上一次信号机表示器灯位代码;" json:"prev_location_signal_post_indicator_light_position_unique_code"`
	PrevLocationSignalPostIndicatorLightPosition           LocationOutdoorSignalPostIndicatorLightPositionModel `gorm:"constraint:OnUpdate:CASCADE;foreignKey:PrevLocationSignalPostIndicatorLightPositionUniqueCode;references:UniqueCode;COMMENT:上一次所属信号灯表示器灯位代码;" json:"prev_location_signal_post_indicator_light_position"`
	OrganizationRailroadGradeCrossUniqueCode               string                                               `gorm:"type:CHAR(5);COMMENT:道口代码;" json:"organization_railroad_grade_cross_unique_code"`
	OrganizationRailroadGradeCross                         OrganizationRailroadGradeCrossModel                  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationRailroadGradeCrossUniqueCode;references:UniqueCode;COMMENT:所属道口;" json:"organization_railroad_grade_cross"`
	PrevOrganizationRailroadGradeCrossUniqueCode           string                                               `gorm:"type:CHAR(5);COMMENT:上一次道口代码;" json:"prev_organization_railroad_grade_cross_unique_code"`
	PrevOrganizationRailroadGradeCross                     OrganizationRailroadGradeCrossModel                  `gorm:"constraint:OnUpdate:CASCADE;foreignKey:PrevOrganizationRailroadGradeCrossUniqueCode;references:UniqueCode;COMMENT:上一次所属道口;" json:"prev_organization_railroad_grade_cross"`
	OrganizationCenterUniqueCode                           string                                               `gorm:"type:CHAR(6);COMMENT:中心代码;" json:"organization_center_unique_code"`
	OrganizationCenter                                     OrganizationCenterModel                              `gorm:"constraint:OnUpdate:CASCADE;foreignKey:OrganizationCenterUniqueCode;references:UniqueCode;COMMENT:所属中心;" json:"organization_center"`
	PrevOrganizationCenterUniqueCode                       string                                               `gorm:"type:CHAR(6);COMMENT:上一次中心代码;" json:"prev_organization_center_unique_code"`
	PrevOrganizationCenter                                 OrganizationCenterModel                              `gorm:"constraint:OnUpdate:CASCADE;foreignKey:PrevOrganizationCenterUniqueCode;references:UniqueCode;COMMENT:上一次所属中心;" json:"prev_organization_center"`
}

// TableName 表名称
func (cls *EntireInstanceUseModel) TableName() string {
	return "entire_instance_uses"
}