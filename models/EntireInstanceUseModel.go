package models

import (
	"time"
)

type EntireInstanceUseModel struct {
	BaseModel
	EntireInstanceIdentityCode                       string                                               `gorm:"type:VARCHAR(20);COMMENT:所属器材;" json:"entire_instance_identity_code"`
	EntireInstance                                   EntireInstanceModel                                  `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:所属器材;" json:"entire_instance"`
	InAt                                             time.Time                                            `gorm:"type:DATETIME;COMMENT:入所时间;" json:"in_at"`
	PrevInAt                                         time.Time                                            `gorm:"type:DATETIME;COMMENT:上一次入所时间;" json:"prev_in_at"`
	OutAt                                            time.Time                                            `gorm:"type:DATETIME;COMMENT:出所时间;" json:"out_at"`
	PrevOutAt                                        time.Time                                            `gorm:"type:DATETIME;COMMENT:上一次出所时间;" json:"prev_out_at"`
	OrganizationLineUuid                             string                                               `gorm:"type:VARCHAR(36);COMMENT:所属线别代码;" json:"organization_line_unique_code"`
	OrganizationLine                                 LocationLineModel                                    `gorm:"foreignKey:OrganizationLineUuid;references:Uuid;COMMENT:所属线别;" json:"organization_line"`
	PrevOrganizationLineUuid                         string                                               `gorm:"type:VARCHAR(36);COMMENT:上一次线别代码;" json:"prev_organization_line_uuid"`
	PrevOrganizationLine                             LocationLineModel                                    `gorm:"foreignKey:PrevOrganizationLineUuid;references:Uuid;COMMENT:上一次所属线别;" json:"prev_organization_line"`
	OrganizationWorkshopUuid                         string                                               `gorm:"type:VARCHAR(36);COMMENT:所属车间代码;" json:"organization_workshop_uuid"`
	OrganizationWorkshop                             OrganizationWorkshopModel                            `gorm:"foreignKey:OrganizationWorkshopUuid;references:Uuid;COMMENT:所属车间;" json:"organization_workshop"`
	PrevOrganizationWorkshopUuid                     string                                               `gorm:"type:VARCHAR(36);COMMENT:上一次所属车间代码;" json:"prev_organization_workshop_uuid"`
	PrevOrganizationWorkshop                         OrganizationWorkshopModel                            `gorm:"foreignKey:PrevOrganizationWorkshopUuid;references:Uuid;COMMENT:上一次所属车间;" json:"prev_organization_workshop"`
	OrganizationStationUuid                          string                                               `gorm:"type:VARCHAR(36);COMMENT:所属车站代码;" json:"organization_station_uuid"`
	OrganizationStation                              LocationStationModel                                 `gorm:"foreignKey:LocationStationUuid;references:Uuid;COMMENT:所属车间;" json:"organization_station"`
	PrevOrganizationStationUuid                      string                                               `gorm:"type:VARCHAR(36);COMMENT:上一次所属车站代码;" json:"prev_organization_station_uuid"`
	PrevOrganizationStation                          LocationStationModel                                 `gorm:"foreignKey:PrevOrganizationWorkshopUuid;references:Uuid;COMMENT:上一次所属车间;" json:"prev_organization_station"`
	OrganizationWorkAreaUuid                         string                                               `gorm:"type:CHAR(8);COMMENT:所属现场工区代码;" json:"organization_work_area_uuid"`
	OrganizationWorkArea                             OrganizationWorkAreaModel                            `gorm:"foreignKey:OrganizationWorkAreaUuid;references:Uuid;COMMENT:所属工区;" json:"organization_work_area"`
	PrevOrganizationWorkAreaUuid                     string                                               `gorm:"type:CHAR(8);COMMENT:上一次所属现场工区代码;" json:"prev_organization_work_area_uuid"`
	PrevOrganizationWorkArea                         OrganizationWorkAreaModel                            `gorm:"foreignKey:PrevOrganizationWorkAreaUuid;references:Uuid;COMMENT:上一次所属工区;" json:"prev_organization_work_area"`
	LocationInstallPositionUuid                      string                                               `gorm:"type:VARCHAR(64);COMMENT:所属室内上道位置代码;" json:"location_install_position_uuid"`
	LocationInstallPosition                          PositionIndoorCellModel                              `gorm:"foreignKey:LocationInstallPositionUuid;references:Uuid;COMMENT:所属室内上道位置;" json:"location_install_position"`
	PrevLocationInstallPositionUuid                  string                                               `gorm:"type:VARCHAR(64);COMMENT:上一次所属室内上道位置代码;" json:"prev_location_install_position_uuid"`
	PrevLocationInstallPosition                      PositionIndoorCellModel                              `gorm:"foreignKey:PrevLocationInstallPositionUuid;references:Uuid;COMMENT:上一次所属室内上道位置;" json:"prev_location_install_position"`
	CrossroadName                                    string                                               `gorm:"type:VARCHAR(64);COMMENT:所属道岔号名称;" json:"crossroad_name"`
	PrevCrossroadName                                string                                               `gorm:"type:VARCHAR(64);COMMENT:上一次所属道岔号名称;" json:"prev_crossroad_name"`
	OpenDirection                                    string                                               `gorm:"type:VARCHAR(64);COMMENT:开向;" json:"open_direction"`
	PrevOpenDirection                                string                                               `gorm:"type:VARCHAR(64);COMMENT:上一次开向;" json:"prev_open_direction"`
	OrganizationSectionName                          string                                               `gorm:"type:VARCHAR(64);COMMENT:所属区间名称;" json:"organization_section_name"`
	PrevOrganizationSectionName                      string                                               `gorm:"type:VARCHAR(64);COMMENT:上一次所属区间名称;" json:"prev_organization_section_name"`
	OrganizationSection                              LocationSectionModel                                 `gorm:"foreignKey:OrganizationSectionUniqueCode;references:Uuid;COMMENT:所属区间;" json:"organization_section"`
	OrganizationSectionUniqueCode                    string                                               `gorm:"type:CHAR(6);COMMENT:所属区间代码;" json:"organization_section_uuid"`
	PrevOrganizationSection                          LocationSectionModel                                 `gorm:"foreignKey:PrevOrganizationSectionUniqueCode;references:Uuid;COMMENT:上一次所属区间;" json:"prev_organization_section"`
	PrevOrganizationSectionUniqueCode                string                                               `gorm:"type:CHAR(6);COMMENT:上一次所属区间代码;" json:"prev_organization_section_uuid"`
	SendOrReceive                                    string                                               `gorm:"type:VARCHAR(64);COMMENT:送/受端;" json:"send_or_receive"`
	PrevSendOrReceive                                string                                               `gorm:"type:VARCHAR(64);COMMENT:上一次送/受端;" json:"prev_send_or_receive"`
	LocationSignalPostMainOrIndicatorUuid            string                                               `gorm:"type:CHAR(6);COMMENT:信号机主机或表示器;" json:"location_signal_post_main_or_indicator_uuid"`
	LocationSignalPostMainOrIndicator                PositionOutdoorSignalPostMainOrIndicatorModel        `gorm:"foreignKey:LocationSignalPostMainOrIndicatorUuid;references:Uuid;COMMENT:所属信号灯主机或表示器;" json:"location_signal_post_main_or_indicator"`
	PrevLocationSignalPostMainOrIndicatorUuid        string                                               `gorm:"type:CHAR(6);COMMENT:上一次信号机主机或表示器;" json:"prev_location_signal_post_main_or_indicator_uuid"`
	PrevLocationSignalPostMainOrIndicator            PositionOutdoorSignalPostMainOrIndicatorModel        `gorm:"foreignKey:PrevLocationSignalPostMainOrIndicatorUuid;references:Uuid;COMMENT:上一次所属信号灯主机或表示器;" json:"prev_location_signal_post_main_or_indicator"`
	LocationSignalPostMainLightPositionUuid          string                                               `gorm:"type:CHAR(2);COMMENT:信号机主机灯位代码;" json:"location_signal_post_main_light_position_uuid"`
	LocationSignalPostMainLightPosition              PositionOutdoorSignalPostMainLightPositionModel      `gorm:"foreignKey:LocationSignalPostMainLightPositionUuid;references:Uuid;COMMENT:所属信号灯主机灯位代码;" json:"location_signal_post_main_light_position"`
	PrevLocationSignalPostMainLightPositionUuid      string                                               `gorm:"type:CHAR(2);COMMENT:上一次信号机主机灯位代码;" json:"prev_location_signal_post_main_light_position_uuid"`
	PrevLocationSignalPostMainLightPosition          PositionOutdoorSignalPostMainLightPositionModel      `gorm:"foreignKey:PrevLocationSignalPostMainLightPositionUuid;references:Uuid;COMMENT:上一次所属信号灯主机灯位代码;" json:"prev_location_signal_post_main_light_position"`
	LocationSignalPostIndicatorLightPositionUuid     string                                               `gorm:"type:CHAR(2);COMMENT:信号机表示器灯位代码;" json:"location_signal_post_indicator_light_position_uuid"`
	LocationSignalPostIndicatorLightPosition         PositionOutdoorSignalPostIndicatorLightPositionModel `gorm:"foreignKey:LocationSignalPostIndicatorLightPositionUuid;references:Uuid;COMMENT:所属信号灯表示器灯位代码;" json:"location_signal_post_indicator_light_position"`
	PrevLocationSignalPostIndicatorLightPositionUuid string                                               `gorm:"type:CHAR(2);COMMENT:上一次信号机表示器灯位代码;" json:"prev_location_signal_post_indicator_light_position_uuid"`
	PrevLocationSignalPostIndicatorLightPosition     PositionOutdoorSignalPostIndicatorLightPositionModel `gorm:"foreignKey:PrevLocationSignalPostIndicatorLightPositionUuid;references:Uuid;COMMENT:上一次所属信号灯表示器灯位代码;" json:"prev_location_signal_post_indicator_light_position"`
	OrganizationRailroadGradeCrossUuid               string                                               `gorm:"type:CHAR(5);COMMENT:道口代码;" json:"organization_railroad_grade_cross_uuid"`
	OrganizationRailroadGradeCross                   LocationRailroadGradeCrossModel                      `gorm:"foreignKey:OrganizationRailroadGradeCrossUuid;references:Uuid;COMMENT:所属道口;" json:"organization_railroad_grade_cross"`
	PrevOrganizationRailroadGradeCrossUuid           string                                               `gorm:"type:CHAR(5);COMMENT:上一次道口代码;" json:"prev_organization_railroad_grade_cross_uuid"`
	PrevOrganizationRailroadGradeCross               LocationRailroadGradeCrossModel                      `gorm:"foreignKey:PrevOrganizationRailroadGradeCrossUuid;references:Uuid;COMMENT:上一次所属道口;" json:"prev_organization_railroad_grade_cross"`
	OrganizationCenterUuid                           string                                               `gorm:"type:CHAR(6);COMMENT:中心代码;" json:"organization_center_uuid"`
	OrganizationCenter                               LocationCenterModel                                  `gorm:"foreignKey:LocationCenterUuid;references:Uuid;COMMENT:所属中心;" json:"organization_center"`
	PrevOrganizationCenterUuid                       string                                               `gorm:"type:CHAR(6);COMMENT:上一次中心代码;" json:"prev_organization_center_uuid"`
	PrevOrganizationCenter                           LocationCenterModel                                  `gorm:"foreignKey:PrevOrganizationCenterUuid;references:Uuid;COMMENT:上一次所属中心;" json:"prev_organization_center"`
}

// TableName 表名称
func (EntireInstanceUseModel) TableName() string {
	return "entire_instance_uses"
}
